#!/bin/bash
# https://github.com/complexorganizations/go-build-script

# Detect Operating System
function dist-check() {
    if [ -e /etc/os-release ]; then
        # shellcheck disable=SC1091
        source /etc/os-release
        DISTRO=${ID}
    fi
}

# Check Operating System
dist-check

# Pre-Checks system requirements
function installing-system-requirements() {
    if { [ "${DISTRO}" == "ubuntu" ] || [ "${DISTRO}" == "debian" ] || [ "${DISTRO}" == "raspbian" ] || [ "${DISTRO}" == "pop" ] || [ "${DISTRO}" == "kali" ] || [ "${DISTRO}" == "linuxmint" ] || [ "${DISTRO}" == "fedora" ] || [ "${DISTRO}" == "centos" ] || [ "${DISTRO}" == "rhel" ] || [ "${DISTRO}" == "arch" ] || [ "${DISTRO}" == "manjaro" ] || [ "${DISTRO}" == "alpine" ] || [ "${DISTRO}" == "freebsd" ]; }; then
        if { [ ! -x "$(command -v sha256sum)" ] || [ ! -x "$(command -v sha512sum)" ] || [ ! -x "$(command -v go)" ]; }; then
            if { [ "${DISTRO}" == "ubuntu" ] || [ "${DISTRO}" == "debian" ] || [ "${DISTRO}" == "raspbian" ] || [ "${DISTRO}" == "pop" ] || [ "${DISTRO}" == "kali" ] || [ "${DISTRO}" == "linuxmint" ]; }; then
                sudo apt-get update && sudo apt-get install coreutils golang-go -y
            elif { [ "${DISTRO}" == "fedora" ] || [ "${DISTRO}" == "centos" ] || [ "${DISTRO}" == "rhel" ]; }; then
                sudo yum update -y && sudo yum install coreutils golang -y
            elif { [ "${DISTRO}" == "arch" ] || [ "${DISTRO}" == "manjaro" ]; }; then
                sudo pacman -Syu --noconfirm --needed coreutils go
            elif [ "${DISTRO}" == "alpine" ]; then
                sudo apk update && sudo apk add coreutils go
            elif [ "${DISTRO}" == "freebsd" ]; then
                sudo pkg update && sudo pkg install coreutils go
            fi
        fi
    else
        echo "Error: ${DISTRO} not supported."
        exit
    fi
}

# Run the function and check for requirements
installing-system-requirements

