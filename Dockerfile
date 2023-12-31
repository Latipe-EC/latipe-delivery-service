FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /out/main ./

EXPOSE 5005
ENTRYPOINT ["/out/main"]