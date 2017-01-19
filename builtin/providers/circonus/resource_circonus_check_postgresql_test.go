package circonus

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCirconusCheckPostgreSQL_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDestroyCirconusCheckBundle,
		Steps: []resource.TestStep{
			{
				Config: testAccCirconusCheckPostgreSQLConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("circonus_check.table_ops", "active", "true"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "collector.#", "1"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "collector.2388330941.id", "/broker/1"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "postgresql.#", "1"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "postgresql.1831600166.dsn", "user=postgres host=pg1.example.org port=5432 password=12345 sslmode=require"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "postgresql.1831600166.query", `SELECT 'tables', sum(n_tup_ins) as inserts, sum(n_tup_upd) as updates, sum(n_tup_del) as deletes, sum(idx_scan)  as index_scans, sum(seq_scan) as seq_scans, sum(idx_tup_fetch) as index_tup_fetch, sum(seq_tup_read) as seq_tup_read from pg_stat_all_tables`),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "name", "PostgreSQL ops per table check"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "period", "300s"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.#", "7"),

					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.997634628.name", "tables`inserts"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.997634628.tags.#", "2"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.997634628.tags.2087084518", "author:terraform"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.997634628.tags.1401442048", "lifecycle:unittest"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.997634628.type", "numeric"),

					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.565883273.name", "tables`updates"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.565883273.tags.#", "2"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.565883273.tags.2087084518", "author:terraform"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.565883273.tags.1401442048", "lifecycle:unittest"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.565883273.type", "numeric"),

					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.3965415003.name", "tables`deletes"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.3965415003.tags.#", "2"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.3965415003.tags.2087084518", "author:terraform"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.3965415003.tags.1401442048", "lifecycle:unittest"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.3965415003.type", "numeric"),

					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.3868690100.name", "tables`index_scans"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.3868690100.tags.#", "2"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.3868690100.tags.2087084518", "author:terraform"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.3868690100.tags.1401442048", "lifecycle:unittest"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.3868690100.type", "numeric"),

					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.2772400178.name", "tables`seq_scans"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.2772400178.tags.#", "2"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.2772400178.tags.2087084518", "author:terraform"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.2772400178.tags.1401442048", "lifecycle:unittest"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.2772400178.type", "numeric"),

					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.3278042831.name", "tables`index_tup_fetch"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.3278042831.tags.#", "2"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.3278042831.tags.2087084518", "author:terraform"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.3278042831.tags.1401442048", "lifecycle:unittest"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.3278042831.type", "numeric"),

					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.2685537214.name", "tables`seq_tup_read"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.2685537214.tags.#", "2"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.2685537214.tags.2087084518", "author:terraform"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.2685537214.tags.1401442048", "lifecycle:unittest"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "stream.2685537214.type", "numeric"),

					resource.TestCheckResourceAttr("circonus_check.table_ops", "tags.#", "2"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "tags.2087084518", "author:terraform"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "tags.1401442048", "lifecycle:unittest"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "target", "pgdb.example.org"),
					resource.TestCheckResourceAttr("circonus_check.table_ops", "type", "postgres"),
				),
			},
		},
	})
}

const testAccCirconusCheckPostgreSQLConfig = `
variable "test_tags" {
  type = "list"
  default = [ "author:terraform", "lifecycle:unittest" ]
}
resource "circonus_check" "table_ops" {
  active = true
  name = "PostgreSQL ops per table check"
  period = "300s"

  collector {
    id = "/broker/1"
  }

  postgresql {
    dsn = "user=postgres host=pg1.example.org port=5432 password=12345 sslmode=require"
    query = <<EOF
SELECT 'tables', sum(n_tup_ins) as inserts, sum(n_tup_upd) as updates, sum(n_tup_del) as deletes, sum(idx_scan)  as index_scans, sum(seq_scan) as seq_scans, sum(idx_tup_fetch) as index_tup_fetch, sum(seq_tup_read) as seq_tup_read from pg_stat_all_tables
EOF
  }

  stream {
    name = "tables` + "`" + `inserts"
    tags = [ "${var.test_tags}" ]
    type = "numeric"
  }

  stream {
    name = "tables` + "`" + `updates"
    tags = [ "${var.test_tags}" ]
    type = "numeric"
  }

  stream {
    name = "tables` + "`" + `deletes"
    tags = [ "${var.test_tags}" ]
    type = "numeric"
  }

  stream {
    name = "tables` + "`" + `index_scans"
    tags = [ "${var.test_tags}" ]
    type = "numeric"
  }

  stream {
    name = "tables` + "`" + `seq_scans"
    tags = [ "${var.test_tags}" ]
    type = "numeric"
  }

  stream {
    name = "tables` + "`" + `index_tup_fetch"
    tags = [ "${var.test_tags}" ]
    type = "numeric"
  }

  stream {
    name = "tables` + "`" + `seq_tup_read"
    tags = [ "${var.test_tags}" ]
    type = "numeric"
  }

  tags = [ "${var.test_tags}" ]
  target = "pgdb.example.org"
}
`