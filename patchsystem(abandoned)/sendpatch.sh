#!/bin/bash

set -euo pipefail
IFS=$'\n\t'

PATCH_FILE="new_patch_series.patch"
EMAIL_TO="9idli5v6u@mozmail.com"

if [[ ! -f "$PATCH_FILE" ]]; then
  echo "🧨 Error: Patch file '$PATCH_FILE' not found!"
  exit 1
fi

if ! command -v git send-email &> /dev/null; then
  echo "🚫 'git send-email' is not installed or not in your PATH!"
  echo "📦 Try installing it via your package manager, like:"
  echo "    Fedora: sudo dnf install git-email"
  echo "    Debian: sudo apt install git-email"
  echo "    Arch:   sudo pacman -S git"
  exit 1
fi

echo "📤 Sending patch to $EMAIL_TO..."
git send-email --to="$EMAIL_TO" "$PATCH_FILE"

echo "✅ Patch sent to $EMAIL_TO successfully!"
