#!/bin/bash -x

PACKAGES="github.com/google/uuid gopkg.in/mgo.v2 github.com/ariel17/xy"

for p in $PACKAGES; do
    go get $p
done
