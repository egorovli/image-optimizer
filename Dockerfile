FROM golang:1.10.2-alpine as builder
# RUN apk update && apk add curl
# RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
# RUN apk update && apk add git
RUN apk add --no-cache --update ca-certificates git
RUN go get -u -v github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/egorovli/image-optimizer
COPY ./src ./Gopkg.lock ./Gopkg.toml ./
RUN dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM egorovli/mozjpeg
LABEL maintainer="Anton Egorov <anton@egorov.li>"

WORKDIR /var/app
COPY --from=builder /go/src/github.com/egorovli/image-optimizer/app ./
CMD [ "./app" ]
