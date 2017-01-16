package circonus

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCirconusCheckJSON_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyCirconusCheckBundle,
		Steps: []resource.TestStep{
			{
				Config: testAccCirconusCheckJSONConfig1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("circonus_check.usage", "active", "true"),
					resource.TestCheckResourceAttr("circonus_check.usage", "collector.#", "1"),
					resource.TestCheckResourceAttr("circonus_check.usage", "collector.2388330941.id", "/broker/1"),
					resource.TestCheckResourceAttr("circonus_check.usage", "json.#", "1"),
					// resource.TestCheckResourceAttr("circonus_check.usage", "json.2883347764.auth_method", ""),
					// resource.TestCheckResourceAttr("circonus_check.usage", "json.2883347764.auth_password", ""),
					// resource.TestCheckResourceAttr("circonus_check.usage", "json.2883347764.auth_user", ""),
					// resource.TestCheckResourceAttr("circonus_check.usage", "json.2883347764.ca_chain", ""),
					// resource.TestCheckResourceAttr("circonus_check.usage", "json.2883347764.certificate_file", ""),
					// resource.TestCheckResourceAttr("circonus_check.usage", "json.2883347764.ciphers", ""),
					// resource.TestCheckResourceAttr("circonus_check.usage", "json.2883347764.key_file", ""),
					// resource.TestCheckResourceAttr("circonus_check.usage", "json.2883347764.payload", ""),
					resource.TestCheckResourceAttr("circonus_check.usage", "json.2883347764.headers.%", "3"),
					resource.TestCheckResourceAttr("circonus_check.usage", "json.2883347764.headers.Accept", "application/json"),
					resource.TestCheckResourceAttr("circonus_check.usage", "json.2883347764.headers.X-Circonus-App-Name", "TerraformCheck"),
					resource.TestCheckResourceAttr("circonus_check.usage", "json.2883347764.headers.X-Circonus-Auth-Token", "<env 'CIRCONUS_API_TOKEN'>"),
					resource.TestCheckResourceAttr("circonus_check.usage", "json.2883347764.version", "1.0"),
					resource.TestCheckResourceAttr("circonus_check.usage", "json.2883347764.method", "GET"),
					resource.TestCheckResourceAttr("circonus_check.usage", "json.2883347764.port", "443"),
					resource.TestCheckResourceAttr("circonus_check.usage", "json.2883347764.read_limit", "1048576"),
					resource.TestCheckResourceAttr("circonus_check.usage", "json.2883347764.url", "https://api.circonus.com/account/current"),
					resource.TestCheckResourceAttr("circonus_check.usage", "name", "Terraform test: api.circonus.com metric usage check"),
					resource.TestCheckResourceAttr("circonus_check.usage", "notes", ""),
					resource.TestCheckResourceAttr("circonus_check.usage", "period", "60s"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.#", "2"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.1992097900.active", "true"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.1992097900.name", "_usage`0`_limit"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.1992097900.tags.#", "1"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.1992097900.tags.3241999189", "source:circonus"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.1992097900.type", "numeric"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.1992097900.unit", "qty"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.3280673139.active", "true"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.3280673139.name", "_usage`0`_used"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.3280673139.tags.#", "1"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.3280673139.tags.3241999189", "source:circonus"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.3280673139.type", "numeric"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.3280673139.unit", "qty"),
					resource.TestCheckResourceAttr("circonus_check.usage", "tags.#", "2"),
					resource.TestCheckResourceAttr("circonus_check.usage", "tags.3241999189", "source:circonus"),
					resource.TestCheckResourceAttr("circonus_check.usage", "tags.3839162439", "source:unittest"),
					resource.TestCheckResourceAttr("circonus_check.usage", "target", "api.circonus.com"),
					resource.TestCheckResourceAttr("circonus_check.usage", "type", "json"),
				),
			},
			{
				Config: testAccCirconusCheckJSONConfig2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("circonus_check.usage", "active", "true"),
					resource.TestCheckResourceAttr("circonus_check.usage", "collector.#", "1"),
					resource.TestCheckResourceAttr("circonus_check.usage", "collector.2388330941.id", "/broker/1"),
					resource.TestCheckResourceAttr("circonus_check.usage", "json.#", "1"),
					// resource.TestCheckResourceAttr("circonus_check.usage", "json.3705361826.auth_method", ""),
					// resource.TestCheckResourceAttr("circonus_check.usage", "json.3705361826.auth_password", ""),
					// resource.TestCheckResourceAttr("circonus_check.usage", "json.3705361826.auth_user", ""),
					// resource.TestCheckResourceAttr("circonus_check.usage", "json.3705361826.ca_chain", ""),
					// resource.TestCheckResourceAttr("circonus_check.usage", "json.3705361826.certificate_file", ""),
					// resource.TestCheckResourceAttr("circonus_check.usage", "json.3705361826.ciphers", ""),
					// resource.TestCheckResourceAttr("circonus_check.usage", "json.3705361826.key_file", ""),
					// resource.TestCheckResourceAttr("circonus_check.usage", "json.3705361826.payload", ""),
					resource.TestCheckResourceAttr("circonus_check.usage", "json.3705361826.headers.%", "3"),
					resource.TestCheckResourceAttr("circonus_check.usage", "json.3705361826.headers.Accept", "application/json"),
					resource.TestCheckResourceAttr("circonus_check.usage", "json.3705361826.headers.X-Circonus-App-Name", "TerraformCheck"),
					resource.TestCheckResourceAttr("circonus_check.usage", "json.3705361826.headers.X-Circonus-Auth-Token", "<env 'CIRCONUS_API_TOKEN'>"),
					resource.TestCheckResourceAttr("circonus_check.usage", "json.3705361826.version", "1.1"),
					resource.TestCheckResourceAttr("circonus_check.usage", "json.3705361826.method", "GET"),
					resource.TestCheckResourceAttr("circonus_check.usage", "json.3705361826.port", "443"),
					resource.TestCheckResourceAttr("circonus_check.usage", "json.3705361826.read_limit", "1048576"),
					resource.TestCheckResourceAttr("circonus_check.usage", "json.3705361826.url", "https://api.circonus.com/account/current"),
					resource.TestCheckResourceAttr("circonus_check.usage", "name", "Terraform test: api.circonus.com metric usage check"),
					resource.TestCheckResourceAttr("circonus_check.usage", "notes", "notes!"),
					resource.TestCheckResourceAttr("circonus_check.usage", "period", "300s"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.#", "2"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.1992097900.active", "true"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.1992097900.name", "_usage`0`_limit"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.1992097900.tags.#", "1"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.1992097900.tags.3241999189", "source:circonus"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.1992097900.type", "numeric"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.1992097900.unit", "qty"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.3280673139.active", "true"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.3280673139.name", "_usage`0`_used"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.3280673139.tags.#", "1"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.3280673139.tags.3241999189", "source:circonus"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.3280673139.type", "numeric"),
					resource.TestCheckResourceAttr("circonus_check.usage", "stream.3280673139.unit", "qty"),
					resource.TestCheckResourceAttr("circonus_check.usage", "tags.#", "2"),
					resource.TestCheckResourceAttr("circonus_check.usage", "tags.3241999189", "source:circonus"),
					resource.TestCheckResourceAttr("circonus_check.usage", "tags.3839162439", "source:unittest"),
					resource.TestCheckResourceAttr("circonus_check.usage", "target", "api.circonus.com"),
					resource.TestCheckResourceAttr("circonus_check.usage", "type", "json"),
				),
			},
		},
	})
}

