package efs

import (
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func Filesystems() *schema.Table {
	return &schema.Table{
		Name:        "aws_efs_filesystems",
		Description: `https://docs.aws.amazon.com/efs/latest/ug/API_FileSystemDescription.html`,
		Resolver:    fetchEfsFilesystems,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticfilesystem"),
		Transform:   transformers.TransformWithStruct(&types.FileSystemDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FileSystemArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "backup_policy_status",
				Type:     schema.TypeString,
				Resolver: ResolveEfsFilesystemBackupPolicyStatus,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
