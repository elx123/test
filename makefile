SHELL := /bin/bash

translate_windows_amd64:
	cd app/translateNew && gox -osarch="windows/amd64"

test:
	go run app/test/main.go