package genesyscloud

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mypurecloud/platform-client-sdk-go/v74/platformclientv2"
	"time"
)

func dataSourceJourneyOutcome() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for Genesys Cloud Journey Outcome. Select a journey outcome by name",
		ReadContext: readWithPooledClient(dataSourceJourneyOutcomeRead),
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Journey Outcome name.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourceJourneyOutcomeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sdkConfig := m.(*providerMeta).ClientConfig
	journeyApi := platformclientv2.NewJourneyApiWithConfig(sdkConfig)

	name := d.Get("name").(string)

	return withRetries(ctx, 15*time.Second, func() *resource.RetryError {
		pageCount := 1 // Needed because of broken journey common paging
		for pageNum := 1; pageNum <= pageCount; pageNum++ {
			const pageSize = 100
			journeyOutcomes, _, getErr := journeyApi.GetJourneyOutcomes(pageNum, pageSize, "", nil, nil, "")
			if getErr != nil {
				return resource.NonRetryableError(fmt.Errorf("failed to get page of journey outcomes: %v", getErr))
			}

			if journeyOutcomes.Entities == nil || len(*journeyOutcomes.Entities) == 0 {
				return resource.RetryableError(fmt.Errorf("no journey outcome found with name %s", name))
			}

			for _, journeyOutcome := range *journeyOutcomes.Entities {
				if journeyOutcome.DisplayName != nil && *journeyOutcome.DisplayName == name {
					d.SetId(*journeyOutcome.Id)
					return nil
				}
			}

			pageCount = *journeyOutcomes.PageCount
		}
		return resource.RetryableError(fmt.Errorf("no journey outcome found with name %s", name))
	})
}