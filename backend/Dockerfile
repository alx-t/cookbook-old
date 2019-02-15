FROM golang
MAINTAINER Alex Tesevich <alex.tesevich@gmail.com>
WORKDIR /go/src/github.com/alx-t/cookbook/backend
ADD . ./
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cookbook main.go
EXPOSE 3007 3007
ENTRYPOINT ./cookbook
