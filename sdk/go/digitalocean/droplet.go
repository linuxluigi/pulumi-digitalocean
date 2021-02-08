// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a DigitalOcean Droplet resource. This can be used to create,
// modify, and delete Droplets.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v3/go/digitalocean/"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := digitalocean.NewDroplet(ctx, "web", &digitalocean.DropletArgs{
// 			Image:  pulumi.String("ubuntu-18-04-x64"),
// 			Region: pulumi.String("nyc2"),
// 			Size:   pulumi.String("s-1vcpu-1gb"),
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
// Droplets can be imported using the Droplet `id`, e.g.
//
// ```sh
//  $ pulumi import digitalocean:index/droplet:Droplet mydroplet 100823
// ```
type Droplet struct {
	pulumi.CustomResourceState

	// Boolean controlling if backups are made. Defaults to
	// false.
	Backups   pulumi.BoolPtrOutput `pulumi:"backups"`
	CreatedAt pulumi.StringOutput  `pulumi:"createdAt"`
	// The size of the instance's disk in GB
	Disk pulumi.IntOutput `pulumi:"disk"`
	// The uniform resource name of the Droplet
	// * `name`- The name of the Droplet
	DropletUrn pulumi.StringOutput `pulumi:"dropletUrn"`
	// The Droplet image ID or slug.
	Image pulumi.StringOutput `pulumi:"image"`
	// The IPv4 address
	Ipv4Address pulumi.StringOutput `pulumi:"ipv4Address"`
	// The private networking IPv4 address
	Ipv4AddressPrivate pulumi.StringOutput `pulumi:"ipv4AddressPrivate"`
	// Boolean controlling if IPv6 is enabled. Defaults to false.
	Ipv6 pulumi.BoolPtrOutput `pulumi:"ipv6"`
	// The IPv6 address
	Ipv6Address pulumi.StringOutput `pulumi:"ipv6Address"`
	// Is the Droplet locked
	Locked pulumi.BoolOutput `pulumi:"locked"`
	Memory pulumi.IntOutput  `pulumi:"memory"`
	// Boolean controlling whether monitoring agent is installed.
	// Defaults to false.
	Monitoring pulumi.BoolPtrOutput `pulumi:"monitoring"`
	// The Droplet name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Droplet hourly price
	PriceHourly pulumi.Float64Output `pulumi:"priceHourly"`
	// Droplet monthly price
	PriceMonthly pulumi.Float64Output `pulumi:"priceMonthly"`
	// Boolean controlling if private networking
	// is enabled. When VPC is enabled on an account, this will provision the
	// Droplet inside of your account's default VPC for the region. Use the
	// `vpcUuid` attribute to specify a different VPC.
	PrivateNetworking pulumi.BoolOutput `pulumi:"privateNetworking"`
	// The region to start in.
	Region pulumi.StringOutput `pulumi:"region"`
	// Boolean controlling whether to increase the disk
	// size when resizing a Droplet. It defaults to `true`. When set to `false`,
	// only the Droplet's RAM and CPU will be resized. **Increasing a Droplet's disk
	// size is a permanent change**. Increasing only RAM and CPU is reversible.
	ResizeDisk pulumi.BoolPtrOutput `pulumi:"resizeDisk"`
	// The unique slug that indentifies the type of Droplet. You can find a list of available slugs on [DigitalOcean API documentation](https://developers.digitalocean.com/documentation/v2/#list-all-sizes).
	Size pulumi.StringOutput `pulumi:"size"`
	// A list of SSH IDs or fingerprints to enable in
	// the format `[12345, 123456]`. To retrieve this info, use a tool such
	// as `curl` with the [DigitalOcean API](https://developers.digitalocean.com/documentation/v2/#ssh-keys),
	// to retrieve them.
	SshKeys pulumi.StringArrayOutput `pulumi:"sshKeys"`
	// The status of the Droplet
	Status pulumi.StringOutput `pulumi:"status"`
	// A list of the tags to be applied to this Droplet.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// A string of the desired User Data for the Droplet.
	UserData pulumi.StringPtrOutput `pulumi:"userData"`
	// The number of the instance's virtual CPUs
	Vcpus pulumi.IntOutput `pulumi:"vcpus"`
	// A list of the IDs of each block storage volume to be attached to the Droplet.
	VolumeIds pulumi.StringArrayOutput `pulumi:"volumeIds"`
	// The ID of the VPC where the Droplet will be located.
	VpcUuid pulumi.StringOutput `pulumi:"vpcUuid"`
}

