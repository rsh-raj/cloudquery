package ecr

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/resources/services/ecr/models"
)

func RepositoryImageScanFindings() *schema.Table {
	return &schema.Table{
		Name:        "aws_ecr_repository_image_scan_findings",
		Description: `https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_ImageScanFindings.html`,
		Resolver:    fetchEcrRepositoryImageScanFindings,
		Transform:   transformers.TransformWithStruct(&models.ImageScanWrapper{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}
