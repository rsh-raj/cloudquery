package ec2

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func EbsSnapshots() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_ebs_snapshots",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Snapshot.html`,
		Resolver:    fetchEc2EbsSnapshots,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
		Transform:   transformers.TransformWithStruct(&types.Snapshot{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveEbsSnapshotArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "attribute",
				Type:     schema.TypeJSON,
				Resolver: resolveEbsSnapshotAttribute,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
