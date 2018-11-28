`kubectl create -f kubernetes/network-example.yaml`

install insturctions:

https://cilium.readthedocs.io/en/latest/gettingstarted/minikube/#gs-minikube

install minikube
$ some instruction
start minikube
$ minikube start --network-plugin=cni --extra-config=kubelet.network-plugin=cni --memory=5120
Install etcd as a dependency of cilium in minikube by running:

$ kubectl create -n kube-system -f https://raw.githubusercontent.com/cilium/cilium/HEAD/examples/kubernetes/addons/etcd/standalone-etcd.yaml
$ kubectl create -f https://raw.githubusercontent.com/cilium/cilium/HEAD/examples/kubernetes/1.8/cilium.yaml
check deployment
$ kubectl get daemonsets -n kube-system
$ kubectl create -f kubernetes/app-a-app-b.yaml
check access
$ kubectl exec microservice-a -- curl -s http://app-a -m 1
$ kubectl exec microservice-b -- curl -s http://app-a -m 1

Apply L4 policy
$ kubectl create -f kubernetes/l4-network-policy.yaml

attempt to call b from a

kubectl -n default delete po,svc --all                                      # Delete all pods and services, including uninitialized ones, in namespace default,
