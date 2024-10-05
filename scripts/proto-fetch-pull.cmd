@echo off

REM Set default branch name
set BRANCH_NAME=master
if not "%1"=="" set BRANCH_NAME=%1

git fetch halal-screen-proto %BRANCH_NAME%
git subtree pull --prefix proto halal-screen-proto %BRANCH_NAME% --squash