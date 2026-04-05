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