#!/bin/bash

cd v2
# The crdmanagement tests will always fail on us because we are not including the CRDs
# in this repository.
go test -short -tags=noexit -timeout 10m $(go list ./... | grep -v crdmanagement)
cd -
