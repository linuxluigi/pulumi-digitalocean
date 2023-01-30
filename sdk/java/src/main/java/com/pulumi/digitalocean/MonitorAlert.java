// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.digitalocean.MonitorAlertArgs;
import com.pulumi.digitalocean.Utilities;
import com.pulumi.digitalocean.inputs.MonitorAlertState;
import com.pulumi.digitalocean.outputs.MonitorAlertAlerts;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a [DigitalOcean Monitoring](https://docs.digitalocean.com/reference/api/api-reference/#tag/Monitoring)
 * resource. Monitor alerts can be configured to alert about, e.g., disk or memory
 * usage exceeding a certain threshold or traffic at a certain limit. Notifications
 * can be sent to either an email address or a Slack channel.
 * 
 * &gt; **Note** Currently, the [DigitalOcean API](https://docs.digitalocean.com/reference/api/api-reference/#operation/create_alert_policy) only supports creating alerts for Droplets.
 * 
 * ## Import
 * 
 * Monitor alerts can be imported using the monitor alert `uuid`, e.g.
 * 
 * ```sh
 *  $ pulumi import digitalocean:index/monitorAlert:MonitorAlert cpu_alert b8ecd2ab-2267-4a5e-8692-cbf1d32583e3
 * ```
 * 
 */
@ResourceType(type="digitalocean:index/monitorAlert:MonitorAlert")
public class MonitorAlert extends com.pulumi.resources.CustomResource {
    /**
     * How to send notifications about the alerts. This is a list with one element, .
     * Note that for Slack, the DigitalOcean app needs to have permissions for your workspace. You can
     * read more in [Slack&#39;s documentation](https://slack.com/intl/en-dk/help/articles/222386767-Manage-app-installation-settings-for-your-workspace)
     * 
     */
    @Export(name="alerts", type=MonitorAlertAlerts.class, parameters={})
    private Output<MonitorAlertAlerts> alerts;

    /**
     * @return How to send notifications about the alerts. This is a list with one element, .
     * Note that for Slack, the DigitalOcean app needs to have permissions for your workspace. You can
     * read more in [Slack&#39;s documentation](https://slack.com/intl/en-dk/help/articles/222386767-Manage-app-installation-settings-for-your-workspace)
     * 
     */
    public Output<MonitorAlertAlerts> alerts() {
        return this.alerts;
    }
    /**
     * The comparison for `value`.
     * This may be either `GreaterThan` or `LessThan`.
     * 
     */
    @Export(name="compare", type=String.class, parameters={})
    private Output<String> compare;

    /**
     * @return The comparison for `value`.
     * This may be either `GreaterThan` or `LessThan`.
     * 
     */
    public Output<String> compare() {
        return this.compare;
    }
    /**
     * The description of the alert.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output<String> description;

    /**
     * @return The description of the alert.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The status of the alert.
     * 
     */
    @Export(name="enabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return The status of the alert.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * A list of IDs for the resources to which the alert policy applies.
     * 
     */
    @Export(name="entities", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> entities;

    /**
     * @return A list of IDs for the resources to which the alert policy applies.
     * 
     */
    public Output<Optional<List<String>>> entities() {
        return Codegen.optional(this.entities);
    }
    /**
     * A list of tags. When an included tag is added to a resource, the alert policy will apply to it.
     * 
     */
    @Export(name="tags", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return A list of tags. When an included tag is added to a resource, the alert policy will apply to it.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The type of the alert.
     * This may be one of `v1/insights/droplet/load_1`, `v1/insights/droplet/load_5`, `v1/insights/droplet/load_15`,
     * `v1/insights/droplet/memory_utilization_percent`, `v1/insights/droplet/disk_utilization_percent`,
     * `v1/insights/droplet/cpu`, `v1/insights/droplet/disk_read`, `v1/insights/droplet/disk_write`,
     * `v1/insights/droplet/public_outbound_bandwidth`, `v1/insights/droplet/public_inbound_bandwidth`,
     * `v1/insights/droplet/private_outbound_bandwidth`, `v1/insights/droplet/private_inbound_bandwidth`,
     * `v1/insights/lbaas/avg_cpu_utilization_percent`, `v1/insights/lbaas/connection_utilization_percent`,
     * `v1/insights/lbaas/droplet_health`, `v1/insights/lbaas/tls_connections_per_second_utilization_percent`,
     * `v1/dbaas/alerts/load_15_alerts`, `v1/dbaas/alerts/cpu_alerts`, `v1/dbaas/alerts/memory_utilization_alerts`, or
     * `v1/dbaas/alerts/disk_utilization_alerts`.
     * 
     */
    @Export(name="type", type=String.class, parameters={})
    private Output<String> type;

    /**
     * @return The type of the alert.
     * This may be one of `v1/insights/droplet/load_1`, `v1/insights/droplet/load_5`, `v1/insights/droplet/load_15`,
     * `v1/insights/droplet/memory_utilization_percent`, `v1/insights/droplet/disk_utilization_percent`,
     * `v1/insights/droplet/cpu`, `v1/insights/droplet/disk_read`, `v1/insights/droplet/disk_write`,
     * `v1/insights/droplet/public_outbound_bandwidth`, `v1/insights/droplet/public_inbound_bandwidth`,
     * `v1/insights/droplet/private_outbound_bandwidth`, `v1/insights/droplet/private_inbound_bandwidth`,
     * `v1/insights/lbaas/avg_cpu_utilization_percent`, `v1/insights/lbaas/connection_utilization_percent`,
     * `v1/insights/lbaas/droplet_health`, `v1/insights/lbaas/tls_connections_per_second_utilization_percent`,
     * `v1/dbaas/alerts/load_15_alerts`, `v1/dbaas/alerts/cpu_alerts`, `v1/dbaas/alerts/memory_utilization_alerts`, or
     * `v1/dbaas/alerts/disk_utilization_alerts`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * The uuid of the alert.
     * 
     */
    @Export(name="uuid", type=String.class, parameters={})
    private Output<String> uuid;

    /**
     * @return The uuid of the alert.
     * 
     */
    public Output<String> uuid() {
        return this.uuid;
    }
    /**
     * The value to start alerting at, e.g., 90% or 85Mbps. This is a floating-point number.
     * DigitalOcean will show the correct unit in the web panel.
     * 
     */
    @Export(name="value", type=Double.class, parameters={})
    private Output<Double> value;

    /**
     * @return The value to start alerting at, e.g., 90% or 85Mbps. This is a floating-point number.
     * DigitalOcean will show the correct unit in the web panel.
     * 
     */
    public Output<Double> value() {
        return this.value;
    }
    /**
     * The time frame of the alert. Either `5m`, `10m`, `30m`, or `1h`.
     * 
     */
    @Export(name="window", type=String.class, parameters={})
    private Output<String> window;

    /**
     * @return The time frame of the alert. Either `5m`, `10m`, `30m`, or `1h`.
     * 
     */
    public Output<String> window() {
        return this.window;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MonitorAlert(String name) {
        this(name, MonitorAlertArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MonitorAlert(String name, MonitorAlertArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MonitorAlert(String name, MonitorAlertArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("digitalocean:index/monitorAlert:MonitorAlert", name, args == null ? MonitorAlertArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MonitorAlert(String name, Output<String> id, @Nullable MonitorAlertState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("digitalocean:index/monitorAlert:MonitorAlert", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static MonitorAlert get(String name, Output<String> id, @Nullable MonitorAlertState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MonitorAlert(name, id, state, options);
    }
}
