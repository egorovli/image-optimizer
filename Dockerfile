FROM golang:1.10.2-alpine as builder
LABEL maintainer="Anton Egorov <anton@egorov.li>"

RUN apk add --no-cache --update ca-certificates curl git
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/github.com/egorovli/image-optimizer
COPY ./src ./Gopkg.lock ./Gopkg.toml ./
RUN dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o image-optimizer .

FROM alpine:latest
LABEL maintainer="Anton Egorov <anton@egorov.li>"

WORKDIR /var/app

COPY --from=egorovli/mozjpeg /usr/local /usr/local
COPY --from=builder /go/src/github.com/egorovli/image-optimizer/image-optimizer /usr/local/bin

CMD [ "image-optimizer" ]
