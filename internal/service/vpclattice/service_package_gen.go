// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package vpclattice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/vpclattice"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory:  newResourceConfigurationResource,
			TypeName: "aws_vpclattice_resource_configuration",
			Name:     "Resource Configuration",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  newResourceGatewayResource,
			TypeName: "aws_vpclattice_resource_gateway",
			Name:     "Resource Gateway",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  newServiceNetworkResourceAssociationResource,
			TypeName: "aws_vpclattice_service_network_resource_association",
			Name:     "Service Network Resource Association",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceAuthPolicy,
			TypeName: "aws_vpclattice_auth_policy",
			Name:     "Auth Policy",
		},
		{
			Factory:  DataSourceListener,
			TypeName: "aws_vpclattice_listener",
			Name:     "Listener",
		},
		{
			Factory:  DataSourceResourcePolicy,
			TypeName: "aws_vpclattice_resource_policy",
			Name:     "Resource Policy",
		},
		{
			Factory:  dataSourceService,
			TypeName: "aws_vpclattice_service",
			Name:     "Service",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  dataSourceServiceNetwork,
			TypeName: "aws_vpclattice_service_network",
			Name:     "Service Network",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAccessLogSubscription,
			TypeName: "aws_vpclattice_access_log_subscription",
			Name:     "Access Log Subscription",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceAuthPolicy,
			TypeName: "aws_vpclattice_auth_policy",
			Name:     "Auth Policy",
		},
		{
			Factory:  resourceListener,
			TypeName: "aws_vpclattice_listener",
			Name:     "Listener",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceListenerRule,
			TypeName: "aws_vpclattice_listener_rule",
			Name:     "Listener Rule",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceResourcePolicy,
			TypeName: "aws_vpclattice_resource_policy",
			Name:     "Resource Policy",
		},
		{
			Factory:  resourceService,
			TypeName: "aws_vpclattice_service",
			Name:     "Service",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceServiceNetwork,
			TypeName: "aws_vpclattice_service_network",
			Name:     "Service Network",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceServiceNetworkServiceAssociation,
			TypeName: "aws_vpclattice_service_network_service_association",
			Name:     "Service Network Service Association",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceServiceNetworkVPCAssociation,
			TypeName: "aws_vpclattice_service_network_vpc_association",
			Name:     "Service Network VPC Association",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceTargetGroup,
			TypeName: "aws_vpclattice_target_group",
			Name:     "Target Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceTargetGroupAttachment,
			TypeName: "aws_vpclattice_target_group_attachment",
			Name:     "Target Group Attachment",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.VPCLattice
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*vpclattice.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))

	return vpclattice.NewFromConfig(cfg,
		vpclattice.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
