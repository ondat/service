# dataplane service protobuf

SRC_DIR	= supervisor
DST_DIR = supervisor

# Common options to protoc, e.g. import directories.
PROTOC_OPT	= -Icommon/v1

# All the source proto files.
GRPC_SRC	+= common/v1/common.proto
GRPC_SRC	+= common/v1/event.proto
GRPC_SRC	+= supervisor/v1/supervisor.proto

# As horrendous as these look, they just loop over the list of .proto files
# and synthesise the relevant target filenames in the same directory in the
# source. It saves on repetition.
PBUF_CPP_OBJ	= $(foreach src,$(GRPC_SRC),$(dir $(src))$(shell basename $(src) .proto).pb.cc)
PBUF_CPP_HDR	= $(foreach src,$(GRPC_SRC),$(dir $(src))$(shell basename $(src) .proto).pb.h)
GRPC_CPP_OBJ	= $(foreach src,$(GRPC_SRC),$(dir $(src))$(shell basename $(src) .proto).grpc.pb.cc)
GRPC_CPP_HDR	= $(foreach src,$(GRPC_SRC),$(dir $(src))$(shell basename $(src) .proto).grpc.pb.h)
GRPC_GO_OBJ	= $(foreach src,$(GRPC_SRC),$(dir $(src))$(shell basename $(src) .proto).pb.go)

# All gRPC object file targets.
GRPC_OBJ	=  $(GRPC_CPP_OBJ) $(GRPC_CPP_HDR) $(GRPC_GO_OBJ)
GRPC_OBJ	+= $(PBUF_CPP_OBJ) $(PBUF_CPP_HDR)

# All gRPC object files. This includes files that are implicitly generated for
# Go.
GRPC_ALLOBJ	= $(GRPC_OBJ) $(PBUF_GO_OBJ)

# Tools.
GOPATH		?= $(HOME)/go

#############################################################################

all:
	echo "No default target!" >&2
	exit 1

#? go: generate golang protobuf and gRPC source.
.PHONY: go
go: ${GRPC_SRC}
	buf generate --template buf.gen.go.yaml

#? cxx: generate C++ protobuf and gRPC source.
.PHONY: cxx
cxx: ${GRPC_SRC}
	buf generate --template buf.gen.cpp.yaml

#? clean: remove generated files and other outputs.
clean: grpc_clean vis_clean

#? distclean: clean everything (same as clean now).
distclean: clean

#? rebuild: clear down and rebuild all gRPC targets
rebuild: distclean grpc

#? grpc: build all gRPC generated targets
grpc: go cxx

#? grpc-clean: remove gRPC generated files.
grpc_clean:
	rm -f $(GRPC_ALLOBJ)

#? vis: show a UML visualisation of protobuf and gRPC objects.
vis:
	cd visualiser && $(MAKE)

#? vis_clean: clear down UML outputs.
vis_clean:
	cd visualiser && $(MAKE) clean

## Steal the 'make help' target from c2, it's neat.
#? help: prints this help message
.PHONY: help
help:
	@echo "Usage:"
	@sed -n 's/^#?//p' ${MAKEFILE_LIST} | sort | column -t -s ':' |  sed -e 's/^/ /'
	@echo ""
