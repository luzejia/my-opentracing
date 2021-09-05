#!/bin/sh

echo start to deploy

cd service_a
docker build -t uhub.service.ucloud.cn/luzejia/tracing-a:v0.0.12 .
docker push uhub.service.ucloud.cn/luzejia/tracing-a:v0.0.12
kubectl delete -f my-tracing.yaml
kubectl apply -f my-tracing.yaml
cd ..

cd service_b
docker build -t uhub.service.ucloud.cn/luzejia/tracing-b:v0.0.12 .
docker push uhub.service.ucloud.cn/luzejia/tracing-b:v0.0.12
kubectl delete -f my-tracing.yaml
kubectl apply -f my-tracing.yaml
cd ..

cd service_c
docker build -t uhub.service.ucloud.cn/luzejia/tracing-c:v0.0.12 .
docker push uhub.service.ucloud.cn/luzejia/tracing-c:v0.0.12
kubectl delete -f my-tracing.yaml
kubectl apply -f my-tracing.yaml
cd ..
