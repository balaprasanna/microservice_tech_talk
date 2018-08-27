# Stage 1

FROM golang:alpine as builder

RUN mkdir /build 

RUN apk add git

RUN go get github.com/gin-gonic/gin

ADD . /build/

WORKDIR /build 

RUN go build -o main .

# Stage 2

FROM alpine

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 5000

CMD ["/app/main"]