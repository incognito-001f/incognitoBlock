FROM golang:1.22-alpine
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o incognito-node main.go
EXPOSE 8080
CMD ["./incognito-node"]