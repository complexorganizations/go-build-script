#!/bin/bash

# Detect Operating System
function dist-check() {
    if [ -e /etc/os-release ]; then
        # shellcheck disable=SC1091
        source /etc/os-release
        DISTRO=$ID
    fi
}

# Check Operating System
dist-check

# Pre-Checks system requirements
function installing-system-requirements() {
    if { [ "$DISTRO" == "ubuntu" ] || [ "$DISTRO" == "debian" ] || [ "$DISTRO" == "raspbian" ] || [ "$DISTRO" == "pop" ] || [ "$DISTRO" == "kali" ] || [ "$DISTRO" == "linuxmint" ] || [ "$DISTRO" == "fedora" ] || [ "$DISTRO" == "centos" ] || [ "$DISTRO" == "rhel" ] || [ "$DISTRO" == "arch" ] || [ "$DISTRO" == "manjaro" ] || [ "$DISTRO" == "alpine" ] || [ "$DISTRO" == "freebsd" ]; }; then
        if [ ! -x "$(command -v sha1sum)" ]; then
            if { [ "$DISTRO" == "ubuntu" ] || [ "$DISTRO" == "debian" ] || [ "$DISTRO" == "raspbian" ] || [ "$DISTRO" == "pop" ] || [ "$DISTRO" == "kali" ] || [ "$DISTRO" == "linuxmint" ]; }; then
                sudo apt-get update && apt-get install coreutils zip unzip -y
            elif { [ "$DISTRO" == "fedora" ] || [ "$DISTRO" == "centos" ] || [ "$DISTRO" == "rhel" ]; }; then
                sudo yum update -y && yum install coreutils zip unzip -y
            elif { [ "$DISTRO" == "arch" ] || [ "$DISTRO" == "manjaro" ]; }; then
                sudo pacman -Syu --noconfirm iptables coreutils zip unzip
            elif [ "$DISTRO" == "alpine" ]; then
                sudo apk update && apk add coreutils zip unzip
            elif [ "$DISTRO" == "freebsd" ]; then
                sudo pkg update && pkg install coreutils zip unzip
            fi
        fi
    else
        echo "Error: $DISTRO not supported."
        exit
    fi
}

# Run the function and check for requirements
installing-system-requirements

