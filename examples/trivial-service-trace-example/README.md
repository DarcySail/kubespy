# Example: What happens when a `Service` starts?

In this example, we:

-   Install a `Deployment` replicating an `nginx` `Pod` 3 times in a Kubernetes cluster.
-   Install a `Service` exposing those `nginx` `Pod`s to the Internet.
-   Use `kubespy trace` to watch what happens to that `Pod` as it starts.

You'll need:

-   **Access to a Kubernetes cluster.** If you are using Pulumi, you can
    trivially boot an
    [GKE](https://github.com/pulumi/examples/tree/master/gcp-ts-gke),
    [AKS](https://github.com/pulumi/examples/tree/master/azure-ts-aks-mean), or
    [EKS](https://github.com/pulumi/examples/tree/master/aws-ts-eks) cluster.
    **NOTE:** If you use [minikube](https://github.com/kubernetes/minikube),
    you'll need to change the `Service`'s type to be `ClusterIP`, as minikube
    does not support type `LoadBalancer`.
-   **Either the Pulumi CLI or `kubectl`.** There is nothing Pulumi-specific in
    this example, so you can use `kubectl`, but we hope you'll give Pulumi a
    shot! The CLI installation instructions
    [here](https://pulumi.io/quickstart/install.html). Pulumi works anywhere
    `kubectl` works (i.e., anywhere you have a kubeconfig file), so it should
    "just work" if you already have a Kubernetes cluster running.
-   **`kubespy`.** Installation is a handful of commands, which you can find in the [README](https://github.com/DarcySail/kubespy#installation).

Once these are complete, you'll want to do two things:

1. **Run `kubespy`.** Once you've installed it, this should be as simple as
   `kubespy status v1 Pod nginx`. `kubespy` will dutifully wait for you to
   deploy a `Pod` called `nginx` to your cluster.
2. **Run the example.** `kubespy` repository contains a tiny example Pod that
   deploys an NGINX container.

    ```sh
    # With Pulumi CLI.
    $ git clone git@github.com:DarcySail/kubespy.git
    $ cd kubespy/examples/trivial-service-trace-example
    $ npm install
    $ pulumi up

    # With kubectl.
    kubectl create -f https://github.com/DarcySail/kubespy/raw/master/examples/trivial-service-trace-example/yaml/nginx.yaml
    ```

Once done, running `kubespy trace service nginx` should display something like this:

![Changes](../../images/trace/trace-success.gif "Changes a Service undergoes as it starts, in real time")
