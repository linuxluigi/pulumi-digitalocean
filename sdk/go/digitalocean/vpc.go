// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a [DigitalOcean VPC](https://docs.digitalocean.com/reference/api/api-reference/#tag/VPCs) resource.
//
// VPCs are virtual networks containing resources that can communicate with each
// other in full isolation, using private IP addresses.
//
// ## Example Usage
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
// 		_, err := digitalocean.NewVpc(ctx, "example", &digitalocean.VpcArgs{
// 			IpRange: pulumi.String("10.10.10.0/24"),
// 			Region:  pulumi.String("nyc3"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Resource Assignment
//
// `Droplet`, `KubernetesCluster`,
// `digitaloceanLoadBalancer`, and `DatabaseCluster` resources
// may be assigned to a VPC by referencing its `id`. For example:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean/index"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleVpc, err := digitalocean.NewVpc(ctx, "exampleVpc", &digitalocean.VpcArgs{
// 			Region: pulumi.String("nyc3"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = digitalocean.NewDroplet(ctx, "exampleDroplet", &digitalocean.DropletArgs{
// 			Size:    pulumi.String("s-1vcpu-1gb"),
// 			Image:   pulumi.String("ubuntu-18-04-x64"),
// 			Region:  pulumi.String("nyc3"),
// 			VpcUuid: exampleVpc.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// A VPC can be imported using its `id`, e.g.
//
// ```sh
//  $ pulumi import digitalocean:index/vpc:Vpc example 506f78a4-e098-11e5-ad9f-000f53306ae1
// ```
type Vpc struct {
	pulumi.CustomResourceState

	// The date and time of when the VPC was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// A boolean indicating whether or not the VPC is the default one for the region.
	Default pulumi.BoolOutput `pulumi:"default"`
	// A free-form text field up to a limit of 255 characters to describe the VPC.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The range of IP addresses for the VPC in CIDR notation. Network ranges cannot overlap with other networks in the same account and must be in range of private addresses as defined in RFC1918. It may not be larger than `/16` or smaller than `/24`.
	IpRange pulumi.StringOutput `pulumi:"ipRange"`
	// A name for the VPC. Must be unique and contain alphanumeric characters, dashes, and periods only.
	Name pulumi.StringOutput `pulumi:"name"`
	// The DigitalOcean region slug for the VPC's location.
	Region pulumi.StringOutput `pulumi:"region"`
	// The uniform resource name (URN) for the VPC.
	VpcUrn pulumi.StringOutput `pulumi:"vpcUrn"`
}

