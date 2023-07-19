> ⚠️ The kubefirst exploration applications metaphor has been merged into one version, and moved inside the [gitops-teplate repository](https://github.com/kubefirst/gitops-template/tree/main/metaphor).

# Metaphor GO API

`metaphor-go` is an example application which serves as a demonstration space for how your applications hook into your 
infrastructure and tooling.

`metaphor-go` is deployed to all of your environments just like your other applications. This means that when you make 
changes to your ci/cd, you can test it out using an application that works just like your applications do.

`metaphor-go` currently demonstrates the following capabilities:

- building a docker container  
- publishing a docker container to ecr  
- publishing a prerelease helm chart  
- gitops delivery of metaphor-go to different namespaces  
- a release process that publishes a release chart and patches the chart version for the next release  
- secrets sourced from vault  
- certificate management using cert-manager  
- automatic dns management using external-dns  

# CI/CD

`metaphor-go` has 5 ci stages defined in its .gitlab-ci.yml file.

- branch: branch jobs run when your branch pushes to origin and report status to your merge requests
- publish: publishes your docker container to ecr and publishes your prerelease chart to chartmuseum
- development: sets the desired chart version for development
- staging: sets the desired chart version for staging
- release: publishes a release chart, sets the desired chart version for production, and patches chart in source for the next release

`argocd` is the gitops tool responsible for autosyncing the desired state in each environment. It follows a pull model so our CI/CD ecosystem doesn't need to know how to connect to our kubernetes clusters.

We have `metaphor-go` set up to run its automation by invoking argo-workflows. Those submitted workflows are in the .argo directory of this repository.
