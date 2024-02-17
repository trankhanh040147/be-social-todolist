FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app

FROM alpine
WORKDIR /app/
COPY --from=builder /app/app /app/
#EXPOSE
ENTRYPOINT ["./app"]

#docker build -t social-todo-service:1.0 .
