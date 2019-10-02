output "homepage_url" {
  value = "${azurerm_storage_blob.homepage.url}"
}

output "storage_name" {
  value = "${azurerm_storage_account.main.name}"
}