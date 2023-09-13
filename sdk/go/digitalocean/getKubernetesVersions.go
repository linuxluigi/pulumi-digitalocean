// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides access to the available DigitalOcean Kubernetes Service versions.
//
// ## Example Usage
// ### Output a list of all available versions
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := digitalocean.GetKubernetesVersions(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("k8s-versions", example.ValidVersions)
//			return nil
//		})
//	}
//
// ```
// ### Create a Kubernetes cluster using the most recent version available
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := digitalocean.GetKubernetesVersions(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = digitalocean.NewKubernetesCluster(ctx, "example-cluster", &digitalocean.KubernetesClusterArgs{
//				Region:  pulumi.String("lon1"),
//				Version: *pulumi.String(example.LatestVersion),
//				NodePool: &digitalocean.KubernetesClusterNodePoolArgs{
//					Name:      pulumi.String("default"),
//					Size:      pulumi.String("s-1vcpu-2gb"),
//					NodeCount: pulumi.Int(3),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Pin a Kubernetes cluster to a specific minor version
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := digitalocean.GetKubernetesVersions(ctx, &digitalocean.GetKubernetesVersionsArgs{
//				VersionPrefix: pulumi.StringRef("1.22."),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = digitalocean.NewKubernetesCluster(ctx, "example-cluster", &digitalocean.KubernetesClusterArgs{
//				Region:  pulumi.String("lon1"),
//				Version: *pulumi.String(example.LatestVersion),
//				NodePool: &digitalocean.KubernetesClusterNodePoolArgs{
//					Name:      pulumi.String("default"),
//					Size:      pulumi.String("s-1vcpu-2gb"),
//					NodeCount: pulumi.Int(3),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetKubernetesVersions(ctx *pulumi.Context, args *GetKubernetesVersionsArgs, opts ...pulumi.InvokeOption) (*GetKubernetesVersionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetKubernetesVersionsResult
	err := ctx.Invoke("digitalocean:index/getKubernetesVersions:getKubernetesVersions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKubernetesVersions.
type GetKubernetesVersionsArgs struct {
	// If provided, the provider will only return versions that match the string prefix. For example, `1.15.` will match all 1.15.x series releases.
	VersionPrefix *string `pulumi:"versionPrefix"`
}

// A collection of values returned by getKubernetesVersions.
type GetKubernetesVersionsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The most recent version available.
	LatestVersion string `pulumi:"latestVersion"`
	// A list of available versions.
	ValidVersions []string `pulumi:"validVersions"`
	VersionPrefix *string  `pulumi:"versionPrefix"`
}

func GetKubernetesVersionsOutput(ctx *pulumi.Context, args GetKubernetesVersionsOutputArgs, opts ...pulumi.InvokeOption) GetKubernetesVersionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetKubernetesVersionsResult, error) {
			args := v.(GetKubernetesVersionsArgs)
			r, err := GetKubernetesVersions(ctx, &args, opts...)
			var s GetKubernetesVersionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetKubernetesVersionsResultOutput)
}

// A collection of arguments for invoking getKubernetesVersions.
type GetKubernetesVersionsOutputArgs struct {
	// If provided, the provider will only return versions that match the string prefix. For example, `1.15.` will match all 1.15.x series releases.
	VersionPrefix pulumi.StringPtrInput `pulumi:"versionPrefix"`
}

func (GetKubernetesVersionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubernetesVersionsArgs)(nil)).Elem()
}

// A collection of values returned by getKubernetesVersions.
type GetKubernetesVersionsResultOutput struct{ *pulumi.OutputState }

func (GetKubernetesVersionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubernetesVersionsResult)(nil)).Elem()
}

func (o GetKubernetesVersionsResultOutput) ToGetKubernetesVersionsResultOutput() GetKubernetesVersionsResultOutput {
	return o
}

func (o GetKubernetesVersionsResultOutput) ToGetKubernetesVersionsResultOutputWithContext(ctx context.Context) GetKubernetesVersionsResultOutput {
	return o
}

func (o GetKubernetesVersionsResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetKubernetesVersionsResult] {
	return pulumix.Output[GetKubernetesVersionsResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetKubernetesVersionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubernetesVersionsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The most recent version available.
func (o GetKubernetesVersionsResultOutput) LatestVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubernetesVersionsResult) string { return v.LatestVersion }).(pulumi.StringOutput)
}

// A list of available versions.
func (o GetKubernetesVersionsResultOutput) ValidVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetKubernetesVersionsResult) []string { return v.ValidVersions }).(pulumi.StringArrayOutput)
}

func (o GetKubernetesVersionsResultOutput) VersionPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetKubernetesVersionsResult) *string { return v.VersionPrefix }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetKubernetesVersionsResultOutput{})
}