// NewVpc registers a new resource with the given unique name, arguments, and options.
func NewVpc(ctx *pulumi.Context,
	name string, args *VpcArgs, opts ...pulumi.ResourceOption) (*Vpc, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource Vpc
	err := ctx.RegisterResource("digitalocean:index/vpc:Vpc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpc gets an existing Vpc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcState, opts ...pulumi.ResourceOption) (*Vpc, error) {
	var resource Vpc
	err := ctx.ReadResource("digitalocean:index/vpc:Vpc", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vpc resources.
type vpcState struct {
	// The date and time of when the VPC was created.
	CreatedAt *string `pulumi:"createdAt"`
	// A boolean indicating whether or not the VPC is the default one for the region.
	Default *bool `pulumi:"default"`
	// A free-form text field up to a limit of 255 characters to describe the VPC.
	Description *string `pulumi:"description"`
	// The range of IP addresses for the VPC in CIDR notation. Network ranges cannot overlap with other networks in the same account and must be in range of private addresses as defined in RFC1918. It may not be larger than `/16` or smaller than `/24`.
	IpRange *string `pulumi:"ipRange"`
	// A name for the VPC. Must be unique and contain alphanumeric characters, dashes, and periods only.
	Name *string `pulumi:"name"`
	// The DigitalOcean region slug for the VPC's location.
	Region *string `pulumi:"region"`
	// The uniform resource name (URN) for the VPC.
	VpcUrn *string `pulumi:"vpcUrn"`
}

type VpcState struct {
	// The date and time of when the VPC was created.
	CreatedAt pulumi.StringPtrInput
	// A boolean indicating whether or not the VPC is the default one for the region.
	Default pulumi.BoolPtrInput
	// A free-form text field up to a limit of 255 characters to describe the VPC.
	Description pulumi.StringPtrInput
	// The range of IP addresses for the VPC in CIDR notation. Network ranges cannot overlap with other networks in the same account and must be in range of private addresses as defined in RFC1918. It may not be larger than `/16` or smaller than `/24`.
	IpRange pulumi.StringPtrInput
	// A name for the VPC. Must be unique and contain alphanumeric characters, dashes, and periods only.
	Name pulumi.StringPtrInput
	// The DigitalOcean region slug for the VPC's location.
	Region pulumi.StringPtrInput
	// The uniform resource name (URN) for the VPC.
	VpcUrn pulumi.StringPtrInput
}

func (VpcState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcState)(nil)).Elem()
}

type vpcArgs struct {
	// A free-form text field up to a limit of 255 characters to describe the VPC.
	Description *string `pulumi:"description"`
	// The range of IP addresses for the VPC in CIDR notation. Network ranges cannot overlap with other networks in the same account and must be in range of private addresses as defined in RFC1918. It may not be larger than `/16` or smaller than `/24`.
	IpRange *string `pulumi:"ipRange"`
	// A name for the VPC. Must be unique and contain alphanumeric characters, dashes, and periods only.
	Name *string `pulumi:"name"`
	// The DigitalOcean region slug for the VPC's location.
	Region string `pulumi:"region"`
}

// The set of arguments for constructing a Vpc resource.
type VpcArgs struct {
	// A free-form text field up to a limit of 255 characters to describe the VPC.
	Description pulumi.StringPtrInput
	// The range of IP addresses for the VPC in CIDR notation. Network ranges cannot overlap with other networks in the same account and must be in range of private addresses as defined in RFC1918. It may not be larger than `/16` or smaller than `/24`.
	IpRange pulumi.StringPtrInput
	// A name for the VPC. Must be unique and contain alphanumeric characters, dashes, and periods only.
	Name pulumi.StringPtrInput
	// The DigitalOcean region slug for the VPC's location.
	Region pulumi.StringInput
}

func (VpcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcArgs)(nil)).Elem()
}

type VpcInput interface {
	pulumi.Input

	ToVpcOutput() VpcOutput
	ToVpcOutputWithContext(ctx context.Context) VpcOutput
}

func (*Vpc) ElementType() reflect.Type {
	return reflect.TypeOf((*Vpc)(nil))
}

func (i *Vpc) ToVpcOutput() VpcOutput {
	return i.ToVpcOutputWithContext(context.Background())
}

func (i *Vpc) ToVpcOutputWithContext(ctx context.Context) VpcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcOutput)
}

func (i *Vpc) ToVpcPtrOutput() VpcPtrOutput {
	return i.ToVpcPtrOutputWithContext(context.Background())
}

func (i *Vpc) ToVpcPtrOutputWithContext(ctx context.Context) VpcPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPtrOutput)
}

type VpcPtrInput interface {
	pulumi.Input

	ToVpcPtrOutput() VpcPtrOutput
	ToVpcPtrOutputWithContext(ctx context.Context) VpcPtrOutput
}

type vpcPtrType VpcArgs

func (*vpcPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Vpc)(nil))
}

func (i *vpcPtrType) ToVpcPtrOutput() VpcPtrOutput {
	return i.ToVpcPtrOutputWithContext(context.Background())
}

func (i *vpcPtrType) ToVpcPtrOutputWithContext(ctx context.Context) VpcPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPtrOutput)
}

// VpcArrayInput is an input type that accepts VpcArray and VpcArrayOutput values.
// You can construct a concrete instance of `VpcArrayInput` via:
//
//          VpcArray{ VpcArgs{...} }
type VpcArrayInput interface {
	pulumi.Input

	ToVpcArrayOutput() VpcArrayOutput
	ToVpcArrayOutputWithContext(context.Context) VpcArrayOutput
}

