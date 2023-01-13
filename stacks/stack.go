package stacks

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// CreateStack creates a new cdktf stack and returns it
func CreateStack(scope constructs.Construct, id string) (stack cdktf.TerraformStack) {
	stack = cdktf.NewTerraformStack(scope, &id)
	return
}
