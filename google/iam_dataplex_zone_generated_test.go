// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccDataplexZoneIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
		"project_name":  acctest.GetTestProjectFromEnv(),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataplexZoneIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_dataplex_zone_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s roles/viewer", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-lake%s", context["random_suffix"]), fmt.Sprintf("tf-test-zone%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccDataplexZoneIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_dataplex_zone_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s roles/viewer", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-lake%s", context["random_suffix"]), fmt.Sprintf("tf-test-zone%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataplexZoneIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
		"project_name":  acctest.GetTestProjectFromEnv(),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccDataplexZoneIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_dataplex_zone_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s roles/viewer user:admin@hashicorptest.com", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-lake%s", context["random_suffix"]), fmt.Sprintf("tf-test-zone%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataplexZoneIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
		"role":          "roles/viewer",
		"project_name":  acctest.GetTestProjectFromEnv(),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataplexZoneIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_dataplex_zone_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_dataplex_zone_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-lake%s", context["random_suffix"]), fmt.Sprintf("tf-test-zone%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDataplexZoneIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_dataplex_zone_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s", acctest.GetTestProjectFromEnv(), acctest.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-lake%s", context["random_suffix"]), fmt.Sprintf("tf-test-zone%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDataplexZoneIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataplex_lake" "example" {
  location     = "us-central1"
  name         = "tf-test-lake%{random_suffix}"
  description  = "Test Lake"
  display_name = "Test Lake"

  labels = {
    my-lake = "exists"
  }

  project = "%{project_name}"
}


resource "google_dataplex_zone" "example" {
  name          = "tf-test-zone%{random_suffix}"
  discovery_spec {
    enabled = false
  }

  lake     = google_dataplex_lake.example.name
  location = "us-central1"

  resource_spec {
    location_type = "MULTI_REGION"
  }

  type         = "RAW"
  description  = "Test Zone"
  display_name = "Test Zone"
  labels       = {}
  project      = "%{project_name}"
}

resource "google_dataplex_zone_iam_member" "foo" {
  project = google_dataplex_zone.example.project
  location = google_dataplex_zone.example.location
  lake = google_dataplex_zone.example.lake
  dataplex_zone = google_dataplex_zone.example.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccDataplexZoneIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataplex_lake" "example" {
  location     = "us-central1"
  name         = "tf-test-lake%{random_suffix}"
  description  = "Test Lake"
  display_name = "Test Lake"

  labels = {
    my-lake = "exists"
  }

  project = "%{project_name}"
}


resource "google_dataplex_zone" "example" {
  name          = "tf-test-zone%{random_suffix}"
  discovery_spec {
    enabled = false
  }

  lake     = google_dataplex_lake.example.name
  location = "us-central1"

  resource_spec {
    location_type = "MULTI_REGION"
  }

  type         = "RAW"
  description  = "Test Zone"
  display_name = "Test Zone"
  labels       = {}
  project      = "%{project_name}"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_dataplex_zone_iam_policy" "foo" {
  project = google_dataplex_zone.example.project
  location = google_dataplex_zone.example.location
  lake = google_dataplex_zone.example.lake
  dataplex_zone = google_dataplex_zone.example.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_dataplex_zone_iam_policy" "foo" {
  project = google_dataplex_zone.example.project
  location = google_dataplex_zone.example.location
  lake = google_dataplex_zone.example.lake
  dataplex_zone = google_dataplex_zone.example.name
  depends_on = [
    google_dataplex_zone_iam_policy.foo
  ]
}
`, context)
}

func testAccDataplexZoneIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataplex_lake" "example" {
  location     = "us-central1"
  name         = "tf-test-lake%{random_suffix}"
  description  = "Test Lake"
  display_name = "Test Lake"

  labels = {
    my-lake = "exists"
  }

  project = "%{project_name}"
}


resource "google_dataplex_zone" "example" {
  name          = "tf-test-zone%{random_suffix}"
  discovery_spec {
    enabled = false
  }

  lake     = google_dataplex_lake.example.name
  location = "us-central1"

  resource_spec {
    location_type = "MULTI_REGION"
  }

  type         = "RAW"
  description  = "Test Zone"
  display_name = "Test Zone"
  labels       = {}
  project      = "%{project_name}"
}

data "google_iam_policy" "foo" {
}

resource "google_dataplex_zone_iam_policy" "foo" {
  project = google_dataplex_zone.example.project
  location = google_dataplex_zone.example.location
  lake = google_dataplex_zone.example.lake
  dataplex_zone = google_dataplex_zone.example.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccDataplexZoneIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataplex_lake" "example" {
  location     = "us-central1"
  name         = "tf-test-lake%{random_suffix}"
  description  = "Test Lake"
  display_name = "Test Lake"

  labels = {
    my-lake = "exists"
  }

  project = "%{project_name}"
}


resource "google_dataplex_zone" "example" {
  name          = "tf-test-zone%{random_suffix}"
  discovery_spec {
    enabled = false
  }

  lake     = google_dataplex_lake.example.name
  location = "us-central1"

  resource_spec {
    location_type = "MULTI_REGION"
  }

  type         = "RAW"
  description  = "Test Zone"
  display_name = "Test Zone"
  labels       = {}
  project      = "%{project_name}"
}

resource "google_dataplex_zone_iam_binding" "foo" {
  project = google_dataplex_zone.example.project
  location = google_dataplex_zone.example.location
  lake = google_dataplex_zone.example.lake
  dataplex_zone = google_dataplex_zone.example.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccDataplexZoneIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataplex_lake" "example" {
  location     = "us-central1"
  name         = "tf-test-lake%{random_suffix}"
  description  = "Test Lake"
  display_name = "Test Lake"

  labels = {
    my-lake = "exists"
  }

  project = "%{project_name}"
}


resource "google_dataplex_zone" "example" {
  name          = "tf-test-zone%{random_suffix}"
  discovery_spec {
    enabled = false
  }

  lake     = google_dataplex_lake.example.name
  location = "us-central1"

  resource_spec {
    location_type = "MULTI_REGION"
  }

  type         = "RAW"
  description  = "Test Zone"
  display_name = "Test Zone"
  labels       = {}
  project      = "%{project_name}"
}

resource "google_dataplex_zone_iam_binding" "foo" {
  project = google_dataplex_zone.example.project
  location = google_dataplex_zone.example.location
  lake = google_dataplex_zone.example.lake
  dataplex_zone = google_dataplex_zone.example.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}