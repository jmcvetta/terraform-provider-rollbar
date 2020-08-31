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
		Schema:        resourceSchema(),
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

	moreDiags := resourceProjectRead(ctx, d, m)

	for _, v := range moreDiags {
		diags = append(diags, v)
	}
	return diags
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	id := d.Id()
	l := log.With().Str("id", id).Logger()
	l.Debug().Msg("Reading project resource")

	c := m.(*client.RollbarApiClient)
	proj, err := c.ReadProject(id)
	if err != nil {
		return diag.FromErr(err)
	}
	mProj, err := model.Map(proj)
	if err != nil {
		log.Err(err).Send()
		return diag.FromErr(err)
	}
	for k, v := range mProj {
		d.Set(k, v)
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