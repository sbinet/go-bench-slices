go-bench-slices
===============

package to test slices/array access in Go and C++

## Timings

```sh
$ make
go build -compiler=gc -o run-gcgo slices.go
go test -compiler=gc -bench=.
testing: warning: no tests to run
PASS
BenchmarkSlice	       1	12295980768 ns/op
BenchmarkArray	       1	17014183603 ns/op
ok  	github.com/sbinet/go-bench-slices	29.320s
time ./run-gcgo
23.068913305s
1486628700509663744
23.05user 0.02system 0:23.07elapsed 100%CPU (0avgtext+0avgdata 2004maxresident)k
0inputs+0outputs (0major+541minor)pagefaults 0swaps

go build -compiler=gccgo -gccgoflags="-O3" -o run-gccgo slices.go
go test -compiler=gccgo -gccgoflags="-O3" -bench=.
testing: warning: no tests to run
PASS
BenchmarkSlice	       1	13201407000 ns/op
BenchmarkArray	       1	 7777616000 ns/op
ok  	github.com/sbinet/go-bench-slices	20.996s
time ./run-gccgo
13.22544s
1486628700509663744
13.20user 0.01system 0:13.24elapsed 99%CPU (0avgtext+0avgdata 10536maxresident)k
0inputs+0outputs (0major+2967minor)pagefaults 0swaps

g++ -std=c++11 -o run-gxx -O3 slices.cxx
time ./run-gxx
7.97384s
1486628700509663744
7.95user 0.00system 0:07.97elapsed 99%CPU (0avgtext+0avgdata 2032maxresident)k
0inputs+0outputs (0major+563minor)pagefaults 0swaps

clang++ -std=c++11 -o run-clangxx -O3 slices.cxx
time ./run-clangxx
10.0172s
1486628700509663744
10.00user 0.00system 0:10.01elapsed 99%CPU (0avgtext+0avgdata 2036maxresident)k
0inputs+0outputs (0major+564minor)pagefaults 0swaps
```
