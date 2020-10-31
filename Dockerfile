FROM vibioh/scratch

ENV NETATMO_PORT 1080
EXPOSE 1080

ENV ZONEINFO /zoneinfo.zip
COPY zoneinfo.zip /zoneinfo.zip
COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

HEALTHCHECK --retries=5 CMD [ "/netatmo", "-url", "http://localhost:1080/health" ]
ENTRYPOINT [ "/netatmo" ]

ARG VERSION
ENV VERSION=${VERSION}

ARG TARGETOS
ARG TARGETARCH

COPY release/netatmo_${TARGETOS}_${TARGETARCH} /netatmo
