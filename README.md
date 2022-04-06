go mod init myapp

go get -d github.com/jensg-st/dbuilder/cmd/dbuilder

replace github.com/jensg-st/dbuilder => /home/jensg/go/src/github.com/jensg-st/dbuilder

go run github.com/jensg-st/dbuilder/cmd/dbuilder init

go generate ./app/

