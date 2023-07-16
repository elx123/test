SHELL := /bin/bash

translate_windows_amd64:
	cd app/translate && gox -os "windows" -arch amd64

test:
	go run app/test/main.go