# Palindrome Checker
The purpose of this project is to create a simple Custom Resource Defintion and a Custom operator watching the resource. 

## Description
PalindromeChecker will be a custom Resource that I will create. It has an input field that contains a string. The Kubernetes Cluster will be configured using this yaml manifest. The Controller/Operator keeps an eye on this resource and extracts the input string from the object if a new resource is created. This is how it goes:-  
* When a resource is first created, its status will be empty because nothing has been done yet.  
* Now, if the string is **abcdcba**, I'll compare the first and last letters, and if they're the same, I'll remove them both and change the resource's status to **bcdcb** in the controller which will then be communicated to cluster. Since the status has been updated, Kubernetes will call the controller again, and the first and last are compared in the same way as before. The status is then changed to **cdc**, followed by **d**, and finally **Palindrome**. If any of the above stages fail, the status is changed to **Not Palindrome**. These calls will be made by Kubernetes control manager indefinitely until the resource's object/status  does not change. This is the logic I'll be implementing in the custom controller. Based on your need, the CRD definition and controller code can be changed.

## Getting Started
1. **Some Key Terms**  
    * **Kubernetes**  
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
        &emsp; After applying the above CRD YAML, your Kubernetes cluster will be aware of the Custom Resource PalindromeChecker and will know what structure of API object to expect from the user of kind PalindromeChecker.  

    * ### Controllers
        &emsp; In Kubernetes, Controllers are control loops that monitor the current state of your cluster and make the necessary changes to move it to the intended state, as specified in the API Object's spec file.

    * ### Operators
        &emsp; Controllers keep an eye on the native Kubernetes resources. However, we'll need an operator to keep an eye on our custom Resources, and Kubernetes does allow us to construct a custom operator to keep a watch on a custom resource. In the next sections, I may use the phrases controller and operator interchangeably.

