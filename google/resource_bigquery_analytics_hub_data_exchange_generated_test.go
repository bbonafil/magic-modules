// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccBigqueryAnalyticsHubDataExchange_bigqueryAnalyticshubDataExchangeBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigqueryAnalyticsHubDataExchangeDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryAnalyticsHubDataExchange_bigqueryAnalyticshubDataExchangeBasicExample(context),
			},
			{
				ResourceName:            "google_bigquery_analytics_hub_data_exchange.data_exchange",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"data_exchange_id", "location"},
			},
		},
	})
}

func testAccBigqueryAnalyticsHubDataExchange_bigqueryAnalyticshubDataExchangeBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_analytics_hub_data_exchange" "data_exchange" {
  location         = "US"
  data_exchange_id = "tf_test_my_data_exchange%{random_suffix}"
  display_name     = "tf_test_my_data_exchange%{random_suffix}"
  description      = "example data exchange%{random_suffix}"
}
`, context)
}

func testAccCheckBigqueryAnalyticsHubDataExchangeDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_bigquery_analytics_hub_data_exchange" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{BigqueryAnalyticsHubBasePath}}projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("BigqueryAnalyticsHubDataExchange still exists at %s", url)
			}
		}

		return nil
	}
}