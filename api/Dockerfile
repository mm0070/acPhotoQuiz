FROM golang:1.15.7-alpine3.13

# ENV GIN_MODE=release
ENV PORT=4000

WORKDIR /go/src/github.com/mm0070/aircraftQuiz

COPY . /go/src/github.com/mm0070/aircraftQuiz

RUN apk update && apk add --no-cache git
RUN go get ./...

RUN go build -o server /go/src/github.com/mm0070/aircraftQuiz

EXPOSE $PORT

ENTRYPOINT ["./server"]
