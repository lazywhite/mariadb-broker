# MariaDB Broker

This is an implementation of a Service Broker that uses Helm to provision
instances of [MariaDB](https://kubeapps.com/charts/stable/mariadb). This is a
**proof-of-concept** for the [Kubernetes Service
Catalog](https://github.com/kubernetes-sigs/service-catalog), and should not
be used in production.

**NOTE**: Works on Kubernetes v1.18.3 + Service Catalog v0.3.0.

## Prerequisite: An existing MariaDB Server

```
docker run -d --name=mariadb-server -p 3306:3306 -e MYSQL_ROOT_PASSWORD=passw0rd mariadb:10.1.16
```

## For development
```
1. start broker server
sh run_dev_broker.sh
2. register broker(change url first)
kubectl apply -f examples/1.mariadb-broker.yaml
```

## For production
```
1. build docker image
docker login
docker build . -t "<yourrepo>:<tag>"
docker push <yourrepo>:<tag>

2. install broker by helm chart
helm install --name=mariadb-broker charts/mariadb-broker

3. register broker(change url first)
kubectl apply -f examples/1.mariadb-broker.yaml

```


## test broker

```
1. create service instance
kubectl apply -f examples/2.jpress/2.1.jpress-instance.yaml

2. create service binding
kubectl apply -f examples/2.jpress/2.2.jpress-binding.yaml

3. check if service working
kubectl apply -f examples/2.jpress/2.3.jpress-blog-system.yaml
```
