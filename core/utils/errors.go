package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const (
	InfoMode = iota
	WarningMode
	PanicMode
	FatalMode
)

const (
	blueFormat   = "\033[1;34m[%d] %s:%d %s\n\033[0m"
	yellowFormat = "\033[1;33m[%d] %s:%d %s\n\033[0m"
	redFormat    = "\033[1;31m[%d] %s:%d %s\n\033[0m"
)

func init() {

	// Remove timestamp from logs
	log.SetFlags(0)
}

func CheckError(err error, mode int) {

	if err != nil {

		// Get error info
		utime := time.Now().UnixMilli()
		_, file, line, _ := runtime.Caller(1)
		file = filepath.Base(file)

		switch mode {
		case InfoMode:
			fmt.Printf(blueFormat, utime, file, line, err)
		case WarningMode:
			fmt.Printf(yellowFormat, utime, file, line, err)
		case PanicMode: // To recover control
			s := fmt.Sprintf(redFormat, utime, file, line, err)
			panic(s)
		case FatalMode:
			fmt.Printf(redFormat, utime, file, line, err)
			os.Exit(1)
		}
	}
}
