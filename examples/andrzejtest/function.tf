terraform {
  required_providers {
    functions = {
      source = "andrzej/test/andrzejtest"
    }
  }
  required_version = ">= 1.8.0"
}

provider "andrzejtest" {}

output "map_product" {
  value = provider::andrzejtest::map_product(
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