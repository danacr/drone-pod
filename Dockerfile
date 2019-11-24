FROM golang:1.13-alpine as build-env

WORKDIR /drone
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/drone

FROM scratch
COPY --from=build-env /go/bin/drone /go/bin/drone
ENTRYPOINT ["/go/bin/drone"]