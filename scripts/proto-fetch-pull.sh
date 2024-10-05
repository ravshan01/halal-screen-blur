#!/bin/bash

# Set default branch name
BRANCH_NAME=${1:-master}

git fetch halal-screen-proto $BRANCH_NAME
git subtree pull --prefix proto halal-screen-proto $BRANCH_NAME --squash