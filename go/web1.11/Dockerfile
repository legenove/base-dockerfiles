FROM legenove/amazonlinux-go-1-11:v0.05

WORKDIR /go/src/app

COPY ./glide.lock ./
COPY ./glide.yaml ./
RUN glide install
RUN rm -rf ./*

CMD ["go", "run", "main.go"]