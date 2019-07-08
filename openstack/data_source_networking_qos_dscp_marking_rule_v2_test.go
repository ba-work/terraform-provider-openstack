package openstack

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccNetworkingV2QoSDSCPMarkingRuleDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckAdminOnly(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkingV2QoSDSCPMarkingRule_dataSource,
			},
			{
				Config: testAccOpenStackNetworkingQoSDSCPMarkingRuleV2DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkingQoSDSCPMarkingRuleV2DataSourceID("data.openstack_networking_qos_dscp_marking_rule_v2.dscp_mark_rule_1"),
					resource.TestCheckResourceAttr(
						"data.openstack_networking_qos_dscp_marking_rule_v2.dscp_mark_rule_1", "dscp_mark", "26"),
				),
			},
		},
	})
}

func testAccCheckNetworkingQoSDSCPMarkingRuleV2DataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find QoS DSCP marking data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("QoS DSCP marking data source ID not set")
		}

		return nil
	}
}

const testAccNetworkingV2QoSDSCPMarkingRule_dataSource = `
resource "openstack_networking_qos_policy_v2" "qos_policy_1" {
  name = "qos_policy_1"
}
resource "openstack_networking_qos_dscp_marking_rule_v2" "dscp_mark_rule_1" {
  qos_policy_id  = "${openstack_networking_qos_policy_v2.qos_policy_1.id}"
  dscp_mark      = 26
}
`

var testAccOpenStackNetworkingQoSDSCPMarkingRuleV2DataSource_basic = fmt.Sprintf(`
%s
data "openstack_networking_qos_dscp_marking_rule_v2" "dscp_mark_rule_1" {
  qos_policy_id = "${openstack_networking_qos_policy_v2.qos_policy_1.id}"
  dscp_mark     = "${openstack_networking_qos_dscp_marking_rule_v2.dscp_mark_rule_1.dscp_mark}"
}
`, testAccNetworkingV2QoSDSCPMarkingRule_dataSource)