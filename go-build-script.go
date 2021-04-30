package main

import (
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	applicationName string
	versionNumber   string
	codePath        string
	binPath         string
	err             error
)

func init() {
	// Make sure we have go installed in the system.
	if !commandExists("go") {
		log.Fatal("The application go was not found in the system.")
	}
	if len(os.Args) > 1 {
		tempApplicationName := flag.String("name", "Example", "The name of your app.")
		tempVersionNumber := flag.String("version", "v1.0.0", "The version of your app.")
		tempCodePath := flag.String("code", "/user/example/folder", "The location of your code.")
		tempBinPath := flag.String("bin", "/user/example/folder/bin/", "The location of your bin folder.")
		flag.Parse()
		applicationName = *tempApplicationName
		versionNumber = *tempVersionNumber
		codePath = *tempCodePath
		binPath = *tempBinPath
	} else {
		log.Fatal("Error: The system path has not been given.")
	}
	// Make sure everything is provided
	if applicationName == "" || versionNumber == "" || codePath == "" || binPath == "" {
		log.Fatal("Error: Required arguments were not given")
	}
}

func main() {
	// Make a list of all the env var and than use a single statment and remove all these copy of the statment
	buildGoApps()
}

func buildGoApps() {
	/* Aix */
	// aix/ppc64
	os.Setenv("GOOS", "aix")
	os.Setenv("GOARCH", "ppc64")
	cmd := exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	/* Android */
	// android/386
	os.Setenv("GOOS", "android")
	os.Setenv("GOARCH", "386")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// android/amd64
	os.Setenv("GOOS", "android")
	os.Setenv("GOARCH", "amd64")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// android/arm
	os.Setenv("GOOS", "android")
	os.Setenv("GOARCH", "arm")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// android/arm64
	os.Setenv("GOOS", "android")
	os.Setenv("GOARCH", "arm64")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	/* Darwin */
	// darwin/amd64
	os.Setenv("GOOS", "darwin")
	os.Setenv("GOARCH", "amd64")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// darwin/arm64
	os.Setenv("GOOS", "darwin")
	os.Setenv("GOARCH", "arm64")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	/* Dragonfly */
	// dragonfly/amd64
	os.Setenv("GOOS", "dragonfly")
	os.Setenv("GOARCH", "amd64")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	/* Freebsd */
	// freebsd/386
	os.Setenv("GOOS", "freebsd")
	os.Setenv("GOARCH", "386")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// freebsd/amd64
	os.Setenv("GOOS", "freebsd")
	os.Setenv("GOARCH", "amd64")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// freebsd/arm
	os.Setenv("GOOS", "freebsd")
	os.Setenv("GOARCH", "arm")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// freebsd/arm64
	os.Setenv("GOOS", "freebsd")
	os.Setenv("GOARCH", "arm64")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	/* Lllumos */
	// illumos/amd64
	os.Setenv("GOOS", "illumos")
	os.Setenv("GOARCH", "amd64")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	/* JS */
	// js/amd64
	os.Setenv("GOOS", "js")
	os.Setenv("GOARCH", "wasm")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	/* iOS */
	// ios/amd64
	os.Setenv("GOOS", "ios")
	os.Setenv("GOARCH", "amd64")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// ios/arm64
	os.Setenv("GOOS", "ios")
	os.Setenv("GOARCH", "arm64")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	/* Linux */
	// linux/386
	os.Setenv("GOOS", "linux")
	os.Setenv("GOARCH", "386")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// linux/amd64
	os.Setenv("GOOS", "linux")
	os.Setenv("GOARCH", "amd64")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// linux/arm
	os.Setenv("GOOS", "linux")
	os.Setenv("GOARCH", "arm")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// linux/arm64
	os.Setenv("GOOS", "linux")
	os.Setenv("GOARCH", "arm64")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// linux/mips
	os.Setenv("GOOS", "linux")
	os.Setenv("GOARCH", "mips")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// linux/mips64
	os.Setenv("GOOS", "linux")
	os.Setenv("GOARCH", "mips64")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// linux/mips64le
	os.Setenv("GOOS", "linux")
	os.Setenv("GOARCH", "mips64le")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// linux/mipsle
	os.Setenv("GOOS", "linux")
	os.Setenv("GOARCH", "mipsle")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// linux/ppc64
	os.Setenv("GOOS", "linux")
	os.Setenv("GOARCH", "ppc64")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// linux/ppc64le
	os.Setenv("GOOS", "linux")
	os.Setenv("GOARCH", "ppc64le")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// linux/riscv64
	os.Setenv("GOOS", "linux")
	os.Setenv("GOARCH", "riscv64")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// linux/s390x
	os.Setenv("GOOS", "linux")
	os.Setenv("GOARCH", "s390x")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	/* Nacl */
	// nacl/386
	os.Setenv("GOOS", "nacl")
	os.Setenv("GOARCH", "386")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// nacl/amd64p32
	os.Setenv("GOOS", "nacl")
	os.Setenv("GOARCH", "amd64p32")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// nacl/nacl-arm
	os.Setenv("GOOS", "nacl")
	os.Setenv("GOARCH", "nacl-arm")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	/* Netbsd */
	// netbsd/386
	os.Setenv("GOOS", "netbsd")
	os.Setenv("GOARCH", "386")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// netbsd/amd64
	os.Setenv("GOOS", "netbsd")
	os.Setenv("GOARCH", "amd64")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// netbsd/arm
	os.Setenv("GOOS", "netbsd")
	os.Setenv("GOARCH", "arm")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// netbsd/arm64
	os.Setenv("GOOS", "netbsd")
	os.Setenv("GOARCH", "arm64")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	/* Openbsd */
	// openbsd/386
	os.Setenv("GOOS", "openbsd")
	os.Setenv("GOARCH", "386")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// openbsd/amd64
	os.Setenv("GOOS", "openbsd")
	os.Setenv("GOARCH", "amd64")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// openbsd/arm
	os.Setenv("GOOS", "openbsd")
	os.Setenv("GOARCH", "arm")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// openbsd/arm64
	os.Setenv("GOOS", "openbsd")
	os.Setenv("GOARCH", "arm64")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// openbsd/mips64
	os.Setenv("GOOS", "openbsd")
	os.Setenv("GOARCH", "mips64")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	/* Plan9 */
	// plan9/386
	os.Setenv("GOOS", "plan9")
	os.Setenv("GOARCH", "386")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// plan9/386
	os.Setenv("GOOS", "plan9")
	os.Setenv("GOARCH", "amd64")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// plan9/386
	os.Setenv("GOOS", "plan9")
	os.Setenv("GOARCH", "arm")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	/* Solaris */
	// solaris/amd64
	os.Setenv("GOOS", "solaris")
	os.Setenv("GOARCH", "amd64")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	/* Windows */
	// windows/386
	os.Setenv("GOOS", "windows")
	os.Setenv("GOARCH", "386")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// windows/amd64
	os.Setenv("GOOS", "windows")
	os.Setenv("GOARCH", "amd64")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	// windows/arm
	os.Setenv("GOOS", "windows")
	os.Setenv("GOARCH", "arm")
	cmd = exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
	err = cmd.Run()
	handleErrors(err)
	/* Get SHA512 and put everything in a register. */
	getSHA512Hash()
}

func getSHA512Hash() {
	if folderExists(binPath) {
		filepath.Walk(binPath, func(path string, info os.FileInfo, err error) error {
			if fileExists(path) {
				file, err := os.Open(path)
				handleErrors(err)
				defer file.Close()
				hash := sha512.New()
				io.Copy(hash, file)
				handleErrors(err)
				directory, fileName := filepath.Split(path)
				_ = directory
				fmt.Printf("%s %x \n", fileName, hash.Sum(nil))
			}
			return nil
		})
	}
}

// Application Check
func commandExists(cmd string) bool {
	appName, err := exec.LookPath(cmd)
	if err != nil {
		return false
	}
	_ = appName
	return true
}

// Folder exists
func folderExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// File exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Handle errors
func handleErrors(err error) {
	if err != nil {
		log.Println(err)
	}
}
