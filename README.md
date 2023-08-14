# yueyuea

## Prepare In Advance
You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) or [K3D](https://k3d.io/v5.5.2) to get a local cluster for testing, or run against a remote cluster.

Tips: 
- Recommended to use k3d to quickly build k8s cluster (https://github.com/yueyueaa/k8s-setup/blob/master/k3d/README.md) `bash k3d.sh`
- If you want to use webhook, you need install cert-manager to your cluster (https://cert-manager.io/docs/installation)

## Getting Started
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster
1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

2. Build and push your image to the location specified by `IMG`:

```sh
make docker-build docker-push IMG=<some-registry>/yueyuea:tag
```

Tips: 
- If you use `bash k3d.sh` create cluster, your registry is `localhost:5000`
- If you can't use `make docker-build` or `make docker-push`, you can try `make ko-image` to quickly build image (https://ko.build/)

**NOTE:** Run `make --help` for more information on all potential `make` targets

3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/yueyuea:tag
```

Tips:
- To make it easier to package and deploy controller/operator to the cluster, you can use `make all`, then you can apply yaml to cluster

### Uninstall CRDs Or Undeploy controller 
To delete the CRDs from the cluster:

```sh
make uninstall  # delete the CRDs from the cluster
make undeploy   # delete the controller from the cluster
```

### How it works
This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/).

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/),
which provide a reconcile function responsible for synchronizing resources until the desired state is reached on the cluster.

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

