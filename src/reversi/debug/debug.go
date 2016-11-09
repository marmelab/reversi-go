package debug

import (
	"flag"
	"io"
	"os"
)

var (
	debugfile = flag.String("debugfile", "/dev/null", "Path to the debug file")
)

func Log(message string) {

	flag.Parse()

	f, _ := os.OpenFile(*debugfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	io.WriteString(f, message+"\n")
	f.Close()

}
