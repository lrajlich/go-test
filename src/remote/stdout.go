package remote

import (
	"bufio"
	"net/http"
	"io"
	"fmt"
	"bytes"
)

var channel = make(chan []byte)

func SendLoop(url string) {
	var buf bytes.Buffer
	fmt.Print("IN SEND LOOP")
	writer := bufio.NewWriter(&buf)
	for {
		entry := <- channel
		writer.Write(entry)
		writer.WriteString("\n")

		writer.Flush()
		if buf.Len() > 1000 {
			_, err := http.Post(url, "text/plain", &buf)
			if err != nil {
				panic(err)
			}

			// Reset buf
			buf.Reset()
		}
	}
}

func ReadLoop(stdout io.Reader) {
	//channel = make(chan []byte)
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		channel <- scanner.Bytes()
	}
}
