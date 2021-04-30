package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var applicationName string
var versionNumber string
var err error

func init() {
	// Make sure we have go installed in the system.
	if !commandExists("go") {
		log.Fatal("The application gos was not found in the system.")
	}
	if len(os.Args) > 1 {
		tempApplicationName := flag.String("name", "Example", "The name of your app.")
		tempVersionNumber := flag.String("version", "v1.0.0", "The version of your app.")
		flag.Parse()
		applicationName = *tempApplicationName
		versionNumber = *tempVersionNumber
	} else {
		log.Fatal("Error: The system path has not been given.")
	}
}

func main() {
	fmt.Println("Hello, playground")
}

func buildGoApps() {
	cmd := exec.Command("GOOS=aix", "go build", "-o", "bin/" + applicationName + "-" + versionNumber, "-aix-ppc64", ".")
	err = cmd.Run()
	handleErrors(err)
}

// Directory Check
func folderExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// File Check
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
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

// Handle errors
func handleErrors(err error) {
	if err != nil {
		log.Println(err)
	}
}
