package compute

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func adminPasswordDiffSuppressFunc(_, old, new string, _ *schema.ResourceData) bool {
	// this is not the greatest hack in the world, this is just a tribute.
	if old == "ignored-as-imported" || new == "ignored-as-imported" {
		return true
	}

	return false
}