type VpcArray []VpcInput

func (VpcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vpc)(nil)).Elem()
}

func (i VpcArray) ToVpcArrayOutput() VpcArrayOutput {
	return i.ToVpcArrayOutputWithContext(context.Background())
}

func (i VpcArray) ToVpcArrayOutputWithContext(ctx context.Context) VpcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcArrayOutput)
}

// VpcMapInput is an input type that accepts VpcMap and VpcMapOutput values.
// You can construct a concrete instance of `VpcMapInput` via:
//
//          VpcMap{ "key": VpcArgs{...} }
type VpcMapInput interface {
	pulumi.Input

	ToVpcMapOutput() VpcMapOutput
	ToVpcMapOutputWithContext(context.Context) VpcMapOutput
}

type VpcMap map[string]VpcInput

func (VpcMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vpc)(nil)).Elem()
}

func (i VpcMap) ToVpcMapOutput() VpcMapOutput {
	return i.ToVpcMapOutputWithContext(context.Background())
}

func (i VpcMap) ToVpcMapOutputWithContext(ctx context.Context) VpcMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcMapOutput)
}

type VpcOutput struct{ *pulumi.OutputState }

func (VpcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Vpc)(nil))
}

func (o VpcOutput) ToVpcOutput() VpcOutput {
	return o
}

func (o VpcOutput) ToVpcOutputWithContext(ctx context.Context) VpcOutput {
	return o
}

func (o VpcOutput) ToVpcPtrOutput() VpcPtrOutput {
	return o.ToVpcPtrOutputWithContext(context.Background())
}

func (o VpcOutput) ToVpcPtrOutputWithContext(ctx context.Context) VpcPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Vpc) *Vpc {
		return &v
	}).(VpcPtrOutput)
}

type VpcPtrOutput struct{ *pulumi.OutputState }

func (VpcPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vpc)(nil))
}

func (o VpcPtrOutput) ToVpcPtrOutput() VpcPtrOutput {
	return o
}

func (o VpcPtrOutput) ToVpcPtrOutputWithContext(ctx context.Context) VpcPtrOutput {
	return o
}

func (o VpcPtrOutput) Elem() VpcOutput {
	return o.ApplyT(func(v *Vpc) Vpc {
		if v != nil {
			return *v
		}
		var ret Vpc
		return ret
	}).(VpcOutput)
}

type VpcArrayOutput struct{ *pulumi.OutputState }

func (VpcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Vpc)(nil))
}

func (o VpcArrayOutput) ToVpcArrayOutput() VpcArrayOutput {
	return o
}

func (o VpcArrayOutput) ToVpcArrayOutputWithContext(ctx context.Context) VpcArrayOutput {
	return o
}

func (o VpcArrayOutput) Index(i pulumi.IntInput) VpcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Vpc {
		return vs[0].([]Vpc)[vs[1].(int)]
	}).(VpcOutput)
}

type VpcMapOutput struct{ *pulumi.OutputState }

func (VpcMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Vpc)(nil))
}

func (o VpcMapOutput) ToVpcMapOutput() VpcMapOutput {
	return o
}

func (o VpcMapOutput) ToVpcMapOutputWithContext(ctx context.Context) VpcMapOutput {
	return o
}

func (o VpcMapOutput) MapIndex(k pulumi.StringInput) VpcOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Vpc {
		return vs[0].(map[string]Vpc)[vs[1].(string)]
	}).(VpcOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcInput)(nil)).Elem(), &Vpc{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPtrInput)(nil)).Elem(), &Vpc{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcArrayInput)(nil)).Elem(), VpcArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcMapInput)(nil)).Elem(), VpcMap{})
	pulumi.RegisterOutputType(VpcOutput{})
	pulumi.RegisterOutputType(VpcPtrOutput{})
	pulumi.RegisterOutputType(VpcArrayOutput{})
	pulumi.RegisterOutputType(VpcMapOutput{})
}
