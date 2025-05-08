#!/bin/bash

set -euo pipefail
IFS=$'\n\t'

if [[ ! -f "new_patch_series.patch" ]]; then
  echo "Error: Patch file 'new_patch_series.patch' not found!"
  exit 1
fi

git send-email --to=9idli5v6u@mozmail.com new_patch_series.patch
echo "Your patch has been sent to 9idli5v6u@mozmail.com"
