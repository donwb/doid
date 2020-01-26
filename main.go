package main

import (
	"fmt"
	"log"
	//"encoding/json"
	"net/http"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", RootRoute)
	e.POST("/register", Register)

	e.Logger.Fatal(e.Start(":8000"))

}


func RootRoute(c echo.Context) error {
	fmt.Println("in root route...")

	return c.String(http.StatusOK, "Moving along... nothing to see here")

}

func Register(c echo.Context) (err error) {
	fmt.Println("in Register....")

	ck := c.FormValue("cookieID")
	ua := c.FormValue("browser")
	cookie := Cookie{
		cookieID: ck,
		browser: ua,
	}

// Can't seem to get binding working atm
/*
	if err = c.Bind(cookie); err != nil {
		fmt.Println("Error occured binding object...")
    	return c.String(http.StatusOK, "Bad shit happened")
  	}
*/
	
	log.Println(cookie)

	Push(cookie)

	return c.JSON(http.StatusOK, cookie)

}