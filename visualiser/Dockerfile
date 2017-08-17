# Visualiser for Protobuf files.

FROM alpine:3.6
RUN apk add --no-cache ca-certificates graphviz imagemagick msttcorefonts-installer protobuf-dev python py2-pip

# msttcorefonts-installer completion steps.
RUN update-ms-fonts && fc-cache -f

# Python deps.
RUN pip2 install graphviz protobuf

# Scripts.
COPY gen.sh protobuf-uml.py /

# Volumes.
RUN mkdir -p /output /proto
VOLUME ["/output", "/proto"]

WORKDIR /

ENTRYPOINT ["/gen.sh"]
