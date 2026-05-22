FROM golang:1.26

WORKDIR /root/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY src/go.mod src/go.sum ./
RUN go mod download

COPY ./src .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]
