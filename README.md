# Azure Terratest sample

An variation of [Test Terraform modules in Azure by using Terratest](https://docs.microsoft.com/en-us/azure/terraform/terratest-in-terraform-modules) as this code stopped working as of terraform 0.12.

Some of the install instructions assume that you are running bash/WSL on Windows 

## Install required components

1. [Install terraform](
https://techcommunity.microsoft.com/t5/Azure-Developer-Community-Blog/Configuring-Terraform-on-Windows-10-Linux-Sub-System/ba-p/393845)
2. [Install go](https://sal.as/post/install-golan-on-wsl/)
    
    I ignored the GOROOT path, the GOPATH should point to the root location for your projects
3. [Install dep](https://github.com/golang/dep)

## Set up the terraform sample

