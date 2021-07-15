package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Ignore(os.Interrupt)
	signal.Notify(c, os.Interrupt)
	gwriter := gzip.NewWriter(os.Stdout)
	buffStdin := bufio.NewReader(os.Stdin)
	eof := false

	buf := make([]byte, 1024)
	for {
		select {
		case <-c:
			fmt.Fprintln(os.Stderr, "Detected an interrupt signal")
		default:
		}

		n, err := buffStdin.Read(buf)
		if err != nil {
			if err == io.EOF {
				eof = true
				fmt.Fprintln(os.Stderr, "Stdin has been closed. Flushing...")
			} else {
				fmt.Fprintln(os.Stderr, err)
				break
			}
		}

		if eof {
			gwriter.Flush()
			gwriter.Close()
			os.Stdout.Close()
			break
		}
		gwriter.Write(buf[:n])
	}

	fmt.Fprintln(os.Stderr, "Compression is done!")
}
