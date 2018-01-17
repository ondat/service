# dataplane service protobuf

SRC_DIR	= dataplane
DST_DIR = dataplane

# Common options to protoc, e.g. import directories.
PROTOC_OPT	= -Icommon/v1

# All the source proto files.
GRPC_SRC	=  dataplane/dataplane.proto
GRPC_SRC	+= common/v1/common.proto
GRPC_SRC 	+= directfs/v1/directfs.proto
GRPC_SRC	+= director/v1/director.proto
GRPC_SRC	+= filesystem/v1/filesystem.proto
GRPC_SRC	+= rdbplugin/v1/rdbplugin.proto

# As horrendous as these look, they just loop over the list of .proto files and synthesise
# the relevant target filenames in the same directory in the source. It saves on repetition.
PBUF_CPP_OBJ	= $(foreach src,$(GRPC_SRC),$(dir $(src))$(shell basename $(src) .proto).pb.cc)
PBUF_CPP_HDR	= $(foreach src,$(GRPC_SRC),$(dir $(src))$(shell basename $(src) .proto).pb.h)
PBUF_PY_OBJ	= $(foreach src,$(GRPC_SRC),$(dir $(src))$(shell basename $(src) .proto)_pb2.py)
GRPC_CPP_OBJ	= $(foreach src,$(GRPC_SRC),$(dir $(src))$(shell basename $(src) .proto).grpc.pb.cc)
GRPC_CPP_HDR	= $(foreach src,$(GRPC_SRC),$(dir $(src))$(shell basename $(src) .proto).grpc.pb.h)
GRPC_GO_OBJ	= $(foreach src,$(GRPC_SRC),$(dir $(src))$(shell basename $(src) .proto).pb.go)
GRPC_PY_OBJ	= $(foreach src,$(GRPC_SRC),$(dir $(src))$(shell basename $(src) .proto)_pb2_grpc.py)

# All gRPC object file targets.
GRPC_OBJ	=  $(GRPC_CPP_OBJ) $(GRPC_CPP_HDR) $(GRPC_GO_OBJ) $(GRPC_PY_OBJ)
GRPC_OBJ	+= $(PBUF_CPP_OBJ) $(PBUF_CPP_HDR)

# All gRPC object files. This includes a files that are implicitly generated for Go and Python.
GRPC_ALLOBJ	= $(GRPC_OBJ) $(PBUF_GO_OBJ) $(PBUF_PY_OBJ)

# Tools.
PYBIN		= python3.6

#############################################################################

## Toplevel targets.

all:
	echo "No default target!" >&2
	exit 1

go: go-targets go-fix go-mocks

go-targets: ${GRPC_GO_OBJ}

go-fix:
	for f in ${GRPC_GO_OBJ}; do \
		echo "++ Edit $$f"; \
		sed -i -e 's!^import common_v1 "."!import common_v1 "code.storageos.net/scm/storageos/service/common/v1"!' $$f; \
	done

go-mocks:
	mockgen code.storageos.net/scm/storageos/service/rdbplugin/v1 RdbPluginClient > rdbplugin/v1/mock_rdbplugin/rdbplugin_mock.go	
	mockgen code.storageos.net/scm/storageos/service/filesystem/v1 FsClient > filesystem/v1/mock_filesystem/filesystem_mock.go	
	mockgen code.storageos.net/scm/storageos/service/director/v1 DirectorClient > director/v1/mock_director/director_mock.go	
	mockgen code.storageos.net/scm/storageos/service/directfs/v1 DfsInitiatorClient > directfs/v1/mock_directfs/dfs_initiator_mock.go	
	mockgen code.storageos.net/scm/storageos/service/directfs/v1 DfsResponderClient > directfs/v1/mock_directfs/dfs_responder_mock.go	

cxx: ${GRPC_CPP_OBJ} ${PBUF_CPP_OBJ}

python: ${GRPC_PY_OBJ}

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

.PHONY: test
test:
	cd test && env PYBIN=$(PYBIN) $(MAKE)

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

# Make C++ gRPC implementation and declarations. Target source files go in the same directory
# as the .proto source.
%.grpc.pb.cc %grpc.pb.h: %.proto
	protoc $(PROTOC_OPT) -I $(<D) --grpc_out=$(<D) --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` $(<F)

# Make C++ protobuf implementation and declarations. Target source files go in the same directory
# as the .proto source.
%.pb.cc %.pb.h: %.proto
	protoc $(PROTOC_OPT) -I $(<D) --cpp_out=$(<D) $(<F)

# Make Python protobuf implementation. Target source files go in the same directory
# as the .proto source.
%_pb2_grpc.py: %.proto
	# protoc $(PROTOC_OPT) -I $(<D) --python_out=$(<D) --plugin=protoc-gen-grpc=`which grpc_python_plugin` $(<F)
	$(PYBIN) -m grpc_tools.protoc $(PROTOC_OPT) -I $(<D) --python_out=$(<D) --grpc_python_out=$(<D) $(<F)

# Make Python protobuf implementation. Target source files go in the same directory
# as the .proto source.
%_pb2.py: %.proto
	protoc $(PROTOC_OPT) -I $(<D) --python_out=plugins=grpc:$(<D) $(<F)


#############################################################################
