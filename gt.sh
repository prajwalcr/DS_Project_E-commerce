#!/bin/bash                                                                     

# Run this script with an argument which will be the commit message             
git pull
git add .
git commit -m "$1"
git push
