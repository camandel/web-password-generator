# Web Password Generator
Simple web application to generate long and complex passwords 
![screenshot](../assets/screenshot.jpg/?raw=true)
## Installation
```
git clone https://github.com/camandel/web-password-generator.git
cd web-password-generator
go get github.com/GeertJohan/go.rice/rice
rice embed-go 
go build .
```
## Run
```
$ web-password-generator --help
Usage of web-password-generator:
  -digits int
    	Specify the the number of digits in the password (default 16)
  -length int
    	Specify the password length (default 64)
  -listen int
    	Specify on which port to listen (default 8080)
  -symbols int
    	Specify the the number of symbols in the password (default 16)
```
## Docker
TODO
