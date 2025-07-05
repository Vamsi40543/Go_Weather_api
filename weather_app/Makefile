# Makefile for Weather_app (Gin + OpenWeather)

APP_NAME := weather-api
MAIN := main.go

.PHONY: all run build tidy clean debug

all: run

run:
	go run $(MAIN)

build:
	go build -o $(APP_NAME) $(MAIN)

tidy:
	go mod tidy

clean:
	rm -f $(APP_NAME)
	rm -f debug.log

debug:
	GIN_MODE=debug go run $(MAIN)
