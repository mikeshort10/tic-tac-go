FROM golang

WORKDIR /usr/src/app

COPY src src

RUN go build -C src/interface -o ${PWD}/dist/interface

CMD ["./dist/interface"]