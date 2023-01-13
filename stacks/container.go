package stacks

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/algolucky/algocdk/config"

	"github.com/cdktf/cdktf-provider-docker-go/docker/v4/container"
	"github.com/cdktf/cdktf-provider-docker-go/docker/v4/image"
	dockerprovider "github.com/cdktf/cdktf-provider-docker-go/docker/v4/provider"
	"github.com/cdktf/cdktf-provider-random-go/random/v4/pet"
	randomprovider "github.com/cdktf/cdktf-provider-random-go/random/v4/provider"
)

// defaults

var (
	defaultToken         string  = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	defaultContainerRepo string  = "docker.io/algorand/algod"
	defaultContainerTag  string  = "latest"
	defaultAlgodPort     float64 = 18080
	defaultKmdPort       float64 = 17833
)

// ContainerStack deploys a single container
func ContainerStackSimple(stack cdktf.TerraformStack, config config.ContainerConfig) cdktf.TerraformStack {
	dockerprovider.NewDockerProvider(stack, jsii.String("docker"), &dockerprovider.DockerProviderConfig{})
	randomprovider.NewRandomProvider(stack, jsii.String("random"), &randomprovider.RandomProviderConfig{})

	algodImage := image.NewImage(stack, jsii.String("algodImage"), &image.ImageConfig{
		Name: jsii.String(config.ContainerRepo + ":" + config.ContainerTag),
	})

	pet := pet.NewPet(stack, jsii.String("slug"), &pet.PetConfig{
		Length: jsii.Number(2),
	})

	containerName := "algod-" + *pet.Id()

	if config.ContainerRepo == "" {
		config.ContainerRepo = defaultContainerRepo
	}

	if config.ContainerTag == "" {
		config.ContainerTag = defaultContainerTag
	}

	containerEnv := []*string{}

	containerPorts := []*container.ContainerPorts{{
		Internal: jsii.Number(8080), External: jsii.Number(config.AlgodPort),
	}}

	// configure algod token
	if config.Token == "" {
		config.Token = defaultToken
	}
	tokenEnvString := "TOKEN=" + config.Token
	containerEnv = append(containerEnv, &tokenEnvString)

	// configure admin token
	if config.AdminToken == "" {
		config.AdminToken = config.Token
	}
	adminTokenEnvString := "ADMIN_TOKEN=" + config.AdminToken
	containerEnv = append(containerEnv, &adminTokenEnvString)

	// configure kmd
	if config.StartKMD {
		startKmdEnvString := "START_KMD=1"
		containerEnv = append(containerEnv, &startKmdEnvString)
		kmdContainerPort := container.ContainerPorts{
			Internal: jsii.Number(7833), External: jsii.Number(config.KmdPort),
		}
		containerPorts = append(containerPorts, &kmdContainerPort)

		if config.KmdToken == "" {
			config.KmdToken = config.AdminToken
			kmdTokenEnvString := "KMD_TOKEN=" + config.KmdToken
			containerEnv = append(containerEnv, &kmdTokenEnvString)
		}
	}

	if config.Network != "" {
		networkEnvString := "NETWORK=" + config.Network
		containerEnv = append(containerEnv, &networkEnvString)
	}

	if config.FastCatchup {
		fastCatchupEnvString := "FAST_CATCHUP=1"
		containerEnv = append(containerEnv, &fastCatchupEnvString)
	}

	algodContainer := container.NewContainer(stack, jsii.String("algodContainer"), &container.ContainerConfig{
		Image: algodImage.Latest(),
		Name:  jsii.String(containerName),
		Env:   &containerEnv,
		Ports: containerPorts,
	})

	// outputs

	cdktf.NewTerraformOutput(stack, jsii.String("id"), &cdktf.TerraformOutputConfig{
		Value: algodContainer.Id(),
	})

	cdktf.NewTerraformOutput(stack, jsii.String("name"), &cdktf.TerraformOutputConfig{
		Value: algodContainer.Name(),
	})

	cdktf.NewTerraformOutput(stack, jsii.String("algod_token"), &cdktf.TerraformOutputConfig{
		Value: config.Token,
	})

	cdktf.NewTerraformOutput(stack, jsii.String("admin_token"), &cdktf.TerraformOutputConfig{
		Value:     config.AdminToken,
		Sensitive: jsii.Bool(true),
	})

	cdktf.NewTerraformOutput(stack, jsii.String("kmd_token"), &cdktf.TerraformOutputConfig{
		Value:     config.KmdToken,
		Sensitive: jsii.Bool(true),
	})

	return stack
}