# Build for all the OS
function build-golang-app() {
    APPLICATION="AppName"
    VERSION="1.0.0"
    if [ -n "$(ls ./*.go)" ]; then
        # Aix
        GOOS=aix GOARCH=ppc64 go build -o bin/${APPLICATION}-${VERSION}-aix-ppc64 .
        # Android
        GOOS=android GOARCH=386 go build -o bin/${APPLICATION}-${VERSION}-android-386 .
        GOOS=android GOARCH=amd64 go build -o bin/${APPLICATION}-${VERSION}-android-amd64 .
        GOOS=android GOARCH=arm go build -o bin/${APPLICATION}-${VERSION}-android-arm .
        GOOS=android GOARCH=arm64 go build -o bin/${APPLICATION}-${VERSION}-android-arm64 .
        # Darwin
        GOOS=darwin GOARCH=amd64 go build -o bin/${APPLICATION}-${VERSION}-darwin-amd64 .
        GOOS=darwin GOARCH=arm64 go build -o bin/${APPLICATION}-${VERSION}-darwin-arm64 .
        # Dragonfly
        GOOS=dragonfly GOARCH=amd64 go build -o bin/${APPLICATION}-${VERSION}-dragonfly-amd64 .
        # Freebsd
        GOOS=freebsd GOARCH=386 go build -o bin/${APPLICATION}-${VERSION}-freebsd-386 .
        GOOS=freebsd GOARCH=amd64 go build -o bin/${APPLICATION}-${VERSION}-freebsd-amd64 .
        GOOS=freebsd GOARCH=arm go build -o bin/${APPLICATION}-${VERSION}-freebsd-arm .
        GOOS=freebsd GOARCH=arm64 go build -o bin/${APPLICATION}-${VERSION}-freebsd-arm64 .
        # Lllumos
        GOOS=illumos GOARCH=amd64 go build -o bin/${APPLICATION}-${VERSION}-amd64 .
        # JS
        GOOS=js GOARCH=wasm go build -o bin/${APPLICATION}-${VERSION}-js-wasm .
        # iOS
        GOOS=ios GOARCH=amd64 go build -o bin/${APPLICATION}-${VERSION}-ios-amd64 .
        GOOS=ios GOARCH=arm64 go build -o bin/${APPLICATION}-${VERSION}-ios-arm64 .
        # Linux
        GOOS=linux GOARCH=386 go build -o bin/${APPLICATION}-${VERSION}-linux-386 .
        GOOS=linux GOARCH=amd64 go build -o bin/${APPLICATION}-${VERSION}-linux-amd64 .
        GOOS=linux GOARCH=arm go build -o bin/${APPLICATION}-${VERSION}-linux-arm .
        GOOS=linux GOARCH=arm64 go build -o bin/${APPLICATION}-${VERSION}-linux-arm64 .
        GOOS=linux GOARCH=mips go build -o bin/${APPLICATION}-${VERSION}-linux-mips .
        GOOS=linux GOARCH=mips64 go build -o bin/${APPLICATION}-${VERSION}-linux-mips64 .
        GOOS=linux GOARCH=mips64le go build -o bin/${APPLICATION}-${VERSION}-linux-mips64le .
        GOOS=linux GOARCH=mipsle go build -o bin/${APPLICATION}-${VERSION}-linux-mipsle .
        GOOS=linux GOARCH=ppc64 go build -o bin/${APPLICATION}-${VERSION}-linux-ppc64 .
        GOOS=linux GOARCH=ppc64le go build -o bin/${APPLICATION}-${VERSION}-linux-ppc64le .
        GOOS=linux GOARCH=riscv64 go build -o bin/${APPLICATION}-${VERSION}-linux-riscv64 .
        GOOS=linux GOARCH=s390x go build -o bin/${APPLICATION}-${VERSION}-linux-s390x .
        # Nacl
        GOOS=nacl GOARCH=386 go build -o bin/${APPLICATION}-${VERSION}-nacl-386 .
        GOOS=nacl GOARCH=amd64p32 go build -o bin/${APPLICATION}-${VERSION}-nacl-amd64p32 .
        GOOS=nacl GOARCH=nacl-arm go build -o bin/${APPLICATION}-${VERSION}-nacl-arm .
        # Netbsd
        GOOS=netbsd GOARCH=386 go build -o bin/${APPLICATION}-${VERSION}-netbsd-386 .
        GOOS=netbsd GOARCH=amd64 go build -o bin/${APPLICATION}-${VERSION}-netbsd-amd64 .
        GOOS=netbsd GOARCH=arm go build -o bin/${APPLICATION}-${VERSION}-netbsd-arm .
        GOOS=netbsd GOARCH=arm64 go build -o bin/${APPLICATION}-${VERSION}-netbsd-arm64 .
        # Openbsd
        GOOS=openbsd GOARCH=386 go build -o bin/${APPLICATION}-${VERSION}-openbsd-386 .
        GOOS=openbsd GOARCH=amd64 go build -o bin/${APPLICATION}-${VERSION}-openbsd-amd64 .
        GOOS=openbsd GOARCH=arm go build -o bin/${APPLICATION}-${VERSION}-openbsd-arm .
        GOOS=openbsd GOARCH=arm64 go build -o bin/${APPLICATION}-${VERSION}-openbsd-arm64 .
        GOOS=openbsd GOARCH=mips64 go build -o bin/${APPLICATION}-${VERSION}-openbsd-mips64 .
        # Plan9
        GOOS=plan9 GOARCH=386 go build -o bin/${APPLICATION}-${VERSION}-plan9-386 .
        GOOS=plan9 GOARCH=amd64 go build -o bin/${APPLICATION}-${VERSION}-plan9-amd64 .
        GOOS=plan9 GOARCH=arm go build -o bin/${APPLICATION}-${VERSION}-plan9-arm .
        # Solaris
        GOOS=solaris GOARCH=amd64 go build -o bin/${APPLICATION}-${VERSION}-solaris-amd64 .
        # Windows
        GOOS=windows GOARCH=386 go build -o bin/${APPLICATION}-${VERSION}-windows-386.exe .
        GOOS=windows GOARCH=amd64 go build -o bin/${APPLICATION}-${VERSION}-windows-amd64.exe .
        GOOS=windows GOARCH=arm go build -o bin/${APPLICATION}-${VERSION}-windows-arm.exe .
        # Get SHA and put everything in a register.
        find bin/ -type f -print0 | xargs -0 sha512sum
    else
        echo "Error: Failed to find \".go\" files."
        exit
    fi
}

# Start the build
build-golang-ap
