#!/bin/bash -e

# cargo install mdbook
mdbook serve

# kill -9 $(lsof -t -i:3000)
