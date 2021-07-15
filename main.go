package main

import (
	"compress/gzip"
	"io"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		os.Stdin.Close()
	}()

	gwriter := gzip.NewWriter(os.Stdout)
	io.Copy(gwriter, os.Stdin)
}
