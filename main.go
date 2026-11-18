package main

import (
	"fmt"
	"github.com/dugedog/gator/internal/config"
)

func main() {
	value, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("The output of Read() is %v", value)
}
