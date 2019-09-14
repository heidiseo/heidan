package main

import (
	"bytes"
	//"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	_ "github.com/heroku/x/hmetrics/onload"
)

type Joke struct {
	Message string `json:"message"`
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

//{
//	body: {
//		message: 'I am a JSON'
//	}
//}

func jokeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		joke := Joke{Message: "댄 is 바보"}
		c.Header("content-type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, joke)
	}
}

//ShowJokes is a function that has http reader and writer and returns JSON
//
/*func ShowJokes(w http.ResponseWriter, r *http.Request) {
	var joke= Joke{}
	err := json.NewDecoder(r.Body).Decode(&joke)
	if err != nil {
		panic(err)
	}
	jokeJson, err := json.Marshal(joke)
	if err != nil {
		panic(err)
	}
	w.Header.Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jokeJson)
}

/*func OwnJokes(w http.ResponseWriter, r *http.Request) {
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

func SayJokes(jokef func (w http.ResponseWriter, r *http.Request)) {

}


/*

import axios from "axios"

export async function handler(event, context) {
  try {
    const response = await axios.get("https://icanhazdadjoke.com", { headers: { Accept: "application/json" } })
    const data = response.data
    return {
      statusCode: 200,
      headers: {
        'content-type': 'application/json',
        'Access-Control-Allow-Origin': '*',
      },
      body: JSON.stringify({ msg: data.joke })
    }
  } catch (err) {
    console.log(err) // output to netlify function log
    return {
      statusCode: 500,
      body: JSON.stringify({ msg: err.message }) // Could be a custom message or object i.e. JSON.stringify(err)
    }
  }
} */

