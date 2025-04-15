package project

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWritter struct{}

func FifthProjectHttp() {
	res, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	if res.StatusCode == 200 {
		fmt.Println("SUCCESS")
	}
	/*
		res: status, statuscode, body

		status > string
		statuscode > int
		body > io.ReadCloser

		io.ReadCloserer
		Reader > read([]byte)(int error)
		Closer >	Close()(error)
	*/
	lw := logWritter{}
	io.Copy(lw, res.Body)
}

func (logWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("DONE->", len(bs))
	return len(bs), nil
}