const testAccCirconusCheckJSONConfig1 = `
variable "usage_default_unit" {
  default = "qty"
}

resource "circonus_metric" "limit" {
  name = "_usage` + "`0`" + `_limit"
  tags = [ "source:circonus" ]
  type = "numeric"
  unit = "${var.usage_default_unit}"
}

resource "circonus_metric" "used" {
  name = "_usage` + "`0`" + `_used"
  tags = [ "source:circonus" ]
  type = "numeric"
  unit = "${var.usage_default_unit}"
}

resource "circonus_check" "usage" {
  active = true
  name = "Terraform test: api.circonus.com metric usage check"
  period = "60s"

  collector {
    id = "/broker/1"
  }

  json {
    url = "https://api.circonus.com/account/current"
    headers = {
      Accept                = "application/json",
      X-Circonus-App-Name   = "TerraformCheck",
      X-Circonus-Auth-Token = "<env 'CIRCONUS_API_TOKEN'>",
    }
    version = "1.0"
    method = "GET"
    port = 443
    read_limit = 1048576
  }

  stream {
    name = "${circonus_metric.used.name}"
    tags = [ "${circonus_metric.used.tags}" ]
    type = "${circonus_metric.used.type}"
    unit = "${coalesce(circonus_metric.used.unit, var.usage_default_unit)}"
  }

  stream {
    name = "${circonus_metric.limit.name}"
    tags = [ "${circonus_metric.limit.tags}" ]
    type = "${circonus_metric.limit.type}"
    unit = "${coalesce(circonus_metric.limit.unit, var.usage_default_unit)}"
  }

  tags = [ "source:circonus", "source:unittest" ]
}
`

const testAccCirconusCheckJSONConfig2 = `
variable "usage_default_unit" {
  default = "qty"
}

resource "circonus_metric" "limit" {
  name = "_usage` + "`0`" + `_limit"
  tags = [ "source:circonus" ]
  type = "numeric"
  unit = "${var.usage_default_unit}"
}

resource "circonus_metric" "used" {
  name = "_usage` + "`0`" + `_used"
  tags = [ "source:circonus" ]
  type = "numeric"
  unit = "${var.usage_default_unit}"
}

resource "circonus_check" "usage" {
  active = true
  name = "Terraform test: api.circonus.com metric usage check"
  notes = "notes!"
  period = "300s"

  collector {
    id = "/broker/1"
  }

  json {
    url = "https://api.circonus.com/account/current"
    headers = {
      Accept                = "application/json",
      X-Circonus-App-Name   = "TerraformCheck",
      X-Circonus-Auth-Token = "<env 'CIRCONUS_API_TOKEN'>",
    }
    version = "1.1"
    method = "GET"
    port = 443
    read_limit = 1048576
  }

  stream {
    name = "${circonus_metric.used.name}"
    tags = [ "${circonus_metric.used.tags}" ]
    type = "${circonus_metric.used.type}"
    unit = "${coalesce(circonus_metric.used.unit, var.usage_default_unit)}"
  }

  stream {
    name = "${circonus_metric.limit.name}"
    tags = [ "${circonus_metric.limit.tags}" ]
    type = "${circonus_metric.limit.type}"
    unit = "${coalesce(circonus_metric.limit.unit, var.usage_default_unit)}"
  }

  tags = [ "source:circonus", "source:unittest" ]
}
`
