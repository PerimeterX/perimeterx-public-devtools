FROM golang:1.10-alpine
RUN apk add --update git
RUN mkdir /app
WORKDIR /app
COPY use.go feed.mmdb ./
RUN go get -d . && go build
EXPOSE 8080
CMD ["/app/app"]
