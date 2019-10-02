package test

import (
	"fmt"
	"strings"
	"testing"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestIT_HelloWorldExample(t *testing.T) {
	t.Parallel()

	// Generate a random website name to prevent a naming conflict
	uniqueID := random.UniqueId()
	websiteName := fmt.Sprintf("Hello-World-%s", uniqueID)

	// Specify the test case folder and "-var" options
	tfOptions := &terraform.Options{
		TerraformDir: "../examples/hello-world",
		Vars: map[string]interface{}{
			"website_name": websiteName,
		},
	}

	// Defer terraform destroy til the end of the test
	defer terraform.Destroy(t, tfOptions)

	// run terraform init, and terraform apply and capture the homepage value
	terraform.InitAndApply(t, tfOptions)
	homepage := terraform.Output(t, tfOptions, "homepage")

	// Validate the provisioned webpage
	http_helper.HttpGetWithCustomValidation(t, homepage, func(status int, content string) bool {
		return status == 200 &&
			strings.Contains(content, "Hi, Terraform Module") &&
			strings.Contains(content, "This is a sample web page to demonstrate Terratest.")
	})
}
