// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Get information on a load balancer for use in other resources. This data source
// provides all of the load balancers properties as configured on your DigitalOcean
// account. This is useful if the load balancer in question is not managed by
// the provider or you need to utilize any of the load balancers data.
//
// An error is triggered if the provided load balancer name does not exist.
//
// ## Example Usage
//
// Get the load balancer:
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
// 		example, err := digitalocean.LookupLoadBalancer(ctx, &digitalocean.LookupLoadBalancerArgs{
// 			Name: "app",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("lbOutput", example.Ip)
// 		return nil
// 	})
// }
// ```
func LookupLoadBalancer(ctx *pulumi.Context, args *LookupLoadBalancerArgs, opts ...pulumi.InvokeOption) (*LookupLoadBalancerResult, error) {
	var rv LookupLoadBalancerResult
	err := ctx.Invoke("digitalocean:index/getLoadBalancer:getLoadBalancer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLoadBalancer.
type LookupLoadBalancerArgs struct {
	// The name of load balancer.
	Name string `pulumi:"name"`
}

// A collection of values returned by getLoadBalancer.
type LookupLoadBalancerResult struct {
	Algorithm              string                          `pulumi:"algorithm"`
	DropletIds             []int                           `pulumi:"dropletIds"`
	DropletTag             string                          `pulumi:"dropletTag"`
	EnableBackendKeepalive bool                            `pulumi:"enableBackendKeepalive"`
	EnableProxyProtocol    bool                            `pulumi:"enableProxyProtocol"`
	ForwardingRules        []GetLoadBalancerForwardingRule `pulumi:"forwardingRules"`
	Healthchecks           []GetLoadBalancerHealthcheck    `pulumi:"healthchecks"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string                         `pulumi:"id"`
	Ip                  string                         `pulumi:"ip"`
	LoadBalancerUrn     string                         `pulumi:"loadBalancerUrn"`
	Name                string                         `pulumi:"name"`
	RedirectHttpToHttps bool                           `pulumi:"redirectHttpToHttps"`
	Region              string                         `pulumi:"region"`
	Size                string                         `pulumi:"size"`
	Status              string                         `pulumi:"status"`
	StickySessions      []GetLoadBalancerStickySession `pulumi:"stickySessions"`
	VpcUuid             string                         `pulumi:"vpcUuid"`
}
