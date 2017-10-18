package helpers

import (
	"log"
	"os"
	"path/filepath"
)

var (
	CurrentWorkingDir string
)

func init() {
	// working dir
	getCWD()
}

func getCWD() {
	binaryf, err := filepath.Abs(os.Args[0])
	if err != nil {
		log.Fatalf("Unable to find current working directory - %v", err)
	}
	path := filepath.Dir(binaryf)

	// cope with symlinks
	// might not work on Windows if has .exe?
	actbinaryf, err := filepath.EvalSymlinks(binaryf)
	if err == nil && actbinaryf != binaryf {
		path = filepath.Dir(actbinaryf)
	}

	CurrentWorkingDir = path
}
