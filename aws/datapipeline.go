package aws

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline"
)

var datapipelineAllowEmptyValues = []string{"tags."}

type DataPipelineGenerator struct {
	AWSService
}

func (g *DataPipelineGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := datapipeline.NewFromConfig(config)
	p := datapipeline.NewListPipelinesPaginator(svc, &datapipeline.ListPipelinesInput{})
	var resources []terraformutils.Resource
	for p.HasMorePages() {
		page, e := p.NextPage(context.TODO())
		if e != nil {
			return e
		}
		for _, pipeline := range page.PipelineIdList {
			pipelineID := StringValue(pipeline.Id)
			pipelineName := StringValue(pipeline.Name)
			resources = append(resources, terraformutils.NewSimpleResource(
				pipelineID,
				pipelineName,
				"aws_datapipeline_pipeline",
				"aws",
				datapipelineAllowEmptyValues))
		}
	}
	g.Resources = resources
	return nil
}
