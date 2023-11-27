export EMAIL_PROXY_LOG_PATH="./logs/email_logs.txt" && \
export EMAIL_PROXY_PORT="9004" && \

app_folder_name=build && \

os=$(uname | awk '{print tolower($0)}') && \
arch=$(uname -m) && \

os_arch_tag=$(echo $app_name)_$(echo $os)_$(echo $arch) && \
exec_name=$(ls ./$app_folder_name| grep "$os_arch_tag") && \
echo "exec_name: $exec_name" && \
"$app_folder_name/$exec_name" && \
