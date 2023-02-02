build:
	go build  -o go-chat

dev:
	go run main.go

clear :
	rm -rf go-chat

add-%:
	cobra-cli add $*

.PHONY: build dev clear
