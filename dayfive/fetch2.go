package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		prefix := "http://"
		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		status := response.Status + "\n\n\n\n"
		os.Stdout.Write([]byte(status))
		_, e := io.Copy(os.Stdout, response.Body)
		response.Body.Close()
		if e != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	}
}
