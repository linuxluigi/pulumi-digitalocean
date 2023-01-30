// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a [DigitalOcean Monitoring](https://docs.digitalocean.com/reference/api/api-reference/#tag/Monitoring)
// resource. Monitor alerts can be configured to alert about, e.g., disk or memory
// usage exceeding a certain threshold or traffic at a certain limit. Notifications
// can be sent to either an email address or a Slack channel.
//
// > **Note** Currently, the [DigitalOcean API](https://docs.digitalocean.com/reference/api/api-reference/#operation/create_alert_policy) only supports creating alerts for Droplets.
//
// ## Import
//
// Monitor alerts can be imported using the monitor alert `uuid`, e.g.
//
// ```sh
//
//	$ pulumi import digitalocean:index/monitorAlert:MonitorAlert cpu_alert b8ecd2ab-2267-4a5e-8692-cbf1d32583e3
//
// ```
type MonitorAlert struct {
	pulumi.CustomResourceState

	// How to send notifications about the alerts. This is a list with one element, .
	// Note that for Slack, the DigitalOcean app needs to have permissions for your workspace. You can
	// read more in [Slack's documentation](https://slack.com/intl/en-dk/help/articles/222386767-Manage-app-installation-settings-for-your-workspace)
	Alerts MonitorAlertAlertsOutput `pulumi:"alerts"`
	// The comparison for `value`.
	// This may be either `GreaterThan` or `LessThan`.
	Compare pulumi.StringOutput `pulumi:"compare"`
	// The description of the alert.
	Description pulumi.StringOutput `pulumi:"description"`
	// The status of the alert.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// A list of IDs for the resources to which the alert policy applies.
	Entities pulumi.StringArrayOutput `pulumi:"entities"`
	// A list of tags. When an included tag is added to a resource, the alert policy will apply to it.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The type of the alert.
	// This may be one of `v1/insights/droplet/load_1`, `v1/insights/droplet/load_5`, `v1/insights/droplet/load_15`,
	// `v1/insights/droplet/memory_utilization_percent`, `v1/insights/droplet/disk_utilization_percent`,
	// `v1/insights/droplet/cpu`, `v1/insights/droplet/disk_read`, `v1/insights/droplet/disk_write`,
	// `v1/insights/droplet/public_outbound_bandwidth`, `v1/insights/droplet/public_inbound_bandwidth`,
	// `v1/insights/droplet/private_outbound_bandwidth`, `v1/insights/droplet/private_inbound_bandwidth`,
	// `v1/insights/lbaas/avg_cpu_utilization_percent`, `v1/insights/lbaas/connection_utilization_percent`,
	// `v1/insights/lbaas/droplet_health`, `v1/insights/lbaas/tls_connections_per_second_utilization_percent`,
	// `v1/dbaas/alerts/load_15_alerts`, `v1/dbaas/alerts/cpu_alerts`, `v1/dbaas/alerts/memory_utilization_alerts`, or
	// `v1/dbaas/alerts/disk_utilization_alerts`.
	Type pulumi.StringOutput `pulumi:"type"`
	// The uuid of the alert.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// The value to start alerting at, e.g., 90% or 85Mbps. This is a floating-point number.
	// DigitalOcean will show the correct unit in the web panel.
	Value pulumi.Float64Output `pulumi:"value"`
	// The time frame of the alert. Either `5m`, `10m`, `30m`, or `1h`.
	Window pulumi.StringOutput `pulumi:"window"`
}

