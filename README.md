# 🚀 QR Code Generator

A simple QR Code Generator using Golang with [gin-gonic](https://github.com/gin-gonic/gin) and [go-qrcode](https://github.com/yeqown/go-qrcode).

## Installation

1. Clone the repository
2. Install the dependencies with `go mod tidy`
3. Create a `.env` file with `cp .env.example .env` and choose a port of your liking
4. Run the server with `go run main.go`

## Usage

1. Open your browser and go to `http://localhost:<PORT>`
2. Enter the text you want to encode in the QR Code
3. Click on the `Generate` button
4. Enjoy your QR Code!

## Deployment

This project can be easily deployed to any cloud provider that supports Docker containers. You can use the provided `Dockerfile` to build the image and deploy it to your favorite cloud provider.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
