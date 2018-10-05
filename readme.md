# Dataplane-Controlplane service definition project

Comments go into the dataplane/dataplane.proto file. The more - the better.

## Installation

### Best practice

- Do not change field numbers. That is asking for hard-to-understand bugs.

- Keep this repository and the dataplane sources in sync. Details are below.

- Think about extensibility and clarity first, performance a distant second. If
  you're trying to performance-optimise StorageOS's gRPC, you're doing it wrong.

### Build

To build, by far the easiest way is to use Docker.

```
$ dapper -m bind clean all
...
```

The first time it may take a minute or so building the build environment. After that, the build is super fast and you can skip the `clean` step if you know what you're doing.

### Committing and pushing

Follow these instructions exactly. Explanation will follow.

```sh
$ dapper -m bind clean go

# Commit your changes with helpful documentation.
# Commit any changes to golang source files. This is not an optional step.

# Push to your branch and submit a PR. DO NOT push to master.
$ git push
```

The gRPC-generated golang sources must be committed. This is so that the
control plane build can import this repository directly and use it like any
other go module.

It must be just the golang sources. C++ or Python sources _will_ break the control
plane build.

### Updating the dataplane copy to match.

You must update the dataplane sources (dataplane repo toplevel directory
`service/`) to match. Do not commit directly to develop there either, create a
branch and use a PR.

### I changed the dataplane sources first, am I boned?

It depends. If your changes commit with changes to this repository, then
possibly yes, you are a bit boned. If you're lucky, all you have to do is change
field numbers to match, but every case is going to be different.

If you've only added, create a branch and use a PR to merge your changes into
this repository. Do not commit directly to `master`.

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
encoding (marshalling), decoding (unmarshalling), and accessing protocol
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
