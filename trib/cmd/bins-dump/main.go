package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"trib"
	"triblab"
)

func noError(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
}

func logError(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
	}
}

type client struct {
	bin string
	s   trib.BinStorage
}

func (self *client) printBin() {
	fmt.Printf("(working on bin %q)\n", self.bin)
}

var (
	frc       = flag.String("rc", trib.DefaultRCPath, "bin storage config file")
	logToFile = flag.Bool("l", false, "log output to file")
	destDir   = flag.String("d", "", "[optional] directory to use with -l")
)

func getFileWriter(addr string) (*os.File, *bufio.Writer) {
	if *destDir != "" {
		os.Mkdir(*destDir, os.ModeDir|os.ModePerm)
	}

	var filename = *destDir + "/backs-" + strings.SplitN(addr, ":", 2)[1] + ".txt"
	f, fError := os.Create(filename)
	if fError != nil {
		log.Fatal("Error creating file ", filename, fError)
	}
	w := bufio.NewWriter(f)
	return f, w
}

func main() {
	flag.Parse()

	var rc *trib.RC
	var e error = nil
	rc, e = trib.LoadRC(*frc)
	noError(e)

	for _, addr := range rc.Backs {
		var client = triblab.NewClient(addr)
		var w *bufio.Writer
		var f *os.File
		if *logToFile {
			f, w = getFileWriter(addr)
		}

		var keys trib.List
		client.Keys(&trib.Pattern{Prefix: "", Suffix: ""}, &keys)
		sort.Strings(keys.L[:])
		for _, key := range keys.L {
			var value string
			client.Get(key, &value)

			var record = key + "=" + value
			if *logToFile {
				w.WriteString(record + "\n")
			} else {
				fmt.Println(record)
			}
		}

		client.ListKeys(&trib.Pattern{Prefix: "", Suffix: ""}, &keys)
		sort.Strings(keys.L[:])
		for _, key := range keys.L {
			var values trib.List
			client.ListGet(key, &values)
			var record = key + "=" + fmt.Sprintf("%q", values.L)
			if *logToFile {
				w.WriteString(record + "\n")
			} else {
				fmt.Println(record)
			}
		}

		if *logToFile {
			w.Flush()
			f.Sync()
			f.Close()
		}
	}
}
