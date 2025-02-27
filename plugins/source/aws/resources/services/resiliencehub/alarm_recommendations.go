package resiliencehub

import (
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func alarmRecommendations() *schema.Table {
	return &schema.Table{
		Name:        "aws_resiliencehub_alarm_recommendations",
		Description: `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_AlarmRecommendation.html`,
		Resolver:    fetchAlarmRecommendations,
		Transform:   transformers.TransformWithStruct(&types.AlarmRecommendation{}, transformers.WithPrimaryKeys("RecommendationId")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("resiliencehub"),
		Columns:     []schema.Column{client.DefaultAccountIDColumn(false), client.DefaultRegionColumn(false), appARN, assessmentARN},
	}
}
