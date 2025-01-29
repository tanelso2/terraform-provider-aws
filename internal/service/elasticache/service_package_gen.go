// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory:  newDataSourceReservedCacheNodeOffering,
			TypeName: "aws_elasticache_reserved_cache_node_offering",
			Name:     "Reserved Cache Node Offering",
		},
		{
			Factory:  newDataSourceServerlessCache,
			TypeName: "aws_elasticache_serverless_cache",
			Name:     "Serverless Cache",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory:  newResourceReservedCacheNode,
			TypeName: "aws_elasticache_reserved_cache_node",
			Name:     "Reserved Cache Node",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  newServerlessCacheResource,
			TypeName: "aws_elasticache_serverless_cache",
			Name:     "Serverless Cache",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceCluster,
			TypeName: "aws_elasticache_cluster",
			Name:     "Cluster",
		},
		{
			Factory:  dataSourceReplicationGroup,
			TypeName: "aws_elasticache_replication_group",
			Name:     "Replication Group",
		},
		{
			Factory:  dataSourceSubnetGroup,
			TypeName: "aws_elasticache_subnet_group",
			Name:     "Subnet Group",
		},
		{
			Factory:  dataSourceUser,
			TypeName: "aws_elasticache_user",
			Name:     "User",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceCluster,
			TypeName: "aws_elasticache_cluster",
			Name:     "Cluster",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceGlobalReplicationGroup,
			TypeName: "aws_elasticache_global_replication_group",
			Name:     "Global Replication Group",
		},
		{
			Factory:  resourceParameterGroup,
			TypeName: "aws_elasticache_parameter_group",
			Name:     "Parameter Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceReplicationGroup,
			TypeName: "aws_elasticache_replication_group",
			Name:     "Replication Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceSubnetGroup,
			TypeName: "aws_elasticache_subnet_group",
			Name:     "Subnet Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceUser,
			TypeName: "aws_elasticache_user",
			Name:     "User",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceUserGroup,
			TypeName: "aws_elasticache_user_group",
			Name:     "User Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceUserGroupAssociation,
			TypeName: "aws_elasticache_user_group_association",
			Name:     "User Group Association",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ElastiCache
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*elasticache.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*elasticache.Options){
		elasticache.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return elasticache.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*elasticache.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*elasticache.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *elasticache.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*elasticache.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
