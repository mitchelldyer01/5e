FROM golang:1.18 as build

WORKDIR /go/src/5e

COPY . .

RUN make build && \
    chmod +x out/5e

FROM scratch as runtime

COPY --from=build /go/src/5e/out/5e /

ENTRYPOINT ["/5e"]