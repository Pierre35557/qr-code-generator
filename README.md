# QR Code Generator

A simple command-line QR code generator written in Go. This tool allows you to generate QR codes from URLs and save them as PNG images.

## Features
- Generates QR codes from URLs
- Customizable output file name
- Adjustable QR code image size
- Uses `go-qrcode` for QR code generation

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
   go build -o qrgen.exe
   ```

#### Linux / macOS:
   ```sh
   go build -o qrgen
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

## Dependencies

This project uses the [go-qrcode](https://github.com/skip2/go-qrcode) package.

## License

This project is licensed under the MIT License.

## Contributing

Feel free to submit pull requests or open issues to suggest improvements.

## Notes

- Ensure that the output directory is writable before running the command.
- The QR code size should be a reasonable value (e.g., between `100` and `1000` for best results).
