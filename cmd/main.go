package main

import (
	"fmt"
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func main() {
	filePath := "test"
	filePath2 := "test12"

	result, err := service.Run(filePath)
	if err != nil {
		log.Fatal(err)
	}

	result2, err := service.Run(filePath2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
	fmt.Println(result2)
}
