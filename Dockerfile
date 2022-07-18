FROM --platform=$BUILDPLATFORM golang:1.18.4-alpine AS build
WORKDIR /src
ARG TARGETOS
ARG TARGETARCH
ENV GOPROXY https://goproxy.cn,direct
RUN --mount=target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /out/ipsw-go .

FROM alpine
RUN apk add --no-cache tzdata
COPY --from=build /out/ipsw-go /
CMD ["./ipsw-go"]
