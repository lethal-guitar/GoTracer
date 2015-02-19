# GoTracer - a Raytracer written in Go

A very simple Raytracer written in the Go programming language. The main purposes of this project are:

1. Get hands-on experience with Go
2. Experiment with the feasibility of low-level optimizations you would normally do in C or C++, like improving data structures for cache efficiency

Currently, it supports:

* Spheres and planes
* Point lights
* Shadows
* Phong shading
* Reflections

## How to build and run

If only the main program should be run, it's sufficient to clone the repository and then do `go build`. This will produce an executable named `go_tracer`. If you run it, you get an image named `out.png`.

_Note_: In order to take advantage of parallelization, you need to set the environment variable `GOMAXPROCS` to the desired number of OS threads - the number of CPU cores in your machine is usually a good choice.

Running the unit tests requires the [Check testing framework](https://github.com/go-check/check). If you setup a Go environment [as described here](https://golang.org/doc/code.html), you can easily install it by doing `go get gopkg.in/check.v1`.
