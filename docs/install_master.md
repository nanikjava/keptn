# Install master branch

**Attention:** This installs an unstable and unreliable version of Keptn. Do not use this unless you are a core developer. 

**DO NOT PERFORM THIS ON A PRODUCTION ENVIRONMENT**

1. Create a new cluster (e.g., using k3s)
1. Download latest CLI: https://storage.cloud.google.com/keptn-cli/latest/keptn-linux.zip
1. Install keptn using `keptn install --chart-repo=https://storage.googleapis.com/keptn-installer/latest/keptn-0.1.0.tgz --use-case=continuous-delivery`


## Warning

Depending on your cluster configuration and the used Keptn Version you might end up with a state that's not the current master branch.

We use `imagePullPolicy: IfNotPresent` and as we install images with the `:latest` tag, you might end up with and old version of Keptn on your Kubernetes cluster (in the Docker cache). 

You can verify this by looking at the log output of one of the services, which should print a datetime header on when it was installed, e.g.:
```
$> kubectl -n keptn logs deployments/api-service api-service

##########
branch: release-0.7.2
repository: https://github.com/keptn/keptn
commitlink: https://github.com/keptn/keptn/commit/d469876ecfe95d467872921ae6e9ce877f2ccca6
repolink: https://github.com/keptn/keptn/tree/d469876ecfe95d467872921ae6e9ce877f2ccca6
travisbuild: https://travis-ci.org/keptn/keptn/jobs/735015708
timestamp: 20201012.1651
##########

```

If this is the case for you, you can manually edit the deployments and set image pull policy to always, or alternatively, create a new Kubernetes cluster (which should clear the Docker cache).
