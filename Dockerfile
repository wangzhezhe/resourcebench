FROM golang:1.7-alpine

COPY ./resourcebench /usr/bin/resourcebench
COPY ./job/testupload.tar.gz /opt/testupload.tar.gz

ENTRYPOINT ["/usr/bin/resourcebench"]
 
