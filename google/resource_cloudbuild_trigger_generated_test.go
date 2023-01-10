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

func TestAccCloudBuildTrigger_cloudbuildTriggerFilenameExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudBuildTriggerDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudBuildTrigger_cloudbuildTriggerFilenameExample(context),
			},
			{
				ResourceName:            "google_cloudbuild_trigger.filename-trigger",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccCloudBuildTrigger_cloudbuildTriggerFilenameExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloudbuild_trigger" "filename-trigger" {
  location = "us-central1"

  trigger_template {
    branch_name = "main"
    repo_name   = "my-repo"
  }

  substitutions = {
    _FOO = "bar"
    _BAZ = "qux"
  }

  filename = "cloudbuild.yaml"
}
`, context)
}

func TestAccCloudBuildTrigger_cloudbuildTriggerBuildExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudBuildTriggerDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudBuildTrigger_cloudbuildTriggerBuildExample(context),
			},
			{
				ResourceName:            "google_cloudbuild_trigger.build-trigger",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccCloudBuildTrigger_cloudbuildTriggerBuildExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloudbuild_trigger" "build-trigger" {
  location = "global"

  trigger_template {
    branch_name = "main"
    repo_name   = "my-repo"
  }

  build {
    step {
      name = "gcr.io/cloud-builders/gsutil"
      args = ["cp", "gs://mybucket/remotefile.zip", "localfile.zip"]
      timeout = "120s"
      secret_env = ["MY_SECRET"]
    }

    step {
      name   = "ubuntu"
      script = "echo hello" # using script field
    }
    
    source {
      storage_source {
        bucket = "mybucket"
        object = "source_code.tar.gz"
      }
    }
    tags = ["build", "newFeature"]
    substitutions = {
      _FOO = "bar"
      _BAZ = "qux"
    }
    queue_ttl = "20s"
    logs_bucket = "gs://mybucket/logs"
    secret {
      kms_key_name = "projects/myProject/locations/global/keyRings/keyring-name/cryptoKeys/key-name"
      secret_env = {
        PASSWORD = "ZW5jcnlwdGVkLXBhc3N3b3JkCg=="
      }
    }
    available_secrets {
      secret_manager {
        env          = "MY_SECRET"
        version_name = "projects/myProject/secrets/mySecret/versions/latest"
      }
    }
    artifacts {
      images = ["gcr.io/$PROJECT_ID/$REPO_NAME:$COMMIT_SHA"]
      objects {
        location = "gs://bucket/path/to/somewhere/"
        paths = ["path"]
      }
    }
    options {
      source_provenance_hash = ["MD5"]
      requested_verify_option = "VERIFIED"
      machine_type = "N1_HIGHCPU_8"
      disk_size_gb = 100
      substitution_option = "ALLOW_LOOSE"
      dynamic_substitutions = true
      log_streaming_option = "STREAM_OFF"
      worker_pool = "pool"
      logging = "LEGACY"
      env = ["ekey = evalue"]
      secret_env = ["secretenv = svalue"]
      volumes {
        name = "v1"
        path = "v1"
      }
    }
  }
}
`, context)
}

