#++++++++++++++++++++++++++++++++#
# GoRPC Tests Docker in Alpine   #
#++++++++++++++++++++++++++++++++#

FROM golang:alpine
MAINTAINER limx <715557344@qq.com>

RUN apk --update add git

RUN go get github.com/limingxinleo/gorpc

CMD ["go", "run", "/${GOPATH}/src/github.com/limingxinleo/gorpc/main.go"]



