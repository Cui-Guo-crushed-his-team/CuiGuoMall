#!/bin/bash
find . -name "go.mod" -execdir go mod tidy \;
