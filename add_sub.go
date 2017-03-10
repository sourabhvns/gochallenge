package main


import (
	"io"
	"net/http"
	"os"
)

func main() {
    config_call() 
}

func config_call() {
	resp, err := http.Get("http://ifconfig.me/ip")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		config_call()
		os.Exit(1)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}