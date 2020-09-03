FROM golang:latest
RUN mkdir /go/src/MyPIPE
WORKDIR /go/src/MyPIPE
ADD .. /go/src/MyPIPE
EXPOSE 8080
CMD ["go","run","/go/src/MyPIPE/main.go"]