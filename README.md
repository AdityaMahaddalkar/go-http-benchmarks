# go-http-benchmarks 🌐

A benchmarking program in Golang to compare various HTTP Client APIs for GoLang

## Benchmark Results

- goos: windows
- goarch: amd64
- pkg: github.com/AdityaMahaddalkar/go-http-benchmarks/bench
- cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz

| Test                                  | #   | Time             | Memory        | Allocations     |
| ------------------------------------- | --- | ---------------- | ------------- | --------------- |
| BenchmarkBaseHttpGet-12               | 1   | 4805452200 ns/op | 3484448 B/op  | 26876 allocs/op |
| BenchmarkBaseHttpGetConcurrent-12     | 2   | 664215600 ns/op  | 851788 B/op   | 3355 allocs/op  |
| BenchmarkHeimdallHttpGet-12           | 1   | 2575437000 ns/op | 776984 B/op   | 3552 allocs/op  |
| BenchmarkHeimdallHttpGetConcurrent-12 | 1   | 1382867100 ns/op | 2042176 B/op  | 15833 allocs/op |
| BenchmarkReqGet-12                    | 1   | 4182439500 ns/op | 15870752 B/op | 27962 allocs/op |
| BenchmarkReqGetConcurrent-12          | 2   | 816695000 ns/op  | 13309200 B/op | 4764 allocs/op  |

- PASS
- coverage: 66.7% of statements
- ok github.com/AdityaMahaddalkar/go-http-benchmarks/bench 18.157s

## How to run benchmarks

1. `go get -d ./...`
2. `go test -bench=./...`

## TODOs

✔️ Write benchmark for Golang's in-built http package

✔️ Write benchmark for req

✔️ Write benchmark for heimdall

❌ Write benchmark for requests

❌ Write benchmark for go-cleanhttp
