swag:
	export GOPATH=$HOME/go
	export PATH=$PATH:$GOPATH/bin
	swag init

run:
	go mod download
	go build -o go-starter
	./go-starter

rm-go-starter:
	rm -rf go-starter