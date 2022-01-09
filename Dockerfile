FROM golang:alpine as builder

RUN apk add --no-cache git make build-base
ENV CGO_ENABLED=0

WORKDIR /go/src/git.ucode.space/goshorly
COPY . .

RUN go get -d -v ./...

RUN chmod +x build-ci.sh && ./build-ci.sh

FROM scratch as production
WORKDIR /
COPY --from=builder /go/src/git.ucode.space/goshorly/app /app
ENTRYPOINT [ "/app" ]