#!/bin/sh

echo start to deploy

cd service_a
docker build -t iamges-address .
docker iamges-address
kubectl delete -f my-tracing.yaml
kubectl apply -f my-tracing.yaml
cd ..

cd service_b
docker build -t amges-address .
docker push amges-address
kubectl delete -f my-tracing.yaml
kubectl apply -f my-tracing.yaml
cd ..

cd service_c
docker build -t amges-address .
docker push amges-address
kubectl delete -f my-tracing.yaml
kubectl apply -f my-tracing.yaml
cd ..

cd service_d
docker build -t amges-address .
docker push amges-address
kubectl delete -f my-tracing.yaml
kubectl apply -f my-tracing.yaml
cd ..
