#!/usr/bin/bash
templ generate
npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css 
go build -o test cmd/test/main.go
./test