// NewDroplet registers a new resource with the given unique name, arguments, and options.
func NewDroplet(ctx *pulumi.Context,
	name string, args *DropletArgs, opts ...pulumi.ResourceOption) (*Droplet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Image == nil {
		return nil, errors.New("invalid value for required argument 'Image'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	var resource Droplet
	err := ctx.RegisterResource("digitalocean:index/droplet:Droplet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDroplet gets an existing Droplet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDroplet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DropletState, opts ...pulumi.ResourceOption) (*Droplet, error) {
	var resource Droplet
	err := ctx.ReadResource("digitalocean:index/droplet:Droplet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Droplet resources.
type dropletState struct {
	// Boolean controlling if backups are made. Defaults to
	// false.
	Backups   *bool   `pulumi:"backups"`
	CreatedAt *string `pulumi:"createdAt"`
	// The size of the instance's disk in GB
	Disk *int `pulumi:"disk"`
	// The uniform resource name of the Droplet
	// * `name`- The name of the Droplet
	DropletUrn *string `pulumi:"dropletUrn"`
	// The Droplet image ID or slug.
	Image *string `pulumi:"image"`
	// The IPv4 address
	Ipv4Address *string `pulumi:"ipv4Address"`
	// The private networking IPv4 address
	Ipv4AddressPrivate *string `pulumi:"ipv4AddressPrivate"`
	// Boolean controlling if IPv6 is enabled. Defaults to false.
	Ipv6 *bool `pulumi:"ipv6"`
	// The IPv6 address
	Ipv6Address *string `pulumi:"ipv6Address"`
	// Is the Droplet locked
	Locked *bool `pulumi:"locked"`
	Memory *int  `pulumi:"memory"`
	// Boolean controlling whether monitoring agent is installed.
	// Defaults to false.
	Monitoring *bool `pulumi:"monitoring"`
	// The Droplet name.
	Name *string `pulumi:"name"`
	// Droplet hourly price
	PriceHourly *float64 `pulumi:"priceHourly"`
	// Droplet monthly price
	PriceMonthly *float64 `pulumi:"priceMonthly"`
	// Boolean controlling if private networking
	// is enabled. When VPC is enabled on an account, this will provision the
	// Droplet inside of your account's default VPC for the region. Use the
	// `vpcUuid` attribute to specify a different VPC.
	PrivateNetworking *bool `pulumi:"privateNetworking"`
	// The region to start in.
	Region *string `pulumi:"region"`
	// Boolean controlling whether to increase the disk
	// size when resizing a Droplet. It defaults to `true`. When set to `false`,
	// only the Droplet's RAM and CPU will be resized. **Increasing a Droplet's disk
	// size is a permanent change**. Increasing only RAM and CPU is reversible.
	ResizeDisk *bool `pulumi:"resizeDisk"`
	// The unique slug that indentifies the type of Droplet. You can find a list of available slugs on [DigitalOcean API documentation](https://developers.digitalocean.com/documentation/v2/#list-all-sizes).
	Size *string `pulumi:"size"`
	// A list of SSH IDs or fingerprints to enable in
	// the format `[12345, 123456]`. To retrieve this info, use a tool such
	// as `curl` with the [DigitalOcean API](https://developers.digitalocean.com/documentation/v2/#ssh-keys),
	// to retrieve them.
	SshKeys []string `pulumi:"sshKeys"`
	// The status of the Droplet
	Status *string `pulumi:"status"`
	// A list of the tags to be applied to this Droplet.
	Tags []string `pulumi:"tags"`
	// A string of the desired User Data for the Droplet.
	UserData *string `pulumi:"userData"`
	// The number of the instance's virtual CPUs
	Vcpus *int `pulumi:"vcpus"`
	// A list of the IDs of each block storage volume to be attached to the Droplet.
	VolumeIds []string `pulumi:"volumeIds"`
	// The ID of the VPC where the Droplet will be located.
	VpcUuid *string `pulumi:"vpcUuid"`
}

type DropletState struct {
	// Boolean controlling if backups are made. Defaults to
	// false.
	Backups   pulumi.BoolPtrInput
	CreatedAt pulumi.StringPtrInput
	// The size of the instance's disk in GB
	Disk pulumi.IntPtrInput
	// The uniform resource name of the Droplet
	// * `name`- The name of the Droplet
	DropletUrn pulumi.StringPtrInput
	// The Droplet image ID or slug.
	Image pulumi.StringPtrInput
	// The IPv4 address
	Ipv4Address pulumi.StringPtrInput
	// The private networking IPv4 address
	Ipv4AddressPrivate pulumi.StringPtrInput
	// Boolean controlling if IPv6 is enabled. Defaults to false.
	Ipv6 pulumi.BoolPtrInput
	// The IPv6 address
	Ipv6Address pulumi.StringPtrInput
	// Is the Droplet locked
	Locked pulumi.BoolPtrInput
	Memory pulumi.IntPtrInput
	// Boolean controlling whether monitoring agent is installed.
	// Defaults to false.
	Monitoring pulumi.BoolPtrInput
	// The Droplet name.
	Name pulumi.StringPtrInput
	// Droplet hourly price
	PriceHourly pulumi.Float64PtrInput
	// Droplet monthly price
	PriceMonthly pulumi.Float64PtrInput
	// Boolean controlling if private networking
	// is enabled. When VPC is enabled on an account, this will provision the
	// Droplet inside of your account's default VPC for the region. Use the
	// `vpcUuid` attribute to specify a different VPC.
	PrivateNetworking pulumi.BoolPtrInput
	// The region to start in.
	Region pulumi.StringPtrInput
	// Boolean controlling whether to increase the disk
	// size when resizing a Droplet. It defaults to `true`. When set to `false`,
	// only the Droplet's RAM and CPU will be resized. **Increasing a Droplet's disk
	// size is a permanent change**. Increasing only RAM and CPU is reversible.
	ResizeDisk pulumi.BoolPtrInput
	// The unique slug that indentifies the type of Droplet. You can find a list of available slugs on [DigitalOcean API documentation](https://developers.digitalocean.com/documentation/v2/#list-all-sizes).
	Size pulumi.StringPtrInput
	// A list of SSH IDs or fingerprints to enable in
	// the format `[12345, 123456]`. To retrieve this info, use a tool such
	// as `curl` with the [DigitalOcean API](https://developers.digitalocean.com/documentation/v2/#ssh-keys),
	// to retrieve them.
	SshKeys pulumi.StringArrayInput
	// The status of the Droplet
	Status pulumi.StringPtrInput
	// A list of the tags to be applied to this Droplet.
	Tags pulumi.StringArrayInput
	// A string of the desired User Data for the Droplet.
	UserData pulumi.StringPtrInput
	// The number of the instance's virtual CPUs
	Vcpus pulumi.IntPtrInput
	// A list of the IDs of each block storage volume to be attached to the Droplet.
	VolumeIds pulumi.StringArrayInput
	// The ID of the VPC where the Droplet will be located.
	VpcUuid pulumi.StringPtrInput
}

func (DropletState) ElementType() reflect.Type {
	return reflect.TypeOf((*dropletState)(nil)).Elem()
}

type dropletArgs struct {
	// Boolean controlling if backups are made. Defaults to
	// false.
	Backups *bool `pulumi:"backups"`
	// The Droplet image ID or slug.
	Image string `pulumi:"image"`
	// Boolean controlling if IPv6 is enabled. Defaults to false.
	Ipv6 *bool `pulumi:"ipv6"`
	// Boolean controlling whether monitoring agent is installed.
	// Defaults to false.
	Monitoring *bool `pulumi:"monitoring"`
	// The Droplet name.
	Name *string `pulumi:"name"`
	// Boolean controlling if private networking
	// is enabled. When VPC is enabled on an account, this will provision the
	// Droplet inside of your account's default VPC for the region. Use the
	// `vpcUuid` attribute to specify a different VPC.
	PrivateNetworking *bool `pulumi:"privateNetworking"`
	// The region to start in.
	Region string `pulumi:"region"`
	// Boolean controlling whether to increase the disk
	// size when resizing a Droplet. It defaults to `true`. When set to `false`,
	// only the Droplet's RAM and CPU will be resized. **Increasing a Droplet's disk
	// size is a permanent change**. Increasing only RAM and CPU is reversible.
	ResizeDisk *bool `pulumi:"resizeDisk"`
	// The unique slug that indentifies the type of Droplet. You can find a list of available slugs on [DigitalOcean API documentation](https://developers.digitalocean.com/documentation/v2/#list-all-sizes).
	Size string `pulumi:"size"`
	// A list of SSH IDs or fingerprints to enable in
	// the format `[12345, 123456]`. To retrieve this info, use a tool such
	// as `curl` with the [DigitalOcean API](https://developers.digitalocean.com/documentation/v2/#ssh-keys),
	// to retrieve them.
	SshKeys []string `pulumi:"sshKeys"`
	// A list of the tags to be applied to this Droplet.
	Tags []string `pulumi:"tags"`
	// A string of the desired User Data for the Droplet.
	UserData *string `pulumi:"userData"`
	// A list of the IDs of each block storage volume to be attached to the Droplet.
	VolumeIds []string `pulumi:"volumeIds"`
	// The ID of the VPC where the Droplet will be located.
	VpcUuid *string `pulumi:"vpcUuid"`
}

// The set of arguments for constructing a Droplet resource.
type DropletArgs struct {
	// Boolean controlling if backups are made. Defaults to
	// false.
	Backups pulumi.BoolPtrInput
	// The Droplet image ID or slug.
	Image pulumi.StringInput
	// Boolean controlling if IPv6 is enabled. Defaults to false.
	Ipv6 pulumi.BoolPtrInput
	// Boolean controlling whether monitoring agent is installed.
	// Defaults to false.
	Monitoring pulumi.BoolPtrInput
	// The Droplet name.
	Name pulumi.StringPtrInput
	// Boolean controlling if private networking
	// is enabled. When VPC is enabled on an account, this will provision the
	// Droplet inside of your account's default VPC for the region. Use the
	// `vpcUuid` attribute to specify a different VPC.
	PrivateNetworking pulumi.BoolPtrInput
	// The region to start in.
	Region pulumi.StringInput
	// Boolean controlling whether to increase the disk
	// size when resizing a Droplet. It defaults to `true`. When set to `false`,
	// only the Droplet's RAM and CPU will be resized. **Increasing a Droplet's disk
	// size is a permanent change**. Increasing only RAM and CPU is reversible.
	ResizeDisk pulumi.BoolPtrInput
	// The unique slug that indentifies the type of Droplet. You can find a list of available slugs on [DigitalOcean API documentation](https://developers.digitalocean.com/documentation/v2/#list-all-sizes).
	Size pulumi.StringInput
	// A list of SSH IDs or fingerprints to enable in
	// the format `[12345, 123456]`. To retrieve this info, use a tool such
	// as `curl` with the [DigitalOcean API](https://developers.digitalocean.com/documentation/v2/#ssh-keys),
	// to retrieve them.
	SshKeys pulumi.StringArrayInput
	// A list of the tags to be applied to this Droplet.
	Tags pulumi.StringArrayInput
	// A string of the desired User Data for the Droplet.
	UserData pulumi.StringPtrInput
	// A list of the IDs of each block storage volume to be attached to the Droplet.
	VolumeIds pulumi.StringArrayInput
	// The ID of the VPC where the Droplet will be located.
	VpcUuid pulumi.StringPtrInput
}

func (DropletArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dropletArgs)(nil)).Elem()
}

type DropletInput interface {
	pulumi.Input

	ToDropletOutput() DropletOutput
	ToDropletOutputWithContext(ctx context.Context) DropletOutput
}

func (*Droplet) ElementType() reflect.Type {
	return reflect.TypeOf((*Droplet)(nil))
}

func (i *Droplet) ToDropletOutput() DropletOutput {
	return i.ToDropletOutputWithContext(context.Background())
}

func (i *Droplet) ToDropletOutputWithContext(ctx context.Context) DropletOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DropletOutput)
}

type DropletOutput struct {
	*pulumi.OutputState
}

func (DropletOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Droplet)(nil))
}

func (o DropletOutput) ToDropletOutput() DropletOutput {
	return o
}

func (o DropletOutput) ToDropletOutputWithContext(ctx context.Context) DropletOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DropletOutput{})
}
