#!/bin/bash
# https://github.com/complexorganizations/go-build-script

# Pre-Checks system requirements
function installing-system-requirements() {
    if [ ! -x "$(command -v checksum)" ]; then
        echo "The application checksum was not found in the system. [ https://github.com/complexorganizations/checksum ]"
        exit
    elif [ ! -x "$(command -v go)" ]; then
        echo "The application go was not found in the system. [ https://go.dev ]"
        exit
    fi
}

# Run the function and check for requirements
installing-system-requirements

# Build for all the OS
function build-golang-app() {
    APPLICATION="AppName"
    VERSION="v1.0.0"
    SOURCE_CODE="./main.go"
    BIN="bin/"
    if [ -n "$(ls ./*.go)" ]; then
        # Aix
        GOOS=aix GOARCH=ppc64 go build -o ${BIN}${APPLICATION}-${VERSION}-aix-ppc64 ${SOURCE_CODE}
        # Android
        GOOS=android GOARCH=386 go build -o ${BIN}${APPLICATION}-${VERSION}-android-386 ${SOURCE_CODE}
        GOOS=android GOARCH=amd64 go build -o ${BIN}${APPLICATION}-${VERSION}-android-amd64 ${SOURCE_CODE}
        GOOS=android GOARCH=arm go build -o ${BIN}${APPLICATION}-${VERSION}-android-arm ${SOURCE_CODE}
        GOOS=android GOARCH=arm64 go build -o ${BIN}${APPLICATION}-${VERSION}-android-arm64 ${SOURCE_CODE}
        # Darwin
        GOOS=darwin GOARCH=amd64 go build -o ${BIN}${APPLICATION}-${VERSION}-darwin-amd64 ${SOURCE_CODE}
        GOOS=darwin GOARCH=arm64 go build -o ${BIN}${APPLICATION}-${VERSION}-darwin-arm64 ${SOURCE_CODE}
        # Dragonfly
        GOOS=dragonfly GOARCH=amd64 go build -o ${BIN}${APPLICATION}-${VERSION}-dragonfly-amd64 ${SOURCE_CODE}
        # Freebsd
        GOOS=freebsd GOARCH=386 go build -o ${BIN}${APPLICATION}-${VERSION}-freebsd-386 ${SOURCE_CODE}
        GOOS=freebsd GOARCH=amd64 go build -o ${BIN}${APPLICATION}-${VERSION}-freebsd-amd64 ${SOURCE_CODE}
        GOOS=freebsd GOARCH=arm go build -o ${BIN}${APPLICATION}-${VERSION}-freebsd-arm ${SOURCE_CODE}
        GOOS=freebsd GOARCH=arm64 go build -o ${BIN}${APPLICATION}-${VERSION}-freebsd-arm64 ${SOURCE_CODE}
        # Lllumos
        GOOS=illumos GOARCH=amd64 go build -o ${BIN}${APPLICATION}-${VERSION}-amd64 ${SOURCE_CODE}
        # JS
        GOOS=js GOARCH=wasm go build -o ${BIN}${APPLICATION}-${VERSION}-js-wasm ${SOURCE_CODE}
        # iOS
        GOOS=ios GOARCH=amd64 go build -o ${BIN}${APPLICATION}-${VERSION}-ios-amd64 ${SOURCE_CODE}
        GOOS=ios GOARCH=arm64 go build -o ${BIN}${APPLICATION}-${VERSION}-ios-arm64 ${SOURCE_CODE}
        # Linux
        GOOS=linux GOARCH=386 go build -o ${BIN}${APPLICATION}-${VERSION}-linux-386 ${SOURCE_CODE}
        GOOS=linux GOARCH=amd64 go build -o ${BIN}${APPLICATION}-${VERSION}-linux-amd64 ${SOURCE_CODE}
        GOOS=linux GOARCH=arm go build -o ${BIN}${APPLICATION}-${VERSION}-linux-arm ${SOURCE_CODE}
        GOOS=linux GOARCH=arm64 go build -o ${BIN}${APPLICATION}-${VERSION}-linux-arm64 ${SOURCE_CODE}
        GOOS=linux GOARCH=mips go build -o ${BIN}${APPLICATION}-${VERSION}-linux-mips ${SOURCE_CODE}
        GOOS=linux GOARCH=mips64 go build -o ${BIN}${APPLICATION}-${VERSION}-linux-mips64 ${SOURCE_CODE}
        GOOS=linux GOARCH=mips64le go build -o ${BIN}${APPLICATION}-${VERSION}-linux-mips64le ${SOURCE_CODE}
        GOOS=linux GOARCH=mipsle go build -o ${BIN}${APPLICATION}-${VERSION}-linux-mipsle ${SOURCE_CODE}
        GOOS=linux GOARCH=ppc64 go build -o ${BIN}${APPLICATION}-${VERSION}-linux-ppc64 ${SOURCE_CODE}
        GOOS=linux GOARCH=ppc64le go build -o ${BIN}${APPLICATION}-${VERSION}-linux-ppc64le ${SOURCE_CODE}
        GOOS=linux GOARCH=riscv64 go build -o ${BIN}${APPLICATION}-${VERSION}-linux-riscv64 ${SOURCE_CODE}
        GOOS=linux GOARCH=s390x go build -o ${BIN}${APPLICATION}-${VERSION}-linux-s390x ${SOURCE_CODE}
        # Nacl
        GOOS=nacl GOARCH=386 go build -o ${BIN}${APPLICATION}-${VERSION}-nacl-386 ${SOURCE_CODE}
        GOOS=nacl GOARCH=amd64p32 go build -o ${BIN}${APPLICATION}-${VERSION}-nacl-amd64p32 ${SOURCE_CODE}
        GOOS=nacl GOARCH=nacl-arm go build -o ${BIN}${APPLICATION}-${VERSION}-nacl-arm ${SOURCE_CODE}
        # Netbsd
        GOOS=netbsd GOARCH=386 go build -o ${BIN}${APPLICATION}-${VERSION}-netbsd-386 ${SOURCE_CODE}
        GOOS=netbsd GOARCH=amd64 go build -o ${BIN}${APPLICATION}-${VERSION}-netbsd-amd64 ${SOURCE_CODE}
        GOOS=netbsd GOARCH=arm go build -o ${BIN}${APPLICATION}-${VERSION}-netbsd-arm ${SOURCE_CODE}
        GOOS=netbsd GOARCH=arm64 go build -o ${BIN}${APPLICATION}-${VERSION}-netbsd-arm64 ${SOURCE_CODE}
        # Openbsd
        GOOS=openbsd GOARCH=386 go build -o ${BIN}${APPLICATION}-${VERSION}-openbsd-386 ${SOURCE_CODE}
        GOOS=openbsd GOARCH=amd64 go build -o ${BIN}${APPLICATION}-${VERSION}-openbsd-amd64 ${SOURCE_CODE}
        GOOS=openbsd GOARCH=arm go build -o ${BIN}${APPLICATION}-${VERSION}-openbsd-arm ${SOURCE_CODE}
        GOOS=openbsd GOARCH=arm64 go build -o ${BIN}${APPLICATION}-${VERSION}-openbsd-arm64 ${SOURCE_CODE}
        GOOS=openbsd GOARCH=mips64 go build -o ${BIN}${APPLICATION}-${VERSION}-openbsd-mips64 ${SOURCE_CODE}
        # Plan9
        GOOS=plan9 GOARCH=386 go build -o ${BIN}${APPLICATION}-${VERSION}-plan9-386 ${SOURCE_CODE}
        GOOS=plan9 GOARCH=amd64 go build -o ${BIN}${APPLICATION}-${VERSION}-plan9-amd64 ${SOURCE_CODE}
        GOOS=plan9 GOARCH=arm go build -o ${BIN}${APPLICATION}-${VERSION}-plan9-arm ${SOURCE_CODE}
        # Solaris
        GOOS=solaris GOARCH=amd64 go build -o ${BIN}${APPLICATION}-${VERSION}-solaris-amd64 ${SOURCE_CODE}
        # Windows
        GOOS=windows GOARCH=386 go build -o ${BIN}${APPLICATION}-${VERSION}-windows-386.exe ${SOURCE_CODE}
        GOOS=windows GOARCH=amd64 go build -o ${BIN}${APPLICATION}-${VERSION}-windows-amd64.exe ${SOURCE_CODE}
        GOOS=windows GOARCH=arm go build -o ${BIN}${APPLICATION}-${VERSION}-windows-arm.exe ${SOURCE_CODE}
        # Get SHA and put everything in a register.
        checksum -path="$BIN"
    else
        echo "Error: The \".go\" files could not be found."
        exit
    fi
}

# Start the build
build-golang-app
