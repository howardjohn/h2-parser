package main

import (
	"golang.org/x/net/http2"
	"io"
	"log"
	"os"
)

func main() {
	const preface = "PRI * HTTP/2.0\r\n\r\nSM\r\n\r\n"
	b := make([]byte, len(preface))
	if _, err := io.ReadFull(os.Stdin, b); err != nil {
		log.Fatalln(err)
	}
	if string(b) != preface {
		log.Fatalln("invalid preface")
	}
	f := http2.NewFramer(os.Stdout, os.Stdin)
	for {
		f1, err := f.ReadFrame()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(f1)
		switch tf := f1.(type) {
		case *http2.DataFrame:
			log.Println(string(tf.Data()))
		}
	}
	return
}
