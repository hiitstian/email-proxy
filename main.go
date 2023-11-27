package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

var (
	log_path *string
	port     *string
)

func init() {
	fmt.Println("args", os.Args)

	log_path = flag.String("log_path", os.Getenv("EMAIL_PROXY_LOG_PATH"), "Path to log file")
	port = flag.String("port", os.Getenv("EMAIL_PROXY_PORT"), "Port to listen on")

	flag.Parse()

}
func main() {

	fmt.Println("[MAIL-PROXY] Starting mail-proxy service")
	// SECTION - setup logging

	logFilePath := path.Join(*log_path)
	fmt.Println("[MAIL-PROXY][Info]log path: ", logFilePath)

	err := os.Mkdir(path.Dir(logFilePath), os.ModePerm)
	if err != nil {
		log.Printf("[MAIL-PROXY][WARN] creating log dir: %v", err)
	}
	logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("[MAIL-PROXY][Fatal] error opening log file: %v", err)
	}
	log.SetOutput(logFile)
	log.Println("[MAIL-PROXY][Info] Starting mail-proxy service")
	defer logFile.Close()

	apiPrefix := os.Getenv("EMAIL_PROXY_API_PREFIX")

	http.HandleFunc(apiPrefix+"/", isAliveHandler)
	http.HandleFunc(apiPrefix+"/send-mail", sendMailHandler)

	log.Println("Mail-proxy starting at port " + *port)
	fmt.Println("Mail-proxy starting at port " + *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

// SECTION - alive
func isAliveHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("email-proxy is alive\n"))
}

// SECTION - utils
type RES_Error struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
type RES_Success struct {
	Status  string `json:"status"`
	Payload any    `json:"payload"`
}

func makeErrorReponder(w http.ResponseWriter, endpoint string) func(err error, message string) {
	return func(err error, message string) {
		log.Printf("[Error] %s %s: %v", endpoint, message, err)
		json.NewEncoder(w).Encode(RES_Error{
			Status:  "error",
			Message: message,
		})
	}

}

// SECTION - Send Mail
func sendMailHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		w.Write([]byte("/send-mail email-proxy is alive\n"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	errorResponder := makeErrorReponder(w, "/send-mail")

	var mailConfig EmailConfig
	err := json.NewDecoder(r.Body).Decode(&mailConfig)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		errorResponder(err, "Failed to parse body")
		return
	}
	log.Println("sendMailHandler", mailConfig.From, mailConfig.To, mailConfig.Subject)

	err = sendEmail(mailConfig)
	if err != nil {
		errorResponder(err, "Failed to send email")
		return
	}
	log.Println("SENT: ", mailConfig.From, mailConfig.To, mailConfig.Subject)

	json.NewEncoder(w).Encode(RES_Success{
		Status:  "success",
		Payload: mailConfig,
	})
}
