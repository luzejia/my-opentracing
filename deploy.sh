#!/bin/sh

echo start to deploy

# platform根据运行的机器的docker平台而定，执行docker version可以查看运行的平台

cd service_a
docker build -t images-address . --platform linux/amd64
docker push images-address
kubectl delete -f my-tracing.yaml
kubectl apply -f my-tracing.yaml
cd ..

cd service_b
docker build -t images-address . --platform linux/amd64
docker push images-address
kubectl delete -f my-tracing.yaml
kubectl apply -f my-tracing.yaml
cd ..

cd service_c
docker build -t images-address . --platform linux/amd64
docker push images-address
kubectl delete -f my-tracing.yaml
kubectl apply -f my-tracing.yaml
cd ..

cd service_d
docker build -t images-address . --platform linux/amd64
docker push images-address
kubectl delete -f my-tracing.yaml
kubectl apply -f my-tracing.yaml
cd ..
