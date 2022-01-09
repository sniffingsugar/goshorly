FROM golang:alpine as builder

RUN apk add --no-cache git make build-base

ENV CGO_ENABLED=0
ENV I_PACKAGE="git.ucode.space/Phil/goshorly/utils"

WORKDIR /go/src/git.ucode.space/goshorly
COPY . .

RUN go get -d -v ./...

RUN echo $I_PACKAGE
RUN echo $CI_COMMIT_SHORT_SHA

RUN go build -a -installsuffix -ldflags="-w -s -X $I_PACKAGE.GitCommitShort=$CI_COMMIT_SHORT_SHA -X $I_PACKAGE.GitBranch=$CI_COMMIT_BRANCH" -o app .

FROM scratch as production
WORKDIR /
COPY --from=builder /go/src/git.ucode.space/goshorly/app /app
ENTRYPOINT [ "/app" ]