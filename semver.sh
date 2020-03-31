#!/bin/bash

# Increment a version string using Semantic Versioning (SemVer) terminology.

# Parse command line options.

while getopts ":Mmp" Option
do
  case $Option in
    M ) major=true;;
    m ) minor=true;;
    p ) patch=true;;
  esac
done

shift $(($OPTIND - 1))

git pull --tags
version=`git tag | sort -V | tail -1`
echo "version found - $version"

# Build array from version string.

a=( ${version//./ } )

# If version string is missing or has the wrong number of members, show usage message.

if [ ${#a[@]} -ne 3 ]
then
  echo "usage: $(basename $0) [-Mmp] major.minor.patch"
  exit 1
fi

# Increment version numbers as requested.

if [ ! -z $major ]
then
  ((a[0]++))
  a[1]=0
  a[2]=0
fi

if [ ! -z $minor ]
then
  ((a[1]++))
  a[2]=0
fi

if [ ! -z $patch ]
then
  ((a[2]++))
fi

git tag "${a[0]}.${a[1]}.${a[2]}"
git push --tags
echo "${a[0]}.${a[1]}.${a[2]}"
