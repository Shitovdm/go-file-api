# STEP 1 build executable binary
ARG  BUILDER=${BUILDER}
FROM ${BUILDER} as builder

FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/localtime /etc/localtime
COPY --from=builder /etc/timezone /etc/timezone

# Copy our static executable
WORKDIR /go/bin
COPY --from=builder /go/bin/go-file-api /go/bin/go-file-api

CMD ["/go/bin/go-file-api", "api"]