curl -X POST \
-H 'Content-Type: application/json' \
-d @req.json \
"http://localhost:9004/send-mail"


curl -X POST \
-H 'Content-Type: application/json' \
-d @req.json \
"https://staging.jinguanwen.com/api/email-proxy/send-mail"