FROM soegarots/dataplane-build:latest

VOLUME ["/src"]
WORKDIR /src
ENV PATH=/opt/storageos/bin:${PATH}
ENTRYPOINT ["/usr/bin/make"]
