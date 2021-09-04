resource "digitalocean_record" "do_domain_record_a" {
  domain = digitalocean_domain.do_domain_record.name
  type   = "A"
  name   = "www"
  value  = "192.168.0.11"
}

resource "digitalocean_record" "do_domain_record_txt" {
  domain = digitalocean_domain.do_domain_record.name
  type   = "TXT"
  name   = "test"
  value  = "do_record_txt${var.test_prefix}-${var.test_suffix}"
}

resource "digitalocean_domain" "do_domain_record" {
  name       = "${var.test_prefix}-${var.test_suffix}.com"
}