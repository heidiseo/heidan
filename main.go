package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	_ "github.com/heroku/x/hmetrics/onload"
)

type DadJoke struct {
	Joke string `json:"joke"`
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

	router.GET("/joke", jokeHandler())

	router.Run(":" + port)
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

func jokeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := http.NewRequest("GET", "https://icanhazdadjoke.com", nil)
		if err != nil {
			log.Fatal(err)
		}
		response.Header.Set("Accept", "application/json")
		resp, err := http.DefaultClient.Do(response)
		if err != nil {
			log.Fatal(err)
		}

		var results map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&results)
		if err != nil {
			log.Fatal(err)
		}

		jokeString := results["joke"].(string)
		joke := DadJoke{Joke: jokeString}
		c.Header("content-type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		//fmt.Printf("joke is: %v", dadJoke.Joke)
		c.JSON(http.StatusOK, joke)
	}
}
