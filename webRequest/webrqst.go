package main

import (
	"fmt"
	"io"
	"net/http"
	
)

const url = "https://www.google.com"

func main() {
	fmt.Println("It's web request")
	response,err:=http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type :%T \n",response)

	defer response.Body.Close()

	byteData , err:=io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("The data inside the url is :",string(byteData))
}