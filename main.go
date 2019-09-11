package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/heroku/x/hmetrics/onload"
)

*type Joke struct {
	Msg string `json:"-"`
}


func repeatHandler(r int) gin.HandlerFunc {
	return func(c *gin.Context) {
		var buffer bytes.Buffer
		for i := 0; i < r; i++ {
			buffer.WriteString("Hello from Go!\n")
		}
		c.String(http.StatusOK, buffer.String())
	}
}

func main() {

	port := os.Getenv("PORT")

		if port == "" {
			log.Fatal("$PORT must be set")
		}

		tStr := os.Getenv("REPEAT")
		repeat, err := strconv.Atoi(tStr)
		if err != nil {
			log.Printf("Error converting $REPEAT to an int: %q - Using default\n", err)
			repeat = 5
		}

		router := gin.New()
		router.Use(gin.Logger())
		router.LoadHTMLGlob("templates/*.tmpl.html")
		router.Static("/static", "static")

		router.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.tmpl.html", nil)
		})

		router.GET("/mark", func(c *gin.Context) {
			c.String(http.StatusOK, string(blackfriday.Run([]byte("**hi!**"))))
		})

		router.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, string(blackfriday.Run([]byte("**hi!**"))))
		})

		router.GET("/repeat", repeatHandler(repeat))

		router.GET("/jokes", showJokes)

		router.Run(":" + port)
}

func showJokes(w http.ResponseWriter, r *http.Request) {
	var joke = Joke{}
	err := json.NewDecoder(r.Body).Decode(&joke)
	if err != nil {
		panic(err)
	}
	jokeJson, err := json.Marshal(joke)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jokeJson)
}

