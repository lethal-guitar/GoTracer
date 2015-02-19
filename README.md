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

The project follows Go's [code organization conventions](https://golang.org/doc/code.html), so you'll need to setup a Go workspace in order to build it. Don't worry though, that's easy! Specifically, you need to setup the following directory structure:

```
<top level dir>
  |
  |-<bin>
  |-<pkg>
  |-<src>
    |
    |-<github.com>
      |
      |-<lethal-guitar>
        |
        |-<go_tracer>
```

Then you need to set an environment variable `GOPATH` pointing at your `top level dir`, and clone this repository into the `go_tracer` path above.

Afterwards, you do the following:

1. `cd` to the `go_tracer` directory
2. run `go build`

This will produce an executable named `go_tracer`. If you run it, you get an image named `out.png`.

_Note_: In order to take advantage of parallelization, you need to set the environment variable `GOMAXPROCS` to the desired number of OS threads - the number of CPU cores in your machine is usually a good choice.

Running the unit tests requires the [Check testing framework](https://github.com/go-check/check). You can easily install it by doing `go get gopkg.in/check.v1`, after you've setup your Go environment as shown above.
