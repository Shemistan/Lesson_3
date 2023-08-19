package utils

import (
	"fmt"
	"os"
)

func Read(input *string) {
	_, err := fmt.Scan(input)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
