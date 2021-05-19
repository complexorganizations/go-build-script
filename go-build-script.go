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
	buildGoApps()
}

func buildGoApps() {
	// Put everything into an array and than pull the data from there and start buiding
	var distList = []string{
		"aix/ppc64",
		"android/386",
		"android/amd64",
		"android/arm",
		"android/arm64",
		"darwin/amd64",
		"darwin/arm64",
		"dragonfly/amd64",
		"freebsd/386",
		"freebsd/amd64",
		"freebsd/arm",
		"freebsd/arm64",
		"illumos/amd64",
		"ios/amd64",
		"ios/arm64",
		"js/wasm",
		"linux/386",
		"linux/amd64",
		"linux/arm",
		"linux/arm64",
		"linux/mips",
		"linux/mips64",
		"linux/mips64le",
		"linux/mipsle",
		"linux/ppc64",
		"linux/ppc64le",
		"linux/riscv64",
		"linux/s390x",
		"netbsd/386",
		"netbsd/amd64",
		"netbsd/arm",
		"netbsd/arm64",
		"openbsd/386",
		"openbsd/amd64",
		"openbsd/arm",
		"openbsd/arm64",
		"openbsd/mips64",
		"plan9/386",
		"plan9/amd64",
		"plan9/arm",
		"solaris/amd64",
		"windows/386",
		"windows/amd64",
		"windows/arm",
	}
	// Split the array and than start the build.
	// https://golang.org/pkg/strings/#Split
	/* Start the build */
	os.Setenv("GOOS", "aix")
	os.Setenv("GOARCH", "ppc64")
	cmd := exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+"-"+os.Getenv("GOOS")+"-"+os.Getenv("GOARCH"), codePath)
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
