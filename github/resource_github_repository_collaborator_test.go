package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubRepositoryCollaborator(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates invitations without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "tf-acc-test-%s"
				auto_init = true
			}

			resource "github_repository_collaborator" "test_repo_collaborator" {
				repository = "${github_repository.test.name}"
				username   = "notjcudit"
				permission = "triage"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository_collaborator.test_repo_collaborator", "permission",
				"triage",
			),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

}

// func TestAccGithubRepositoryCollaboratorUpdates(t *testing.T) {

// 	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

// 	t.Run("creates and updates labels without error", func(t *testing.T) {

// 		description := "label_description"
// 		updatedDescription := "updated_label_description"

// 		config := fmt.Sprintf(`
// 			resource "github_repository" "test" {
// 				name = "tf-acc-test-%s"
// 				auto_init = true
// 			}

// 			resource "github_repository_collaborator" "test_repo_collaborator" {
// 				repository = "${github_repository.test.name}"
// 				username   = "%s"
// 				permission = "%s"
// 			}

// 		`, randomID, description)

// 		checks := map[string]resource.TestCheckFunc{
// 			"before": resource.ComposeTestCheckFunc(
// 				resource.TestCheckResourceAttr(
// 					"github_issue_label.test", "description",
// 					description,
// 				),
// 			),
// 			"after": resource.ComposeTestCheckFunc(
// 				resource.TestCheckResourceAttr(
// 					"github_issue_label.test", "description",
// 					updatedDescription,
// 				),
// 			),
// 		}

// 		testCase := func(t *testing.T, mode string) {
// 			resource.Test(t, resource.TestCase{
// 				PreCheck:  func() { skipUnlessMode(t, mode) },
// 				Providers: testAccProviders,
// 				Steps: []resource.TestStep{
// 					{
// 						Config: config,
// 						Check:  checks["before"],
// 					},
// 					{
// 						Config: strings.Replace(config,
// 							description,
// 							updatedDescription, 1),
// 						Check: checks["after"],
// 					},
// 				},
// 			})
// 		}

// 		t.Run("with an anonymous account", func(t *testing.T) {
// 			t.Skip("anonymous account not supported for this operation")
// 		})

// 		t.Run("with an individual account", func(t *testing.T) {
// 			testCase(t, individual)
// 		})

// 		t.Run("with an organization account", func(t *testing.T) {
// 			testCase(t, organization)
// 		})
// 	})

// }
