package resource

import (
	"fmt"
	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

//type LoadBalancer struct {
//	LoadBalancerType string `json:"loadBalancerType"`
//}

type Iam struct {
}

func (d *Deployment) createNewNewUser(
	ctx *pulumi.Context,
	region *Region,
) (*iam.User, error) {
	newUser, err := iam.NewUser(ctx, "lbUser", &iam.UserArgs{
		Path: pulumi.String("/system/"),
		Tags: pulumi.StringMap{
			"tag-key": pulumi.String("tag-value"),
		},
	})

	if err != nil {
		return nil, err
	}

	_, err = iam.NewAccessKey(ctx, "lbAccessKey", &iam.AccessKeyArgs{
		User: newUser.Name,
	})

	if err != nil {
		return nil, err
	}

	_, err = iam.NewUserPolicy(ctx, "lbRo", &iam.UserPolicyArgs{
		User:   newUser.Name,
		Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": [\n", "        \"ec2:Describe*\"\n", "      ],\n", "      \"Effect\": \"Allow\",\n", "      \"Resource\": \"*\"\n", "    }\n", "  ]\n", "}\n")),
	})

	if err != nil {
		return nil, err
	}

	return newUser, nil
}


