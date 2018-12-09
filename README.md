# Web Password Generator
Simple web application to generate long and complex passwords 
![screenshot](../assets/screenshot.jpg/?raw=true)
## Build from source
```
git clone https://github.com/camandel/web-password-generator.git
cd web-password-generator
go get github.com/GeertJohan/go.rice/rice
rice embed-go 
go build .
```
## Run with Docker image
```
docker run --rm -p 8080:8080 camandel/web-password-generator
```
## Run from CLI
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
