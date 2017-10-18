package helpers

import (
	"log"
)

var debug bool

// SetDebug turns on debugging
func SetDebug(on bool) {
	debug = on
}

// Debugmsg prints a passed in debug message if debugging is turned on
func DebugMsg(msg string) {
	if debug {
		log.Println(msg)
	}
}
