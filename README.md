# Palindrome Checker
The purpose of this to create a simple Custom Resource Defintion and a Custom operator watching the resource. 

## Description
PalindromeChecker will be a custom Resource that I will create. It has an input field that contains a string. The Kubernetes Cluster will be configured using this yaml manifest. The Controller/Operator keeps an eye on this resource and extracts the input string from the object if a new resource is created. This is how it goes:-  
* When a resource is first created, its status will be empty because nothing has been done yet.  
* Now, if the string is **abcdcba**, I'll compare the first and last letters, and if they're the same, I'll remove them both and change the resource's status to **bcdcb** in the controller which will then be communicated to cluster. Since the status has been updated, Kubernetes will call the controller again, and the first and last are compared in the same way as before. The status is then changed to **cdc**, followed by **d**, and finally **Palindrome**. If any of the above stages fail, the status is changed to **Not Palindrome**. These calls will be made by Kubernetes control manager indefinitely until the object/status resource's does not change. This is the logic I'll be implementing in the custom controller. Based on your need, the CRD definition and controller code can be changed.

## Getting Started
1. Some Key Terms  
    * ### Kubernetes  
        **Kubernetes** is an open-source container orchestration platform that automates deployment, management and scaling of applications.Before we understand what Custom Resource Definition(CRD) is, there are some concepts that you need to get acquainted with.  
        * **Resource** is an endpoint in Kubernetes API that allows you to store an API Object of any kind.
        * A custom Resource allows you to create your own API objects and define their kind just like Pod, Deployment and Replicaset, etc.
        ```yaml
        apiVersion: apiextensions.k8s.io/v1
        kind: CustomResourceDefinition
        metadata:
          # name must match the spec fields below, and be in the form: <plural>.<group>
          name: palindromecheckers.demo.pavan.com
        spec:
          # group name to use for REST API: /apis/<group>/<version>
          group: demo.pavan.com
          # list of versions supported by this CustomResourceDefinition
          versions:
            - name: v1
              # Each version can be enabled/disabled by Served flag.
              served: true
              # One and only one version must be marked as the storage version.
              storage: true
              schema:
                openAPIV3Schema:
                  type: object
                  properties:
                    spec:
                      description: PalindromeCheckerSpec defines the desired state of PalindromeChecker
                      properties:
                        input:
                          type: string
                      type: object
          # either Namespaced or Cluster
          scope: Namespaced
          names:
            # plural name to be used in the URL: /apis/<group>/<version>/<plural>
            plural: palindromecheckers
            # singular name to be used as an alias on the CLI and for display
            singular: palindromechecker
            # kind is normally the CamelCased singular type. Your resource manifests use this.
            kind: PalindromeChecker
            # shortNames allow shorter string to match your resource on the CLI
            shortNames:
            - pc
        ```
        Now, when the above CRD YAML is applied, your Kubernetes cluster will become aware of the Custom Resource PalindromeChecker and knows the structure of the API object to expect from the user of kind PalindromeChecker
    * 

    * ###

You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster
1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

2. Build and push your image to the location specified by `IMG`:
	
```sh
make docker-build docker-push IMG=<some-registry>/palindromechecker:tag
```
	
3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/palindromechecker:tag
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller
UnDeploy the controller to the cluster:

```sh
make undeploy
```

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

### How it works
This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/) 
which provides a reconcile function responsible for synchronizing resources untile the desired state is reached on the cluster 

### Test It Out
1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and

https://user-images.githubusercontent.com/86822039/173196733-d8fa7bad-add4-4418-82fa-9b60abfbd3ca.mp4


limitations under the License.



https://user-images.githubusercontent.com/86822039/173197238-ce4faec1-d1e3-40af-8a0d-451e1ee4aba6.mp4