// NewMonitorAlert registers a new resource with the given unique name, arguments, and options.
func NewMonitorAlert(ctx *pulumi.Context,
	name string, args *MonitorAlertArgs, opts ...pulumi.ResourceOption) (*MonitorAlert, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Alerts == nil {
		return nil, errors.New("invalid value for required argument 'Alerts'")
	}
	if args.Compare == nil {
		return nil, errors.New("invalid value for required argument 'Compare'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	if args.Window == nil {
		return nil, errors.New("invalid value for required argument 'Window'")
	}
	var resource MonitorAlert
	err := ctx.RegisterResource("digitalocean:index/monitorAlert:MonitorAlert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMonitorAlert gets an existing MonitorAlert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMonitorAlert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MonitorAlertState, opts ...pulumi.ResourceOption) (*MonitorAlert, error) {
	var resource MonitorAlert
	err := ctx.ReadResource("digitalocean:index/monitorAlert:MonitorAlert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MonitorAlert resources.
type monitorAlertState struct {
	// How to send notifications about the alerts. This is a list with one element, .
	// Note that for Slack, the DigitalOcean app needs to have permissions for your workspace. You can
	// read more in [Slack's documentation](https://slack.com/intl/en-dk/help/articles/222386767-Manage-app-installation-settings-for-your-workspace)
	Alerts *MonitorAlertAlerts `pulumi:"alerts"`
	// The comparison for `value`.
	// This may be either `GreaterThan` or `LessThan`.
	Compare *string `pulumi:"compare"`
	// The description of the alert.
	Description *string `pulumi:"description"`
	// The status of the alert.
	Enabled *bool `pulumi:"enabled"`
	// A list of IDs for the resources to which the alert policy applies.
	Entities []string `pulumi:"entities"`
	// A list of tags. When an included tag is added to a resource, the alert policy will apply to it.
	Tags []string `pulumi:"tags"`
	// The type of the alert.
	// This may be one of `v1/insights/droplet/load_1`, `v1/insights/droplet/load_5`, `v1/insights/droplet/load_15`,
	// `v1/insights/droplet/memory_utilization_percent`, `v1/insights/droplet/disk_utilization_percent`,
	// `v1/insights/droplet/cpu`, `v1/insights/droplet/disk_read`, `v1/insights/droplet/disk_write`,
	// `v1/insights/droplet/public_outbound_bandwidth`, `v1/insights/droplet/public_inbound_bandwidth`,
	// `v1/insights/droplet/private_outbound_bandwidth`, `v1/insights/droplet/private_inbound_bandwidth`,
	// `v1/insights/lbaas/avg_cpu_utilization_percent`, `v1/insights/lbaas/connection_utilization_percent`,
	// `v1/insights/lbaas/droplet_health`, `v1/insights/lbaas/tls_connections_per_second_utilization_percent`,
	// `v1/dbaas/alerts/load_15_alerts`, `v1/dbaas/alerts/cpu_alerts`, `v1/dbaas/alerts/memory_utilization_alerts`, or
	// `v1/dbaas/alerts/disk_utilization_alerts`.
	Type *string `pulumi:"type"`
	// The uuid of the alert.
	Uuid *string `pulumi:"uuid"`
	// The value to start alerting at, e.g., 90% or 85Mbps. This is a floating-point number.
	// DigitalOcean will show the correct unit in the web panel.
	Value *float64 `pulumi:"value"`
	// The time frame of the alert. Either `5m`, `10m`, `30m`, or `1h`.
	Window *string `pulumi:"window"`
}

type MonitorAlertState struct {
	// How to send notifications about the alerts. This is a list with one element, .
	// Note that for Slack, the DigitalOcean app needs to have permissions for your workspace. You can
	// read more in [Slack's documentation](https://slack.com/intl/en-dk/help/articles/222386767-Manage-app-installation-settings-for-your-workspace)
	Alerts MonitorAlertAlertsPtrInput
	// The comparison for `value`.
	// This may be either `GreaterThan` or `LessThan`.
	Compare pulumi.StringPtrInput
	// The description of the alert.
	Description pulumi.StringPtrInput
	// The status of the alert.
	Enabled pulumi.BoolPtrInput
	// A list of IDs for the resources to which the alert policy applies.
	Entities pulumi.StringArrayInput
	// A list of tags. When an included tag is added to a resource, the alert policy will apply to it.
	Tags pulumi.StringArrayInput
	// The type of the alert.
	// This may be one of `v1/insights/droplet/load_1`, `v1/insights/droplet/load_5`, `v1/insights/droplet/load_15`,
	// `v1/insights/droplet/memory_utilization_percent`, `v1/insights/droplet/disk_utilization_percent`,
	// `v1/insights/droplet/cpu`, `v1/insights/droplet/disk_read`, `v1/insights/droplet/disk_write`,
	// `v1/insights/droplet/public_outbound_bandwidth`, `v1/insights/droplet/public_inbound_bandwidth`,
	// `v1/insights/droplet/private_outbound_bandwidth`, `v1/insights/droplet/private_inbound_bandwidth`,
	// `v1/insights/lbaas/avg_cpu_utilization_percent`, `v1/insights/lbaas/connection_utilization_percent`,
	// `v1/insights/lbaas/droplet_health`, `v1/insights/lbaas/tls_connections_per_second_utilization_percent`,
	// `v1/dbaas/alerts/load_15_alerts`, `v1/dbaas/alerts/cpu_alerts`, `v1/dbaas/alerts/memory_utilization_alerts`, or
	// `v1/dbaas/alerts/disk_utilization_alerts`.
	Type pulumi.StringPtrInput
	// The uuid of the alert.
	Uuid pulumi.StringPtrInput
	// The value to start alerting at, e.g., 90% or 85Mbps. This is a floating-point number.
	// DigitalOcean will show the correct unit in the web panel.
	Value pulumi.Float64PtrInput
	// The time frame of the alert. Either `5m`, `10m`, `30m`, or `1h`.
	Window pulumi.StringPtrInput
}

func (MonitorAlertState) ElementType() reflect.Type {
	return reflect.TypeOf((*monitorAlertState)(nil)).Elem()
}

type monitorAlertArgs struct {
	// How to send notifications about the alerts. This is a list with one element, .
	// Note that for Slack, the DigitalOcean app needs to have permissions for your workspace. You can
	// read more in [Slack's documentation](https://slack.com/intl/en-dk/help/articles/222386767-Manage-app-installation-settings-for-your-workspace)
	Alerts MonitorAlertAlerts `pulumi:"alerts"`
	// The comparison for `value`.
	// This may be either `GreaterThan` or `LessThan`.
	Compare string `pulumi:"compare"`
	// The description of the alert.
	Description string `pulumi:"description"`
	// The status of the alert.
	Enabled *bool `pulumi:"enabled"`
	// A list of IDs for the resources to which the alert policy applies.
	Entities []string `pulumi:"entities"`
	// A list of tags. When an included tag is added to a resource, the alert policy will apply to it.
	Tags []string `pulumi:"tags"`
	// The type of the alert.
	// This may be one of `v1/insights/droplet/load_1`, `v1/insights/droplet/load_5`, `v1/insights/droplet/load_15`,
	// `v1/insights/droplet/memory_utilization_percent`, `v1/insights/droplet/disk_utilization_percent`,
	// `v1/insights/droplet/cpu`, `v1/insights/droplet/disk_read`, `v1/insights/droplet/disk_write`,
	// `v1/insights/droplet/public_outbound_bandwidth`, `v1/insights/droplet/public_inbound_bandwidth`,
	// `v1/insights/droplet/private_outbound_bandwidth`, `v1/insights/droplet/private_inbound_bandwidth`,
	// `v1/insights/lbaas/avg_cpu_utilization_percent`, `v1/insights/lbaas/connection_utilization_percent`,
	// `v1/insights/lbaas/droplet_health`, `v1/insights/lbaas/tls_connections_per_second_utilization_percent`,
	// `v1/dbaas/alerts/load_15_alerts`, `v1/dbaas/alerts/cpu_alerts`, `v1/dbaas/alerts/memory_utilization_alerts`, or
	// `v1/dbaas/alerts/disk_utilization_alerts`.
	Type string `pulumi:"type"`
	// The value to start alerting at, e.g., 90% or 85Mbps. This is a floating-point number.
	// DigitalOcean will show the correct unit in the web panel.
	Value float64 `pulumi:"value"`
	// The time frame of the alert. Either `5m`, `10m`, `30m`, or `1h`.
	Window string `pulumi:"window"`
}

// The set of arguments for constructing a MonitorAlert resource.
type MonitorAlertArgs struct {
	// How to send notifications about the alerts. This is a list with one element, .
	// Note that for Slack, the DigitalOcean app needs to have permissions for your workspace. You can
	// read more in [Slack's documentation](https://slack.com/intl/en-dk/help/articles/222386767-Manage-app-installation-settings-for-your-workspace)
	Alerts MonitorAlertAlertsInput
	// The comparison for `value`.
	// This may be either `GreaterThan` or `LessThan`.
	Compare pulumi.StringInput
	// The description of the alert.
	Description pulumi.StringInput
	// The status of the alert.
	Enabled pulumi.BoolPtrInput
	// A list of IDs for the resources to which the alert policy applies.
	Entities pulumi.StringArrayInput
	// A list of tags. When an included tag is added to a resource, the alert policy will apply to it.
	Tags pulumi.StringArrayInput
	// The type of the alert.
	// This may be one of `v1/insights/droplet/load_1`, `v1/insights/droplet/load_5`, `v1/insights/droplet/load_15`,
	// `v1/insights/droplet/memory_utilization_percent`, `v1/insights/droplet/disk_utilization_percent`,
	// `v1/insights/droplet/cpu`, `v1/insights/droplet/disk_read`, `v1/insights/droplet/disk_write`,
	// `v1/insights/droplet/public_outbound_bandwidth`, `v1/insights/droplet/public_inbound_bandwidth`,
	// `v1/insights/droplet/private_outbound_bandwidth`, `v1/insights/droplet/private_inbound_bandwidth`,
	// `v1/insights/lbaas/avg_cpu_utilization_percent`, `v1/insights/lbaas/connection_utilization_percent`,
	// `v1/insights/lbaas/droplet_health`, `v1/insights/lbaas/tls_connections_per_second_utilization_percent`,
	// `v1/dbaas/alerts/load_15_alerts`, `v1/dbaas/alerts/cpu_alerts`, `v1/dbaas/alerts/memory_utilization_alerts`, or
	// `v1/dbaas/alerts/disk_utilization_alerts`.
	Type pulumi.StringInput
	// The value to start alerting at, e.g., 90% or 85Mbps. This is a floating-point number.
	// DigitalOcean will show the correct unit in the web panel.
	Value pulumi.Float64Input
	// The time frame of the alert. Either `5m`, `10m`, `30m`, or `1h`.
	Window pulumi.StringInput
}

func (MonitorAlertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*monitorAlertArgs)(nil)).Elem()
}

type MonitorAlertInput interface {
	pulumi.Input

	ToMonitorAlertOutput() MonitorAlertOutput
	ToMonitorAlertOutputWithContext(ctx context.Context) MonitorAlertOutput
}

func (*MonitorAlert) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorAlert)(nil)).Elem()
}

