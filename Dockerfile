FROM golang:1.18

WORKDIR /app
COPY . .

RUN go get

RUN go build -o api .

CMD ["./api"]

