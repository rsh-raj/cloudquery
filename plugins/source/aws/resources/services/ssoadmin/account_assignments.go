package ssoadmin

import (
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
)

func AccountAssignments() *schema.Table {
	return &schema.Table{
		Name:        "aws_ssoadmin_account_assignments",
		Description: `https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_AccountAssignment.html`,
		Resolver:    fetchSsoadminAccountAssignments,
		Transform:   transformers.TransformWithStruct(&types.AccountAssignment{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("identitystore"),
	}
}
