// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// > **Deprecated:** DigitalOcean Floating IPs have been renamed reserved IPs. This data source will be removed in a future release. Please use `ReservedIp` instead.
//
// Get information on a floating ip. This data source provides the region and Droplet id
// as configured on your DigitalOcean account. This is useful if the floating IP
// in question is not managed by the provider or you need to find the Droplet the IP is
// attached to.
//
// An error is triggered if the provided floating IP does not exist.
//
// ## Example Usage
//
// Get the floating IP:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			publicIp := cfg.RequireObject("publicIp")
//			example, err := digitalocean.LookupFloatingIp(ctx, &digitalocean.LookupFloatingIpArgs{
//				IpAddress: publicIp,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("fipOutput", example.DropletId)
//			return nil
//		})
//	}
//
// ```
func LookupFloatingIp(ctx *pulumi.Context, args *LookupFloatingIpArgs, opts ...pulumi.InvokeOption) (*LookupFloatingIpResult, error) {
	var rv LookupFloatingIpResult
	err := ctx.Invoke("digitalocean:index/getFloatingIp:getFloatingIp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFloatingIp.
type LookupFloatingIpArgs struct {
	// The allocated IP address of the specific floating IP to retrieve.
	IpAddress string `pulumi:"ipAddress"`
}

// A collection of values returned by getFloatingIp.
type LookupFloatingIpResult struct {
	// The Droplet id that the floating IP has been assigned to.
	DropletId int `pulumi:"dropletId"`
	// The uniform resource name of the floating IP.
	FloatingIpUrn string `pulumi:"floatingIpUrn"`
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	IpAddress string `pulumi:"ipAddress"`
	// The region that the floating IP is reserved to.
	Region string `pulumi:"region"`
}

func LookupFloatingIpOutput(ctx *pulumi.Context, args LookupFloatingIpOutputArgs, opts ...pulumi.InvokeOption) LookupFloatingIpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFloatingIpResult, error) {
			args := v.(LookupFloatingIpArgs)
			r, err := LookupFloatingIp(ctx, &args, opts...)
			var s LookupFloatingIpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFloatingIpResultOutput)
}

// A collection of arguments for invoking getFloatingIp.
type LookupFloatingIpOutputArgs struct {
	// The allocated IP address of the specific floating IP to retrieve.
	IpAddress pulumi.StringInput `pulumi:"ipAddress"`
}

func (LookupFloatingIpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFloatingIpArgs)(nil)).Elem()
}

// A collection of values returned by getFloatingIp.
type LookupFloatingIpResultOutput struct{ *pulumi.OutputState }

func (LookupFloatingIpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFloatingIpResult)(nil)).Elem()
}

func (o LookupFloatingIpResultOutput) ToLookupFloatingIpResultOutput() LookupFloatingIpResultOutput {
	return o
}

func (o LookupFloatingIpResultOutput) ToLookupFloatingIpResultOutputWithContext(ctx context.Context) LookupFloatingIpResultOutput {
	return o
}

// The Droplet id that the floating IP has been assigned to.
func (o LookupFloatingIpResultOutput) DropletId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFloatingIpResult) int { return v.DropletId }).(pulumi.IntOutput)
}

// The uniform resource name of the floating IP.
func (o LookupFloatingIpResultOutput) FloatingIpUrn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFloatingIpResult) string { return v.FloatingIpUrn }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFloatingIpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFloatingIpResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFloatingIpResultOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFloatingIpResult) string { return v.IpAddress }).(pulumi.StringOutput)
}

// The region that the floating IP is reserved to.
func (o LookupFloatingIpResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFloatingIpResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFloatingIpResultOutput{})
}
