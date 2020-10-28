# STEP 1 build executable binary
FROM golang:1.15.3-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh curl tzdata

ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
# Create appuser
RUN adduser -D -g '' appuser
RUN mkdir /go-file-api
WORKDIR /go-file-api
COPY . .

RUN go env -w GOPROXY=direct
RUN go env -w GOSUMDB=off

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/go-file-api