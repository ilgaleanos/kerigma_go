#! /bin/bash

# gcloud init
# docker run -d --restart always -p 8080:8080 gcr.io/kerigma/back/kerigma:v1.0.2

read -p "Version? : " VERSION_KDD

# gcloud auth configure-docker
export RUTA_KDD="fasthttp.dockerfile"
export ZONA_KDD="us-east1-d"
export NOMBRE_KDD="backend"
export CLUSTER_KDD="cluster-kerigma"
export PROJECT_KDD="kerigma"

docker build -t "gcr.io/${PROJECT_KDD}/${NOMBRE_KDD}:${VERSION_KDD}" -f "${RUTA_KDD}" .
docker push gcr.io/${PROJECT_KDD}/${NOMBRE_KDD}
gcloud container clusters get-credentials --zone ${ZONA_KDD} ${CLUSTER_KDD}
kubectl set image "deployment/${NOMBRE_KDD}" "${NOMBRE_KDD}=gcr.io/${PROJECT_KDD}/${NOMBRE_KDD}:${VERSION_KDD}"
