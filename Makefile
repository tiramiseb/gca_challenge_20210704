challenge:
	CGO_ENABLED=0 go build -o challenge

all:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o compiled/challenge_linux
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o compiled/challenge_win.exe
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o compiled/challenge_mac
