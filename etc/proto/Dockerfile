FROM golang:1.16.1
LABEL maintainer="msteffen@pachyderm.io"

ENV GOPROXY https://proxy.golang.org

RUN apt-get update -yq
RUN apt-get install -yq unzip

RUN \
  PROTO_COMPILER_VERSION=3.11.4; \
  wget "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTO_COMPILER_VERSION}/protoc-${PROTO_COMPILER_VERSION}-linux-x86_64.zip" -O protoc.zip
RUN unzip protoc.zip -d /
RUN cp -r /include /bin

RUN go get -u -v github.com/golang/protobuf/proto

# if you modify the version of gogo/protobuf. you also need to update the path in run.sh
RUN go get -u -v github.com/gogo/protobuf/protoc-gen-gofast@v1.3.2
RUN mkdir -p ${GOPATH}/src/github.com/pachyderm/pachyderm
RUN date +%s >/last_run_time

ADD run.sh /
CMD ["/run.sh"]
