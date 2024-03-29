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
	"runtime"
	"strings"
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
	commandExists("go")
	// Flags
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
	if len(applicationName) < 1 || len(versionNumber) < 1 || len(codePath) < 1 || len(binPath) < 1 {
		log.Fatal("Error: Required arguments were not given")
	}
	// Make sure we're using the right version of the runtime.
	if runtime.Version() == "go1.16.5" {
		log.Fatal("Error: Please upgrade your Go version to 1.16.5 on your local machine.")
	}
}

func main() {
	buildGoApps()
}

func buildGoApps() {
	// Put everything into an array and than pull the data from there and start buiding
	distList := []string{
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
	/* Start the build */
	for i := 0; i < len(distList); i++ {
		splitDistList := strings.Split(distList[i], "/")
		// Delete all environment variables
		os.Clearenv()
		// Determine how to construct based on the operating system.
		switch runtime.GOOS {
		case "darwin", "linux":
			err = os.Setenv("GOOS", splitDistList[0])
			handleErrors(err)
			err = os.Setenv("GOARCH", splitDistList[1])
			handleErrors(err)
		case "windows":
			cmd := exec.Command("$env"+":"+"GOOS", "=", splitDistList[0])
			err = cmd.Run()
			handleErrors(err)
			cmd = exec.Command("$env"+":"+"GOARCH", "=", splitDistList[1])
			err = cmd.Run()
			handleErrors(err)
		}
		cmd := exec.Command("go", "build", "-o", binPath+applicationName+"-"+versionNumber+"-"+splitDistList[0]+"-"+splitDistList[1], codePath)
		err = cmd.Run()
		handleErrors(err)
	}
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
func commandExists(cmd string) {
	cmd, err := exec.LookPath(cmd)
	if err != nil {
		log.Fatalf("Error: The application %s was not found in the system.\n", cmd)
	}
}

// Folder exists
func folderExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// File exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
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
