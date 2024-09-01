#!/usr/bin/bash
templ generate
npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css 
go build -o start cmd/prod/main.go
./start
