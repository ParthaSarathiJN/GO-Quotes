package main

import "net/http"
import "github.com/labstack/echo/v4"
import "math/rand"
import "os"

type quote struct {
	Title string
	Author string
}

var quotes []quote = [] quote {
	{"Learn to 0", "0 Vidya"},
	{"Learn to 1", "1 Vidya"},
	{"Learn to 2", "2 Vidya"},
	{"Learn to 3", "3 Vidya"},
}


func main() {

	api := echo.New()
	api.GET("/home", hello)
	api.GET("/home/random", getRandom)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	api.Start(":" + port)
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, quotes)
}

func getRandom(c echo.Context) error {
	index := rand.Intn(len(quotes))
	return c.JSON(http.StatusOK, quotes[index])
}