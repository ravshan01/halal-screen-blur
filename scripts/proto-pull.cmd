@echo off

set BRANCH_NAME=master
if not "%1"=="" set BRANCH_NAME=%1

REM git subtree pull ... is working only from the root of the repository
REM so we need to change the directory to the root of the repository
REM and then back to the original directory
set DIR=%CD%
cd ..

git subtree pull --prefix proto halal-screen-proto %BRANCH_NAME% --squash

cd %DIR%