func (i *MonitorAlert) ToMonitorAlertOutput() MonitorAlertOutput {
	return i.ToMonitorAlertOutputWithContext(context.Background())
}

func (i *MonitorAlert) ToMonitorAlertOutputWithContext(ctx context.Context) MonitorAlertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorAlertOutput)
}

// MonitorAlertArrayInput is an input type that accepts MonitorAlertArray and MonitorAlertArrayOutput values.
// You can construct a concrete instance of `MonitorAlertArrayInput` via:
//
//	MonitorAlertArray{ MonitorAlertArgs{...} }
type MonitorAlertArrayInput interface {
	pulumi.Input

	ToMonitorAlertArrayOutput() MonitorAlertArrayOutput
	ToMonitorAlertArrayOutputWithContext(context.Context) MonitorAlertArrayOutput
}

type MonitorAlertArray []MonitorAlertInput

func (MonitorAlertArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MonitorAlert)(nil)).Elem()
}

func (i MonitorAlertArray) ToMonitorAlertArrayOutput() MonitorAlertArrayOutput {
	return i.ToMonitorAlertArrayOutputWithContext(context.Background())
}

func (i MonitorAlertArray) ToMonitorAlertArrayOutputWithContext(ctx context.Context) MonitorAlertArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorAlertArrayOutput)
}

