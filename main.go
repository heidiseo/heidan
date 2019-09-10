package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/heroku/x/hmetrics/onload"
)

/*type Joke struct {
	msg string
}

type Jokes []Joke

func repeatHandler(r int) gin.HandlerFunc {
	return func(c *gin.Context) {
		var buffer bytes.Buffer
		for i := 0; i < r; i++ {
			buffer.WriteString("Hello from Go!\n")
		}
		c.String(http.StatusOK, buffer.String())
	}
}*/

func main() {
	client := &http
	response, err := http.Get("https://icanhazdadjoke.com")
	response.Header.Add()
headers:
	{
	Accept:
		"application/json"
	}
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	/*port := os.Getenv("PORT")

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

		router.GET("/jokes", showJokes())
	)

		router.Run(":" + port)*/
}

/*func showJokes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Body()
	logging.Info(ctx, logging.Data{}, jokes accessed)
    jokes := Jokes{
        Joke{msg: ctx.msg},
    }
    json.NewEncoder(w).Encode(jokes)
}
*/
