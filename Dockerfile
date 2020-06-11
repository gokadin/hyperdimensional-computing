FROM iron/go:dev

ENV SRC_DIR=/go/src/github.com/gokadin/hyperdimensional-computing

WORKDIR $SRC_DIR

ADD . .

RUN go build -o app

ENTRYPOINT ["./app"]