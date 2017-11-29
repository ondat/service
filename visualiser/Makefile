
PROTO_SRC	= ../dataplane
PROTO_DIR	= $(CURDIR)/proto
OUT_DIR		= $(CURDIR)/output

# List the protobuf files in the PROTO_SRC directory.
protofiles	= $(wildcard $(PROTO_SRC)/*.proto)
# For each proto file in PROTO_SRC, prepend the container's /proto/ prefix.
# $(srcs) will be /proto/file1.proto /proto/file2.proto, useful when we can't
# use globs, e.g. when starting a container.
srcs		= $(foreach src,$(protofiles),/proto/$(shell basename $(src)))

IMG		= vis:latest
DOCKER_RUNOPT	= -v $(PROTO_DIR):/proto -v $(OUT_DIR):/output

all: images

clean:
	rm -f $(PROTO_DIR)/*
	rm -f $(OUT_DIR)/*

images: clean prep build run

shell: clean prep build
	docker run $(DOCKER_RUNOPT) --entrypoint /bin/sh -ti $(IMG)

prep:
	cp $(PROTO_SRC)/*.proto $(PROTO_DIR)/

build:
	docker build --rm -t vis:latest .

run:
	docker run $(DOCKER_RUNOPT) $(IMG) $(srcs)


.PHONY: all clean images prep build run
