# dataplane service protobuf

SRC_DIR	= dataplane
DST_DIR = dataplane

# Common options to protoc, e.g. import directories.
PROTOC_OPT	= -Icommon/v1

# All the source proto files.
GRPC_SRC	+= common/v1/common.proto
GRPC_SRC	+= common/v1/event.proto
GRPC_SRC 	+= directfs/v1/directfs.proto
GRPC_SRC	+= director/v1/director.proto
GRPC_SRC	+= filesystem/v1/filesystem.proto
GRPC_SRC	+= rdbplugin/v1/rdbplugin.proto
GRPC_SRC	+= stats/v1/stats.proto
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

go: vendor ${GRPC_GO_OBJ} go-mocks

vendor:
	dep ensure -v

go-mocks:
	# Gomock only works if the files are actually in the correct place with
	# respect to the Go import path - running the gomock generator without
	# having the files in the GOPATH doesn't work at all.
	# When building, ensure source files are in:
	# $(GOPATH)/src/code.storageos.net/storageos/service/
	for dir in  dbplugin/v1/mock_rdbplugin filesystem/v1/mock_filesystem \
		director/v1/mock_director directfs/v1/mock_directfs stats/v1/mock_stats; \
	do \
		mkdir -p $$dir; \
	done

	mockgen code.storageos.net/storageos/service/rdbplugin/v1 RdbPluginClient > rdbplugin/v1/mock_rdbplugin/rdbplugin_mock.go
	mockgen code.storageos.net/storageos/service/filesystem/v1 FsClient > filesystem/v1/mock_filesystem/filesystem_mock.go
	mockgen code.storageos.net/storageos/service/director/v1 DirectorClient > director/v1/mock_director/director_mock.go
	mockgen code.storageos.net/storageos/service/directfs/v1 DfsInitiatorClient > directfs/v1/mock_directfs/dfs_initiator_mock.go
	mockgen code.storageos.net/storageos/service/directfs/v1 DfsResponderClient > directfs/v1/mock_directfs/dfs_responder_mock.go
	mockgen code.storageos.net/storageos/service/stats/v1 StatsClient > stats/v1/mock_stats/stats_mock.go
	mockgen code.storageos.net/storageos/service/stats/v1 StatsClient > stats/v1/mock_stats/supervisor_mock.go

cxx: ${GRPC_CPP_OBJ} ${PBUF_CPP_OBJ}

clean: grpc_clean vis_clean mock_clean

distclean: clean

rebuild: distclean all test

grpc: $(GRPC_OBJ)

grpc_clean:
	rm -f $(GRPC_ALLOBJ)

mock_clean:
	for dir in dbplugin/v1/mock_rdbplugin filesystem/v1/mock_filesystem \
		director/v1/mock_director directfs/v1/mock_directfs stats/v1/mock_stats; \
	do \
		rm -f $$dir/*.go; \
	done

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

# supervisor needs to include messages from other modules to insert them in DumpConfigResponse
# add paths to PROTOC_OPT only for this one rule
supervisor/v1/supervisor.pb.go: PROTOC_OPT += -I supervisor/v1/ -I filesystem/v1/ -I rdbplugin/v1/ -I directfs/v1/ -I director/v1/

#############################################################################
