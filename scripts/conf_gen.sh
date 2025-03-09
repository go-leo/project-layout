protoc \
--proto_path=. \
--go_out=. \
--go_opt=paths=source_relative \
pkg/*/config.proto \
configs/config.proto
