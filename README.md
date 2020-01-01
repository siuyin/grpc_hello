# grpc examples in Dart and Go

The dart folder has code for dart client and server.

```sh
export PATH=~/flutter/bin/cache/dart-sdk/bin:~/.pub-cache/bin:$PATH

cd dart
protoc -I.. hello.proto --dart_out=grpc:lib

```

To generate go code.

```sh
cd go/cmd/server
go generate

# does the following
//go:generate protoc -I ../../.. --go_out=plugins=grpc:../../hello hello.proto
```
