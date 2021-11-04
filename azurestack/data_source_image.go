package azurestack

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceArmImage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceArmImageRead,
		Schema: map[string]*schema.Schema{
			"resource_group_name": resourceGroupNameSchema(),
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceArmImageRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).imagesClient
	ctx := meta.(*ArmClient).StopContext

	resGroup := d.Get("resource_group_name").(string)
	name := d.Get("name").(string)

	img, err := client.Get(ctx, resGroup, name, "")
	if err != nil {
		return fmt.Errorf("image %q was not found in resource group %q", name, resGroup)
	}

	d.SetId(*img.ID)
	d.Set("name", img.Name)
	d.Set("resource_group_name", resGroup)
	if location := img.Location; location != nil {
		d.Set("location", azureStackNormalizeLocation(*location))
	}

	return nil
}
