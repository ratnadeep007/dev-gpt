clean:
	rm devgpt-*

build: clean
	GOOS=linux GOARCH=amd64 go build -o devgpt-linux
	GOOS=darwin GOARCH=amd64 go build -o devgpt-mac
	GOOS=windows GOARCH=amd64 go build -o devgpt-windows.exe