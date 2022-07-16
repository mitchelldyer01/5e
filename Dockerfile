FROM golang:1.18 as build

WORKDIR /go/src/characters-5e

COPY . .

RUN make build && \
    chmod +x out/characters-5e

FROM scratch as runtime

COPY --from=build /go/src/characters-5e/out/characters-5e /

ENTRYPOINT ["/characters-5e"]