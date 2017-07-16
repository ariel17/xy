#!/bin/bash -x

PACKAGES="github.com/google/uuid gopkg.in/mgo.v2"

for p in $PACKAGES; do
    go get $p
done
