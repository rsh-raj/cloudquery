package iam

import (
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func VirtualMfaDevices() *schema.Table {
	return &schema.Table{
		Name:        "aws_iam_virtual_mfa_devices",
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_VirtualMFADevice.html`,
		Resolver:    fetchIamVirtualMfaDevices,
		Transform:   transformers.TransformWithStruct(&types.VirtualMFADevice{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name: "serial_number",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