2. Prequisites for running your custom resource and controllers.
    * **[Golang](https://golang.org/doc/install)**, the programming language that our definitions and controllers are written in.
    * **[kubectl](https://kubernetes.io/docs/tasks/tools/)** , the Kubernetes command line tool for running commands against Kubernetes Cluster.  
    * **[kind](https://kind.sigs.k8s.io/docs/user/quick-start/)**, **[minikube](https://minikube.sigs.k8s.io/docs/start/)** or **[kubeadm](https://kubernetes.io/docs/tasks/tools/#kubeadm)** for creating and managing Kubernetes cluster.  Type `kubectl version` to verify that everything is working properly. If the installation went well, you should obtain both the client and server versions.
    * **[Kubebuilder](https://book.kubebuilder.io/quick-start.html#:~:text=Installation%20Install%20kubebuilder%3A%20%23%20download%20kubebuilder%20and%20install,%2Bx%20kubebuilder%20%26%26%20mv%20kubebuilder%20%2Fusr%2F%20local%20%2Fbin%2F)** or similar SDKs for helping use quicken the process of writing operators. But why do we need this? Watching, creating, and listing of K8s API objects need libraries like **[client-go](https://github.com/kubernetes/client-go)** and **[controller-runtime](https://github.com/kubernetes-sigs/controller-runtime).** However, because we're using  custom Resources, the aforementioned operations require a lot of code. Kubebuilder and other SDKs relieve this tension by bootstrapping the code for us, allowing us to edit only the files that are required. Make sure you have the compatible version of Go installed. For Kubebuilder, I'm using 3.4.1, and for Go, I'm using 1.17.7. Kubebuilder also requires a linux environment. Install WSL (Windows Subsystem for Linux) and make sure GOOS = "linux" if you're using Windows.  
  
3. Modifying the API definitions
    *  If you are editing the API definitions, generate the manifests such as CRs or CRDs using:
        ```sh
        make manifests
       ```
  
4. Test It Out
    * Install the CRDs into the cluster:
        ```sh
        make install
        ```

    * Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):
        ```sh
        make run
        ```  

5. Running on the cluster
    * Install Instances of Custom Resources into our cluster:
        ```sh
        kubectl apply -f config/samples/
        ```
    
    * Build and push your image to the location specified by `IMG`:
        ```sh
        make docker-build docker-push IMG=<some-registry>/palindromechecker:tag
        ```

    * Deploy the controller to the cluster with the image specified by `IMG`:
        ```sh
        make deploy IMG=<some-registry>/palindromechecker:tag
        ```
6. Uninstall CRDs
    * To delete the CRDs from the cluster:
        ```sh
        make uninstall
        ```
          
7. Undeploy controller
   * UnDeploy the controller to the cluster:
        ```sh
        make undeploy
        ```

## Coding CRD
* Create a new folder for our project. I will call it PalindromeChecker.  
    ```sh 
    mkdir -p PalindromeChecker
    cd PalindromeChecker
    ```  
  
* Bootstrap the code for the project by running the following.  
    ```sh
    kubebuilder init --domain pavan.com --repo github.com/pavanneerudu/PalindromeChecker --licence apache2 --skip-go-version-check
    ```  
    **NOTE:** If your project is initialized within **[GOPATH](https://go.dev/doc/code#GOPATH)**, the implicitly called go mod init will interpolate the module path for you. Otherwise `--repo=<module path>` must be set.

* Create a new **API(group/version)** as demo/v1 and a new kind(CRD) PalindromeChecker on it with the following command. Enter y for Create Resource and Create Controller, and then demo/v1/palindromechecker types.go and controllers/palindromechecker controller will be created for us.   
    ```sh
    kubebuilder create api --group demo --version v1 --kind PalindromeChecker --namespaced true
    ```  
    You should now have some files in the folder after running the aforementioned instructions. We're concerned about the `demo/v1/palindromechecker_types.go` and `controllers/palindromechecker_controller.go` files, despite the fact that there are a lot of them.  
  
* Here's a quick rundown of our project's files and directories.  
    * `api/v1/*.go` contains files related to all the resources.  
    * `bin/*` contains binaries such as kustomize and controller-gen.
    * `config/*` has YAML files for various actions. These include the YAML manifest of our custom resource in `./bases/` . This file when applied lets our Kubernetes cluster to know about our PlaindromeChecker. Also, there is a sample yaml file in `./samples` which can be used to create a resource in our cluster.  
    * `controllers/*` contains the source code of our operator.  
    * `main.go` is the entry point of our project. Here the **Manager** is started, and every operator is instantiated and set up with manager (method defined in controller) and then manager is started. Manager wraps up the controllers and registers them with a Kubernetes Cluster. 
    * `Makefile` contains the make targets which build and deploy our operator.  To know about different make targets, go through Getting Started section.  
    * `Dockerfile` has docker scripts for packing our manager into a docker image which can be later deployed into a Kubernetes cluster.  
  
* #### Change `api/v1/palindromechecker_types.go` file
    This file was already created for us by the kubebuilder. The PlaindromeChecker struct is defined as
    ```go
    //+kubebuilder:object:root=true
    //+kubebuilder:printcolumn:name="Initial_Input",type=string,JSONPath=`.spec.input`
    //+kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.status`
    //+kubebuilder:subresource:status
    //+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
    //PalindromeChecker is the Schema for the palindromecheckers API
    type PalindromeChecker struct {
    	metav1.TypeMeta   `json:",inline"`
    	metav1.ObjectMeta `json:"metadata,omitempty"`

    	Spec   PalindromeCheckerSpec   `json:"spec,omitempty"`
    	Status PalindromeCheckerStatus `json:"status,omitempty"`
    }
    ```
    * `metav1.TypeMeta` is a struct exposing the type and APIVersion of API Object. APIVersion of PalindromeChecker is `demo.pavan.com/v1` . Here demo.pavan.com is Group and v1 is the version. Also, the kind is `PalindromeChecker`. These are collectively called **GVK (Group-Version-Kind)**. Every resource in Kubernetes can be identified by a unique combination of GVK.  
    * `metav1.ObjectMeta` lets you work with object metadata of the API Object. It has many fields. Name, Namespace, CreationTimeStamp, and DeletionTimeStamp are a few examples.
    * `Spec` of type `PalindromeCheckerSpec` struct defines the spec of our Object. Here we define the spec of our Resource. Let’s go ahead and do that.
        ```go
        type PalindromeCheckerSpec struct {
	        Input string `json:"input,omitempty"`
        }
        ```  
    * `Status` of type `PalindromeCheckerStatus` struct defines the observed status of Palindrome Checker. It is with this value that we know if any change in the status has been seen.  
        ```go
        type PalindromeCheckerStatus struct {
	        Status string `json:"status,omitempty"`
        }
        ```
        Initially, the status of our resource is “” (empty).
    * Above the PalindromeCheckerSpec struct, I'd like to mention kubebuilder **validation markers**. This allows us to display additional columns in the resource's command line. The third statement, for example, allows us to print the input under the Initial_Input column. Also worth mentioning is the use of `//+kubebuilder:subresource:status` to enable the status subresource. Updates to the main resource's status will not change if this option is enabled. Updates to the status subresource, meanwhile, can only affect the status field.
    * `PalindromeCheckerList` is another significant type in this file. We can establish many PlaindromeChecker resources in the Kubernetes cluster, each with its own state. This List contains the PalindromeChecker resource list.  

* Now that we have defined out PalindromeChecker, let’s go and apply our CRD YAML manifiest.  
        Before, applying the YAML manifest, PalindromeChecker is not recognised by k8s cluster.
        <p align="center"><img src="https://user-images.githubusercontent.com/86822039/173192043-9327f29c-8067-4084-aeb7-733e250cab83.png"></p>
        Applying YAML manifets after genearating manifest by running the following command  
        ```ssh
        make manifests;make install
        ```
        <p align="center"><img src="https://user-images.githubusercontent.com/86822039/173192208-aa083dab-53c9-4417-a90c-25e1c9fd1bca.png"></p>Now, k8s cluster can recongnise PalindromeChecker.

* Let’s create a resource and apply it into our cluster. There is a sample YAML file in config/sample. I am editing it in accordance with our definition	and applying it to our cluster. The yaml file is 
    ```go
    apiVersion: demo.pavan.com/v1
    kind: PalindromeChecker
    metadata:
      name: palindrome
    spec:
      input: abcdedcba
    ```
    Later, run `kubectl get PalindromeChecker` to get the resources
    <p align="center"><img src="https://user-images.githubusercontent.com/86822039/173192426-aa98cb88-2c26-43e4-b532-479a8d62f17d.png"></p>
    The Status field is empty, and it is not changing, but how to make the status of our resource change? 
    Enter <b>controllers!!!</b>  

## Coding Controller  
* Let’s take a look at PalindromeCheckerReconciler struct.
    ```go
    type PalindromeCheckerReconciler struct {
    	client.Client
    	Scheme *runtime.Scheme
    }
    ```
    * Our controller uses the `client.Client` attribute from container-runtime to communicate with the K8s cluster and perform any operations on it. Creating, modifying, and deleting a resource are examples of these actions. We don't create any new functions; instead, we use the ones that have already been defined in controller-runtime. 
    * The PalindromeChecker struct is registered to its relevant GVK via manager using `Scheme`. Various clients use this feature largely behind the scenes to correlate and connect go types/structs with their GVKs.  
    As previously stated, each operator/controller after being defined is set up with manager using SetupWithManager.
    ```go
    func (r *PalindromeCheckerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	    return ctrl.NewControllerManagedBy(mgr).
		    For(&demov1.PalindromeChecker{}).
		    Complete(r)
    }
    ```
    &nbsp;&nbsp;&nbsp;&nbsp; For any change related to PalindromeChecker `(For(&demov1.PalindromeChecker{}))`, call Reconcile `(Complete(r))`.  

* Now that we've learned about the attributes, let's create the Reconciliation Loops logic in the Reconcile() method.
    ```go
    func (r *PalindromeCheckerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	    logger := log.FromContext(ctx)
	    logger.Info("Entered PalindromeChecker Controller")
    ```  
    * `log.FromContext` defined in `controller-runtime/pkg` returns a logger with predefined values from a context.Context
    * `logger.Info` logs a non-error message with the given key/value pairs as context. Here, we logged a simple message for us to see it the logs.  
       
    It’s time to get our PalindromeCheckerObject.  
    ```go  
	    PalindromeCheckerObject := &demov1.PalindromeChecker{}  
	    err := r.Get(ctx, req.NamespacedName, PalindromeCheckerObject)
    ```
    &emsp; Here, we initialize PalindromeCheckerObject and obtain our resource by calling r. Get(). This method implements client.Client behind the scenes and returns our object.

    Since the Reconcile() is called for even updation and deletion, we check if our object was deleted in the previous step and if that is the case, we return nil as error, so that Reconcile() will not be called again.
    ```go  
    	if err != nil {  
	    //Resource was not found.  
	    if errors.IsNotFound(err) {  
	    return ctrl.Result{}, nil  
	    }  
	    //Other errors  
	    logger.Error(err, "Error looking up for PalindromeChecker Object")  
	    return ctrl.Result{}, err  
	}
    ```  
    We then use PalindromeChecherObject.Status to get the status of our object.  We take different steps depending on the status. This is where the true logic lies. Depending on your functionality, you can alter your logic.
    ```go  
    	    status := PalindromeCheckerObject.Status.Status

	    if status == "" {
	    	PalindromeCheckerObject.Status.Status = PalindromeCheckerObject.Spec.Input
	    	logger.Info(fmt.Sprintf("Current Status: %s", PalindromeCheckerObject.Status.Status))
	    } else {
	    	if status == "Palindrome" || status == "Not Palindrome" {
	    		logger.Info("Palindrome Checker process done!!")
	    		return ctrl.Result{}, nil //No more updation to etcd
	    	} else {
	    		if len(status) == 0 || len(status) == 1 {
	    			PalindromeCheckerObject.Status.Status = "Palindrome"
	    			logger.Info(fmt.Sprintf("%s is a Palindrome", PalindromeCheckerObject.Spec.Input))
	    		} else {
	    			time.Sleep(6 * time.Second)
	    			len := len(status)
	    			if status[0] == status[len-1] {
	    				PalindromeCheckerObject.Status.Status = status[1 : len-1]
	    				logger.Info(fmt.Sprintf("Current Status: %s", PalindromeCheckerObject.Status.Status))
	    			} else {
	    				PalindromeCheckerObject.Status.Status = "Not Palindrome"
	    				logger.Info(fmt.Sprintf("%s is not Palindrome", PalindromeCheckerObject.Spec.Input))
	    			}
	    		}
	    	}
	    }
    ```
    * if status is empty, that means, the object was just created now. So, we copy the input from `Spec.Input` to `Status.Status`
    * The object's final status would be Palidrome or Non-Palindrome. The reconcile() function will not be called after that.  
    * The logic's next steps are self-explanatory. I am directing the programme to **sleep** for 6 seconds every time the initial and last letters are the same. That's because Reconcile() will be called so quickly that we won't be able to see the change in resource state.<br></br>  

    The last step in the code is to update the resource's status. Isn't that what we've already done? We've changed it in our code, but the <u>Kubernetes cluster is unaware of it</u>. The following code snippet can be used to accomplish this.
    ```go
        err = r.Status().Update(ctx, PalindromeCheckerObject)
	    if err != nil {
	       	logger.Error(err, "Error updating the status of PalindromeCheckerObject")
	       	return ctrl.Result{}, err
    	}
	    return ctrl.Result{}, nil
    }
    ```
* Now that we've finished developing the controller, it's time to put it to use. Run the controller by executing the following command
    ```ssh
    make run
    ```
* While the controller is run, I will watch the PalindromeChecker resource in one terminal and apply the sample YAML file from other terminal. The following video shows the same 

https://user-images.githubusercontent.com/86822039/173200082-edb98c1c-01ad-40a3-9041-1e8f02761701.mp4

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
limitations under the License.

