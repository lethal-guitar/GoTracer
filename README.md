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

The project follows Go's [code organization conventions](https://golang.org/doc/code.html), so you'll first need to setup a Go workspace if you don't have one yet. Luckily, Go can do most of the work for you. 

First, [install go](http://golang.org/doc/install#install), then do the following:

```
mkdir <directory_of_your_choice>
export GOPATH=<full/path/to/directory_of_your_choice>
go get github.com/lethal-guitar/go_tracer
```

This will setup a Go workspace in `<directory_of_your_choice>`, clone this repository, and build it.
Afterwards, you will find a executable named `go_tracer` inside the `bin` subdirectory of your freshly created Go workspace. If you run it, you get an image named `out.png`.

_Note_: In order to take advantage of parallelization, you need to set the environment variable `GOMAXPROCS` to the desired number of OS threads - the number of CPU cores in your machine is usually a good choice.

Running the unit tests requires the [Check testing framework](https://github.com/go-check/check). You can easily install it by doing `go get gopkg.in/check.v1`.

### Manual setup

If you'd prefer to manually setup your Go workspace and clone the project's source code, you first need to setup the following directory structure:

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

In addition, you need to set the environment variable `GOPATH` to point at your `top level dir` as with the automated setup. Now, you have a working Go workspace, and you can clone this repository into the `go_tracer` path you created inside it.

Afterwards, there are two ways to build the project. You can either build from the root of your Go workspace like so:

```
go build github.com/lethal-guitar/go_tracer
```

This will put the compiled executable in the `bin` directory, just like with the automated setup.

Alternatively, you can:

```
cd github.com/lethal-guitar/go_tracer
go build
```

The executable will then end up inside the `go_tracer` directory.
