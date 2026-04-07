output "valid_netcidr" {
  value = provider::functions::validate_netcidr("10.1.1.0/24") // true
}

output "invalid_netcidr" {
  value = provider::functions::validate_netcidr("10.1.1.0/16") // false
}
