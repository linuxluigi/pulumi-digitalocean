// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information on a domain. This data source provides the name, TTL, and zone
// file as configured on your DigitalOcean account. This is useful if the domain
// name in question is not managed by this provider or you need to utilize TTL or zone
// file data.
//
// An error is triggered if the provided domain name is not managed with your
// DigitalOcean account.
func LookupDomain(ctx *pulumi.Context, args *LookupDomainArgs, opts ...pulumi.InvokeOption) (*LookupDomainResult, error) {
	var rv LookupDomainResult
	err := ctx.Invoke("digitalocean:index/getDomain:getDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDomain.
type LookupDomainArgs struct {
	// The name of the domain.
	Name string `pulumi:"name"`
}

// A collection of values returned by getDomain.
type LookupDomainResult struct {
	// The uniform resource name of the domain
	DomainUrn string `pulumi:"domainUrn"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The TTL of the domain.
	Ttl int `pulumi:"ttl"`
	// The zone file of the domain.
	ZoneFile string `pulumi:"zoneFile"`
}

func LookupDomainOutput(ctx *pulumi.Context, args LookupDomainOutputArgs, opts ...pulumi.InvokeOption) LookupDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDomainResult, error) {
			args := v.(LookupDomainArgs)
			r, err := LookupDomain(ctx, &args, opts...)
			var s LookupDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDomainResultOutput)
}

// A collection of arguments for invoking getDomain.
type LookupDomainOutputArgs struct {
	// The name of the domain.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainArgs)(nil)).Elem()
}

// A collection of values returned by getDomain.
type LookupDomainResultOutput struct{ *pulumi.OutputState }

func (LookupDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainResult)(nil)).Elem()
}

func (o LookupDomainResultOutput) ToLookupDomainResultOutput() LookupDomainResultOutput {
	return o
}

func (o LookupDomainResultOutput) ToLookupDomainResultOutputWithContext(ctx context.Context) LookupDomainResultOutput {
	return o
}

// The uniform resource name of the domain
func (o LookupDomainResultOutput) DomainUrn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.DomainUrn }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

// The TTL of the domain.
func (o LookupDomainResultOutput) Ttl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDomainResult) int { return v.Ttl }).(pulumi.IntOutput)
}

// The zone file of the domain.
func (o LookupDomainResultOutput) ZoneFile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.ZoneFile }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainResultOutput{})
}
