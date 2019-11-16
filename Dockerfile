FROM golang:1.13

COPY . /src/

RUN go get github.com/bmizerany/pq
