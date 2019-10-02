variable "website_name" {
  description = "The name of your static website."
  default     = "Hello-World"
}

module "staticwebpage" {
  source       = "../../modules/web-site/"
  location     = "West US"
  website_name = "${var.website_name}"
  html_path    = "index.html"
}

output "homepage" {
  value = "${module.staticwebpage.homepage_url}"
}