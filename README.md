# interactive web cmd app
This app was build with the intention to later on develop a full interactive unix cmd,
passing commands to the terminal and getting the terminal response as a json.
This will later on be used with a javascript framework like react, angular or vue

## Motivation
Testing out the go-lang api system for go-lang, used bootstrap and jquery to do some testing as quick as possible

## Tech/framework used
<b>Built with</b>
* [go-lang v1.12](https://golang.org/)
* [gorilla/mux](https://github.com/gorilla/mux)
* [bootstrap](https://getbootstrap.com/)
* [jQuery](https://jquery.com/)

## Installation
[install go-lang](https://golang.org/doc/install)
```
go get -u github.com/gorilla/mux
```
## API Reference
|  Url               | Method |
|------------------  |--------|
| /api/command/{cmd} | GET    |
| /home              | GET    |

## How to use?
1. Run the app
```
go run main.go
```
2. Go to home page
```
http://localhost:8000/home
```
2. Type any command you like
```
go version

go version go1.12.5 windows/amd64
```

© [beltrd](https://github.com/beltrd)