FROM golang:alpine
WORKDIR /go/src/github.com/than-os/dante-apis
COPY . /go/src/github.com/than-os/dante-apis/
RUN apk add git shadow && go get github.com/labstack/echo/... && apk add ca-certificates
RUN cd /go/src/github.com/than-os/dante-apis && go build -o dante main.go
EXPOSE 30002
ENTRYPOINT ["./dante"]
