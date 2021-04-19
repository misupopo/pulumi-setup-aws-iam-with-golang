package resource

import (
	"encoding/json"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"io/ioutil"
)

func Setup(ctx *pulumi.Context) error {
	config, _ := ioutil.ReadFile("./config.json")

	var r Region

	err := json.Unmarshal(config, &r)

	if err != nil {
		ctx.Export("Unmarshal error", pulumi.Printf("%v", err))
		return err
	}

	region := newRegion(r)

	deployment := new(Deployment)

	_, err = deployment.createNewNewUser(ctx, region)

	if err != nil {
		ctx.Export("createNewNewUser error", pulumi.Printf("%v", err))
		return err
	}

	return nil
}

