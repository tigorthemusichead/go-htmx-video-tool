FROM golang:1.20.5-alpine3.18
WORKDIR /app
COPY . .
RUN apk add  --no-cache ffmpeg
RUN go mod download
#RUN CGO_ENABLED=0 GOOS=linux go build -o /app/ms-auth /app/services/msAuth/cmd/main.go
#RUN ls -la /app
EXPOSE 8081
#CMD ["go", "run", "/app/cmd/main.go"]
