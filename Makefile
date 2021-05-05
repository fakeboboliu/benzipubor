build:
	GOOS=linux GOARCH=386 go build -o dist/benzipubor-linux-386 -ldflags="-s -w" github.com/fakeboboliu/benzipubor/cmd
	GOOS=linux GOARCH=amd64 go build -o dist/benzipubor-linux-amd64 -ldflags="-s -w" github.com/fakeboboliu/benzipubor/cmd
	GOOS=windows GOARCH=386 go build -o dist/benzipubor-win32.exe -ldflags="-s -w" github.com/fakeboboliu/benzipubor/cmd
	GOOS=darwin GOARCH=amd64 go build -o dist/benzipubor-macos -ldflags="-s -w" github.com/fakeboboliu/benzipubor/cmd
	upx dist/benzipubor-*
	
clean:
	rm -rf dist/benzipubor-*