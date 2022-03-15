FROM golang:1.16-alpine

WORKDIR /blogorm

COPY go.mod ./
COPY go.sum ./

RUN go build -o ./blogorm

EXPOSE 80

CMD [ "/blogorm" ]