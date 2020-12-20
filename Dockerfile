FROM golang:1.15.6-alpine3.12

WORKDIR /code

COPY ./ .

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build .

EXPOSE 8080

CMD ["darmowe"]