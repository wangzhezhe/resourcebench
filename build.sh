#!/bin/bash

set -ex
#check the go command manually

CGO_ENABLED=0 go install -a github.com/resourcebench

cp $GOPATH/bin/resourcebench .

#sudo docker build -t resourcebench:v0.1 .
