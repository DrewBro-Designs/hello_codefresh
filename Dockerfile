FROM golang

ADD . /go/src/github.com/DrewBro-Designs/hello_codefresh
RUN go install github.com/DrewBro-Designs/hello_codefresh
ENTRYPOINT /go/bin/hello_codefresh

EXPOSE 3000

