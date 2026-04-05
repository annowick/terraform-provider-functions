output "map_product" {
  value = provider::functions::validate_netcidr("10.1.1.0/24")
}