FROM golang:1.13

COPY . /src/

RUN go get github.com/bmizerany/pq
RUN go get github.com/lib/pq
