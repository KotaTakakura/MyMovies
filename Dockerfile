FROM golang:latest
WORKDIR /go/src/MyPIPE
ADD ./go /go
EXPOSE 8080
#CMD ["go","run","/go/src/MyPIPE/main.go"]