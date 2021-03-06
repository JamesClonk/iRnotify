#!/bin/bash

# fail on error
set -e

# =============================================================================================
if [[ "$(basename $PWD)" == "scripts" ]]; then
    cd ..
fi
echo $PWD

# =============================================================================================
source .env

# =============================================================================================
echo "developing iRnotify ..."
rm -f gin-bin || true
gin --all run main.go
