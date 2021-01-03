FROM golang:1.15-alpine
WORKDIR /go/src/github.com/derhabicht/rose-park-api

RUN apk add git make

COPY ./ ./
RUN make
RUN make install
RUN rm -rf /go/src/

RUN apk del git make

EXPOSE 8080
ENV URL 0.0.0.0:8080

ENTRYPOINT ["/go/bin/rose-park"]
