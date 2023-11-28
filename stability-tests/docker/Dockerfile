ARG KASPAD_IMAGE
ARG KASPAMINER_IMAGE

FROM ${KASPAD_IMAGE} as freed
FROM ${KASPAMINER_IMAGE} as freeminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/hungyu99/freed/

WORKDIR /go/src/github.com/hungyu99/freed/

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=freed /app/ /app/
COPY --from=freeminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/hungyu99/freed/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]