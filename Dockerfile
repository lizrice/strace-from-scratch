ARG GO_VERSION

FROM golang:${GO_VERSION}-alpine AS build

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux

WORKDIR /build

#RUN apt-get update \
#    && apt-get install libseccomp-dev
RUN apk add gcc musl-dev libseccomp-dev pkgconfig

COPY go.mod go.sum *.go LICENSE ./

RUN go build -o main .

FROM alpine:3.18.2

WORKDIR /dist

RUN apk add libseccomp

COPY --from=build /build/main /build/LICENSE .

ENTRYPOINT ["/dist/main"]

