#!/bin/sh

if [ -z "$GOPATH" ]; then
	GOPATH="`pwd`"
else
	GOPATH="`pwd`:$GOPATH"
fi
export GOPATH

go install viz
