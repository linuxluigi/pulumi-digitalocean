// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DigitalOcean Volume Snapshot which can be used to create a snapshot from an existing volume.
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
// 		foobarVolume, err := digitalocean.NewVolume(ctx, "foobarVolume", &digitalocean.VolumeArgs{
// 			Region:      pulumi.String("nyc1"),
// 			Size:        pulumi.Int(100),
// 			Description: pulumi.String("an example volume"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = digitalocean.NewVolumeSnapshot(ctx, "foobarVolumeSnapshot", &digitalocean.VolumeSnapshotArgs{
// 			VolumeId: foobarVolume.ID(),
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
// Volume Snapshots can be imported using the `snapshot id`, e.g.
//
// ```sh
//  $ pulumi import digitalocean:index/volumeSnapshot:VolumeSnapshot snapshot 506f78a4-e098-11e5-ad9f-000f53306ae1
// ```
type VolumeSnapshot struct {
	pulumi.CustomResourceState

	// The date and time the volume snapshot was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The minimum size in gigabytes required for a volume to be created based on this volume snapshot.
	MinDiskSize pulumi.IntOutput `pulumi:"minDiskSize"`
	// A name for the volume snapshot.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of DigitalOcean region "slugs" indicating where the volume snapshot is available.
	Regions pulumi.StringArrayOutput `pulumi:"regions"`
	// The billable size of the volume snapshot in gigabytes.
	Size pulumi.Float64Output `pulumi:"size"`
	// A list of the tags to be applied to this volume snapshot.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The ID of the volume from which the volume snapshot originated.
	VolumeId pulumi.StringOutput `pulumi:"volumeId"`
}

// NewVolumeSnapshot registers a new resource with the given unique name, arguments, and options.
func NewVolumeSnapshot(ctx *pulumi.Context,
	name string, args *VolumeSnapshotArgs, opts ...pulumi.ResourceOption) (*VolumeSnapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VolumeId == nil {
		return nil, errors.New("invalid value for required argument 'VolumeId'")
	}
	var resource VolumeSnapshot
	err := ctx.RegisterResource("digitalocean:index/volumeSnapshot:VolumeSnapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolumeSnapshot gets an existing VolumeSnapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolumeSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeSnapshotState, opts ...pulumi.ResourceOption) (*VolumeSnapshot, error) {
	var resource VolumeSnapshot
	err := ctx.ReadResource("digitalocean:index/volumeSnapshot:VolumeSnapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VolumeSnapshot resources.
type volumeSnapshotState struct {
	// The date and time the volume snapshot was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The minimum size in gigabytes required for a volume to be created based on this volume snapshot.
	MinDiskSize *int `pulumi:"minDiskSize"`
	// A name for the volume snapshot.
	Name *string `pulumi:"name"`
	// A list of DigitalOcean region "slugs" indicating where the volume snapshot is available.
	Regions []string `pulumi:"regions"`
	// The billable size of the volume snapshot in gigabytes.
	Size *float64 `pulumi:"size"`
	// A list of the tags to be applied to this volume snapshot.
	Tags []string `pulumi:"tags"`
	// The ID of the volume from which the volume snapshot originated.
	VolumeId *string `pulumi:"volumeId"`
}

type VolumeSnapshotState struct {
	// The date and time the volume snapshot was created.
	CreatedAt pulumi.StringPtrInput
	// The minimum size in gigabytes required for a volume to be created based on this volume snapshot.
	MinDiskSize pulumi.IntPtrInput
	// A name for the volume snapshot.
	Name pulumi.StringPtrInput
	// A list of DigitalOcean region "slugs" indicating where the volume snapshot is available.
	Regions pulumi.StringArrayInput
	// The billable size of the volume snapshot in gigabytes.
	Size pulumi.Float64PtrInput
	// A list of the tags to be applied to this volume snapshot.
	Tags pulumi.StringArrayInput
	// The ID of the volume from which the volume snapshot originated.
	VolumeId pulumi.StringPtrInput
}

func (VolumeSnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeSnapshotState)(nil)).Elem()
}

type volumeSnapshotArgs struct {
	// A name for the volume snapshot.
	Name *string `pulumi:"name"`
	// A list of the tags to be applied to this volume snapshot.
	Tags []string `pulumi:"tags"`
	// The ID of the volume from which the volume snapshot originated.
	VolumeId string `pulumi:"volumeId"`
}

// The set of arguments for constructing a VolumeSnapshot resource.
type VolumeSnapshotArgs struct {
	// A name for the volume snapshot.
	Name pulumi.StringPtrInput
	// A list of the tags to be applied to this volume snapshot.
	Tags pulumi.StringArrayInput
	// The ID of the volume from which the volume snapshot originated.
	VolumeId pulumi.StringInput
}

func (VolumeSnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeSnapshotArgs)(nil)).Elem()
}

