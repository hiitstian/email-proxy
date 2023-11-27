version="0.0.5" && \
build_dir="./build" && \
app_name="mail-proxy" && \
export EMAIL_PROXY_LOG_PATH="./email-logs.txt"
export EMAIL_PROXY_PORT="9004"

doBuild () {
    # goos: linux | darwin | windows
    # arch: amd64 | arm64
    os=$1 && \
    arch=$2 && \
    mkdir -p "./$build_dir" && \
    app_version_name=$(echo $app_name)_$(echo $os)_$(echo $arch)_$(echo $version)
    CGO_ENABLED=0  GOOS=$os  GOARCH=$arch  go build -o "./$build_dir/$app_version_name"  && \
    echo " build complete. Check ./$build_dir/$app_version_name" 
}

rm $build_dir/$app_name*  &&\

if [ "$1" == "" ]
then
    doBuild $(uname | awk '{print tolower($0)}') $(uname -m)
    exit 1
else
    if [ "$1" == "all" ]
    then
        doBuild 'linux' 'amd64'
        doBuild 'darwin' 'arm64'
    else
        doBuild $1 $2
    fi
fi