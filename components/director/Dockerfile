FROM golang:1.16.3-alpine3.13 as builder

ENV BASE_APP_DIR /go/src/github.com/kyma-incubator/compass/components/director
WORKDIR ${BASE_APP_DIR}

#
# Download dependencies
#

COPY go.mod go.sum ${BASE_APP_DIR}/
RUN go mod download -x

#
# Copy files
#

COPY . .

#
# Build app
#

RUN go build -v -o director ./cmd/director/main.go \
  && go build -v -o tenantfetcher ./cmd/tenantfetcher/main.go \
  && go build -v -o tenantloader ./cmd/tenantloader/main.go \
  && go build -v -o ordaggregator ./cmd/ordaggregator/main.go \
  && go build -v -o scopessynchronizer ./cmd/scopessynchronizer/main.go
RUN mkdir /app && mv ./director /app/director \
  && mv ./tenantfetcher /app/tenantfetcher \
  && mv ./tenantloader /app/tenantloader \
  && mv ./ordaggregator /app/ordaggregator \
  && mv ./scopessynchronizer /app/scopessynchronizer \
  && mv ./licenses /app/licenses

FROM alpine:3.13.5
LABEL source = git@github.com:kyma-incubator/compass.git
WORKDIR /app

#
# Copy binary
#

RUN apk --no-cache add curl ca-certificates

COPY --from=builder /app /app
COPY ./examples/ /app/examples/

#
# Run app
#

CMD ["/app/director"]
