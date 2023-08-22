#!/bin/bash

# Go uygulamasını Docker imajına build et
docker build --rm -t meeting-app .

# Mevcut Docker imajlarını listele
docker images

# Docker konteynerını çalıştır
docker run -e PORT=3000 -p 3000:3000 --name=meeting-app meeting-app

# Minikube durumunu kontrol et
minikube status

# Minikube Docker ortamını yapılandır
eval $(minikube docker-env)

# Go uygulamasını tekrar Docker imajına build et (Minikube ortamında)
docker build --rm -t meeting-app .

# Uygulamayı Kubernetes üzerine dağıt
kubectl apply -f meeting-app.yaml

# Pod durumunu kontrol et
kubectl get pods

# Deployment durumunu kontrol et
kubectl get deployment meeting-app

# Servis oluştur
kubectl expose deployment meeting-app --type=LoadBalancer --name=meeting-app-service

# Servis hakkında detaylı bilgi al
kubectl describe services meeting-app-service

# Servisi web tarayıcıda çalıştır
minikube service meeting-app-service