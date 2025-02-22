# QR Code Generator

QR Code Generator is a versatile tool written in Go that allows you to generate QR codes from URLs. It provides both a command-line interface (CLI) for generating QR code images as PNG files and a REST API service for generating QR codes as a base64-encoded string, complete with Swagger documentation for easy API exploration.


## Features
- **Command-Line Tool**: Generate QR codes from a URL and save them as PNG images.
- **REST API Service**: Generate QR codes from a URL and retrieve them as a base64-encoded string.
- **Swagger Documentation**: Explore and test the API endpoints via Swagger UI.
- **Customizable Output**: Adjust the QR code image size.

## Installation

You need to have Go installed. If you haven’t installed Go yet, you can download it from [Go's official website](https://go.dev/dl/).

1. Clone the repository:
   ```sh
   git clone https://github.com/Pierre35557/qr-code-generator.git
   cd qr-code-generator
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

## Building and Running

### Build the executable (based on your OS)

#### Windows:
   ```sh
  go build -o qrgen-cli.exe ./cmd/tools
   ```

#### Linux / macOS:
   ```sh
   go build -o qrgen-cli ./cmd/tools
   chmod +x qrgen  # Required for execution
   ```

### Run the application
After building, use the following command to generate a QR code:
   ```sh
   ./qrgen -url=https://www.example.com -out=my_qr.png -size=300
   ```
On Windows, use:
   ```sh
   .\qrgen.exe -url=https://www.example.com -out=my_qr.png -size=300
   ```

### REST API Service

#### Build
Build the REST API executable from the cmd directory.
- Windows:
   ```sh
  go build -o qrgen-api.exe ./cmd
   ```
- Linux / macOS:
   ```sh
   go build -o qrgen-api ./cmd
   chmod +x qrgen-api
   ```

#### Generate Swagger Documentation
If you haven't already installed the swag tool, install it with:
   ```sh
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

Then, from the project root, generate the Swagger docs:
   ```sh
   swag init
   ```

This will create a `docs` folder containing the Swagger documentation.

#### Run
- Windows:
Start the REST API service:
   ```sh
   .\qrgen-api.exe
   ```
- Linux / macOS:
   ```sh
   ./qrgen-api
   ```
The server will start on port 8080.

#### Run
Using the REST API
- **Generate a QR Code:**
Send a GET request to:
   ```sh
   http://localhost:8080/api/v1/qrcode?url=https://www.example.com&size=300
   ```

The response will be a JSON object containing the base64-encoded QR code image:
   ```json
   {
      "imageBase64": "iVBORw0KGgoAAAANSUhEUgAA..."
   }
   ```

- Swagger UI:
Explore and test the API documentation at:
   ```bash
   http://localhost:8080/swagger/index.html
   ```

## Dependencies

This project uses the [go-qrcode](https://github.com/skip2/go-qrcode) package.

## License

This project is licensed under the MIT License.

## Contributing

Feel free to submit pull requests or open issues to suggest improvements.

## Notes

- Ensure that the output directory is writable before running the command.
- The QR code size should be a reasonable value (e.g., between `100` and `1000` for best results).
