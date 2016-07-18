package main

import (
	"fmt"
	"os"
)

import (
	"github.com/mrcrilly/go-cm-test/user"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// _options := []string{
	// 	"-u 9001",
	// 	"-k key1=something,key2=yeah",
	// }

	_exit_condition := user.CreateUser("michaelc", []string{"-u 9001", "-k key1=yeah"})

	if _exit_condition.RawError != nil {
		fmt.Printf("Error: (%d) %s\n", _exit_condition.Code, _exit_condition.Message)
		os.Exit(1)
	}

	fmt.Println("Success.")
}
