env GOOS=darwin GOARCH=amd64 go build -o build/walgur-darwin-amd64
env GOOS=darwin GOARCH=386 go build -o build/walgur-darwin-i386
env GOOS=linux GOARCH=amd64 go build -o build/walgur-linux-amd64
env GOOS=linux GOARCH=386 go build -o build/walgur-linux-i386
env GOOS=windows GOARCH=386 go build -o build/walgur-windows-i386.exe
env GOOS=windows GOARCH=amd64 go build -o build/walgur-windows-amd64.exe
