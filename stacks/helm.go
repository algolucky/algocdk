package stacks

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	helmprovider "github.com/cdktf/cdktf-provider-helm-go/helm/v4/provider"
	"github.com/cdktf/cdktf-provider-helm-go/helm/v4/release"
)

func HelmModule(stack cdktf.TerraformStack, id string) cdktf.TerraformStack {
	helmprovider.NewHelmProvider(stack, jsii.String("helm"), &helmprovider.HelmProviderConfig{
		Kubernetes: &helmprovider.HelmProviderKubernetes{
			ConfigPath:    jsii.String(""),
			ConfigContext: jsii.String(""),
		},
	})

	releaseName := "algod-" + id

	release := release.NewRelease(stack, jsii.String(releaseName), &release.ReleaseConfig{
		Repository:      jsii.String(""),
		Chart:           jsii.String(""),
		Name:            jsii.String(""),
		Namespace:       jsii.String(""),
		CreateNamespace: jsii.Bool(true),
	})

	cdktf.NewTerraformOutput(stack, jsii.String("name"), &cdktf.TerraformOutputConfig{
		Value: release.Name(),
	})

	return stack
}
