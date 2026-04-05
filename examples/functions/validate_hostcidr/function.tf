output "map_product" {
  value = provider::functions::validate_hostcidr("10.1.1.1/24")
}