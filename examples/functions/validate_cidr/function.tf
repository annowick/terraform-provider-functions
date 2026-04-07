output "map_product" {
  value = provider::functions::validate_cidr("10.1.1.1/24")
}
