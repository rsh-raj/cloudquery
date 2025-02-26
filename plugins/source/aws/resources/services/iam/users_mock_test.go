package iam

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client"
	"github.com/rsh-raj/cloudquery/plugins/source/aws/client/mocks"
)

func buildIamUsers(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	u := types.User{}
	err := faker.FakeObject(&u)
	if err != nil {
		t.Fatal(err)
	}
	g := types.Group{}
	err = faker.FakeObject(&g)
	if err != nil {
		t.Fatal(err)
	}
	km := types.AccessKeyMetadata{}
	err = faker.FakeObject(&km)
	if err != nil {
		t.Fatal(err)
	}
	aup := types.AttachedPolicy{}
	err = faker.FakeObject(&aup)
	if err != nil {
		t.Fatal(err)
	}
	akl := iam.GetAccessKeyLastUsedOutput{}
	err = faker.FakeObject(&akl)
	if err != nil {
		t.Fatal(err)
	}

	sshPublicKey := types.SSHPublicKeyMetadata{}
	err = faker.FakeObject(&sshPublicKey)
	if err != nil {
		t.Fatal(err)
	}

	var tags []types.Tag
	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListUsers(gomock.Any(), gomock.Any()).Return(
		&iam.ListUsersOutput{
			Users: []types.User{u},
		}, nil)
	m.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(
		&iam.GetUserOutput{
			User: &u,
		}, nil)
	m.EXPECT().ListGroupsForUser(gomock.Any(), gomock.Any()).Return(
		&iam.ListGroupsForUserOutput{
			Groups: []types.Group{g},
		}, nil)
	m.EXPECT().ListAccessKeys(gomock.Any(), gomock.Any()).Return(
		&iam.ListAccessKeysOutput{
			AccessKeyMetadata: []types.AccessKeyMetadata{km},
		}, nil)
	m.EXPECT().ListAttachedUserPolicies(gomock.Any(), gomock.Any()).Return(
		&iam.ListAttachedUserPoliciesOutput{
			AttachedPolicies: []types.AttachedPolicy{aup},
		}, nil)
	m.EXPECT().GetAccessKeyLastUsed(gomock.Any(), gomock.Any()).Return(
		&akl, nil)

	//list user inline policies
	var l []string
	err = faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListUserPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListUserPoliciesOutput{
			PolicyNames: l,
		}, nil)

	//get policy
	p := iam.GetUserPolicyOutput{}
	err = faker.FakeObject(&p)
	if err != nil {
		t.Fatal(err)
	}
	document := "{\"test\": {\"t1\":1}}"
	p.PolicyDocument = &document
	m.EXPECT().GetUserPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&p, nil)

	m.EXPECT().ListSSHPublicKeys(gomock.Any(), gomock.Any()).Return(
		&iam.ListSSHPublicKeysOutput{
			SSHPublicKeys: []types.SSHPublicKeyMetadata{sshPublicKey},
		}, nil)

	// get signing key

	sc := types.SigningCertificate{}
	err = faker.FakeObject(&sc)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListSigningCertificates(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListSigningCertificatesOutput{
			Certificates: []types.SigningCertificate{sc},
		}, nil)

	return client.Services{
		Iam: m,
	}
}

func TestIamUsers(t *testing.T) {
	client.AwsMockTestHelper(t, Users(), buildIamUsers, client.TestOptions{})
}
