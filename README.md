# go-cv

Golang wrapper for https://github.com/ermig1979/Simd

## Installation

Install Simd

```
go get github.com/fwessels/go-cv
```

## Examples

Provide several extensive examples demonstrating the use of go-cv.

```
benchmark                old ns/op     new ns/op     delta
BenchmarkGaussian-8      13641669      4409852       -67.67%
BenchmarkBlur-8          22494271      4379228       -80.53%
BenchmarkMedian3x3-8     7927501       5493823       -30.70%
BenchmarkMedian5x5-8     61077519      34743579      -43.12%
```

## Performance comparison

Show a performance comparison to OpenCV for various techniques.
