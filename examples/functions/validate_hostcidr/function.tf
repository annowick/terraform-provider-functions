output "valid_hostcidr" {
  value = provider::functions::validate_hostcidr("10.1.1.1/24") // true
}

output "invalid_hostcidr_net" {
  value = provider::functions::validate_hostcidr("10.1.1.0/24") // false
}

output "invalid_hostcidr_bcast" {
  value = provider::functions::validate_hostcidr("10.1.1.255/24") // false
}
