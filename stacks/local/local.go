package local

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/cdktf/cdktf-provider-docker-go/docker/v4/container"
	"github.com/cdktf/cdktf-provider-docker-go/docker/v4/image"
	dockerprovider "github.com/cdktf/cdktf-provider-docker-go/docker/v4/provider"
)

var LocalStackConfig struct {
	ContainerRepo string
	ContainerTag  string
	AlgodPort     string
}

func LocalStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	dockerprovider.NewDockerProvider(stack, jsii.String("docker"), &dockerprovider.DockerProviderConfig{})

	algodImage := image.NewImage(stack, jsii.String("algodImage"), &image.ImageConfig{
		Name: jsii.String("algolucky/algod:container-update-docs"),
	})

	container.NewContainer(stack, jsii.String("algodContainer"), &container.ContainerConfig{
		Image: algodImage.Latest(),
		Name:  jsii.String("algod"),
		Ports: &[]*container.ContainerPorts{{
			Internal: jsii.Number(8080), External: jsii.Number(18080),
		}},
	})

	return stack
}
