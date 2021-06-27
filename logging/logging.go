package logging

import (
	"fmt"
	"time"
)

var debug bool

func Debug(d bool) {
	debug = d
}

func Log(message string) {
	if !debug {
		return
	}

	fmt.Printf("%s %s", time.Now().Format("01-02-2006"), message)
}
