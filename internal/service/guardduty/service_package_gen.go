// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package guardduty

import (
	"context"

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
			Factory:  DataSourceDetector,
			TypeName: "aws_guardduty_detector",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceDetector,
			TypeName: "aws_guardduty_detector",
		},
		{
			Factory:  ResourceFilter,
			TypeName: "aws_guardduty_filter",
		},
		{
			Factory:  ResourceInviteAccepter,
			TypeName: "aws_guardduty_invite_accepter",
		},
		{
			Factory:  ResourceIPSet,
			TypeName: "aws_guardduty_ipset",
		},
		{
			Factory:  ResourceMember,
			TypeName: "aws_guardduty_member",
		},
		{
			Factory:  ResourceOrganizationAdminAccount,
			TypeName: "aws_guardduty_organization_admin_account",
		},
		{
			Factory:  ResourceOrganizationConfiguration,
			TypeName: "aws_guardduty_organization_configuration",
		},
		{
			Factory:  ResourcePublishingDestination,
			TypeName: "aws_guardduty_publishing_destination",
		},
		{
			Factory:  ResourceThreatIntelSet,
			TypeName: "aws_guardduty_threatintelset",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.GuardDuty
}

var ServicePackage = &servicePackage{}
