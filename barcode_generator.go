package main

import (
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	qrCode, _ := qr.Encode("Your text, or link on picture or website, whatever you want", qr.M, qr.Auto)

	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	file, _ := os.Create("qrcode.png")
	defer file.Close()

	png.Encode(file, qrCode)
}
