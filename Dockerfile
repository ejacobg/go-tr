FROM golang:1.18

WORKDIR /usr/src/app

COPY . .

RUN go build -o go-tr

ENTRYPOINT ["./go-tr"]

CMD []