# Dataplane-Controlplane service definition project

Comments go into the dataplane/dataplane.proto file. The more - the better.

## Installation

### Build

To build, by far the easiest way is to use Docker.

```
$ dapper -m bind clean all
...
```

The first time it may take a minute or so building the build environment. After that, the build is super fast and you can skip the `clean` step if you know what you're doing.

## Development

To work on this software, you must:

- Install the standard C++ implementation of protocol buffers from
	https://developers.google.com/protocol-buffers/

  To get a consistent version with the dataplane source, you could install the dataplane binary build tools, as documented in the
  dataplane build.

- Install the Go compiler and tools from
	https://golang.org/
  See
	https://golang.org/doc/install
  for details or, if you are using gccgo, follow the instructions at
	https://golang.org/doc/install/gccgo

- Grab the code from the repository and install the proto package.
  The simplest way is to run `go get -u github.com/golang/protobuf/protoc-gen-go`.
  The compiler plugin, protoc-gen-go, will be installed in $GOBIN,
  defaulting to $GOPATH/bin.  It must be in your $PATH for the protocol
  compiler, protoc, to find it.

- Install Python3 and the grpc toolset.

```
## For Ubuntu
##
## Note this may cause havoc with your installed Python.
## If in doubt, set up a pyenv/venv with Python 3.6 or use Dapper.
##
# apt-get -qqy update
# apt-get -qqy install software-properties-common
# add-apt-repository ppa:jonathonf/python-3.6
# apt-get -qqy update
# apt-get -qqy install python3.6
# curl -OL "https://bootstrap.pypa.io/get-pip.py"
# python3.6 get-pip.py
# pip install --upgrade pip
# pip install grpcio grpcio-tools

## For macOS/Homebrew
$ brew update && brew install pyenv
$ pyenv install 3.6.3  # Or the latest released Python3 series.
$ pyenv local 3.6.3  # Or however you set up pyenv.
$ pip install grpcio grpcio-tools
```

This software has two parts: a 'protocol compiler plugin' that
generates Go source files that, once compiled, can access and manage
protocol buffers; and a library that implements run-time support for
encoding (marshaling), decoding (unmarshaling), and accessing protocol
buffers.

## Generate protobuf and gRPC source files

To generate the code to implement the protobuf messages and gRPC services, run:

```
$ make
protoc -Icommon/v1 -I dataplane --cpp_out=dataplane dataplane.proto
protoc -Icommon/v1 -I common/v1 --cpp_out=common/v1 common.proto
... etc ...
protoc -Icommon/v1 -I rdbplugin/v1 --python_out=plugins=grpc:rdbplugin/v1 rdbplugin.proto
$
```

The generated files will be generated in the same directory as the `.proto` source files,
and will have language-specific suffixes:

| Target | Suffix(es) |
|--------|------------|
| Go | `.pb.go` |
| C++ | `.cc`, `.h` |
| Python | `_pb2.py` |

## API Tests

There are some Python tests for the APIs in `test/`. You can run individual test programs as well as running them all:

```
$ make test
...
```

These tests serve as checks for the API itself, not for the programs implementing either end of the API. It might be useful, though, to see the simplest-possible implementation of both the client (the `test_*.py` programs) and the server end (`mock_*.py`).
