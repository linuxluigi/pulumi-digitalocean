// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Get information on Droplets for use in other resources, with the ability to filter and sort the results.
// If no filters are specified, all Droplets will be returned.
//
// This data source is useful if the Droplets in question are not managed by the provider or you need to
// utilize any of the Droplets' data.
//
// Note: You can use the `Droplet` data source to obtain metadata
// about a single Droplet if you already know the `id`, unique `name`, or unique `tag` to retrieve.
//
// ## Example Usage
//
// Use the `filter` block with a `key` string and `values` list to filter images.
//
// For example to find all Droplets with size `s-1vcpu-1gb`:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v3/go/digitalocean"
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v3/go/digitalocean/"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := digitalocean.GetDroplets(ctx, &digitalocean.GetDropletsArgs{
// 			Filters: []digitalocean.GetDropletsFilter{
// 				digitalocean.GetDropletsFilter{
// 					Key: "size",
// 					Values: []string{
// 						"s-1vcpu-1gb",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// You can filter on multiple fields and sort the results as well:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v3/go/digitalocean"
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v3/go/digitalocean/"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := digitalocean.GetDroplets(ctx, &digitalocean.GetDropletsArgs{
// 			Filters: []digitalocean.GetDropletsFilter{
// 				digitalocean.GetDropletsFilter{
// 					Key: "size",
// 					Values: []string{
// 						"s-1vcpu-1gb",
// 					},
// 				},
// 				digitalocean.GetDropletsFilter{
// 					Key: "backups",
// 					Values: []string{
// 						"true",
// 					},
// 				},
// 			},
// 			Sorts: []digitalocean.GetDropletsSort{
// 				digitalocean.GetDropletsSort{
// 					Direction: "desc",
// 					Key:       "created_at",
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDroplets(ctx *pulumi.Context, args *GetDropletsArgs, opts ...pulumi.InvokeOption) (*GetDropletsResult, error) {
	var rv GetDropletsResult
	err := ctx.Invoke("digitalocean:index/getDroplets:getDroplets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDroplets.
type GetDropletsArgs struct {
	// Filter the results.
	// The `filter` block is documented below.
	Filters []GetDropletsFilter `pulumi:"filters"`
	// Sort the results.
	// The `sort` block is documented below.
	Sorts []GetDropletsSort `pulumi:"sorts"`
}

// A collection of values returned by getDroplets.
type GetDropletsResult struct {
	// A list of Droplets satisfying any `filter` and `sort` criteria. Each Droplet has the following attributes:
	Droplets []GetDropletsDroplet `pulumi:"droplets"`
	Filters  []GetDropletsFilter  `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id    string            `pulumi:"id"`
	Sorts []GetDropletsSort `pulumi:"sorts"`
}
