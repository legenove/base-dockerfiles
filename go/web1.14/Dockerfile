FROM legenove/amazonlinux-go-1-14:v0.02

RUN mkdir -p /go/src/github.com/legenove/base-dockerfiles/go/web1.14
WORKDIR /go/src/github.com/legenove/base-dockerfiles/go/web1.14
COPY go/web1.14/go.mod ./
COPY go/web1.14/go.sum ./
COPY go/web1.14/main.go ./
RUN go mod download
RUN rm -rf ./*

CMD ["go", "run", "main.go"]