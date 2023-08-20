FROM vibioh/scratch

ENV ZONEINFO /zoneinfo.zip
COPY zoneinfo.zip /zoneinfo.zip
COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT [ "/netatmo" ]

ARG VERSION
ENV VERSION=${VERSION}

ARG TARGETOS
ARG TARGETARCH

COPY release/netatmo_${TARGETOS}_${TARGETARCH} /netatmo
