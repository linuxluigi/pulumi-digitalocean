// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a [DigitalOcean Uptime Alerts](https://docs.digitalocean.com/reference/api/api-reference/#operation/uptime_alert_create)
// resource. Uptime Alerts provide the ability to add alerts to your [DigitalOcean Uptime Checks](https://docs.digitalocean.com/reference/api/api-reference/#tag/Uptime) when your endpoints are slow, unavailable, or SSL certificates are expiring.
//
// ## Import
//
// Uptime checks can be imported using the uptime alert's `id`, e.g.
//
// ```sh
//
//	$ pulumi import digitalocean:index/uptimeAlert:UptimeAlert target 5a4981aa-9653-4bd1-bef5-d6bff52042e4
//
// ```
type UptimeAlert struct {
	pulumi.CustomResourceState

	// A unique identifier for a check
	CheckId pulumi.StringOutput `pulumi:"checkId"`
	// The comparison operator used against the alert's threshold. Must be one of `greaterThan` or `lessThan`.
	Comparison pulumi.StringPtrOutput `pulumi:"comparison"`
	// A human-friendly display name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The notification settings for a trigger alert.
	Notifications UptimeAlertNotificationArrayOutput `pulumi:"notifications"`
	// Period of time the threshold must be exceeded to trigger the alert. Must be one of `2m`, `3m`, `5m`, `10m`, `15m`, `30m` or `1h`.
	Period pulumi.StringPtrOutput `pulumi:"period"`
	// The threshold at which the alert will enter a trigger state. The specific threshold is dependent on the alert type.
	Threshold pulumi.IntPtrOutput `pulumi:"threshold"`
	// The type of health check to perform. Must be one of `latency`, `down`, `downGlobal` or `sslExpiry`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewUptimeAlert registers a new resource with the given unique name, arguments, and options.
func NewUptimeAlert(ctx *pulumi.Context,
	name string, args *UptimeAlertArgs, opts ...pulumi.ResourceOption) (*UptimeAlert, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CheckId == nil {
		return nil, errors.New("invalid value for required argument 'CheckId'")
	}
	if args.Notifications == nil {
		return nil, errors.New("invalid value for required argument 'Notifications'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UptimeAlert
	err := ctx.RegisterResource("digitalocean:index/uptimeAlert:UptimeAlert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUptimeAlert gets an existing UptimeAlert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUptimeAlert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UptimeAlertState, opts ...pulumi.ResourceOption) (*UptimeAlert, error) {
	var resource UptimeAlert
	err := ctx.ReadResource("digitalocean:index/uptimeAlert:UptimeAlert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UptimeAlert resources.
type uptimeAlertState struct {
	// A unique identifier for a check
	CheckId *string `pulumi:"checkId"`
	// The comparison operator used against the alert's threshold. Must be one of `greaterThan` or `lessThan`.
	Comparison *string `pulumi:"comparison"`
	// A human-friendly display name.
	Name *string `pulumi:"name"`
	// The notification settings for a trigger alert.
	Notifications []UptimeAlertNotification `pulumi:"notifications"`
	// Period of time the threshold must be exceeded to trigger the alert. Must be one of `2m`, `3m`, `5m`, `10m`, `15m`, `30m` or `1h`.
	Period *string `pulumi:"period"`
	// The threshold at which the alert will enter a trigger state. The specific threshold is dependent on the alert type.
	Threshold *int `pulumi:"threshold"`
	// The type of health check to perform. Must be one of `latency`, `down`, `downGlobal` or `sslExpiry`.
	Type *string `pulumi:"type"`
}

type UptimeAlertState struct {
	// A unique identifier for a check
	CheckId pulumi.StringPtrInput
	// The comparison operator used against the alert's threshold. Must be one of `greaterThan` or `lessThan`.
	Comparison pulumi.StringPtrInput
	// A human-friendly display name.
	Name pulumi.StringPtrInput
	// The notification settings for a trigger alert.
	Notifications UptimeAlertNotificationArrayInput
	// Period of time the threshold must be exceeded to trigger the alert. Must be one of `2m`, `3m`, `5m`, `10m`, `15m`, `30m` or `1h`.
	Period pulumi.StringPtrInput
	// The threshold at which the alert will enter a trigger state. The specific threshold is dependent on the alert type.
	Threshold pulumi.IntPtrInput
	// The type of health check to perform. Must be one of `latency`, `down`, `downGlobal` or `sslExpiry`.
	Type pulumi.StringPtrInput
}

func (UptimeAlertState) ElementType() reflect.Type {
	return reflect.TypeOf((*uptimeAlertState)(nil)).Elem()
}

type uptimeAlertArgs struct {
	// A unique identifier for a check
	CheckId string `pulumi:"checkId"`
	// The comparison operator used against the alert's threshold. Must be one of `greaterThan` or `lessThan`.
	Comparison *string `pulumi:"comparison"`
	// A human-friendly display name.
	Name *string `pulumi:"name"`
	// The notification settings for a trigger alert.
	Notifications []UptimeAlertNotification `pulumi:"notifications"`
	// Period of time the threshold must be exceeded to trigger the alert. Must be one of `2m`, `3m`, `5m`, `10m`, `15m`, `30m` or `1h`.
	Period *string `pulumi:"period"`
	// The threshold at which the alert will enter a trigger state. The specific threshold is dependent on the alert type.
	Threshold *int `pulumi:"threshold"`
	// The type of health check to perform. Must be one of `latency`, `down`, `downGlobal` or `sslExpiry`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a UptimeAlert resource.
type UptimeAlertArgs struct {
	// A unique identifier for a check
	CheckId pulumi.StringInput
	// The comparison operator used against the alert's threshold. Must be one of `greaterThan` or `lessThan`.
	Comparison pulumi.StringPtrInput
	// A human-friendly display name.
	Name pulumi.StringPtrInput
	// The notification settings for a trigger alert.
	Notifications UptimeAlertNotificationArrayInput
	// Period of time the threshold must be exceeded to trigger the alert. Must be one of `2m`, `3m`, `5m`, `10m`, `15m`, `30m` or `1h`.
	Period pulumi.StringPtrInput
	// The threshold at which the alert will enter a trigger state. The specific threshold is dependent on the alert type.
	Threshold pulumi.IntPtrInput
	// The type of health check to perform. Must be one of `latency`, `down`, `downGlobal` or `sslExpiry`.
	Type pulumi.StringInput
}

func (UptimeAlertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*uptimeAlertArgs)(nil)).Elem()
}

type UptimeAlertInput interface {
	pulumi.Input

	ToUptimeAlertOutput() UptimeAlertOutput
	ToUptimeAlertOutputWithContext(ctx context.Context) UptimeAlertOutput
}

func (*UptimeAlert) ElementType() reflect.Type {
	return reflect.TypeOf((**UptimeAlert)(nil)).Elem()
}

func (i *UptimeAlert) ToUptimeAlertOutput() UptimeAlertOutput {
	return i.ToUptimeAlertOutputWithContext(context.Background())
}

func (i *UptimeAlert) ToUptimeAlertOutputWithContext(ctx context.Context) UptimeAlertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UptimeAlertOutput)
}

// UptimeAlertArrayInput is an input type that accepts UptimeAlertArray and UptimeAlertArrayOutput values.
// You can construct a concrete instance of `UptimeAlertArrayInput` via:
//
//	UptimeAlertArray{ UptimeAlertArgs{...} }
type UptimeAlertArrayInput interface {
	pulumi.Input

	ToUptimeAlertArrayOutput() UptimeAlertArrayOutput
	ToUptimeAlertArrayOutputWithContext(context.Context) UptimeAlertArrayOutput
}

type UptimeAlertArray []UptimeAlertInput

func (UptimeAlertArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UptimeAlert)(nil)).Elem()
}

func (i UptimeAlertArray) ToUptimeAlertArrayOutput() UptimeAlertArrayOutput {
	return i.ToUptimeAlertArrayOutputWithContext(context.Background())
}

func (i UptimeAlertArray) ToUptimeAlertArrayOutputWithContext(ctx context.Context) UptimeAlertArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UptimeAlertArrayOutput)
}

