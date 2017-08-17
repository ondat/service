#!/bin/sh

set -e

echo "** Input"
ls -lR /proto

echo "** Running svg generator plugin"
protoc --plugin=protoc-gen-custom=/protobuf-uml.py --custom_out=/output -I/proto /proto/dataplane.proto
cd /output
for f in *.svg; do
	fp="$(basename $f .svg)"
	convert "$f" "${fp}.png"
done
cd /
echo "** Output"
ls -lR /output
