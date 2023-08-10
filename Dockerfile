FROM docker.io/library/golang:1.21.0-alpine3.18 AS build

RUN apk add git make

WORKDIR /workspace
ADD ./ /workspace

RUN make install \
      DESTDIR=/cache \
      PREFIX=/usr \
      VERSION=${VERSION}

FROM docker.io/library/alpine:3.18.3

COPY --from=build /cache /

WORKDIR /workspace
VOLUME [ "/workspace" ]

ENTRYPOINT [ "/usr/bin/dcmerge" ]