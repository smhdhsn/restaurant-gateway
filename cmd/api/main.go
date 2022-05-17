package main

import (
	"fmt"
	"log"

	"github.com/smhdhsn/restaurant-api-gateway/internal/config"
)

// main is the application's kernel.
func main() {
	// read configurations
	conf, err := config.LoadConf()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conf)
}
