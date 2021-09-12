FROM golang:1.15.7-alpine
FROM golang:1.15.7 as builder

ENV ROOT=/go/src/app
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64
# ENV GOPROXY=off

WORKDIR ${ROOT}

# RUN apk update && apk add git
COPY . ${ROOT}

RUN go mod download

EXPOSE 8080

# RUN go get -u github.com/cosmtrek/air

# CMD ["air"]

# 作ったGoのバイナリを実行する
RUN go build main.go
FROM alpine:latest  
COPY --from=builder /go/src/app /go/src/app

CMD /go/src/app/main $PORT