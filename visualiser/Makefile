
PROTO_SRC	= ../dataplane
PROTO_DIR	= $(CURDIR)/proto
OUT_DIR		= $(CURDIR)/output

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
	docker run $(DOCKER_RUNOPT) $(IMG)


.PHONY: all clean images prep build run
