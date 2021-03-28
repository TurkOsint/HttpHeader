package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		usage(os.Args[2])
	}

	url := os.Args[2]


	response, err := http.Head(url)
	

	if err != nil {
		fmt.Println("Project Name : HttpHeader")
		fmt.Println("Created : https://github.com/EyupErgin")
		fmt.Println("Github  : https://github.com/IntelSights/HttpHeader/")
		fmt.Println("  ")
		log.Fatal("Error Occurred While Getting URL Please Enter Correct URL.")
	}

	fmt.Println("Project Name : HttpHeader")
	fmt.Println("Created : https://github.com/EyupErgin")
	fmt.Println("Github  : https://github.com/IntelSights/HttpHeader/")
	fmt.Println("  ")	
	for key, value := range response.Header {
		fmt.Printf("%s: %s\n", key, value[0])
	}
}


func usage(name string) {
	fmt.Println("Project Name : HttpHeader")
	fmt.Println("Created : https://github.com/EyupErgin")
	fmt.Println("Github  : https://github.com/IntelSights/HttpHeader/")
	fmt.Println("  ")
	fmt.Printf("Please Just enter URL.\n")
	fmt.Printf("Using : ./HttpHeader https:example.com\n")
	os.Exit(1)
}

