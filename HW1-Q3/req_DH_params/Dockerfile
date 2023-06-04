FROM docker.arvancloud.ir/golang:latest
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
CMD ["go","run","main.go"]
EXPOSE 5000