type VolumeSnapshotInput interface {
	pulumi.Input

	ToVolumeSnapshotOutput() VolumeSnapshotOutput
	ToVolumeSnapshotOutputWithContext(ctx context.Context) VolumeSnapshotOutput
}

func (*VolumeSnapshot) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeSnapshot)(nil))
}

func (i *VolumeSnapshot) ToVolumeSnapshotOutput() VolumeSnapshotOutput {
	return i.ToVolumeSnapshotOutputWithContext(context.Background())
}

func (i *VolumeSnapshot) ToVolumeSnapshotOutputWithContext(ctx context.Context) VolumeSnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeSnapshotOutput)
}

func (i *VolumeSnapshot) ToVolumeSnapshotPtrOutput() VolumeSnapshotPtrOutput {
	return i.ToVolumeSnapshotPtrOutputWithContext(context.Background())
}

func (i *VolumeSnapshot) ToVolumeSnapshotPtrOutputWithContext(ctx context.Context) VolumeSnapshotPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeSnapshotPtrOutput)
}

type VolumeSnapshotPtrInput interface {
	pulumi.Input

	ToVolumeSnapshotPtrOutput() VolumeSnapshotPtrOutput
	ToVolumeSnapshotPtrOutputWithContext(ctx context.Context) VolumeSnapshotPtrOutput
}

type volumeSnapshotPtrType VolumeSnapshotArgs

func (*volumeSnapshotPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeSnapshot)(nil))
}

func (i *volumeSnapshotPtrType) ToVolumeSnapshotPtrOutput() VolumeSnapshotPtrOutput {
	return i.ToVolumeSnapshotPtrOutputWithContext(context.Background())
}

func (i *volumeSnapshotPtrType) ToVolumeSnapshotPtrOutputWithContext(ctx context.Context) VolumeSnapshotPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeSnapshotPtrOutput)
}

// VolumeSnapshotArrayInput is an input type that accepts VolumeSnapshotArray and VolumeSnapshotArrayOutput values.
// You can construct a concrete instance of `VolumeSnapshotArrayInput` via:
//
//          VolumeSnapshotArray{ VolumeSnapshotArgs{...} }
type VolumeSnapshotArrayInput interface {
	pulumi.Input

	ToVolumeSnapshotArrayOutput() VolumeSnapshotArrayOutput
	ToVolumeSnapshotArrayOutputWithContext(context.Context) VolumeSnapshotArrayOutput
}

type VolumeSnapshotArray []VolumeSnapshotInput

func (VolumeSnapshotArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*VolumeSnapshot)(nil))
}

func (i VolumeSnapshotArray) ToVolumeSnapshotArrayOutput() VolumeSnapshotArrayOutput {
	return i.ToVolumeSnapshotArrayOutputWithContext(context.Background())
}

func (i VolumeSnapshotArray) ToVolumeSnapshotArrayOutputWithContext(ctx context.Context) VolumeSnapshotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeSnapshotArrayOutput)
}

// VolumeSnapshotMapInput is an input type that accepts VolumeSnapshotMap and VolumeSnapshotMapOutput values.
// You can construct a concrete instance of `VolumeSnapshotMapInput` via:
//
//          VolumeSnapshotMap{ "key": VolumeSnapshotArgs{...} }
type VolumeSnapshotMapInput interface {
	pulumi.Input

	ToVolumeSnapshotMapOutput() VolumeSnapshotMapOutput
	ToVolumeSnapshotMapOutputWithContext(context.Context) VolumeSnapshotMapOutput
}

