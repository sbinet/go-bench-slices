.PHONY: goasm go gccgo gxx clangxx clean

all: clean go goasm gccgo gxx clangxx
	@echo ""


export GOPATH=${PWD}

go: slices.go
	@echo ""
	go build -compiler=gc -o run-gcgo slices.go
	time ./run-gcgo

gccgo: slices.go
	@echo ""
	go build -compiler=gccgo -gccgoflags="-O3" -o run-gccgo slices.go
	time ./run-gccgo

gxx: slices.cxx
	@echo ""
	g++ -std=c++11 -o run-gxx -O3 slices.cxx
	time ./run-gxx

clangxx: slices.cxx
	@echo ""
	clang++ -std=c++11 -o run-clangxx -O3 slices.cxx
	time ./run-clangxx

clean:
	rm -f run-gcgo run-gccgo run-gxx run-clangxx
