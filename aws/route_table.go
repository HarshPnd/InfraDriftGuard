package aws

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

var rtbAllowEmptyValues = []string{"tags."}

type RouteTableGenerator struct {
	AWSService
}

func (g *RouteTableGenerator) createRouteTablesResources(svc *ec2.Client) []terraformutils.Resource {
	var resources []terraformutils.Resource
	p := ec2.NewDescribeRouteTablesPaginator(svc, &ec2.DescribeRouteTablesInput{})
	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			log.Println(err)
			return resources
		}
		for _, table := range page.RouteTables {
			// route table
			resources = append(resources, terraformutils.NewSimpleResource(
				StringValue(table.RouteTableId),
				StringValue(table.RouteTableId),
				"aws_route_table",
				"aws",
				rtbAllowEmptyValues,
			))

			for _, assoc := range table.Associations {
				if *assoc.Main {
					// main route table association
					resources = append(resources, terraformutils.NewResource(
						StringValue(assoc.RouteTableAssociationId),
						StringValue(table.VpcId),
						"aws_main_route_table_association",
						"aws",
						map[string]string{
							"vpc_id":         StringValue(table.VpcId),
							"route_table_id": StringValue(table.RouteTableId),
						},
						rtbAllowEmptyValues,
						map[string]interface{}{},
					))
				} else if v := assoc.SubnetId; v != nil {
					// subnet-specific route table association
					resources = append(resources, terraformutils.NewResource(
						StringValue(assoc.RouteTableAssociationId),
						StringValue(v),
						"aws_route_table_association",
						"aws",
						map[string]string{
							"subnet_id":      StringValue(v),
							"route_table_id": StringValue(table.RouteTableId),
						},
						rtbAllowEmptyValues,
						map[string]interface{}{},
					))
				} else if v := assoc.GatewayId; v != nil {
					resources = append(resources, terraformutils.NewResource(
						StringValue(assoc.RouteTableAssociationId),
						StringValue(v),
						"aws_route_table_association",
						"aws",
						map[string]string{
							"gateway_id":     StringValue(v),
							"route_table_id": StringValue(table.RouteTableId),
						},
						rtbAllowEmptyValues,
						map[string]interface{}{},
					))

				}
			}
		}
	}
	return resources
}

// Generate TerraformResources from AWS API,
// create terraform resource for each route tables
func (g *RouteTableGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := ec2.NewFromConfig(config)

	g.Resources = g.createRouteTablesResources(svc)
	return nil
}
