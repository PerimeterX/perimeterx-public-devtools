FROM golang:1.10-alpine
RUN mkdir /app
WORKDIR /app
COPY collect.go ./
RUN go build
ENTRYPOINT ["/app/app"]
