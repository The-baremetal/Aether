#!/bin/bash

set -euo pipefail
IFS=$'\n\t'

REPO_DIR=$(pwd)
CURRENT_VERSION_FILE="currentversion.vortex"
TAG_PREFIX="v"

update_version_file() {
  echo "Updating version in $CURRENT_VERSION_FILE..."
  echo "Vortex-${CORE}.${FLOW}.${PATCH}_${TIMESTAMP}${CUSTOM_TAGS}" > "$CURRENT_VERSION_FILE"
}

ask_for_update_type() {
  echo "What type of update is this?"
  echo "1. Patch (Increment Patch)"
  echo "2. Flow (Increment Flow, reset Patch to 0)"
  echo "3. Core (Increment Core, reset Flow and Patch to 0)"
  echo "Enter the number of the update type (1, 2, 3):"
  read update_type

  case $update_type in
    1)
      PATCH=$((PATCH + 1))
      ;;
    2)
      FLOW=$((FLOW + 1))
      PATCH=0
      ;;
    3)
      CORE=$((CORE + 1))
      FLOW=0
      PATCH=0
      ;;
    *)
      echo "Invalid selection, exiting."
      exit 1
      ;;
  esac
}

create_tag() {
  git tag -a "$VERSION" -m "Release version ${VERSION}"
  echo "New tag created: $VERSION"
}
create_release() {
  echo "Do you want to mark this as a prerelease? (y/n)"
  read is_prerelease
  PRERELEASE_FLAG=""
  if [ "$is_prerelease" == "y" ]; then
    PRERELEASE_FLAG="--prerelease"
  fi
  gh release create "$VERSION" --title "Release ${VERSION}" --notes "Release version ${VERSION}" $PRERELEASE_FLAG
  echo "Release created for version ${VERSION}"
}
echo "Pulling the latest changes from the repository..."
git pull origin main || { echo "Failed to pull from repository"; exit 1; }

if [ -f "$CURRENT_VERSION_FILE" ]; then
  VERSION_STRING=$(cat "$CURRENT_VERSION_FILE")
  CORE=$(echo $VERSION_STRING | cut -d '.' -f 1 | sed 's/Vortex-//')
  FLOW=$(echo $VERSION_STRING | cut -d '.' -f 2)
  PATCH=$(echo $VERSION_STRING | cut -d '.' -f 3 | cut -d '_' -f 1)
  TIMESTAMP=$(date +%s%N)
else
  echo "No current version file found, starting from Vortex-0.0.0."
  CORE=0
  FLOW=0
  PATCH=0
  TIMESTAMP=$(date +%s%N)
fi

ask_for_update_type

echo "Do you want to add any custom tags (e.g., _alpha, _beta)? (y/n)"
read add_custom_tag
CUSTOM_TAGS=""
if [ "$add_custom_tag" == "y" ]; then
  echo "Enter the custom tag (e.g., _alpha, _beta):"
  read CUSTOM_TAGS
fi

VERSION="Vortex-${CORE}.${FLOW}.${PATCH}_${TIMESTAMP}${CUSTOM_TAGS}"
echo "Generated version: $VERSION"

update_version_file

git add "$CURRENT_VERSION_FILE"
git commit -m "Updated version to $VERSION"
echo "Committing changes..."
git push origin main || { echo "Failed to push changes to repository"; exit 1; }

create_tag

git push origin "$VERSION" || { echo "Failed to push tag to repository"; exit 1; }

create_release

echo "Sync complete!"
