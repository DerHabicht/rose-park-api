FROM golang:1.15-alpine
WORKDIR /go/src/github.com/derhabicht/rose-park-api

RUN apk add git make

COPY ./ ./
RUN make clean
RUN make
RUN make install
RUN rm -rf /go/src/

EXPOSE 8080
ENV URL 0.0.0.0:8080
ENV GIN_MODE release

CMD ["/go/bin/rose-park"]
