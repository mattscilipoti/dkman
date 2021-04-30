FROM --platform=${BUILDPLATFORM} golang:1.16-alpine AS base

FROM base AS dev
# docker-compose creates shared volume at /go/src
WORKDIR /go/src

RUN apk add --no-cache --update --virtual shell-dependencies \
  bash \
  git \
  vim

CMD /bin/bash

FROM base AS prep-build
WORKDIR /go/src
ENV CGO_ENABLED=0
COPY go.* .
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

FROM prep-build AS build
ARG TARGETOS
ARG TARGETARCH
RUN --mount=target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/dkman .

# prep for cross-compiling, with sections named "bin-OSNAME"
FROM scratch AS bin-unix
COPY --from=build /out/dkman /

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM scratch AS bin-windows
COPY --from=build /out/dkman /dkman.exe

# This builds using the appopriately named section from above
FROM bin-${TARGETOS} as bin