#!/bin/sh

set -e

if [ "`git status -s`" ]
then
    echo "The working directory is dirty. Please commit any pending changes."
    exit 1;
fi

echo "Deleting old publication"
rm -rf public
mkdir public
git worktree prune
rm -rf .git/worktrees/public/

echo "Checking out gh-pages branch into public"
git worktree add -B gh-pages public origin/gh-pages

echo "Removing existing files"
rm -rf public/*

echo "Generating site"
hugo

echo "Updating gh-pages branch"
cd public && git add --all && git commit -m "Deploy $(git rev-parse --short HEAD) to gh-pages (deploy.sh)"

echo "Pushing to github"
git push --all

echo "\033[0;32m\xE2\x9C\x94 Site deployed! \033[0m"