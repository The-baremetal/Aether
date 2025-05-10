#!/bin/bash

set -euo pipefail
IFS=$'\n\t'

if [[ ! -f "rootmaindirAETHER.proj" ]]; then
  echo "🧨 Error: rootmaindirAETHER.proj not found in the current directory!"
  exit 1
fi

current_branch=$(git symbolic-ref --short HEAD)
if [[ "$current_branch" != "main" ]]; then
  echo "🚨 You are currently on branch '$current_branch'."
  read -p "🌱 Do you want to switch to 'main' before continuing? (y/n): " switch_answer
  if [[ "$switch_answer" =~ ^[Yy]$ ]]; then
    git checkout main
  else
    echo "🛑 Please switch to 'main' to generate patch series correctly."
    exit 1
  fi
fi

if ! git diff --quiet || ! git diff --cached --quiet; then
  echo "⚠️ Warning: You have uncommitted changes."
  read -p "💾 Do you want to commit them before proceeding with the rebase? (y/n): " commit_answer
  if [[ "$commit_answer" =~ ^[Yy]$ ]]; then
    echo "📦 Adding all files (including untracked)..."
    git add -A
    read -p "📝 Enter your commit message: " commit_msg
    git commit -m "$commit_msg"
  else
    echo "🤔 You chose not to commit."
    read -p "🧺 Do you want to stash your changes instead? (y/n): " stash_answer
    if [[ "$stash_answer" =~ ^[Yy]$ ]]; then
      echo "📦 Stashing your changes..."
      git stash push -m "Pre-rebase stash"
    else
      echo "🛑 Please commit or stash your changes before proceeding."
      exit 1
    fi
  fi
fi

echo "📡 Fetching latest changes from origin..."
git fetch origin

BASE_COMMIT=$(git merge-base HEAD origin/main)
commits_ahead=$(git log "$BASE_COMMIT"..HEAD --oneline)

if [[ -z "$commits_ahead" ]]; then
  echo "✅ Your branch is up to date with no new commits to patch!"
  exit 0
fi

echo "🔁 Rebasing on origin/main..."
git rebase origin/main

if git stash list | grep -q "Pre-rebase stash"; then
  echo "📥 Applying stashed changes safely (with apply)..."
  git stash apply
  echo "🧹 If all looks good, run 'git stash drop' to remove it."
fi

echo "🧵 Generating patch series..."
git format-patch "$BASE_COMMIT" --cover-letter --stdout > new_patch_series.patch

if [[ ! -s new_patch_series.patch ]]; then
  echo "📭 The patch series is empty. Nothing to save."
  exit 0
fi

echo "🎉 Your patch series is ready: new_patch_series.patch"
