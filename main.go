package main

import "fmt"

func main() {
	hello := "Hello, World!"
	if err := notableInterview(hello); err != nil {
		fmt.Println(err.Error())
	}
}

func notableInterview(hello string) error {
	fmt.Println(hello)
	return nil
}
