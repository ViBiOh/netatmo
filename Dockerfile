FROM scratch

ENV ZONEINFO zoneinfo.zip
EXPOSE 1080

ENV NETATMO_CSP "default-src 'self'; base-uri 'self'; script-src 'self' 'unsafe-inline' unpkg.com/swagger-ui-dist@3/; style-src 'self' 'unsafe-inline' unpkg.com/swagger-ui-dist@3/; img-src 'self' data:"
ENV NETATMO_PORT 1080

HEALTHCHECK --retries=5 CMD [ "/netatmo", "-url", "http://localhost:1080/health" ]
ENTRYPOINT [ "/netatmo" ]

ARG VERSION
ENV VERSION=${VERSION}

ARG TARGETOS
ARG TARGETARCH

COPY cacert.pem /etc/ssl/certs/ca-certificates.crt
COPY zoneinfo.zip /
COPY release/netatmo_${TARGETOS}_${TARGETARCH} /netatmo
