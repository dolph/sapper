FROM golang:1.9

# Install requirements.
RUN go get gopkg.in/urfave/cli.v1

CMD \
    go test \
    && go build -v \
    && ./sapper
