#!/bin/bash

set -e

NAME="gokc"

ROOT=$(dirname $0)/..

goxc -tasks='xc archive' -bc 'linux,!arm windows darwin' -d .
cp -p "$ROOT"/snapshot/linux_amd64/"$NAME" "$ROOT"/snapshot/"$NAME"_linux_amd64
cp -p "$ROOT"/snapshot/linux_386/"$NAME" "$ROOT"/snapshot/"$NAME"_linux_386
cp -p "$ROOT"/snapshot/darwin_amd64/"$NAME" "$ROOT"/snapshot/"$NAME"_darwin_amd64
cp -p "$ROOT"/snapshot/darwin_386/"$NAME" "$ROOT"/snapshot/"$NAME"_darwin_386
cp -p "$ROOT"/snapshot/windows_amd64/"$NAME".exe "$ROOT"/snapshot/"$NAME"_windows_amd64.exe
cp -p "$ROOT"/snapshot/windows_386/"$NAME".exe "$ROOT"/snapshot/"$NAME"_windows_386.exe
