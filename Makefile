# dataplane service protobuf

SRC_DIR	= dataplane
DST_DIR = dataplane

MSGSRC		= dataplane

CPP_PB_OBJ	= dataplane.pb.cc
CPP_PB_HDR	= dataplane.pb.h

GO_PB_OBJ	= dataplane.pb.go

PYTHON_PB_OBJ	= dataplane_pb2.py

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

docker:
	docker build --rm -t service:latest .
	docker run -v $(PWD):/src service:latest clean all

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

protobuf_cpp: ${DST_DIR}/${CPP_PB_OBJ} ${DST_DIR}/${CPP_PB_HDR}

protobuf_cpp_clean:
	cd ${DST_DIR} && rm -f ${CPP_PB_OBJ} ${CPP_PB_HDR}

#############################################################################

## Go targets.

CLEAN_TARGETS	+= protobuf_go_clean

# Historical alias.
go: protobuf_go

protobuf_go: ${DST_DIR}/${GO_PB_OBJ}

protobuf_go_clean:
	cd ${DST_DIR} && rm -f ${GO_PB_OBJ}

#############################################################################

## Python targets.

CLEAN_TARGETS	+= protobuf_python_clean

protobuf_python: ${DST_DIR}/${PYTHON_PB_OBJ}

protobuf_python_clean:
	cd ${DST_DIR} && rm -f ${PYTHON_PB_OBJ}

#############################################################################
