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

func TestAccMemcacheInstance_memcacheInstanceBasicExample(t *testing.T) {
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"network_name":  BootstrapSharedTestNetwork(t, "memcache-private"),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMemcacheInstanceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccMemcacheInstance_memcacheInstanceBasicExample(context),
			},
			{
				ResourceName:            "google_memcache_instance.instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "region"},
			},
		},
	})
}

func testAccMemcacheInstance_memcacheInstanceBasicExample(context map[string]interface{}) string {
	return Nprintf(`
// This example assumes this network already exists.
// The API creates a tenant network per network authorized for a
// Memcache instance and that network is not deleted when the user-created
// network (authorized_network) is deleted, so this prevents issues
// with tenant network quota.
// If this network hasn't been created and you are using this example in your
// config, add an additional network resource or change
// this from "data"to "resource"
data "google_compute_network" "memcache_network" {
  name = "%{network_name}"
}

resource "google_compute_global_address" "service_range" {
  name          = "address%{random_suffix}"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = data.google_compute_network.memcache_network.id
}

resource "google_service_networking_connection" "private_service_connection" {
  network                 = data.google_compute_network.memcache_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.service_range.name]
}

resource "google_memcache_instance" "instance" {
  name = "tf-test-test-instance%{random_suffix}"
  authorized_network = google_service_networking_connection.private_service_connection.network

  node_config {
    cpu_count      = 1
    memory_size_mb = 1024
  }
  node_count = 1
  memcache_version = "MEMCACHE_1_5"

  maintenance_policy {
    weekly_maintenance_window {
      day      = "SATURDAY"
      duration = "14400s"
      start_time {
        hours = 0
        minutes = 30
        seconds = 0
        nanos = 0
      }
    }
  }
}
`, context)
}

func testAccCheckMemcacheInstanceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_memcache_instance" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{MemcacheBasePath}}projects/{{project}}/locations/{{region}}/instances/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("MemcacheInstance still exists at %s", url)
			}
		}

		return nil
	}
}