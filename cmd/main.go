package main

import (
	"encoding/base64"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/skip2/go-qrcode"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/Pierre35557/qr-code-generator/docs"
)

// @title qr-code-generator API
// @version 1.0
// @description API service for generating QR codes using Echo.
// @BasePath /api/v1
func main() {
	// Create a new Echo instance using the v4 package.
	e := echo.New()

	// Define the API route for generating QR codes.
	e.GET("/api/v1/qrcode", generateQRCode)
	// Register the Swagger endpoint. Ensure this is using echo v4.
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start the server on port 8080.
	e.Logger.Fatal(e.Start(":8080"))
}

// generateQRCode godoc
// @Summary Generate a QR code image
// @Description Generate a QR code from a provided URL and return it as a PNG image.
// @Tags qrcode
// @Accept json
// @Produce png
// @Param url query string true "URL to encode"
// @Param size query int false "Size of the QR code image (default 256)"
// @Success 200 {file} file "PNG image of the QR code"
// @Failure 400 {object} map[string]string "Error message"
// @Router /qrcode [get]
func generateQRCode(c echo.Context) error {
	url := c.QueryParam("url")
	if url == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing url parameter"})
	}

	sizeParam := c.QueryParam("size")
	size := 256
	if sizeParam != "" {
		if s, err := strconv.Atoi(sizeParam); err == nil {
			size = s
		}
	}

	pngData, err := qrcode.Encode(url, qrcode.Medium, size)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error generating QR code"})
	}

	base64Str := base64.StdEncoding.EncodeToString(pngData)
	return c.JSON(http.StatusOK, map[string]string{"imageBase64": base64Str})
}
