// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package signer

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
			Factory:  DataSourceSigningJob,
			TypeName: "aws_signer_signing_job",
		},
		{
			Factory:  DataSourceSigningProfile,
			TypeName: "aws_signer_signing_profile",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceSigningJob,
			TypeName: "aws_signer_signing_job",
		},
		{
			Factory:  ResourceSigningProfile,
			TypeName: "aws_signer_signing_profile",
		},
		{
			Factory:  ResourceSigningProfilePermission,
			TypeName: "aws_signer_signing_profile_permission",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Signer
}

var ServicePackage = &servicePackage{}
