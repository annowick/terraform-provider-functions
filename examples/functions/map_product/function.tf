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
    ":"
  )
}

# map_product = tomap({
#   "privateDns1|vNet1" = tolist(["storage", "office"])
#   "privateDns1|vNet2" = tolist(["storage", "field"])
#   "privateDns2|vNet1" = tolist(["webapps", "office"])
#   "privateDns2|vNet2" = tolist(["webapps", "field"])
# })
