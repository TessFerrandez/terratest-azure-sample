package test

import (
	"encoding/json"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

type plan struct {
	Version string `json:"format_version"`
	Planned struct {
		Outputs struct {
			StorageName struct {
				Value string `json:"value"`
			} `json:"website_storage_name"`
		} `json:"outputs"`
	} `json:"planned_values"`
}

func TestUT_StorageAccountName_Plan(t *testing.T) {
	t.Parallel()

	testCases := map[string]string{
		"TestWebsiteName": "testwebsitenamedata001",
		"ALLCAPS":         "allcapsdata001",
		"S_p-e(c)i.a_l":   "specialdata001",
		"A1phaNum321":     "a1phanum321data001",
		"E5e-y7h_ng":      "e5ey7hngdata001",
	}

	for input, expected := range testCases {
		tfOptions := &terraform.Options{
			TerraformDir: "./fixtures/storage-account-name",
			Vars: map[string]interface{}{
				"website_name": input,
			},
		}

		tfPlanOutput := "terraform.tfplan"

		terraform.Init(t, tfOptions)
		terraform.RunTerraformCommand(t, tfOptions, terraform.FormatArgs(tfOptions, "plan", "-out="+tfPlanOutput)...)

		tfOptionsEmpty := &terraform.Options{}
		planJSON, err := terraform.RunTerraformCommandAndGetStdoutE(
			t, tfOptions, terraform.FormatArgs(tfOptionsEmpty, "show", "-json", tfPlanOutput)...,
		)

		if err != nil {
			t.Fatal(err)
		}

		res := plan{}
		json.Unmarshal([]byte(planJSON), &res)

		actualStorageName := res.Planned.Outputs.StorageName.Value
		if actualStorageName != expected {
			t.Fatalf("Expected %v, but found %v", expected, actualStorageName)
		} else {
			t.Logf("Success: Input:%v, StorageName:%v", input, actualStorageName)
		}
	}
}
