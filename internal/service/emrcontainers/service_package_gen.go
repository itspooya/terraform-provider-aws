// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package emrcontainers

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	emrcontainers_sdkv2 "github.com/aws/aws-sdk-go-v2/service/emrcontainers"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceVirtualCluster,
			TypeName: "aws_emrcontainers_virtual_cluster",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceJobTemplate,
			TypeName: "aws_emrcontainers_job_template",
			Name:     "Job Template",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceVirtualCluster,
			TypeName: "aws_emrcontainers_virtual_cluster",
			Name:     "Virtual Cluster",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.EMRContainers
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*emrcontainers_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return emrcontainers_sdkv2.NewFromConfig(cfg,
		emrcontainers_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
