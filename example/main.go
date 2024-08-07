package main

import (
	"fmt"
	"github.com/michaelwp/go-totp"
	"log"
	"time"
)

func main() {
	// Get the current Unix time
	timestamp := time.Now().Unix()

	// define the variables
	totp := go_totp.Totp{
		Secret:    "rahasia",
		Digits:    8,
		Period:    15,
		Algorithm: go_totp.SHA256,
		T0:        0,
	}

	// Generate the TOTP
	pass, err := totp.GenerateTOTP(timestamp)
	if err != nil {
		log.Println("GenerateTOTP error:", err)
	}

	// Print the TOTP
	fmt.Println("TOTP:", pass)
}