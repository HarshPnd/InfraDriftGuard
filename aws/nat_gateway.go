package aws

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var ngwAllowEmptyValues = []string{"tags."}

type NatGatewayGenerator struct {
	AWSService
}

func (g *NatGatewayGenerator) createResources(ngws *ec2.DescribeNatGatewaysOutput) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, ngw := range ngws.NatGateways {
		resources = append(resources, terraformutils.NewSimpleResource(
			StringValue(ngw.NatGatewayId),
			StringValue(ngw.NatGatewayId),
			"aws_nat_gateway",
			"aws",
			ngwAllowEmptyValues,
		))
	}

	return resources
}

// Generate TerraformResources from AWS API,
// create terraform resource for each NAT Gateways
func (g *NatGatewayGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := ec2.NewFromConfig(config)
	p := ec2.NewDescribeNatGatewaysPaginator(svc, &ec2.DescribeNatGatewaysInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		g.Resources = append(g.Resources, g.createResources(page)...)
	}
	return nil
}
