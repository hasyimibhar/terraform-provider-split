package split

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSplitUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSplitUserRead,
		Schema: map[string]*schema.Schema{
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},

			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"2fa": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSplitUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Config).API

	email := d.Get("email").(string)

	user, _, findErr := client.Users.FindByEmail(email)
	if findErr != nil {
		return diag.FromErr(findErr)
	}

	d.SetId(user.GetID())
	d.Set("name", user.GetName())
	d.Set("email", user.GetEmail())
	d.Set("2fa", user.GetTFA())
	d.Set("status", user.GetStatus())

	return nil
}
