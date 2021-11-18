// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information on a container registry. This data source provides the name as
// configured on your DigitalOcean account. This is useful if the container
// registry name in question is not managed by this provider or you need validate if
// the container registry exists in the account.
//
// An error is triggered if the provided container registry name does not exist.
//
// ## Example Usage
// ### Basic Example
//
// Get the container registry:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := digitalocean.LookupContainerRegistry(ctx, &GetContainerRegistryArgs{
// 			Name: "example",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupContainerRegistry(ctx *pulumi.Context, args *LookupContainerRegistryArgs, opts ...pulumi.InvokeOption) (*LookupContainerRegistryResult, error) {
	var rv LookupContainerRegistryResult
	err := ctx.Invoke("digitalocean:index/getContainerRegistry:getContainerRegistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContainerRegistry.
type LookupContainerRegistryArgs struct {
	// The name of the container registry.
	Name string `pulumi:"name"`
}

// A collection of values returned by getContainerRegistry.
type LookupContainerRegistryResult struct {
	Endpoint string `pulumi:"endpoint"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the container registry
	Name      string `pulumi:"name"`
	ServerUrl string `pulumi:"serverUrl"`
	// The slug identifier for the subscription tier
	// * `endpoint`: The URL endpoint of the container registry. Ex: `registry.digitalocean.com/my_registry`
	// * `serverUrl`: The domain of the container registry. Ex: `registry.digitalocean.com`
	SubscriptionTierSlug string `pulumi:"subscriptionTierSlug"`
}

func LookupContainerRegistryOutput(ctx *pulumi.Context, args LookupContainerRegistryOutputArgs, opts ...pulumi.InvokeOption) LookupContainerRegistryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContainerRegistryResult, error) {
			args := v.(LookupContainerRegistryArgs)
			r, err := LookupContainerRegistry(ctx, &args, opts...)
			return *r, err
		}).(LookupContainerRegistryResultOutput)
}

// A collection of arguments for invoking getContainerRegistry.
type LookupContainerRegistryOutputArgs struct {
	// The name of the container registry.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupContainerRegistryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerRegistryArgs)(nil)).Elem()
}

// A collection of values returned by getContainerRegistry.
type LookupContainerRegistryResultOutput struct{ *pulumi.OutputState }

func (LookupContainerRegistryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerRegistryResult)(nil)).Elem()
}

func (o LookupContainerRegistryResultOutput) ToLookupContainerRegistryResultOutput() LookupContainerRegistryResultOutput {
	return o
}

func (o LookupContainerRegistryResultOutput) ToLookupContainerRegistryResultOutputWithContext(ctx context.Context) LookupContainerRegistryResultOutput {
	return o
}

func (o LookupContainerRegistryResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerRegistryResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupContainerRegistryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerRegistryResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the container registry
func (o LookupContainerRegistryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerRegistryResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupContainerRegistryResultOutput) ServerUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerRegistryResult) string { return v.ServerUrl }).(pulumi.StringOutput)
}

// The slug identifier for the subscription tier
// * `endpoint`: The URL endpoint of the container registry. Ex: `registry.digitalocean.com/my_registry`
// * `serverUrl`: The domain of the container registry. Ex: `registry.digitalocean.com`
func (o LookupContainerRegistryResultOutput) SubscriptionTierSlug() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerRegistryResult) string { return v.SubscriptionTierSlug }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerRegistryResultOutput{})
}
