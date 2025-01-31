FROM golang:1.22-alpine as build

ENV GOBIN /usr/local/go/bin
ENV PATH /go/bin:$PATH
ENV GO111MODULE on
ENV GOPROXY direct
ENV GOSUMDB off

RUN apk update && \
   apk --no-cache upgrade && \
   apk add --no-cache tzdata gcc musl-dev git make bash && \
   cp /usr/share/zoneinfo/Brazil/East /etc/localtime && \
   echo 'Brazil/East' > /etc/timezone && \
   rm -rf /var/cache/apk/*

WORKDIR /app

COPY . .

RUN rm -rf go.sum

RUN go clean --modcache
RUN go mod tidy
RUN go mod download
RUN go mod verify
RUN go mod vendor

RUN GOOS=linux go build -a -buildvcs=false -installsuffix cgo -o main .

FROM alpine:3.17.0

RUN apk update && \
   apk --no-cache upgrade && \
   apk add --no-cache tzdata gcc musl-dev git make bash && \
   cp /usr/share/zoneinfo/Brazil/East /etc/localtime && \
   echo 'Brazil/East' > /etc/timezone && \
   rm -rf /var/cache/apk/*

WORKDIR /app
COPY --from=build /app/main .


ENTRYPOINT ["./main"]
