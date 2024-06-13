CGO_LDFLAGS="-luuid" \
GOARCH=arm64 \
CGO_ENABLED=1 \
  go install -v
