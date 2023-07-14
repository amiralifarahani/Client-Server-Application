FROM golang:1.19
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY gen ./gen
RUN go get github.com/fatih/color
RUN go get go.uber.org/ratelimit
# COPY service/*.go ./service
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
EXPOSE 6433
CMD ["/docker-gs-ping"]
# main.go auth_grpc.pb.go auth.pb.go server.go
# CMD ["/app/main"]

