package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	len := 0
	for range os.Args {
		len++
	}
	if len > 2 {
		fmt.Println("Too many arguments")
	} else if len < 2 {
		fmt.Println("File name missing")
	} else if len == 2 {
		data, err := ioutil.ReadFile(os.Args[1])
		fmt.Print(string(data))
		if err != nil {
			fmt.Println(err)
		}
	}
}
