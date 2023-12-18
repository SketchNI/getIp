package main

import (
	"fmt"
	"io"
	"net/http"
)

func httpCall(url string) (body string) {
	fmt.Println("")
	Response, Error := http.Get(url)

	if Error != nil {
		fmt.Errorf("[error] %v", Error)
		return
	}

	defer Response.Body.Close()
	Body, Error := io.ReadAll(Response.Body)

	if Error != nil {
		fmt.Errorf("[error] %v", Error)
		return
	}

	return string(Body)
}

func main() {
	fmt.Printf("Your IP address is: \n%v\n", httpCall("https://ipecho.net/plain"))
	fmt.Println()
	fmt.Println("Press [Enter] to close window.")
	fmt.Scanln()
}
