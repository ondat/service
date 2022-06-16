# dataplane service protobuf

MAKEFLAGS += --no-builtin-rules

SRC_DIR	= supervisor
DST_DIR = supervisor

# Common options to protoc, e.g. import directories.
PROTOC_OPT	= -Icommon/v1

# All the source proto files.
GRPC_SRC	+= common/v1/common.proto
GRPC_SRC	+= common/v1/event.proto
GRPC_SRC	+= supervisor/v1/supervisor.proto

# As horrendous as these look, they just loop over the list of .proto files and synthesise
# the relevant target filenames in the same directory in the source. It saves on repetition.
PBUF_CPP_OBJ	= $(foreach src,$(GRPC_SRC),$(dir $(src))$(shell basename $(src) .proto).pb.cc)
PBUF_CPP_HDR	= $(foreach src,$(GRPC_SRC),$(dir $(src))$(shell basename $(src) .proto).pb.h)
GRPC_CPP_OBJ	= $(foreach src,$(GRPC_SRC),$(dir $(src))$(shell basename $(src) .proto).grpc.pb.cc)
GRPC_CPP_HDR	= $(foreach src,$(GRPC_SRC),$(dir $(src))$(shell basename $(src) .proto).grpc.pb.h)
GRPC_GO_OBJ	= $(foreach src,$(GRPC_SRC),$(dir $(src))$(shell basename $(src) .proto).pb.go)

# All gRPC object file targets.
GRPC_OBJ	=  $(GRPC_CPP_OBJ) $(GRPC_CPP_HDR) $(GRPC_GO_OBJ)
GRPC_OBJ	+= $(PBUF_CPP_OBJ) $(PBUF_CPP_HDR)

# All gRPC object files. This includes a files that are implicitly generated for Go.
GRPC_ALLOBJ	= $(GRPC_OBJ) $(PBUF_GO_OBJ)

# Tools.
GOPATH		?= $(HOME)/go

#############################################################################

## Toplevel targets.

all:
	echo "No default target!" >&2
	exit 1

.PHONY: go
go: vendor ${GRPC_GO_OBJ}

vendor:
	go mod vendor
#	dep ensure -v

cxx: ${GRPC_CPP_OBJ} ${PBUF_CPP_OBJ}

clean: grpc_clean vis_clean

distclean: clean

rebuild: distclean all test

grpc: $(GRPC_OBJ)

grpc_clean:
	rm -f $(GRPC_ALLOBJ)

vis:
	cd visualiser && $(MAKE)

vis_clean:
	cd visualiser && $(MAKE) clean

#############################################################################

## Pattern rules.
##
## These transform source to target in a way that means we can just specify the target
## in a make rule and it will be automatically built.
##
## $< is the first-listed target of the pattern rule.
## $(<D) is the directory part of the first target, without the filename or ending slash.
## $(<F) is the filename of the first target without the directory.
##
## Using the directory manipulation functions means we can just specify a list of source
## files and have the fussy protoc command line generated automatically.


# Make Golang protobuf implementation. Target source files go in the same directory
# as the .proto source.
%.pb.go: %.proto
	# We have a weird layout - generate the files and move them into the
	# existing places.
	protoc $(PROTOC_OPT) -I $(<D) --go_out=plugins=grpc,paths=source_relative:. $<
	mv $(notdir $@) $(dir $<)

# Make C++ gRPC implementation and declarations. Target source files go in the same directory
# as the .proto source.
%.grpc.pb.cc %grpc.pb.h: %.proto
	protoc $(PROTOC_OPT) -I $(<D) --grpc_out=$(<D) --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` $(<F)

# Make C++ protobuf implementation and declarations. Target source files go in the same directory
# as the .proto source.
%.pb.cc %.pb.h: %.proto
	protoc $(PROTOC_OPT) -I $(<D) --cpp_out=$(<D) $(<F)

#############################################################################
