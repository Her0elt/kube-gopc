FROM golang:1.16-alpine

ENV PORT 8080
ENV HOST 0.0.0.0

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./

RUN go build

EXPOSE 8080/tcp

CMD [ "./main" ]
