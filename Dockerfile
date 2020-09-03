FROM golang:latest
ADD ./go /
WORKDIR /go/src/MyPIPE
EXPOSE 8080
CMD ["go","run","/go/src/MyPIPE/main.go"]