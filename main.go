package main

import (
	"net/http"
	"strconv"
	"flag"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sethvargo/go-password/password"
	"github.com/foolin/echo-template/supports/gorice"
	"github.com/GeertJohan/go.rice"
)

type passwordData struct {
	Length     int
	Digits     int
	Symbols    int
	NoUpper    bool
	DenyRepeat bool
	Password   string
}

var d, defaults passwordData

func GeneratePassword(c echo.Context) error {

	var err error

	d.Length, err = strconv.Atoi(c.FormValue("length"))
	if err != nil {
		d.Length = defaults.Length
	}

	d.Digits, err = strconv.Atoi(c.FormValue("digits"))
	if err != nil {
		d.Digits = defaults.Digits
	}

	d.Symbols, err = strconv.Atoi(c.FormValue("symbols"))
	if err != nil {
		d.Symbols = defaults.Symbols
	}

	if c.FormValue("noupper") == "on" {
		d.NoUpper = true
	}

	if c.FormValue("denyrepeat") == "on" {
		d.DenyRepeat = true
	}

	// Generate a password
	d.Password, err = password.Generate(d.Length, d.Digits, d.Symbols, d.NoUpper, !d.DenyRepeat)
	if err != nil {
		d.Password = "Error: " + err.Error()
	}
	if d.Length == 0 {
		d.Password = "Error: password can not have zero length"
	}


	return c.Render(http.StatusOK, "index", d)
}

func CheckHealth(c echo.Context) error {
	s := `{"status":"OK"}`
	return c.String(http.StatusOK, s)
}

func main() {

	// Set defaults
        portPtr := flag.Int("listen", 8080, "Specify on which port to listen")
        lengthPtr := flag.Int("length", 64, "Specify the password length")
        digitsPtr := flag.Int("digits", 16, "Specify the the number of digits in the password")
        symbolsPtr := flag.Int("symbols", 16, "Specify the the number of symbols in the password")
	flag.Parse()

	defaults.Length = *lengthPtr
	defaults.Digits = *digitsPtr
	defaults.Symbols = *symbolsPtr

	e := echo.New()
	e.HideBanner = true

	// Middleware
        e.Use(middleware.Logger())
        e.Use(middleware.Recover())

	// servers other static files
	staticFileServer := http.StripPrefix("/static/", http.FileServer(rice.MustFindBox("static").HTTPBox()))
	e.GET("/static/*", echo.WrapHandler(staticFileServer))

	//Set Renderer
	e.Renderer = gorice.New(rice.MustFindBox("views"))

	e.GET("/", GeneratePassword)
	e.POST("/", GeneratePassword)

	e.GET("/health", CheckHealth)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(*portPtr)))
}
