// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information on a DNS record. This data source provides the name, TTL, and zone
// file as configured on your DigitalOcean account. This is useful if the record
// in question is not managed by the provider.
//
// An error is triggered if the provided domain name or record are not managed with
// your DigitalOcean account.
func GetRecord(ctx *pulumi.Context, args *GetRecordArgs, opts ...pulumi.InvokeOption) (*GetRecordResult, error) {
	var rv GetRecordResult
	err := ctx.Invoke("digitalocean:index/getRecord:getRecord", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRecord.
type GetRecordArgs struct {
	// The domain name of the record.
	Domain string `pulumi:"domain"`
	// The name of the record.
	Name string `pulumi:"name"`
}

// A collection of values returned by getRecord.
type GetRecordResult struct {
	Data   string `pulumi:"data"`
	Domain string `pulumi:"domain"`
	Flags  int    `pulumi:"flags"`
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	Name     string `pulumi:"name"`
	Port     int    `pulumi:"port"`
	Priority int    `pulumi:"priority"`
	Tag      string `pulumi:"tag"`
	Ttl      int    `pulumi:"ttl"`
	Type     string `pulumi:"type"`
	Weight   int    `pulumi:"weight"`
}
