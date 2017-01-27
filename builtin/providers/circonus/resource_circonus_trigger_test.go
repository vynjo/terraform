package circonus

import (
	"fmt"
	"strings"
	"testing"

	"github.com/circonus-labs/circonus-gometrics/api"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccCirconusTrigger_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyCirconusTrigger,
		Steps: []resource.TestStep{
			{
				Config: testAccCirconusTriggerConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("circonus_trigger.icmp-latency-alarm", "check"),
					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "stream_name", "maximum"),
					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "metric_type", "numeric"),
					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "notes", "Simple check to create notifications based on ICMP performance."),
					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "link", "https://wiki.example.org/playbook/what-to-do-when-high-latency-strikes"),
					// resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "parent", "some check ID"),
					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.#", "4"),

					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.0.value.#", "1"),
					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.0.value.4211424620.absent", "70s"),
					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.0.value.4211424620.over.#", "0"),
					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.0.then.#", "1"),
					// Computed:
					// resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.0.then.<computed>.notify.#", "1"),
					// resource.TestCheckResourceAttrSet("circonus_trigger.icmp-latency-alarm", "if.0.then.<computed>.notify.0"),
					// resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.0.then.<computed>.severity", "1"),

					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.1.value.#", "1"),
					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.1.value.3306741531.over.#", "1"),
					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.1.value.3306741531.over.689776960.last", "120s"),
					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.1.value.3306741531.over.689776960.using", "average"),
					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.1.value.3306741531.less", "2"),
					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.1.then.#", "1"),
					// Computed:
					// resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.1.then.<computed>.notify.#", "1"),
					// resource.TestCheckResourceAttrSet("circonus_trigger.icmp-latency-alarm", "if.1.then.<computed>.notify.0"),
					// resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.1.then.<computed>.severity", "2"),

					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.2.value.#", "1"),
					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.2.value.2294297249.over.#", "1"),
					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.2.value.2294297249.over.999877839.last", "180s"),
					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.2.value.2294297249.over.999877839.using", "average"),
					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.2.value.2294297249.more", "300"),
					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.2.then.#", "1"),
					// Computed:
					// resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.2.then.<computed>.notify.#", "1"),
					// resource.TestCheckResourceAttrSet("circonus_trigger.icmp-latency-alarm", "if.2.then.<computed>.notify.0"),
					// resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.2.then.<computed>.severity", "3"),

					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.3.value.#", "1"),
					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.3.value.76621936.over.#", "0"),
					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.3.value.76621936.more", "400"),
					// Computed:
					// resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.3.then.<computed>.notify.#", "1"),
					// resource.TestCheckResourceAttrSet("circonus_trigger.icmp-latency-alarm", "if.3.then.<computed>.notify.0"),
					// resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.3.then.<computed>.after", "2400s"),
					// resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "if.3.then.<computed>.severity", "4"),

					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "tags.#", "2"),
					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "tags.2087084518", "author:terraform"),
					resource.TestCheckResourceAttr("circonus_trigger.icmp-latency-alarm", "tags.1401442048", "lifecycle:unittest"),
				),
			},
		},
	})
}

func testAccCheckDestroyCirconusTrigger(s *terraform.State) error {
	ctxt := testAccProvider.Meta().(*_ProviderContext)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "circonus_trigger" {
			continue
		}

		cid := rs.Primary.ID
		exists, err := checkTriggerExists(ctxt, api.CIDType(&cid))
		switch {
		case !exists:
			// noop
		case exists:
			return fmt.Errorf("stream group still exists after destroy")
		case err != nil:
			return fmt.Errorf("Error checking stream group %s", err)
		}
	}

	return nil
}

func testAccTriggerExists(n string, streamGroupID api.CIDType) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		ctxt := testAccProvider.Meta().(*_ProviderContext)
		cid := rs.Primary.ID
		exists, err := checkTriggerExists(ctxt, api.CIDType(&cid))
		switch {
		case !exists:
			// noop
		case exists:
			return fmt.Errorf("stream group still exists after destroy")
		case err != nil:
			return fmt.Errorf("Error checking stream group %s", err)
		}

		return nil
	}
}

func checkTriggerExists(c *_ProviderContext, streamGroupID api.CIDType) (bool, error) {
	sg, err := c.client.FetchMetricCluster(streamGroupID, "")
	if err != nil {
		if strings.Contains(err.Error(), defaultCirconus404ErrorString) {
			return false, nil
		}

		return false, err
	}

	if api.CIDType(&sg.CID) == streamGroupID {
		return true, nil
	}

	return false, nil
}

const testAccCirconusTriggerConfig = `
variable "test_tags" {
  type = "list"
  default = [ "author:terraform", "lifecycle:unittest" ]
}

resource "circonus_contact_group" "test-trigger" {
  name = "ops-staging-sev3"
  tags = [ "${var.test_tags}" ]
}

resource "circonus_check" "api_latency" {
  active = true
  name = "ICMP Ping check"
  period = "60s"

  collector {
    id = "/broker/1"
  }

  icmp_ping {
    count = 1
  }

  stream {
    name = "maximum"
    tags = [ "${var.test_tags}" ]
    type = "numeric"
    unit = "seconds"
  }

  tags = [ "${var.test_tags}" ]
  target = "api.circonus.com"
}

resource "circonus_trigger" "icmp-latency-alarm" {
  check = "${circonus_check.api_latency.checks[0]}"
  stream_name = "maximum"
  // stream_name = "${circonus_check.api_latency.stream["maximum"].name}"
  // metric_type = "${circonus_check.api_latency.stream["maximum"].type}"
  notes = <<EOF
Simple check to create notifications based on ICMP performance.
EOF
  link = "https://wiki.example.org/playbook/what-to-do-when-high-latency-strikes"
#  parent = "${check cid}"

  if {
    value {
      absent = "70s"
    }

    then {
      notify = [ "${circonus_contact_group.test-trigger.id}" ]
      severity = 1
    }
  }

  if {
    value {
      over {
        last = "120s"
        using = "average"
      }

      less = 2
    }

    then {
      notify = [ "${circonus_contact_group.test-trigger.id}" ]
      severity = 2
    }
  }

  if {
    value {
      over {
        last = "180s"
        using = "average"
      }

      more = 300
    }

    then {
      notify = [ "${circonus_contact_group.test-trigger.id}" ]
      severity = 3
    }
  }

  if {
    value {
      more = 400
    }

    then {
      notify = [ "${circonus_contact_group.test-trigger.id}" ]
      after = "2400s"
      severity = 4
    }
  }

  tags = [ "${var.test_tags}" ]
}
`