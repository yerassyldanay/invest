FROM ubuntu:16.04

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update -y
RUN apt-get install curl -y

#CMD tail -f /dev/null
ENV GOLANG_VERSION 1.13
RUN curl -sSL https://storage.googleapis.com/golang/go1.13.linux-amd64.tar.gz | tar -v -C /usr/local -xz

#RUN echo $PATH
RUN mkdir -p /go/src /go/bin && chmod -R 777 /go
ENV PATH /usr/local/go/bin:$PATH

#CMD tail -f /dev/null
ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

WORKDIR /go

RUN apt-get update -y
RUN apt-get install git -y
RUN apt-get update -y

# to solve the following problem:
# exec: "gcc": executable file not found in %PATH%
RUN apt-get install build-essential -y

WORKDIR /go/src
COPY . /go/src/

#RUN source environment/.environment
COPY migrate /usr/bin

RUN source ./en
RUN migrate -path ./db/postgre/migrate -database postgres://spkuser:c8acb720063d4eb75b56drg@178.170.221.116:7001/invest?sslmode=disable -verbose up

RUN go build -o main
#CMD tail -f /dev/null

