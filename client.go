package main

import (
	b64 "encoding/base64"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/tyayers/openapigen/openapigenlib"
)

func main() {

	r := gin.Default()
	r.Use(cors.Default())

	r.Use(static.Serve("/", static.LocalFile("./public", true)))

	r.GET("/openapi/spec", func(c *gin.Context) {
		url := c.Query("url")
		decodedUrl, _ := b64.StdEncoding.DecodeString(url)
		fmt.Println(string(decodedUrl))
		spec := openapigenlib.GenerateSpec(string(decodedUrl))

		c.String(200, spec)
	})

	r.Run(":8080")
}
