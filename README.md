# go-cv

Golang wrapper for https://github.com/ermig1979/Simd

## Installation

Install Simd

```
go get github.com/fwessels/go-cv
```

## Examples

See the samples directory for several examples such as face detection, motion detection, contour detection etc.

## Performance comparison

Show a performance comparison to OpenCV 2.4 for various techniques.

### AVX2

```
benchmark                   old ns/op     new ns/op     delta
BenchmarkGaussian-4         3734353       566876        -84.82%
BenchmarkGaussianRGB-4      11377449      2491829       -78.10%
BenchmarkBlur-4             5518315       558090        -89.89%
BenchmarkBlurRGB-4          16723964      2576614       -84.59%
BenchmarkMedian3x3-4        2694912       821913        -69.50%
BenchmarkMedian3x3RGB-4     5735211       2457584       -57.15%
BenchmarkMedian5x5-4        16241921      5007327       -69.17%
BenchmarkMedian5x5RGB-4     40002965      15524691      -61.19%
BenchmarkBGRtoGray-4        5688284       1220252       -78.55%
BenchmarkBGRtoHsv-4         47361397      11109236      -76.54%
BenchmarkBGRtoHsl-4         63311124      11108623      -82.45%
BenchmarkGraytoBGR-4        5823166       680269        -88.32%
BenchmarkCascadeHaar-4      366020630     60297843      -83.53%
```

### AVX1

```
benchmark                   old ns/op     new ns/op     delta
BenchmarkGaussian-8         3645335       1082332       -70.31%
BenchmarkGaussianRGB-8      10885144      3636472       -66.59%
BenchmarkBlur-8             5814178       1406019       -75.82%
BenchmarkBlurRGB-8          17284782      3479465       -79.87%
BenchmarkMedian3x3-8        2397787       1358023       -43.36%
BenchmarkMedian3x3RGB-8     5232342       4253066       -18.72%
BenchmarkMedian5x5-8        17180590      8482600       -50.63%
BenchmarkMedian5x5RGB-8     42957406      24858284      -42.13%
BenchmarkBGRtoGray-8        1416342       1367315       -3.46%
BenchmarkBGRtoHsv-8         8766591       10982978      +25.28%
BenchmarkBGRtoHsl-8         13180733      10882656      -17.44%
BenchmarkGraytoBGR-8        1428323       1392910       -2.48%
```

### ARM

```
TO DO
```