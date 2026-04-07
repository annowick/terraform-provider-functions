terraform {
  required_providers {
    functions = {
      source = "andrzej/test/andrzejtest"
      # source  = "annowick/functions"
      # version = "~> 0.0.1"
    }
  }
  required_version = "~> 1.8"
}

provider "functions" {}

output "map_product" {
  value = provider::functions::map_product(
    {
      privateDns1 = "storage"
      privateDns2 = "webapps"
    },
    {
      vNet1 = "office"
      vNet2 = "field"
    },
    "|"
  )
}

output "valid_cidr" {
  value = provider::functions::validate_cidr("10.1.1.1/12") // true
}

output "invalid_cidr" {
  value = provider::functions::validate_cidr("not-a-cidr") // false
}

output "valid_netcidr" {
  value = provider::functions::validate_netcidr("10.1.1.0/24") // true
}

output "invalid_netcidr" {
  value = provider::functions::validate_netcidr("10.1.1.0/16") // false
}

output "valid_hostcidr" {
  value = provider::functions::validate_hostcidr("10.1.1.1/24") // true
}

output "invalid_hostcidr_net" {
  value = provider::functions::validate_hostcidr("10.1.1.0/24") // false
}

output "invalid_hostcidr_bcast" {
  value = provider::functions::validate_hostcidr("10.1.1.255/24") // false
}