env GOOS=darwin GOARCH=amd64 go build -o walgur-darwin=amd64
env GOOS=darwin GOARCH=386 go build -o walgur-darwin-i386
env GOOS=linux GOARCH=amd64 go build -o walgur-linux-amd64
env GOOS=linux GOARCH=386 go build -o walgur-linux-i386
env GOOS=windows GOARCH=386 go build -o walgur-windows-i386.exe
env GOOS=windows GOARCH=amd64 go build -o walgur-windows-amd64.exe
