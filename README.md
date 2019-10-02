# Azure Terratest sample

An variation of [Test Terraform modules in Azure by using Terratest](https://docs.microsoft.com/en-us/azure/terraform/terratest-in-terraform-modules) as this code stopped working as of terraform 0.12.

Some of the install instructions assume that you are running bash/WSL on Windows 

## Install required components

1. [terraform](
https://techcommunity.microsoft.com/t5/Azure-Developer-Community-Blog/Configuring-Terraform-on-Windows-10-Linux-Sub-System/ba-p/393845)
2. [go](https://sal.as/post/install-golan-on-wsl/)
    
    I ignored the GOROOT path, the GOPATH should point to the root location for your projects
3. [dep](https://github.com/golang/dep)
4. [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli-apt?view=azure-cli-latest)

## The Terraform scripts
The scripts we want to test are stored in the **/modules/web-site** directory.  

The following resources are created as part of the script (main.tf)

1. A resource group ([website_name]-staging-rg)
2. A storage account with clean storage name based on [website_name]data001 
3. A storage container (wwwroot)
4. A blob, containing a copied html file

#### Inputs: (variables.tf) 
- location (azure region)
- website_name
- html_path
     
#### Outputs: (outputs.tf)
- homepage_url (the url of the webpage)
- storage_name (the applied name of the storage account)

## Try the terraform scripts

From bash, in the modules/website directory

1. Log in with azure cli
~~~
az login
~~~
2. Initialize the terraform providers
~~~
terraform init
~~~
3. Optionally create a terraform plan (to verify the scripts) - this will prompt you for the input variables (empty.html, westus, mytestweb) - the plan will show what resources will be added
~~~
terraform plan
~~~
4. Apply the terraform script to create the resources (you can now go into the portal and inspect the resources)
~~~
terraform apply
~~~

At this point you should see output that looks like this:
~~~
Outputs:

homepage_url = https://mytestwebdata001.blob.core.windows.net/wwwroot/index.html
storage_name = mytestwebdata001
~~~

If you don't see both the outputs, run a terraform refresh

You should also be able to browse to the url, and get an empty page.

5. Finish by destroying all the resources you created
~~~
terraform destroy
~~~

## Run a basic go test to verify everything is working
In the folder /sample_go/math you can find some sample code based on the [golang book](https://www.golang-book.com/books/intro)

1. Navigate to the math folder
2. Inspect the math_test.go tests - the first and the 3rd should pass, the second should fail
3. Run the tests
~~~ 
go test
~~~
4. Fix the test by setting the expected average to 1 and rerun the test