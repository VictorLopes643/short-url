FROM golang:latest

WORKDIR /app

COPY . .

WORKDIR /app/go/pkg/generateID

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /app/generateID .

COPY . .

WORKDIR /app/go/cmd/urls

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /app/shortUrl .

CMD [ "/app/shortUrl" ]