package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func init() {
	log.Println("Starting server...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLFiles("public/index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	
	r.POST("/", createQrCodeHandler)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	})

	r.Use(cors.Default())

	log.Println("Server started on port " + os.Getenv("PORT"))
	r.Run()
}

func createQrCodeHandler(c *gin.Context) {
	form, _ := c.MultipartForm()
	link := form.Value["link"][0]

	qrc, err := qrcode.NewWith(link,
		qrcode.WithEncodingMode(qrcode.EncModeByte),
		qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionQuart),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating QR code"})
		return
	}

	buf := bytes.NewBuffer(nil)
	wr := nopCloser{Writer: buf}
	w2 := standard.NewWithWriter(wr, standard.WithQRWidth(40))
	if err := qrc.Save(w2); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating QR code"})
		return
	}

	c.Data(
		http.StatusOK,
		"image/png",
		buf.Bytes(),
	)
}

type nopCloser struct {
	io.Writer
}

func (nopCloser) Close() error { return nil }
