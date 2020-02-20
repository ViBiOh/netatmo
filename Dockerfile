FROM scratch

ENV ZONEINFO zoneinfo.zip
EXPOSE 1080

HEALTHCHECK --retries=5 CMD [ "/netatmo", "-url", "http://localhost:1080/health" ]
ENTRYPOINT [ "/netatmo" ]

ARG VERSION
ENV VERSION=${VERSION}

ARG TARGETOS
ARG TARGETARCH

COPY cacert.pem /etc/ssl/certs/ca-certificates.crt
COPY zoneinfo.zip /
COPY release/netatmo_${TARGETOS}_${TARGETARCH} /netatmo
