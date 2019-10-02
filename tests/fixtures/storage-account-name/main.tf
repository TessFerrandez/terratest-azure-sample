variable "website_name" {
  description = "The name of your static website."
}

module "staticwebpage" {
  source       = "../../../modules/web-site/"
  location     = "westus"
  website_name = "${var.website_name}"
  html_path    = "empty.html"
}

output "website_storage_name" {
  value = "${module.staticwebpage.storage_name}"
}