// MonitorAlertMapInput is an input type that accepts MonitorAlertMap and MonitorAlertMapOutput values.
// You can construct a concrete instance of `MonitorAlertMapInput` via:
//
//	MonitorAlertMap{ "key": MonitorAlertArgs{...} }
type MonitorAlertMapInput interface {
	pulumi.Input

	ToMonitorAlertMapOutput() MonitorAlertMapOutput
	ToMonitorAlertMapOutputWithContext(context.Context) MonitorAlertMapOutput
}

type MonitorAlertMap map[string]MonitorAlertInput

func (MonitorAlertMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MonitorAlert)(nil)).Elem()
}

func (i MonitorAlertMap) ToMonitorAlertMapOutput() MonitorAlertMapOutput {
	return i.ToMonitorAlertMapOutputWithContext(context.Background())
}

func (i MonitorAlertMap) ToMonitorAlertMapOutputWithContext(ctx context.Context) MonitorAlertMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorAlertMapOutput)
}

type MonitorAlertOutput struct{ *pulumi.OutputState }

func (MonitorAlertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorAlert)(nil)).Elem()
}

func (o MonitorAlertOutput) ToMonitorAlertOutput() MonitorAlertOutput {
	return o
}

func (o MonitorAlertOutput) ToMonitorAlertOutputWithContext(ctx context.Context) MonitorAlertOutput {
	return o
}

// How to send notifications about the alerts. This is a list with one element, .
// Note that for Slack, the DigitalOcean app needs to have permissions for your workspace. You can
// read more in [Slack's documentation](https://slack.com/intl/en-dk/help/articles/222386767-Manage-app-installation-settings-for-your-workspace)
func (o MonitorAlertOutput) Alerts() MonitorAlertAlertsOutput {
	return o.ApplyT(func(v *MonitorAlert) MonitorAlertAlertsOutput { return v.Alerts }).(MonitorAlertAlertsOutput)
}

