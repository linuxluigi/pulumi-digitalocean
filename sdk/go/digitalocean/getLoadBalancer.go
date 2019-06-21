// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Get information on a load balancer for use in other resources. This data source
// provides all of the load balancers properties as configured on your DigitalOcean
// account. This is useful if the load balancer in question is not managed by
// Terraform or you need to utilize any of the load balancers data.
// 
// An error is triggered if the provided load balancer name does not exist.
func LookupLoadBalancer(ctx *pulumi.Context, args *GetLoadBalancerArgs) (*GetLoadBalancerResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
	}
	outputs, err := ctx.Invoke("digitalocean:index/getLoadBalancer:getLoadBalancer", inputs)
	if err != nil {
		return nil, err
	}
	return &GetLoadBalancerResult{
		Algorithm: outputs["algorithm"],
		DropletIds: outputs["dropletIds"],
		DropletTag: outputs["dropletTag"],
		EnableProxyProtocol: outputs["enableProxyProtocol"],
		ForwardingRules: outputs["forwardingRules"],
		Healthcheck: outputs["healthcheck"],
		Ip: outputs["ip"],
		Name: outputs["name"],
		RedirectHttpToHttps: outputs["redirectHttpToHttps"],
		Region: outputs["region"],
		Status: outputs["status"],
		StickySessions: outputs["stickySessions"],
		Urn: outputs["urn"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getLoadBalancer.
type GetLoadBalancerArgs struct {
	// The name of load balancer.
	Name interface{}
}

// A collection of values returned by getLoadBalancer.
type GetLoadBalancerResult struct {
	Algorithm interface{}
	DropletIds interface{}
	DropletTag interface{}
	EnableProxyProtocol interface{}
	ForwardingRules interface{}
	Healthcheck interface{}
	Ip interface{}
	Name interface{}
	RedirectHttpToHttps interface{}
	Region interface{}
	Status interface{}
	StickySessions interface{}
	Urn interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}