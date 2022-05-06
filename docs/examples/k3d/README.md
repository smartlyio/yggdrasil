Start by creating and ssh'ing to the Vagrant VM
```
vagrant up
vagrant ssh
```

Build yggdrasil and it's docker image, in the repo root mounted at /yggdrasil:
```
sudo su
cd /yggdrasil
docker run -it -w "/app" -v "$(pwd)/:/app" --network bridge golang:1.18.1-buster bash
go get
go mod tidy
make
exit
docker build . -t yggdrasil:latest
```

Go to this examples directory and configure your k3d clusters

```
cd /vagrant
k3d cluster create cluster1 --k3s-arg "--disable=traefik@server:0" --k3s-arg "--disable=servicelb@server:0" --k3s-arg "--cluster-cidr=10.118.0.0/17@server:*" --k3s-arg "--service-cidr=10.118.128.0/17@server:*"
k3d cluster create cluster2 --k3s-arg "--disable=traefik@server:0" --k3s-arg "--disable=servicelb@server:0" --k3s-arg "--cluster-cidr=10.119.0.0/17@server:*" --k3s-arg "--service-cidr=10.119.128.0/17@server:*"

for cluster_name in $(docker network list --format "{{ .Name}}" | grep k3d); do
kubectl config use-context $cluster_name
kubectl config view --minify --raw --output "jsonpath={.clusters.name==\"$cluster_name\"}{..cluster.certificate-authority-data}" | base64 -d > yggdrasil/$cluster_name-ca.crt
kubectl apply -f kube-manifests/yggdrasil.yml
kubectl get secrets -o jsonpath="{.items[?(@.metadata.annotations['kubernetes\.io/service-account\.name']=='yggdrasil-sa')].data.token}"|base64 --decode > config/$cluster_name-token
kubectl apply -f kube-manifests/nginx-ingress-controller.yml
done
```

Run yggdrasil and envoy from the docker-compose.yml and test that envoy is serving example app
```
docker-compose up -d
curl -H host:example.com http://localhost:10000
```
