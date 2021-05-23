package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/tyayers/openapigen"
)

func main() {

	r := gin.Default()
	r.Use(cors.Default())

	r.Use(static.Serve("/", static.LocalFile("./public", true)))

	r.GET("/openapi/spec", func(c *gin.Context) {
		url := c.Query("url")
		fmt.Println(url)
		spec := openapigen.GenerateSpec(url)

		c.String(200, spec)
	})

	r.Run(":8080")
}