// UptimeAlertMapInput is an input type that accepts UptimeAlertMap and UptimeAlertMapOutput values.
// You can construct a concrete instance of `UptimeAlertMapInput` via:
//
//	UptimeAlertMap{ "key": UptimeAlertArgs{...} }
type UptimeAlertMapInput interface {
	pulumi.Input

	ToUptimeAlertMapOutput() UptimeAlertMapOutput
	ToUptimeAlertMapOutputWithContext(context.Context) UptimeAlertMapOutput
}

type UptimeAlertMap map[string]UptimeAlertInput

func (UptimeAlertMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UptimeAlert)(nil)).Elem()
}

func (i UptimeAlertMap) ToUptimeAlertMapOutput() UptimeAlertMapOutput {
	return i.ToUptimeAlertMapOutputWithContext(context.Background())
}

func (i UptimeAlertMap) ToUptimeAlertMapOutputWithContext(ctx context.Context) UptimeAlertMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UptimeAlertMapOutput)
}

type UptimeAlertOutput struct{ *pulumi.OutputState }

func (UptimeAlertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UptimeAlert)(nil)).Elem()
}

func (o UptimeAlertOutput) ToUptimeAlertOutput() UptimeAlertOutput {
	return o
}

func (o UptimeAlertOutput) ToUptimeAlertOutputWithContext(ctx context.Context) UptimeAlertOutput {
	return o
}

