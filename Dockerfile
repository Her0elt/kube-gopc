FROM golang:1.18-alpine

ENV PORT 8080
ENV HOST 0.0.0.0

WORKDIR /app

COPY ./ ./
RUN go mod download

RUN go build -o bin/main cmd/user/main.go

EXPOSE 8080/tcp

CMD [ "bin/main" ]
