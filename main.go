package main

import (
	"fmt"
)

func main(){
	var wordsList []string

	for {
		var s string
		_, err := fmt.Scan(&s)
		if err != nil {
			fmt.Println(err)
		}
		if s == "close"{
			return
		}
		
	}
}