// A unique identifier for a check
func (o UptimeAlertOutput) CheckId() pulumi.StringOutput {
	return o.ApplyT(func(v *UptimeAlert) pulumi.StringOutput { return v.CheckId }).(pulumi.StringOutput)
}

// The comparison operator used against the alert's threshold. Must be one of `greaterThan` or `lessThan`.
func (o UptimeAlertOutput) Comparison() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UptimeAlert) pulumi.StringPtrOutput { return v.Comparison }).(pulumi.StringPtrOutput)
}

// A human-friendly display name.
func (o UptimeAlertOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UptimeAlert) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The notification settings for a trigger alert.
func (o UptimeAlertOutput) Notifications() UptimeAlertNotificationArrayOutput {
	return o.ApplyT(func(v *UptimeAlert) UptimeAlertNotificationArrayOutput { return v.Notifications }).(UptimeAlertNotificationArrayOutput)
}

// Period of time the threshold must be exceeded to trigger the alert. Must be one of `2m`, `3m`, `5m`, `10m`, `15m`, `30m` or `1h`.
func (o UptimeAlertOutput) Period() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UptimeAlert) pulumi.StringPtrOutput { return v.Period }).(pulumi.StringPtrOutput)
}

// The threshold at which the alert will enter a trigger state. The specific threshold is dependent on the alert type.
func (o UptimeAlertOutput) Threshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *UptimeAlert) pulumi.IntPtrOutput { return v.Threshold }).(pulumi.IntPtrOutput)
}

// The type of health check to perform. Must be one of `latency`, `down`, `downGlobal` or `sslExpiry`.
func (o UptimeAlertOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *UptimeAlert) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type UptimeAlertArrayOutput struct{ *pulumi.OutputState }

func (UptimeAlertArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UptimeAlert)(nil)).Elem()
}

func (o UptimeAlertArrayOutput) ToUptimeAlertArrayOutput() UptimeAlertArrayOutput {
	return o
}

func (o UptimeAlertArrayOutput) ToUptimeAlertArrayOutputWithContext(ctx context.Context) UptimeAlertArrayOutput {
	return o
}

func (o UptimeAlertArrayOutput) Index(i pulumi.IntInput) UptimeAlertOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UptimeAlert {
		return vs[0].([]*UptimeAlert)[vs[1].(int)]
	}).(UptimeAlertOutput)
}

type UptimeAlertMapOutput struct{ *pulumi.OutputState }

func (UptimeAlertMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UptimeAlert)(nil)).Elem()
}

func (o UptimeAlertMapOutput) ToUptimeAlertMapOutput() UptimeAlertMapOutput {
	return o
}

func (o UptimeAlertMapOutput) ToUptimeAlertMapOutputWithContext(ctx context.Context) UptimeAlertMapOutput {
	return o
}

func (o UptimeAlertMapOutput) MapIndex(k pulumi.StringInput) UptimeAlertOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UptimeAlert {
		return vs[0].(map[string]*UptimeAlert)[vs[1].(string)]
	}).(UptimeAlertOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UptimeAlertInput)(nil)).Elem(), &UptimeAlert{})
	pulumi.RegisterInputType(reflect.TypeOf((*UptimeAlertArrayInput)(nil)).Elem(), UptimeAlertArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UptimeAlertMapInput)(nil)).Elem(), UptimeAlertMap{})
	pulumi.RegisterOutputType(UptimeAlertOutput{})
	pulumi.RegisterOutputType(UptimeAlertArrayOutput{})
	pulumi.RegisterOutputType(UptimeAlertMapOutput{})
}
