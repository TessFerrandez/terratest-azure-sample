# Azure Terratest sample

This repo is a variation of [Test Terraform modules in Azure by using Terratest](https://docs.microsoft.com/en-us/azure/terraform/terratest-in-terraform-modules) as this code stopped working as of terraform 0.12.

[Terratest](https://github.com/gruntwork-io/terratest) is a go library that allows you to write automated tests.

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

## Integration test

### What are we testing?
In the first test, we'll test the whole install script.
1. Run terraform init and apply
2. Browse to the web page
3. Test that the web page is deployed
4. Run terraform destroy to revert to the state before the test.

### Setting up the basics
To do this, we create an "example" where we call the terraform module we created earlier, setting the input parameters, and an output parameter holding the webpage url that we can browse to and examine.

1. Browse to /examples/hello-world
2. Examine the main.tf file - note that we copy the index.html file and output the url
3. Run **terraform init**, **terraform apply** and finally **terraform destroy** in this directory to see how it works

You should see the output
~~~
Outputs:

homepage = https://helloworlddata001.blob.core.windows.net/wwwroot/index.html
~~~
If not, run terraform refresh

### The test
To create an automated test using terratest, we create a go file (hello_world_example_test.go) in the /tests folder

The key parts here are

1. Set up the tf options, i.e. the folder for the terraform scripts, and the input variables
2. Call terraform init and apply, and capture the homepage output variable
3. Validate the web page
4. Run terraform destroy (defer to end of function)

### Run the test

In order to run the tests, we need to download the dependencies

1. **dep init -v** (run only once at the beginning)
2. **dep ensure -v** (needs to be rerun if you import new packages)

Run the test
~~~
go test ./tests/ -v -timeout 30m | tee test_outputlog.log
~~~

Optionally: Install the [terratest_log_parser](https://github.com/gruntwork-io/terratest#installing-the-utility-binaries) to parse the tests so you can integrate the results in the CI/CD pipeline

to run the tests with the terratest_log_parser, see run_tests.sh

*NOTE: I have some issues with this test, the output variables are empty, but leaving it here in case it is just my machine*

## Unit test

### What are we testing?
In this test we will test the logic of creating a storage safe name.  

We want to test this a number of times with different website names, but don't neccesarily want to deploy.  It is enough that we know what the resulting storage name would be.

1. Run terraform init and plan
2. Retrieve the plan and read the storage name
3. Verify the storage name

### Setting up the basics
We create a test fixture where we call the terraform module we created earlier, setting the input parameters

1. Examine main.tf in the /modules/web-site directory and notice the logic for the azurerm_storage_account name -- this is what we are testing
1. Browse to /tests/fixtures/storage-account-name
2. Examine the main.tf file - note that we have an input variable website_name that we will pass in from the test
3. Optionally: Run **terraform init** in this directory to prep the needed terraform modules

### The test
To create an automated test using terratest, we create a go file (storage_account_name_unit_test.go) in the /tests folder

The key parts here are

1. Set up 5 test cases with website names and corresponding storage safew names 
2. Set up the tf options, i.e. the folder for the terraform scripts, and the input variables
3. Call terraform plan using terraform.RunTerraformCommand to allow for a plan output parameter
4. Get the plan file as json, by running **terraform show terraform.tfplan -json**
5. Parse the json to extract the storage name and compare to expected

### Run the test

Run all the tests
~~~
go test ./tests/ -v -timeout 30m | tee test_outputlog.log
~~~
or just this specific test
~~~
go test ./tests/storage_account_name_unit_test.go -v -timeout 30m | tee test_outputlog.log
~~~

*NOTE: This is much faster than InitAndApply since we don't need to deploy any real resources*

