CGO_CFLAGS="-I/opt/homebrew/opt/rocksdb/include" \
CGO_LDFLAGS="-L/opt/homebrew/opt/rocksdb/lib -luuid -lrocksdb -lstdc++ -lm -lz -L/opt/homebrew/opt/snappy/lib -L/opt/homebrew/opt/lz4/lib -L/opt/homebrew/opt/zstd/lib" \
  go install -v