// The comparison for `value`.
// This may be either `GreaterThan` or `LessThan`.
func (o MonitorAlertOutput) Compare() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitorAlert) pulumi.StringOutput { return v.Compare }).(pulumi.StringOutput)
}

// The description of the alert.
func (o MonitorAlertOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitorAlert) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The status of the alert.
func (o MonitorAlertOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MonitorAlert) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// A list of IDs for the resources to which the alert policy applies.
func (o MonitorAlertOutput) Entities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MonitorAlert) pulumi.StringArrayOutput { return v.Entities }).(pulumi.StringArrayOutput)
}

// A list of tags. When an included tag is added to a resource, the alert policy will apply to it.
func (o MonitorAlertOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MonitorAlert) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The type of the alert.
// This may be one of `v1/insights/droplet/load_1`, `v1/insights/droplet/load_5`, `v1/insights/droplet/load_15`,
// `v1/insights/droplet/memory_utilization_percent`, `v1/insights/droplet/disk_utilization_percent`,
// `v1/insights/droplet/cpu`, `v1/insights/droplet/disk_read`, `v1/insights/droplet/disk_write`,
// `v1/insights/droplet/public_outbound_bandwidth`, `v1/insights/droplet/public_inbound_bandwidth`,
// `v1/insights/droplet/private_outbound_bandwidth`, `v1/insights/droplet/private_inbound_bandwidth`,
// `v1/insights/lbaas/avg_cpu_utilization_percent`, `v1/insights/lbaas/connection_utilization_percent`,
// `v1/insights/lbaas/droplet_health`, `v1/insights/lbaas/tls_connections_per_second_utilization_percent`,
// `v1/dbaas/alerts/load_15_alerts`, `v1/dbaas/alerts/cpu_alerts`, `v1/dbaas/alerts/memory_utilization_alerts`, or
// `v1/dbaas/alerts/disk_utilization_alerts`.
func (o MonitorAlertOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitorAlert) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The uuid of the alert.
func (o MonitorAlertOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitorAlert) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// The value to start alerting at, e.g., 90% or 85Mbps. This is a floating-point number.
// DigitalOcean will show the correct unit in the web panel.
func (o MonitorAlertOutput) Value() pulumi.Float64Output {
	return o.ApplyT(func(v *MonitorAlert) pulumi.Float64Output { return v.Value }).(pulumi.Float64Output)
}

// The time frame of the alert. Either `5m`, `10m`, `30m`, or `1h`.
func (o MonitorAlertOutput) Window() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitorAlert) pulumi.StringOutput { return v.Window }).(pulumi.StringOutput)
}

type MonitorAlertArrayOutput struct{ *pulumi.OutputState }

func (MonitorAlertArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MonitorAlert)(nil)).Elem()
}

func (o MonitorAlertArrayOutput) ToMonitorAlertArrayOutput() MonitorAlertArrayOutput {
	return o
}

func (o MonitorAlertArrayOutput) ToMonitorAlertArrayOutputWithContext(ctx context.Context) MonitorAlertArrayOutput {
	return o
}

func (o MonitorAlertArrayOutput) Index(i pulumi.IntInput) MonitorAlertOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MonitorAlert {
		return vs[0].([]*MonitorAlert)[vs[1].(int)]
	}).(MonitorAlertOutput)
}

type MonitorAlertMapOutput struct{ *pulumi.OutputState }

func (MonitorAlertMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MonitorAlert)(nil)).Elem()
}

func (o MonitorAlertMapOutput) ToMonitorAlertMapOutput() MonitorAlertMapOutput {
	return o
}

func (o MonitorAlertMapOutput) ToMonitorAlertMapOutputWithContext(ctx context.Context) MonitorAlertMapOutput {
	return o
}

func (o MonitorAlertMapOutput) MapIndex(k pulumi.StringInput) MonitorAlertOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MonitorAlert {
		return vs[0].(map[string]*MonitorAlert)[vs[1].(string)]
	}).(MonitorAlertOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorAlertInput)(nil)).Elem(), &MonitorAlert{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorAlertArrayInput)(nil)).Elem(), MonitorAlertArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorAlertMapInput)(nil)).Elem(), MonitorAlertMap{})
	pulumi.RegisterOutputType(MonitorAlertOutput{})
	pulumi.RegisterOutputType(MonitorAlertArrayOutput{})
	pulumi.RegisterOutputType(MonitorAlertMapOutput{})
}
