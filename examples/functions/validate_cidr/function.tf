output "valid_cidr" {
  value = provider::functions::validate_cidr("10.1.1.1/12") # true
}

output "invalid_cidr" {
  value = provider::functions::validate_cidr("not-a-cidr") # false
}
