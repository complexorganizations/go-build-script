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
        if { [ ! -x "$(command -v sha1sum)" ] || [ ! -x "$(command -v shasum)" ]; }; then
            if { [ "$DISTRO" == "ubuntu" ] || [ "$DISTRO" == "debian" ] || [ "$DISTRO" == "raspbian" ] || [ "$DISTRO" == "pop" ] || [ "$DISTRO" == "kali" ] || [ "$DISTRO" == "linuxmint" ]; }; then
                sudo apt-get update && apt-get install coreutils -y
            elif { [ "$DISTRO" == "fedora" ] || [ "$DISTRO" == "centos" ] || [ "$DISTRO" == "rhel" ]; }; then
                sudo yum update -y && yum install coreutils -y
            elif { [ "$DISTRO" == "arch" ] || [ "$DISTRO" == "manjaro" ]; }; then
                sudo pacman -Syu --noconfirm iptables coreutils
            elif [ "$DISTRO" == "alpine" ]; then
                sudo apk update && apk add coreutils
            elif [ "$DISTRO" == "freebsd" ]; then
                sudo pkg update && pkg install coreutils
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
        GOOS=aix GOARCH=ppc64 go build -o build/$APPLICATION-aix-ppc64-$VERSION .
        GOOS=android GOARCH=386 go build -o build/$APPLICATION-android-386-$VERSION .
        GOOS=android GOARCH=amd64 go build -o build/$APPLICATION-android-amd64-$VERSION .
        GOOS=android GOARCH=arm go build -o build/$APPLICATION-android-arm-$VERSION .
        GOOS=android GOARCH=arm64 go build -o build/$APPLICATION-android-arm64-$VERSION .
        GOOS=darwin GOARCH=386 go build -o build/$APPLICATION-darwin-386-$VERSION .
        GOOS=darwin GOARCH=amd64 go build -o build/$APPLICATION-darwin-amd64-$VERSION .
        GOOS=darwin GOARCH=arm go build -o build/$APPLICATION-darwin-arm-$VERSION .
        GOOS=darwin GOARCH=arm64 go build -o build/$APPLICATION-darwin-arm64-$VERSION .
        GOOS=dragonfly GOARCH=amd64 go build -o build/$APPLICATION-dragonfly-amd64-$VERSION .
        GOOS=freebsd GOARCH=386 go build -o build/$APPLICATION-freebsd-386-$VERSION .
        GOOS=freebsd GOARCH=amd64 go build -o build/$APPLICATION-freebsd-amd64-$VERSION .
        GOOS=freebsd GOARCH=arm go build -o build/$APPLICATION-freebsd-arm-$VERSION .
        GOOS=freebsd GOARCH=arm64 go build -o build/$APPLICATION-freebsd-arm64-$VERSION .
        GOOS=js GOARCH=wasm go build -o build/$APPLICATION-js-wasm-$VERSION .
        GOOS=illumos GOARCH=amd64 go build -o build/$APPLICATION-amd64-$VERSION .
        GOOS=linux GOARCH=386 go build -o build/$APPLICATION-linux-386-$VERSION .
        GOOS=linux GOARCH=amd64 go build -o build/$APPLICATION-linux-amd64-$VERSION .
        GOOS=linux GOARCH=arm go build -o build/$APPLICATION-linux-arm-$VERSION .
        GOOS=linux GOARCH=arm64 go build -o build/$APPLICATION-linux-arm64-$VERSION .
        GOOS=linux GOARCH=mips go build -o build/$APPLICATION-linux-mips-$VERSION .
        GOOS=linux GOARCH=mips64 go build -o build/$APPLICATION-linux-mips64-$VERSION .
        GOOS=linux GOARCH=mips64le go build -o build/$APPLICATION-linux-mips64le-$VERSION .
        GOOS=linux GOARCH=mipsle go build -o build/$APPLICATION-linux-mipsle-$VERSION .
        GOOS=linux GOARCH=ppc64 go build -o build/$APPLICATION-linux-ppc64-$VERSION .
        GOOS=linux GOARCH=ppc64le go build -o build/$APPLICATION-linux-ppc64le-$VERSION .
        GOOS=linux GOARCH=riscv64 go build -o build/$APPLICATION-linux-riscv64-$VERSION .
        GOOS=linux GOARCH=s390x go build -o build/$APPLICATION-linux-s390x-$VERSION .
        GOOS=nacl GOARCH=386 go build -o build/$APPLICATION-nacl-386-$VERSION .
        GOOS=nacl GOARCH=amd64p32 go build -o build/$APPLICATION-nacl-amd64p32-$VERSION .
        GOOS=nacl GOARCH=nacl-arm go build -o build/$APPLICATION-nacl-arm-$VERSION .
        GOOS=netbsd GOARCH=386 go build -o build/$APPLICATION-netbsd-386-$VERSION .
        GOOS=netbsd GOARCH=amd64 go build -o build/$APPLICATION-netbsd-amd64-$VERSION .
        GOOS=netbsd GOARCH=arm go build -o build/$APPLICATION-netbsd-arm-$VERSION .
        GOOS=netbsd GOARCH=arm64 go build -o build/$APPLICATION-netbsd-arm64-$VERSION .
        GOOS=openbsd GOARCH=386 go build -o build/$APPLICATION-openbsd-386-$VERSION .
        GOOS=openbsd GOARCH=amd64 go build -o build/$APPLICATION-openbsd-amd64-$VERSION .
        GOOS=openbsd GOARCH=arm go build -o build/$APPLICATION-openbsd-arm-$VERSION .
        GOOS=openbsd GOARCH=arm64 go build -o build/$APPLICATION-openbsd-arm64-$VERSION .
        GOOS=plan9 GOARCH=386 go build -o build/$APPLICATION-plan9-386-$VERSION .
        GOOS=plan9 GOARCH=amd64 go build -o build/$APPLICATION-plan9-amd64-$VERSION .
        GOOS=plan9 GOARCH=arm go build -o build/$APPLICATION-plan9-arm-$VERSION .
        GOOS=solaris GOARCH=amd64 go build -o build/$APPLICATION-solaris-amd64-$VERSION .
        GOOS=windows GOARCH=386 go build -o build/$APPLICATION-windows-386-$VERSION.exe .
        GOOS=windows GOARCH=amd64 go build -o build/$APPLICATION-windows-amd64-$VERSION.exe .
        GOOS=windows GOARCH=arm go build -o build/$APPLICATION-windows-arm-$VERSION.exe .
        echo "$(find build/ -type f -print0 | xargs -0 sha1sum)" >>SHA-1
    else
        echo "Error: In your system, Go wasn't found."
        exit
    fi
}

# Start the build
build-golang-app
