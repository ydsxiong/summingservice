######## get the latest golang for building stage #######
FROM golang:latest as builder

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/app

COPY ./cmd/gRPC/server/main.go .

ARG main-version-tag
RUN dep ensure -v \
&& CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags ${main-version-tag} -o main .

######## Start a new base from scratch for runtime #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /srv/app/

# Copy the Pre-built binary file from the previous build stage
COPY --from=builder /go/src/app/main .
ARG port
ENV SERVICE_TCP_PORT=${port}

CMD ["./main"] 
