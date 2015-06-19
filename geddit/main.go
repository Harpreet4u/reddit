package main

import (
	"fmt"
	"github.com/harpreet4u/reddit"
	"log"
)

func main() {
	items, err := reddit.Get("golang")
	if err != nil {
		log.Fatal(err)
	}
	//_, err = io.Copy(os.Stdout, resp.Body)

	for _, item := range items {
		fmt.Println(item)
	}

}