func TestAccCloudBuildTrigger_cloudbuildTriggerServiceAccountExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudBuildTriggerDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudBuildTrigger_cloudbuildTriggerServiceAccountExample(context),
			},
			{
				ResourceName:            "google_cloudbuild_trigger.service-account-trigger",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccCloudBuildTrigger_cloudbuildTriggerServiceAccountExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_project" "project" {}

resource "google_cloudbuild_trigger" "service-account-trigger" {
  trigger_template {
    branch_name = "main"
    repo_name   = "my-repo"
  }

  service_account = google_service_account.cloudbuild_service_account.id
  filename        = "cloudbuild.yaml"
  depends_on = [
    google_project_iam_member.act_as,
    google_project_iam_member.logs_writer
  ]
}

resource "google_service_account" "cloudbuild_service_account" {
  account_id = "my-service-account"
}

resource "google_project_iam_member" "act_as" {
  project = data.google_project.project.project_id
  role    = "roles/iam.serviceAccountUser"
  member  = "serviceAccount:${google_service_account.cloudbuild_service_account.email}"
}

resource "google_project_iam_member" "logs_writer" {
  project = data.google_project.project.project_id
  role    = "roles/logging.logWriter"
  member  = "serviceAccount:${google_service_account.cloudbuild_service_account.email}"
}
`, context)
}

func TestAccCloudBuildTrigger_cloudbuildTriggerPubsubConfigExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudBuildTriggerDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudBuildTrigger_cloudbuildTriggerPubsubConfigExample(context),
			},
			{
				ResourceName:            "google_cloudbuild_trigger.pubsub-config-trigger",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccCloudBuildTrigger_cloudbuildTriggerPubsubConfigExample(context map[string]interface{}) string {
	return Nprintf(`

resource "google_pubsub_topic" "mytopic" {
  name = "mytopic"
}

resource "google_cloudbuild_trigger" "pubsub-config-trigger" {
  location    = "us-central1"
  name        = "pubsub-trigger"
  description = "acceptance test example pubsub build trigger"

  pubsub_config {
    topic = google_pubsub_topic.mytopic.id
  }

  source_to_build {
    uri       = "https://hashicorp/terraform-provider-google-beta"
    ref       = "refs/heads/main"
    repo_type = "GITHUB"
  }

  git_file_source {
    path      = "cloudbuild.yaml"
    uri       = "https://hashicorp/terraform-provider-google-beta"
    revision  = "refs/heads/main"
    repo_type = "GITHUB"
  }

  substitutions = {
    _ACTION       = "$(body.message.data.action)"
  }

  filter = "_ACTION.matches('INSERT')"
}
`, context)
}

func TestAccCloudBuildTrigger_cloudbuildTriggerWebhookConfigExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudBuildTriggerDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudBuildTrigger_cloudbuildTriggerWebhookConfigExample(context),
			},
			{
				ResourceName:            "google_cloudbuild_trigger.webhook-config-trigger",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccCloudBuildTrigger_cloudbuildTriggerWebhookConfigExample(context map[string]interface{}) string {
	return Nprintf(`

resource "google_secret_manager_secret" "webhook_trigger_secret_key" {
  secret_id = "webhook_trigger-secret-key-1"

  replication {
    user_managed {
      replicas {
        location = "us-central1"
      }
    }
  }
}

resource "google_secret_manager_secret_version" "webhook_trigger_secret_key_data" {
  secret = google_secret_manager_secret.webhook_trigger_secret_key.id

  secret_data = "secretkeygoeshere"
}

data "google_project" "project" {}

data "google_iam_policy" "secret_accessor" {
  binding {
    role = "roles/secretmanager.secretAccessor"
    members = [
      "serviceAccount:service-${data.google_project.project.number}@gcp-sa-cloudbuild.iam.gserviceaccount.com",
    ]
  }
}

resource "google_secret_manager_secret_iam_policy" "policy" {
  project = google_secret_manager_secret.webhook_trigger_secret_key.project
  secret_id = google_secret_manager_secret.webhook_trigger_secret_key.secret_id
  policy_data = data.google_iam_policy.secret_accessor.policy_data
}


resource "google_cloudbuild_trigger" "webhook-config-trigger" {
  name        = "webhook-trigger"
  description = "acceptance test example webhook build trigger"
 
 webhook_config {
    secret = google_secret_manager_secret_version.webhook_trigger_secret_key_data.id
  }

  source_to_build {
    uri       = "https://hashicorp/terraform-provider-google-beta"
    ref       = "refs/heads/main"
    repo_type = "GITHUB"
  }

  git_file_source {
    path      = "cloudbuild.yaml"
    uri       = "https://hashicorp/terraform-provider-google-beta"
    revision  = "refs/heads/main"
    repo_type = "GITHUB"
  }
}
`, context)
}

func TestAccCloudBuildTrigger_cloudbuildTriggerManualExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudBuildTriggerDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudBuildTrigger_cloudbuildTriggerManualExample(context),
			},
			{
				ResourceName:            "google_cloudbuild_trigger.manual-trigger",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccCloudBuildTrigger_cloudbuildTriggerManualExample(context map[string]interface{}) string {
	return Nprintf(`

resource "google_cloudbuild_trigger" "manual-trigger" {
  name        = "manual-build"

  source_to_build {
    uri       = "https://hashicorp/terraform-provider-google-beta"
    ref       = "refs/heads/main"
    repo_type = "GITHUB"
  }

  git_file_source {
    path      = "cloudbuild.yaml"
    uri       = "https://hashicorp/terraform-provider-google-beta"
    revision  = "refs/heads/main"
    repo_type = "GITHUB"
  }

  
  // If this is set on a build, it will become pending when it is run, 
  // and will need to be explicitly approved to start.
  approval_config {
     approval_required = true 
  }
   
  
}
`, context)
}

func testAccCheckCloudBuildTriggerDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_cloudbuild_trigger" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{CloudBuildBasePath}}projects/{{project}}/locations/{{location}}/triggers/{{trigger_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("CloudBuildTrigger still exists at %s", url)
			}
		}

		return nil
	}
}