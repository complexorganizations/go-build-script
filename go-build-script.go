package main

import (
	"flag"
	//"fmt"
	"log"
	"os"
	"os/exec"
)

var (
	applicationName string
	versionNumber   string
	codePath        string
	err             error
)

func init() {
	// Make sure we have go installed in the system.
	if !commandExists("go") {
		log.Fatal("The application gos was not found in the system.")
	}
	if len(os.Args) > 1 {
		tempApplicationName := flag.String("name", "Example", "The name of your app.")
		tempVersionNumber := flag.String("version", "v1.0.0", "The version of your app.")
		tempCodePath := flag.String("path", "/user/example/folder", "The location of your code.")
		flag.Parse()
		applicationName = *tempApplicationName
		versionNumber = *tempVersionNumber
		codePath = *tempCodePath
	} else {
		log.Fatal("Error: The system path has not been given.")
	}
}

func main() {
	buildGoApps()
}

func buildGoApps() {
	cmd := exec.Command("go", "build", "-o", "bin/"+applicationName+"-"+versionNumber+"-aix-ppc64", codePath)
	err = cmd.Run()
	handleErrors(err)
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
