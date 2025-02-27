package resiliencehub

import (
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func recommendationTemplates() *schema.Table {
	return &schema.Table{
		Name:        "aws_resiliencehub_recommendation_templates",
		Description: `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_RecommendationTemplate.html`,
		Resolver:    fetchRecommendationTemplates,
		Transform:   transformers.TransformWithStruct(&types.RecommendationTemplate{}, transformers.WithPrimaryKeys("AppArn", "AssessmentArn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("resiliencehub"),
		Columns:     []schema.Column{client.DefaultAccountIDColumn(false), client.DefaultRegionColumn(false), arnColumn("RecommendationTemplateArn")},
	}
}
