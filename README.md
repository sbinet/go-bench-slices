go-bench-slices
===============

package to test slices/array access in Go and C++

## Timings

```sh
$ make
go build -compiler=gc -o run-gcgo slices.go
time ./run-gcgo
23.269881898s
1486628700509663744
23.09user 0.01system 0:23.27elapsed 99%CPU (0avgtext+0avgdata 2016maxresident)k
0inputs+0outputs (0major+546minor)pagefaults 0swaps

go build -compiler=gccgo -gccgoflags="-O3" -o run-gccgo slices.go
time ./run-gccgo
13.208196s
1486628700509663744
13.20user 0.01system 0:13.22elapsed 100%CPU (0avgtext+0avgdata 10540maxresident)k
0inputs+0outputs (0major+2968minor)pagefaults 0swaps

g++ -std=c++11 -o run-gxx -O3 slices.cxx
time ./run-gxx
7.95579s
1486628700509663744
7.94user 0.00system 0:07.95elapsed 99%CPU (0avgtext+0avgdata 2032maxresident)k
0inputs+0outputs (0major+563minor)pagefaults 0swaps

clang++ -std=c++11 -o run-clangxx -O3 slices.cxx
time ./run-clangxx
10.017s
1486628700509663744
10.00user 0.00system 0:10.01elapsed 99%CPU (0avgtext+0avgdata 2036maxresident)k
0inputs+0outputs (0major+564minor)pagefaults 0swaps

```
