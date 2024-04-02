package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://www.google.com"

func main() {
	fmt.Println("Handling Urls in GO ")

	URL,err:=url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	fmt.Println("The data is : ",URL.Scheme)
	fmt.Println("The data is : ",URL.Host)
	fmt.Println("The data is : ",URL.Port())
	fmt.Println("The data is : ",URL.Path)
	fmt.Println("The data is : ",URL.RawQuery)


	myUrl:=url.URL{
		Scheme: "http",
		Host: "www.google.com",

	}
	another:=myUrl.String()

	fmt.Println("s",another)
}