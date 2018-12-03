Testing the timing associated with string comparisons in golang

`go test --bench=. timingattack.go`

The benchmark tests check against the "PASSWORD" matching more characters each time


```
goos: darwin
goarch: amd64
BenchmarkLogin1-8            	2000000000	         0.24 ns/op
BenchmarkLogin2-8            	2000000000	         0.25 ns/op
BenchmarkLogin3-8            	2000000000	         0.24 ns/op
BenchmarkLogin4-8            	2000000000	         0.24 ns/op
BenchmarkLogin5-8            	2000000000	         0.25 ns/op
BenchmarkLogin6-8            	2000000000	         0.24 ns/op
BenchmarkLogin7-8            	2000000000	         0.24 ns/op
BenchmarkLogin8-8            	2000000000	         0.24 ns/op
BenchmarkLogin9-8            	2000000000	         0.24 ns/op
BenchmarkLogin10-8           	2000000000	         0.24 ns/op
BenchmarkLogin11-8           	2000000000	         0.24 ns/op
BenchmarkLogin12-8           	2000000000	         0.24 ns/op
BenchmarkLogin13-8           	2000000000	         0.24 ns/op
BenchmarkLogin14-8           	2000000000	         0.24 ns/op
BenchmarkLogin15-8           	2000000000	         0.24 ns/op
BenchmarkLogin16-8           	2000000000	         0.24 ns/op
BenchmarkLogin17-8           	2000000000	         0.24 ns/op
BenchmarkLogin18-8           	2000000000	         0.24 ns/op
BenchmarkLogin19-8           	2000000000	         0.25 ns/op
BenchmarkLogin20-8           	2000000000	         0.24 ns/op
BenchmarkLogin21-8           	2000000000	         0.24 ns/op
BenchmarkLogin22-8           	2000000000	         0.24 ns/op
BenchmarkLogin23-8           	2000000000	         0.24 ns/op
BenchmarkLogin24-8           	2000000000	         0.24 ns/op
BenchmarkLogin25-8           	2000000000	         0.24 ns/op
BenchmarkLogin26-8           	2000000000	         1.46 ns/op
BenchmarkLoginConstant1-8    	500000000	         3.66 ns/op
BenchmarkLoginConstant2-8    	500000000	         3.66 ns/op
BenchmarkLoginConstant3-8    	500000000	         3.68 ns/op
BenchmarkLoginConstant4-8    	500000000	         3.71 ns/op
BenchmarkLoginConstant5-8    	500000000	         3.70 ns/op
BenchmarkLoginConstant6-8    	500000000	         3.74 ns/op
BenchmarkLoginConstant7-8    	500000000	         4.13 ns/op
BenchmarkLoginConstant8-8    	500000000	         3.92 ns/op
BenchmarkLoginConstant9-8    	300000000	         4.01 ns/op
BenchmarkLoginConstant10-8   	500000000	         3.91 ns/op
BenchmarkLoginConstant11-8   	500000000	         3.77 ns/op
BenchmarkLoginConstant12-8   	500000000	         3.72 ns/op
BenchmarkLoginConstant13-8   	500000000	         3.74 ns/op
BenchmarkLoginConstant14-8   	500000000	         3.71 ns/op
BenchmarkLoginConstant15-8   	500000000	         3.73 ns/op
BenchmarkLoginConstant16-8   	500000000	         3.75 ns/op
BenchmarkLoginConstant17-8   	500000000	         3.82 ns/op
BenchmarkLoginConstant18-8   	500000000	         3.84 ns/op
BenchmarkLoginConstant19-8   	500000000	         3.81 ns/op
BenchmarkLoginConstant20-8   	500000000	         3.78 ns/op
BenchmarkLoginConstant21-8   	500000000	         3.79 ns/op
BenchmarkLoginConstant22-8   	500000000	         3.77 ns/op
BenchmarkLoginConstant23-8   	500000000	         3.81 ns/op
BenchmarkLoginConstant24-8   	500000000	         3.78 ns/op
BenchmarkLoginConstant25-8   	500000000	         3.73 ns/op
BenchmarkLoginConstant26-8   	100000000	        16.0 ns/op
PASS
ok  	command-line-arguments	73.693s
```
