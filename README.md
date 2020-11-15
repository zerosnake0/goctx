[![Go Report Card](https://goreportcard.com/badge/github.com/zerosnake0/goctx)](https://goreportcard.com/report/github.com/zerosnake0/goctx)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/zerosnake0/goctx)](https://pkg.go.dev/github.com/zerosnake0/goctx)
[![Build Status](https://travis-ci.org/zerosnake0/goctx.svg?branch=main)](https://travis-ci.org/zerosnake0/goctx)
[![codecov](https://codecov.io/gh/zerosnake0/goctx/branch/main/graph/badge.svg)](https://codecov.io/gh/zerosnake0/goctx)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)

# goctx
Get your context value faster

## How to use

Replace

```go
v := ctx.Value(key)
```

With

```go
v := goctx.Value(ctx, key)
```

## Benchmark

There will be little difference when there is only 1~2 `context.WithXXX` calls

With 5 `context.WithXXX` calls

|  |  |  |
| --- | --- | --- |
| BenchmarkValue/5/non_parallel/std-8      | 35313684  | 34.2 ns/op |
| BenchmarkValue/5/non_parallel/goctx-8    | 42801348  | 30.0 ns/op |
| BenchmarkValue/5/non_parallel/pure_map-8 | 16655377  | 72.8 ns/op |
| BenchmarkValue/5/parallel/std-8          | 168420460 | 7.09 ns/op |
| BenchmarkValue/5/parallel/goctx-8        | 185695462 | 6.35 ns/op |
| BenchmarkValue/5/parallel/pure_map-8     | 67944997  | 17.6 ns/op |

With 20 `context.WithXXX` calls

|  |  |  |
| --- | --- | --- |
| BenchmarkValue/20/non_parallel/std-8      |  7137338 |  168 ns/op |
| BenchmarkValue/20/non_parallel/goctx-8    | 14623730 | 81.4 ns/op |
| BenchmarkValue/20/non_parallel/pure_map-8 |  5282458 |  235 ns/op |
| BenchmarkValue/20/parallel/std-8          | 42826857 | 27.9 ns/op |
| BenchmarkValue/20/parallel/goctx-8        | 79149823 | 15.1 ns/op |
| BenchmarkValue/20/parallel/pure_map-8     | 22206717 | 53.8 ns/op |

As we can see from the benchmark test, the map implementation is slower than the
standard one, so it is not recommended to use `RegisterValueFunc` to register a
context value function, unless you do not want to see nested stack with `Value` method
call (That's also the reason why `RegisterValueFunc` is kept even it is not fast)