type VolumeSnapshotMap map[string]VolumeSnapshotInput

func (VolumeSnapshotMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*VolumeSnapshot)(nil))
}

func (i VolumeSnapshotMap) ToVolumeSnapshotMapOutput() VolumeSnapshotMapOutput {
	return i.ToVolumeSnapshotMapOutputWithContext(context.Background())
}

func (i VolumeSnapshotMap) ToVolumeSnapshotMapOutputWithContext(ctx context.Context) VolumeSnapshotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeSnapshotMapOutput)
}

type VolumeSnapshotOutput struct {
	*pulumi.OutputState
}

func (VolumeSnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeSnapshot)(nil))
}

func (o VolumeSnapshotOutput) ToVolumeSnapshotOutput() VolumeSnapshotOutput {
	return o
}

func (o VolumeSnapshotOutput) ToVolumeSnapshotOutputWithContext(ctx context.Context) VolumeSnapshotOutput {
	return o
}

func (o VolumeSnapshotOutput) ToVolumeSnapshotPtrOutput() VolumeSnapshotPtrOutput {
	return o.ToVolumeSnapshotPtrOutputWithContext(context.Background())
}

func (o VolumeSnapshotOutput) ToVolumeSnapshotPtrOutputWithContext(ctx context.Context) VolumeSnapshotPtrOutput {
	return o.ApplyT(func(v VolumeSnapshot) *VolumeSnapshot {
		return &v
	}).(VolumeSnapshotPtrOutput)
}

type VolumeSnapshotPtrOutput struct {
	*pulumi.OutputState
}

func (VolumeSnapshotPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeSnapshot)(nil))
}

func (o VolumeSnapshotPtrOutput) ToVolumeSnapshotPtrOutput() VolumeSnapshotPtrOutput {
	return o
}

func (o VolumeSnapshotPtrOutput) ToVolumeSnapshotPtrOutputWithContext(ctx context.Context) VolumeSnapshotPtrOutput {
	return o
}

type VolumeSnapshotArrayOutput struct{ *pulumi.OutputState }

func (VolumeSnapshotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeSnapshot)(nil))
}

func (o VolumeSnapshotArrayOutput) ToVolumeSnapshotArrayOutput() VolumeSnapshotArrayOutput {
	return o
}

func (o VolumeSnapshotArrayOutput) ToVolumeSnapshotArrayOutputWithContext(ctx context.Context) VolumeSnapshotArrayOutput {
	return o
}

func (o VolumeSnapshotArrayOutput) Index(i pulumi.IntInput) VolumeSnapshotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumeSnapshot {
		return vs[0].([]VolumeSnapshot)[vs[1].(int)]
	}).(VolumeSnapshotOutput)
}

type VolumeSnapshotMapOutput struct{ *pulumi.OutputState }

func (VolumeSnapshotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VolumeSnapshot)(nil))
}

func (o VolumeSnapshotMapOutput) ToVolumeSnapshotMapOutput() VolumeSnapshotMapOutput {
	return o
}

func (o VolumeSnapshotMapOutput) ToVolumeSnapshotMapOutputWithContext(ctx context.Context) VolumeSnapshotMapOutput {
	return o
}

func (o VolumeSnapshotMapOutput) MapIndex(k pulumi.StringInput) VolumeSnapshotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VolumeSnapshot {
		return vs[0].(map[string]VolumeSnapshot)[vs[1].(string)]
	}).(VolumeSnapshotOutput)
}

func init() {
	pulumi.RegisterOutputType(VolumeSnapshotOutput{})
	pulumi.RegisterOutputType(VolumeSnapshotPtrOutput{})
	pulumi.RegisterOutputType(VolumeSnapshotArrayOutput{})
	pulumi.RegisterOutputType(VolumeSnapshotMapOutput{})
}
