package circonus

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCirconusCheckCAQL_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyCirconusCheckBundle,
		Steps: []resource.TestStep{
			{
				Config: testAccCirconusCheckCAQLConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("circonus_check.go_gc_latency", "active", "true"),
					resource.TestCheckResourceAttr("circonus_check.go_gc_latency", "collector.#", "1"),
					resource.TestCheckResourceAttr("circonus_check.go_gc_latency", "collector.36214388.id", "/broker/1490"),
					resource.TestCheckResourceAttr("circonus_check.go_gc_latency", "caql.#", "1"),
					resource.TestCheckResourceAttr("circonus_check.go_gc_latency", "caql.4060628048.query", `search:metric:histogram("*consul*runtime`+"`"+`gc_pause_ns* (active:1)") | histogram:merge() | histogram:percentile(99)`),
					resource.TestCheckResourceAttr("circonus_check.go_gc_latency", "name", "Consul's Go GC latency (Merged Histogram)"),
					resource.TestCheckResourceAttr("circonus_check.go_gc_latency", "period", "60s"),
					resource.TestCheckResourceAttr("circonus_check.go_gc_latency", "stream.#", "1"),

					resource.TestCheckResourceAttr("circonus_check.go_gc_latency", "tags.#", "4"),
					resource.TestCheckResourceAttr("circonus_check.go_gc_latency", "tags.3728194417", "app:consul"),
					resource.TestCheckResourceAttr("circonus_check.go_gc_latency", "tags.2087084518", "author:terraform"),
					resource.TestCheckResourceAttr("circonus_check.go_gc_latency", "tags.1401442048", "lifecycle:unittest"),
					resource.TestCheckResourceAttr("circonus_check.go_gc_latency", "tags.3480593708", "source:goruntime"),
					resource.TestCheckResourceAttr("circonus_check.go_gc_latency", "target", "q._caql"),
					resource.TestCheckResourceAttr("circonus_check.go_gc_latency", "type", "caql"),
				),
			},
		},
	})
}

const testAccCirconusCheckCAQLConfig = `
variable "test_tags" {
  type = "list"
  default = [ "app:consul", "author:terraform", "lifecycle:unittest", "source:goruntime" ]
}

resource "circonus_check" "go_gc_latency" {
  active = true
  name = "Consul's Go GC latency (Merged Histogram)"
  period = "60s"

  collector {
    id = "/broker/1490"
  }

  caql {
    query = <<EOF
search:metric:histogram("*consul*runtime` + "`" + `gc_pause_ns* (active:1)") | histogram:merge() | histogram:percentile(99)
EOF
  }

  stream {
    name = "output[1]"
    tags = [ "${var.test_tags}" ]
    type = "histogram"
    unit = "seconds"
  }

  tags = [ "${var.test_tags}" ]
}
`