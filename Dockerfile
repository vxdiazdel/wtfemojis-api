FROM golang:1.22.5

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/wtfemojis-api cmd/server/main.go

RUN groupadd appgroup && useradd -ms /bin/bash -g appgroup appuser
RUN chown -R appuser:appgroup /app
RUN chmod +x /app/bin/wtfemojis-api
USER appuser
EXPOSE 8080

ENTRYPOINT ["/app/bin/wtfemojis-api"]