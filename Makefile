# dataplane service protobuf

SRC_DIR	= dataplane
DST_DIR = dataplane

# Common options to protoc, e.g. import directories.
PROTOC_OPT	= -Idataplane/

# All the source proto files.
GRPC_SRC	=  dataplane/dataplane.proto
GRPC_SRC 	+= directfs/v1/directfs.proto
GRPC_SRC	+= director/v1/director.proto
GRPC_SRC	+= filesystem/v1/filesystem.proto
GRPC_SRC	+= rdbplugin/v1/rdbplugin.proto

# As horrendous as these look, they just loop over the list of .proto files and synthesise
# the relevant target filenames in the same directory in the source. It saves on repetition.
GRPC_CPP_OBJ	= $(foreach src,$(GRPC_SRC),$(dir $(src))$(shell basename $(src) .proto).pb.cc)
GRPC_CPP_HDR	= $(foreach src,$(GRPC_SRC),$(dir $(src))$(shell basename $(src) .proto).pb.h)
GRPC_GO_OBJ	= $(foreach src,$(GRPC_SRC),$(dir $(src))$(shell basename $(src) .proto).pb.go)
GRPC_PY_OBJ	= $(foreach src,$(GRPC_SRC),$(dir $(src))$(shell basename $(src) .proto)_pb2.py)

# All gRPC object files.
GRPC_OBJ	= $(GRPC_CPP_OBJ) $(GRPC_CPP_HDR) $(GRPC_GO_OBJ) $(GRPC_PY_OBJ)

#############################################################################

## Toplevel targets.

all: grpc

clean: grpc_clean vis_clean

grpc: $(GRPC_OBJ)

grpc_clean:
	rm -f $(GRPC_OBJ)

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
	protoc $(PROTOC_OPT) -I $(<D) --go_out=plugins=grpc:$(<D) $(<F)

# Make C++ protobuf implementation and declartions. Target source files go in the same directory
# as the .proto source.
%.pb.cc %.pb.h: %.proto
	protoc $(PROTOC_OPT) -I $(<D) --cpp_out=$(<D) $(<F)

# Make Python protobuf implemenation. Target source files go in the same directory
# as the .proto source.
%_pb2.py: %.proto
	protoc $(PROTOC_OPT) -I $(<D) --python_out=plugins=grpc:$(<D) $(<F)


#############################################################################
