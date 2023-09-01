.PHONY: all clean

user     = gavsidhu
binary   = templ8
version := 0.0.1
build := $(shell git rev-parse HEAD)
pkd := "github.com/gavsidhu/templ8/cmd"
ldflags := -ldflags="github.com/gavsidhu/templ8/cmd.Version=$(version) github.com/gavsidhu/templ8/cmd.Build=$(build)"

all:
	go build -o $(binary) $(ldflags)

test:
	go test ./... -cover -coverprofile c.out
	go tool cover -html=c.out -o coverage.html

clean:
	rm -rf $(binary) c.out coverage.html