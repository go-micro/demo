# NetworkPolicy + OnlineBoutique

This guide contains instructions for configuring `NetworkPolicy` resources for the OnlineBoutique apps by leveraging the [GKE Dataplane V2](https://cloud.google.com/kubernetes-engine/docs/concepts/dataplane-v2) feature.

## Steps
 
1. You need to have a GKE cluster created with [GKE Dataplane V2 enabled](https://cloud.google.com/kubernetes-engine/docs/how-to/dataplane-v2).

As an example, here is the option you need to use when creating your cluster:
```sh
gcloud container clusters create \
    --enable-dataplane-v2
```

2. Apply all the manifests.

After you applied the Kubernetes manifests of the Online Boutique sample apps, you could apply the `NetworkPolicy` resources:
```sh
kubectl apply -f ./docs/network-policies/
```

3. Verify the resources deployed.

If you run:
```
kubectl get networkpolicy
```

You should see:
```
NAME                    POD-SELECTOR                AGE
ad               app=ad               2m58s
cart             app=cart             2m58s
checkout         app=checkout         2m58s
currency         app=currency         2m58s
deny-all                <none>                      2m58s
email            app=email            2m58s
frontend                app=frontend                2m58s
loadgenerator           app=loadgenerator           2m58s
payment          app=payment          2m58s
productcatalog   app=productcatalog   2m58s
recommendation   app=recommendation   2m58s
shipping         app=shipping         2m58s
```

_Note: `Egress` is wide open in these `NetworkPolicy` resources. In our case, we do this is on purpose because there are multiple egress destinations to take into consideration like the Kubernetes DNS, Istio control plane (`istiod`), Cloud Trace API, Cloud Profiler API, Cloud Debugger API, etc._

4. Verify you could still access the frontend's `EXTERNAL_IP` with no issues.

```
kubectl get service frontend-external | awk '{print $4}'
```

## Resources

- [GKE Dataplane V2 announcement](https://cloud.google.com/blog/products/containers-kubernetes/bringing-ebpf-and-cilium-to-google-kubernetes-engine)
- [Kubernetes Network Policies](https://kubernetes.io/docs/concepts/services-networking/network-policies/)
- [Kubernetes Network Policy Recipes](https://github.com/ahmetb/kubernetes-network-policy-recipes)
- [Network policy logging](https://cloud.google.com/kubernetes-engine/docs/how-to/network-policy-logging)