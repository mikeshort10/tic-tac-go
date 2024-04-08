FROM golang

WORKDIR /usr/src/app

COPY src .

RUN go build -C board

CMD ["./board/board"]