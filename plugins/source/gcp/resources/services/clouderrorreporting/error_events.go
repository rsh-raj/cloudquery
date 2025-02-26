package clouderrorreporting

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/errorreporting/apiv1beta1/errorreportingpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	errorreporting "cloud.google.com/go/errorreporting/apiv1beta1"
)

func ErrorEvents() *schema.Table {
	return &schema.Table{
		Name:        "gcp_clouderrorreporting_error_events",
		Description: `https://cloud.google.com/error-reporting/reference/rest/v1beta1/ErrorEvent`,
		Resolver:    fetchErrorEvents,
		Multiplex:   client.ProjectMultiplexEnabledServices("clouderrorreporting.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.ErrorEvent{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
		},
	}
}

func fetchErrorEvents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListEventsRequest{
		ProjectName: "projects/" + c.ProjectId, GroupId: parent.Item.(*pb.ErrorGroupStats).Group.GroupId,
	}
	gcpClient, err := errorreporting.NewErrorStatsClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListEvents(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp
	}
	return nil
}
