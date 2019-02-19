FROM amazonlinux:1

ENV GOLANG_VERSION 1.11.5

RUN yum update -y -q --security && \
        yum install -y -q wget \
        yum install -y -q git \
        yum install -y -q gzip \
        yum install -y -q tar.x86_64 \
        yum install -y -q make \
        yum clean all

RUN wget -nv https://redirector.gvt1.com/edgedl/go/go${GOLANG_VERSION}.linux-amd64.tar.gz -O /tmp/go${GOLANG_VERSION}.linux-amd64.tar.gz && \
    tar -xzf /tmp/go${GOLANG_VERSION}.linux-amd64.tar.gz -C /usr/local/ && \
    rm /tmp/go${GOLANG_VERSION}.linux-amd64.tar.gz

ENV GOROOT /usr/local/go
ENV GOPATH /gopath
ENV PATH   ${GOPATH}/bin:${GOROOT}/bin:/usr/local/bin:${PATH}

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR $GOPATH