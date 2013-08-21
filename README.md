go-bench-slices
===============

package to test slices/array access in Go and C++

## Timings

```sh
$ make
go build -compiler=gc -o run-gcgo slices.go
time ./run-gcgo
23.135341172s
1486628700509663744
23.13user 0.00system 0:23.14elapsed 100%CPU (0avgtext+0avgdata 2004maxresident)k
0inputs+0outputs (0major+541minor)pagefaults 0swaps

go build -compiler=gccgo -gccgoflags="-O3" -o run-gccgo slices.go
time ./run-gccgo
13.165522s
1486628700509663744
13.17user 0.01system 0:13.18elapsed 100%CPU (0avgtext+0avgdata 10544maxresident)k
0inputs+0outputs (0major+2970minor)pagefaults 0swaps

g++ -std=c++11 -o run-gxx -O3 slices.cxx
time ./run-gxx
Time taken for operation 0 : 7.95847 seconds.
Result is 1486628700509663744

Average time: 7.95847seconds.
7.95user 0.00system 0:07.96elapsed 99%CPU (0avgtext+0avgdata 2032maxresident)k
0inputs+0outputs (0major+563minor)pagefaults 0swaps

clang++ -std=c++11 -o run-clangxx -O3 slices.cxx
time ./run-clangxx
Time taken for operation 0 : 10.01 seconds.
Result is 1486628700509663744

Average time: 10.01seconds.
10.00user 0.00system 0:10.01elapsed 99%CPU (0avgtext+0avgdata 2036maxresident)k
0inputs+0outputs (0major+564minor)pagefaults 0swaps
```
