FROM legenove/amazonlinux-1-base:v0.0.1

ENV GOLANG_VERSION 1.11.5

RUN wget -nv https://redirector.gvt1.com/edgedl/go/go${GOLANG_VERSION}.linux-amd64.tar.gz -O /tmp/go${GOLANG_VERSION}.linux-amd64.tar.gz && \
    tar -xzf /tmp/go${GOLANG_VERSION}.linux-amd64.tar.gz -C /usr/local/ && \
    rm /tmp/go${GOLANG_VERSION}.linux-amd64.tar.gz

ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH   ${GOPATH}/bin:${GOROOT}/bin:/usr/local/bin:${PATH}

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

RUN go get github.com/kardianos/govendor

RUN curl https://glide.sh/get | sh

WORKDIR $GOPATH