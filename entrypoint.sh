#!/bin/sh

trap 'exit 0' TERM
while :
do
  sleep 6h & wait $!
done
