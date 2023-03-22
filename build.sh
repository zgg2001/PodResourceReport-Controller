#!/bin/bash

username='zgg2001'
tag='reporter:1.0'

# docker build
cp $HOME/.kube/config kube_config
docker build -t $username/$tag .
#docker push $username/$tag
rm kube_config