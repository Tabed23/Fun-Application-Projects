FROM golang as dev

WORKDIR /api
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

FROM alpine:3.15
WORKDIR /app
COPY --from=dev /api/main ./
EXPOSE 8080
CMD ["./main"]
