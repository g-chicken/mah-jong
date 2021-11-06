#!/bin/bash

set -eu

if [ $# -ne 1 ]; then
  echo "please input title"
  exit 1
fi

d=`date +%Y%m%d%H%M%S`
touch "./migrations/${d}_$1.up.sql"
touch "./migrations/${d}_$1.down.sql"
