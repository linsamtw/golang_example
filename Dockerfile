# https://hub.docker.com/_/golang/
FROM golang

#  Go compilation environment
ENV GOOS=linux
ENV CGO_ENABLED=0

WORKDIR /opt/test/

COPY main.go .
RUN go mod init test && go mod tidy
RUN go build main.go

CMD ["./main"]
