package iot

import (
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func SecurityProfiles() *schema.Table {
	return &schema.Table{
		Name:        "aws_iot_security_profiles",
		Description: `https://docs.aws.amazon.com/iot/latest/apireference/API_DescribeSecurityProfile.html`,
		Resolver:    fetchIotSecurityProfiles,
		Transform:   transformers.TransformWithStruct(&iot.DescribeSecurityProfileOutput{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("iot"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "targets",
				Type:     schema.TypeStringArray,
				Resolver: ResolveIotSecurityProfileTargets,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: ResolveIotSecurityProfileTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecurityProfileArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
