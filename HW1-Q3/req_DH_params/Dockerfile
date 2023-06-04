FROM golang:alpine
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
CMD ["go","run","main.go"]
EXPOSE 4000