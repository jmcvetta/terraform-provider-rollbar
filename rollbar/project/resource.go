package project

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jmcvetta/terraform-provider-rollbar/client"
	"github.com/rs/zerolog/log"
	"gopkg.in/jeevatkm/go-model.v1"
	"strconv"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		//UpdateContext: resourceProjectUpdate,
		DeleteContext: resourceProjectDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"account_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"date_created": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"date_modified": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	name := d.Get("name").(string)
	log.Debug().Str("name", name).
		Msg("Creating new project")

	c := m.(*client.RollbarApiClient)
	p, err := c.CreateProject(name)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(p.Id))
	d.Set("id", p.Id)

	readDiags := resourceProjectRead(ctx, d, m)

	for _, v := range readDiags {
		diags = append(diags, v)
	}
	return diags
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	id := d.Id()
	l := log.With().Str("id", id).Logger().With().Stack().Logger()
	l.Debug().Msg("Reading project resource")

	c := m.(*client.RollbarApiClient)
	proj, err := c.ReadProject(id)
	if err != nil {
		return diag.FromErr(err)
	}
	mProj, err := model.Map(proj)
	if err != nil {
		l.Err(err).Send()
		return diag.FromErr(err)
	}
	for k, v := range mProj {
		l.Debug().Interface(k, v).Msg("Set resource data")
		err = d.Set(k, v)
		if err != nil {
			l.Err(err).Send()
			return diag.FromErr(err)
		}
	}
	d.SetId(strconv.Itoa(proj.Id))

	return diags
}

func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceProjectRead(ctx, d, m)
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}