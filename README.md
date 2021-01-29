# jsonvsprotobuf
Code to demonstrate the benchmarks of using json and protobuf

## Steps before running the benchmark:
 - Install [Go](https://golang.org/dl/)
 - Install [Protocol Buffers](https://github.com/google/protobuf/releases)
 - Install protoc plugin: `go get github.com/golang/protobuf/proto github.com/golang/protobuf/protoc-gen-go`

## Running the benchmarks
- Compile the proto file - `./scripts/compileproto.sh`
- Run the benchmarks - `./scripts/run_test.sh`
- Run the CPU Profiling - `./scripts/cpu-profiling.sh`
- Run the pprof - `./scripts/pprof_data.sh`

## Data to compare
```powershell
goos: linux
goarch: amd64
pkg: github.com/yashrsharma44/jsonvsprotobuf

BenchmarkOrderProto3Marshal-4             941736              1188 ns/op             256 B/op          1 allocs/op

BenchmarkOrderJSONMarshal-4               463663              2621 ns/op             512 B/op          1 allocs/op
BenchmarkOrderXMLMarshal-4                 60094             20040 ns/op            5600 B/op         16 allocs/op

BenchmarkOrderProto3Unmarshal-4           413136              2931 ns/op            1448 B/op         37 allocs/op

BenchmarkOrderJSONUnmarshal-4              88654             13890 ns/op            2000 B/op         50 allocs/op
BenchmarkOrderXMLUnmarshal-4               18980             63592 ns/op           14160 B/op        386 allocs/op
PASS
ok      github.com/yashrsharma44/jsonvsprotobuf 8.244s
```

## Some Production level benchmarked data
Refer for more details - https://github.com/brunokrebs/auth0-speed-test

