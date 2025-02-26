package hooks

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func deliveries() *schema.Table {
	return &schema.Table{
		Name:                "github_hook_deliveries",
		Resolver:            fetchDeliveries,
		PreResourceResolver: hooksGet,
		Transform: transformers.TransformWithStruct(&github.HookDelivery{},
			append(client.SharedTransformers(), transformers.WithPrimaryKeys("ID"))...),
		Columns: []schema.Column{
			client.OrgColumn,
			{
				Name:            "hook_id",
				Type:            schema.TypeInt,
				Resolver:        schema.ParentColumnResolver("id"),
				Description:     `Hook ID`,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	}
}

func fetchDeliveries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	id := *parent.Item.(*github.Hook).ID

	c := meta.(*client.Client)
	opts := &github.ListCursorOptions{PerPage: 100}

	for {
		hookDeliveries, resp, err := c.Github.Organizations.ListHookDeliveries(ctx, c.Org, id, opts)
		if err != nil {
			return err
		}
		res <- hookDeliveries

		opts.Cursor = resp.NextPageToken
		if resp.NextPageToken == "" {
			return nil
		}
	}
}

func hooksGet(ctx context.Context, meta schema.ClientMeta, r *schema.Resource) error {
	hook := *r.Parent.Item.(*github.Hook)
	delivery := r.Item.(*github.HookDelivery)
	c := meta.(*client.Client)

	deliveryWithRequestResponse, _, err := c.Github.Organizations.GetHookDelivery(ctx, c.Org, *hook.ID, *delivery.ID)
	if err != nil {
		return err
	}

	r.SetItem(deliveryWithRequestResponse)
	return nil
}
