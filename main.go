package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/skip2/go-qrcode"
)

func main() {
	url := flag.String("url", "", "URL to encode in the QR code")
	out := flag.String("out", "qr.png", "Output file for the QR code image")
	size := flag.Int("size", 256, "Size of the QR code image")
	flag.Parse()

	if *url == "" {
		fmt.Println("Usage: go run main.go -url=https://www.example.com [-out=qr.png] [-size=256]")
		return
	}

	err := qrcode.WriteFile(*url, qrcode.Medium, *size, *out)
	if err != nil {
		log.Fatalf("Failed to generate QR code: %v", err)
	}

	fmt.Printf("QR code generated and saved to %s\n", *out)
}
