# Start from golang base image
FROM golang:alpine as builder

WORKDIR /go/src/app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/src/app/bin/hello -a -installsuffix cgo .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder  /go/src/app/bin /

EXPOSE 8080

CMD ["hello"]