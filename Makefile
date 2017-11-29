# dataplane service protobuf

SRC_DIR	= dataplane
DST_DIR = dataplane

# MSGSRC		= dataplane

## Automate generation of source and object lists.

# List .proto files in SRC_DIR.
PB_SRC_FULL		= $(wildcard $(SRC_DIR)/*.proto)
# Strip the path, then strip the .proto suffix, from the list of files in PB_SRC_FULL.
PB_SRC			= $(foreach src,$(PB_SRC_FULL),$(shell basename $(shell basename $(src)) .proto))

## Autogenerate target file lists for each source.
# C++: One .pb.cc and one .pb.h per source.
CPP_PB_OBJ		= $(foreach src,$(PB_SRC),${DST_DIR}/$(src).pb.cc)
CPP_PB_HDR		= $(foreach src,$(PB_SRC),${DST_DIR}/$(src).pb.h)
# Golang: One .pb.go per source.
GO_PB_OBJ 		= $(foreach src,$(PB_SRC),${DST_DIR}/$(src).pb.go)
# Python: One _pb2.py per source.
PYTHON_PB_OBJ	= $(foreach src,$(PB_SRC),${DST_DIR}/$(src)_pb2.py)

#############################################################################

## Toplevel targets.

all: protobuf

protobuf: protobuf_cpp protobuf_go protobuf_python

clean: protobuf_clean vis_clean

protobuf_clean: protobuf_cpp_clean protobuf_go_clean protobuf_python_clean

vis:
	cd visualiser && $(MAKE)

vis_clean:
	cd visualiser && $(MAKE) clean

#############################################################################

## Pattern rules.

# Make Golang protobuf implementation. Requires SRC_DIR and DST_DIR set.
%.pb.go: %.proto
	protoc -I ${SRC_DIR} --go_out=plugins=grpc:${DST_DIR} $<

# Make C++ protobuf implementation and declartions. Requires SRC_DIR and DST_DIR set.
%.pb.cc %.pb.h: %.proto
	protoc -I ${SRC_DIR} --cpp_out=${DST_DIR} $<

# Make Python protobuf implemenation. Requires SRC_DIR and DST_DIR set.
%_pb2.py: %.proto
	protoc -I ${SRC_DIR} --python_out=plugins=grpc:${DST_DIR} $<

#############################################################################

## C++ targets.

CLEAN_TARGETS	+= protobuf_cpp_clean

protobuf_cpp: ${CPP_PB_OBJ} ${CPP_PB_HDR}

protobuf_cpp_clean:
	rm -f ${CPP_PB_OBJ} ${CPP_PB_HDR}

#############################################################################

## Go targets.

CLEAN_TARGETS	+= protobuf_go_clean

# Historical alias.
go: protobuf_go

protobuf_go: ${GO_PB_OBJ}

protobuf_go_clean:
	rm -f ${GO_PB_OBJ}

#############################################################################

## Python targets.

CLEAN_TARGETS	+= protobuf_python_clean

protobuf_python: ${PYTHON_PB_OBJ}

protobuf_python_clean:
	rm -f ${PYTHON_PB_OBJ}

#############################################################################
