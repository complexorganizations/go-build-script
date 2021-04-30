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
		tempBuildPath := flag.String("bin", "/user/example/folder/bin/", "The location of your code.")
		flag.Parse()
		applicationName = *tempApplicationName
		versionNumber = *tempVersionNumber
		codePath = *tempCodePath
		binPath = *tempBuildPath
	} else {
		log.Fatal("Error: The system path has not been given.")
	}
	if applicationName == "" || versionNumber == "" || codePath == "" || binPath == "" {
		log.Fatal("Error: Required arguments were not given")
	}
}

func main() {
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
