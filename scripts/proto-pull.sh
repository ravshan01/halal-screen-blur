#!/bin/bash

# Set default branch name
BRANCH_NAME=${1:-master}

git subtree pull --prefix proto halal-screen-proto $BRANCH_NAME --squash