# Build for all the OS
function build-golang-app() {
    APPLICATION="AppName"
    VERSION="1.0.0"
    if [ -x "$(command -v go)" ]; then
        GOOS=aix GOARCH=ppc64 go build -o build/$APPLICATION-$VERSION-aix-ppc64 .
        GOOS=android GOARCH=386 go build -o build/$APPLICATION-$VERSION-android-386 .
        GOOS=android GOARCH=amd64 go build -o build/$APPLICATION-$VERSION-android-amd64 .
        GOOS=android GOARCH=arm go build -o build/$APPLICATION-$VERSION-android-arm .
        GOOS=android GOARCH=arm64 go build -o build/$APPLICATION-$VERSION-android-arm64 .
        GOOS=darwin GOARCH=386 go build -o build/$APPLICATION-$VERSION-darwin-386 .
        GOOS=darwin GOARCH=amd64 go build -o build/$APPLICATION-$VERSION-darwin-amd64 .
        GOOS=darwin GOARCH=arm go build -o build/$APPLICATION-$VERSION-darwin-arm .
        GOOS=darwin GOARCH=arm64 go build -o build/$APPLICATION-$VERSION-darwin-arm64 .
        GOOS=dragonfly GOARCH=amd64 go build -o build/$APPLICATION-$VERSION-dragonfly-amd64 .
        GOOS=freebsd GOARCH=386 go build -o build/$APPLICATION-$VERSION-freebsd-386 .
        GOOS=freebsd GOARCH=amd64 go build -o build/$APPLICATION-$VERSION-freebsd-amd64 .
        GOOS=freebsd GOARCH=arm go build -o build/$APPLICATION-$VERSION-freebsd-arm .
        GOOS=freebsd GOARCH=arm64 go build -o build/$APPLICATION-$VERSION-freebsd-arm64 .
        GOOS=js GOARCH=wasm go build -o build/$APPLICATION-$VERSION-js-wasm .
        GOOS=illumos GOARCH=amd64 go build -o build/$APPLICATION-$VERSION-amd64 .
        GOOS=linux GOARCH=386 go build -o build/$APPLICATION-$VERSION-linux-386 .
        GOOS=linux GOARCH=amd64 go build -o build/$APPLICATION-$VERSION-linux-amd64 .
        GOOS=linux GOARCH=arm go build -o build/$APPLICATION-$VERSION-linux-arm .
        GOOS=linux GOARCH=arm64 go build -o build/$APPLICATION-$VERSION-linux-arm64 .
        GOOS=linux GOARCH=mips go build -o build/$APPLICATION-$VERSION-linux-mips .
        GOOS=linux GOARCH=mips64 go build -o build/$APPLICATION-$VERSION-linux-mips64 .
        GOOS=linux GOARCH=mips64le go build -o build/$APPLICATION-$VERSION-linux-mips64le .
        GOOS=linux GOARCH=mipsle go build -o build/$APPLICATION-$VERSION-linux-mipsle .
        GOOS=linux GOARCH=ppc64 go build -o build/$APPLICATION-$VERSION-linux-ppc64 .
        GOOS=linux GOARCH=ppc64le go build -o build/$APPLICATION-$VERSION-linux-ppc64le .
        GOOS=linux GOARCH=riscv64 go build -o build/$APPLICATION-$VERSION-linux-riscv64 .
        GOOS=linux GOARCH=s390x go build -o build/$APPLICATION-$VERSION-linux-s390x .
        GOOS=nacl GOARCH=386 go build -o build/$APPLICATION-$VERSION-nacl-386 .
        GOOS=nacl GOARCH=amd64p32 go build -o build/$APPLICATION-$VERSION-nacl-amd64p32 .
        GOOS=nacl GOARCH=nacl-arm go build -o build/$APPLICATION-$VERSION-nacl-arm .
        GOOS=netbsd GOARCH=386 go build -o build/$APPLICATION-$VERSION-netbsd-386 .
        GOOS=netbsd GOARCH=amd64 go build -o build/$APPLICATION-$VERSION-netbsd-amd64 .
        GOOS=netbsd GOARCH=arm go build -o build/$APPLICATION-$VERSION-netbsd-arm .
        GOOS=netbsd GOARCH=arm64 go build -o build/$APPLICATION-$VERSION-netbsd-arm64 .
        GOOS=openbsd GOARCH=386 go build -o build/$APPLICATION-$VERSION-openbsd-386 .
        GOOS=openbsd GOARCH=amd64 go build -o build/$APPLICATION-$VERSION-openbsd-amd64 .
        GOOS=openbsd GOARCH=arm go build -o build/$APPLICATION-$VERSION-openbsd-arm .
        GOOS=openbsd GOARCH=arm64 go build -o build/$APPLICATION-$VERSION-openbsd-arm64 .
        GOOS=plan9 GOARCH=386 go build -o build/$APPLICATION-$VERSION-plan9-386 .
        GOOS=plan9 GOARCH=amd64 go build -o build/$APPLICATION-$VERSION-plan9-amd64 .
        GOOS=plan9 GOARCH=arm go build -o build/$APPLICATION-$VERSION-plan9-arm .
        GOOS=solaris GOARCH=amd64 go build -o build/$APPLICATION-$VERSION-solaris-amd64 .
        GOOS=windows GOARCH=386 go build -o build/$APPLICATION-$VERSION-windows-386.exe .
        GOOS=windows GOARCH=amd64 go build -o build/$APPLICATION-$VERSION-windows-amd64.exe .
        GOOS=windows GOARCH=arm go build -o build/$APPLICATION-$VERSION-windows-arm.exe .
        echo "$(find build/ -type f -print0 | xargs -0 sha1sum)" >>SHA-1
    else
        echo "Error: In your system, Go wasn't found."
        exit
    fi
}

# Start the build
build-golang-app
