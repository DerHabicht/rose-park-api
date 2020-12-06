FROM golang:1.14-alpine
WORKDIR /root

RUN apk add git make

RUN mkdir -p ./temp/rose-park/
COPY ./ ./temp/rose-park/
RUN cd ./temp/rose-park/ && make clean
RUN cd ./temp/rose-park/ && make
RUN cd ./temp/rose-park/ && make install
RUN rm -rf ./temp/
RUN rm -rf /go/src/

EXPOSE 8080
ENV URL 0.0.0.0:8080

CMD ["/go/bin/rose-park"]
