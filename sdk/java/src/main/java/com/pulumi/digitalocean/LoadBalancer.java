// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.digitalocean.LoadBalancerArgs;
import com.pulumi.digitalocean.Utilities;
import com.pulumi.digitalocean.inputs.LoadBalancerState;
import com.pulumi.digitalocean.outputs.LoadBalancerForwardingRule;
import com.pulumi.digitalocean.outputs.LoadBalancerHealthcheck;
import com.pulumi.digitalocean.outputs.LoadBalancerStickySessions;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a DigitalOcean Load Balancer resource. This can be used to create,
 * modify, and delete Load Balancers.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.digitalocean.Droplet;
 * import com.pulumi.digitalocean.DropletArgs;
 * import com.pulumi.digitalocean.LoadBalancer;
 * import com.pulumi.digitalocean.LoadBalancerArgs;
 * import com.pulumi.digitalocean.inputs.LoadBalancerForwardingRuleArgs;
 * import com.pulumi.digitalocean.inputs.LoadBalancerHealthcheckArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var web = new Droplet(&#34;web&#34;, DropletArgs.builder()        
 *             .size(&#34;s-1vcpu-1gb&#34;)
 *             .image(&#34;ubuntu-18-04-x64&#34;)
 *             .region(&#34;nyc3&#34;)
 *             .build());
 * 
 *         var public_ = new LoadBalancer(&#34;public&#34;, LoadBalancerArgs.builder()        
 *             .region(&#34;nyc3&#34;)
 *             .forwardingRules(LoadBalancerForwardingRuleArgs.builder()
 *                 .entryPort(80)
 *                 .entryProtocol(&#34;http&#34;)
 *                 .targetPort(80)
 *                 .targetProtocol(&#34;http&#34;)
 *                 .build())
 *             .healthcheck(LoadBalancerHealthcheckArgs.builder()
 *                 .port(22)
 *                 .protocol(&#34;tcp&#34;)
 *                 .build())
 *             .dropletIds(web.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * When managing certificates attached to the load balancer, make sure to add the `create_before_destroy`
 * lifecycle property in order to ensure the certificate is correctly updated when changed. The order of
 * operations will then be: `Create new certificate` &gt; `Update loadbalancer with new certificate` -&gt;
 * `Delete old certificate`. When doing so, you must also change the name of the certificate,
 * as there cannot be multiple certificates with the same name in an account.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.digitalocean.Certificate;
 * import com.pulumi.digitalocean.CertificateArgs;
 * import com.pulumi.digitalocean.Droplet;
 * import com.pulumi.digitalocean.DropletArgs;
 * import com.pulumi.digitalocean.LoadBalancer;
 * import com.pulumi.digitalocean.LoadBalancerArgs;
 * import com.pulumi.digitalocean.inputs.LoadBalancerForwardingRuleArgs;
 * import com.pulumi.digitalocean.inputs.LoadBalancerHealthcheckArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var cert = new Certificate(&#34;cert&#34;, CertificateArgs.builder()        
 *             .privateKey(&#34;file(&#39;key.pem&#39;)&#34;)
 *             .leafCertificate(&#34;file(&#39;cert.pem&#39;)&#34;)
 *             .build());
 * 
 *         var web = new Droplet(&#34;web&#34;, DropletArgs.builder()        
 *             .size(&#34;s-1vcpu-1gb&#34;)
 *             .image(&#34;ubuntu-18-04-x64&#34;)
 *             .region(&#34;nyc3&#34;)
 *             .build());
 * 
 *         var public_ = new LoadBalancer(&#34;public&#34;, LoadBalancerArgs.builder()        
 *             .region(&#34;nyc3&#34;)
 *             .forwardingRules(LoadBalancerForwardingRuleArgs.builder()
 *                 .entryPort(443)
 *                 .entryProtocol(&#34;https&#34;)
 *                 .targetPort(80)
 *                 .targetProtocol(&#34;http&#34;)
 *                 .certificateName(cert.name())
 *                 .build())
 *             .healthcheck(LoadBalancerHealthcheckArgs.builder()
 *                 .port(22)
 *                 .protocol(&#34;tcp&#34;)
 *                 .build())
 *             .dropletIds(web.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Load Balancers can be imported using the `id`, e.g.
 * 
 * ```sh
 *  $ pulumi import digitalocean:index/loadBalancer:LoadBalancer myloadbalancer 4de7ac8b-495b-4884-9a69-1050c6793cd6
 * ```
 * 
 */
@ResourceType(type="digitalocean:index/loadBalancer:LoadBalancer")
public class LoadBalancer extends com.pulumi.resources.CustomResource {
    /**
     * The load balancing algorithm used to determine
     * which backend Droplet will be selected by a client. It must be either `round_robin`
     * or `least_connections`. The default value is `round_robin`.
     * 
     */
    @Export(name="algorithm", type=String.class, parameters={})
    private Output</* @Nullable */ String> algorithm;

    /**
     * @return The load balancing algorithm used to determine
     * which backend Droplet will be selected by a client. It must be either `round_robin`
     * or `least_connections`. The default value is `round_robin`.
     * 
     */
    public Output<Optional<String>> algorithm() {
        return Codegen.optional(this.algorithm);
    }
    /**
     * A boolean value indicating whether to disable automatic DNS record creation for Let&#39;s Encrypt certificates that are added to the load balancer. Default value is `false`.
     * 
     */
    @Export(name="disableLetsEncryptDnsRecords", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> disableLetsEncryptDnsRecords;

    /**
     * @return A boolean value indicating whether to disable automatic DNS record creation for Let&#39;s Encrypt certificates that are added to the load balancer. Default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> disableLetsEncryptDnsRecords() {
        return Codegen.optional(this.disableLetsEncryptDnsRecords);
    }
    /**
     * A list of the IDs of each droplet to be attached to the Load Balancer.
     * 
     */
    @Export(name="dropletIds", type=List.class, parameters={Integer.class})
    private Output<List<Integer>> dropletIds;

    /**
     * @return A list of the IDs of each droplet to be attached to the Load Balancer.
     * 
     */
    public Output<List<Integer>> dropletIds() {
        return this.dropletIds;
    }
    /**
     * The name of a Droplet tag corresponding to Droplets to be assigned to the Load Balancer.
     * 
     */
    @Export(name="dropletTag", type=String.class, parameters={})
    private Output</* @Nullable */ String> dropletTag;

    /**
     * @return The name of a Droplet tag corresponding to Droplets to be assigned to the Load Balancer.
     * 
     */
    public Output<Optional<String>> dropletTag() {
        return Codegen.optional(this.dropletTag);
    }
    /**
     * A boolean value indicating whether HTTP keepalive connections are maintained to target Droplets. Default value is `false`.
     * 
     */
    @Export(name="enableBackendKeepalive", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enableBackendKeepalive;

    /**
     * @return A boolean value indicating whether HTTP keepalive connections are maintained to target Droplets. Default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> enableBackendKeepalive() {
        return Codegen.optional(this.enableBackendKeepalive);
    }
    /**
     * A boolean value indicating whether PROXY
     * Protocol should be used to pass information from connecting client requests to
     * the backend service. Default value is `false`.
     * 
     */
    @Export(name="enableProxyProtocol", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enableProxyProtocol;

    /**
     * @return A boolean value indicating whether PROXY
     * Protocol should be used to pass information from connecting client requests to
     * the backend service. Default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> enableProxyProtocol() {
        return Codegen.optional(this.enableProxyProtocol);
    }
    /**
     * A list of `forwarding_rule` to be assigned to the
     * Load Balancer. The `forwarding_rule` block is documented below.
     * 
     */
    @Export(name="forwardingRules", type=List.class, parameters={LoadBalancerForwardingRule.class})
    private Output<List<LoadBalancerForwardingRule>> forwardingRules;

    /**
     * @return A list of `forwarding_rule` to be assigned to the
     * Load Balancer. The `forwarding_rule` block is documented below.
     * 
     */
    public Output<List<LoadBalancerForwardingRule>> forwardingRules() {
        return this.forwardingRules;
    }
    /**
     * A `healthcheck` block to be assigned to the
     * Load Balancer. The `healthcheck` block is documented below. Only 1 healthcheck is allowed.
     * 
     */
    @Export(name="healthcheck", type=LoadBalancerHealthcheck.class, parameters={})
    private Output<LoadBalancerHealthcheck> healthcheck;

    /**
     * @return A `healthcheck` block to be assigned to the
     * Load Balancer. The `healthcheck` block is documented below. Only 1 healthcheck is allowed.
     * 
     */
    public Output<LoadBalancerHealthcheck> healthcheck() {
        return this.healthcheck;
    }
    /**
     * Specifies the idle timeout for HTTPS connections on the load balancer in seconds.
     * 
     */
    @Export(name="httpIdleTimeoutSeconds", type=Integer.class, parameters={})
    private Output<Integer> httpIdleTimeoutSeconds;

    /**
     * @return Specifies the idle timeout for HTTPS connections on the load balancer in seconds.
     * 
     */
    public Output<Integer> httpIdleTimeoutSeconds() {
        return this.httpIdleTimeoutSeconds;
    }
    /**
     * The ip of the Load Balancer
     * 
     */
    @Export(name="ip", type=String.class, parameters={})
    private Output<String> ip;

    /**
     * @return The ip of the Load Balancer
     * 
     */
    public Output<String> ip() {
        return this.ip;
    }
    /**
     * The uniform resource name for the Load Balancer
     * 
     */
    @Export(name="loadBalancerUrn", type=String.class, parameters={})
    private Output<String> loadBalancerUrn;

    /**
     * @return The uniform resource name for the Load Balancer
     * 
     */
    public Output<String> loadBalancerUrn() {
        return this.loadBalancerUrn;
    }
    /**
     * The Load Balancer name
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The Load Balancer name
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID of the project that the load balancer is associated with. If no ID is provided at creation, the load balancer associates with the user&#39;s default project.
     * 
     */
    @Export(name="projectId", type=String.class, parameters={})
    private Output<String> projectId;

    /**
     * @return The ID of the project that the load balancer is associated with. If no ID is provided at creation, the load balancer associates with the user&#39;s default project.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * A boolean value indicating whether
     * HTTP requests to the Load Balancer on port 80 will be redirected to HTTPS on port 443.
     * Default value is `false`.
     * 
     */
    @Export(name="redirectHttpToHttps", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> redirectHttpToHttps;

    /**
     * @return A boolean value indicating whether
     * HTTP requests to the Load Balancer on port 80 will be redirected to HTTPS on port 443.
     * Default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> redirectHttpToHttps() {
        return Codegen.optional(this.redirectHttpToHttps);
    }
    /**
     * The region to start in
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return The region to start in
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The size of the Load Balancer. It must be either `lb-small`, `lb-medium`, or `lb-large`. Defaults to `lb-small`. Only one of `size` or `size_unit` may be provided.
     * 
     */
    @Export(name="size", type=String.class, parameters={})
    private Output</* @Nullable */ String> size;

    /**
     * @return The size of the Load Balancer. It must be either `lb-small`, `lb-medium`, or `lb-large`. Defaults to `lb-small`. Only one of `size` or `size_unit` may be provided.
     * 
     */
    public Output<Optional<String>> size() {
        return Codegen.optional(this.size);
    }
    /**
     * The size of the Load Balancer. It must be in the range (1, 100). Defaults to `1`. Only one of `size` or `size_unit` may be provided.
     * 
     */
    @Export(name="sizeUnit", type=Integer.class, parameters={})
    private Output<Integer> sizeUnit;

    /**
     * @return The size of the Load Balancer. It must be in the range (1, 100). Defaults to `1`. Only one of `size` or `size_unit` may be provided.
     * 
     */
    public Output<Integer> sizeUnit() {
        return this.sizeUnit;
    }
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    public Output<String> status() {
        return this.status;
    }
    /**
     * A `sticky_sessions` block to be assigned to the
     * Load Balancer. The `sticky_sessions` block is documented below. Only 1 sticky_sessions block is allowed.
     * 
     */
    @Export(name="stickySessions", type=LoadBalancerStickySessions.class, parameters={})
    private Output<LoadBalancerStickySessions> stickySessions;

    /**
     * @return A `sticky_sessions` block to be assigned to the
     * Load Balancer. The `sticky_sessions` block is documented below. Only 1 sticky_sessions block is allowed.
     * 
     */
    public Output<LoadBalancerStickySessions> stickySessions() {
        return this.stickySessions;
    }
    /**
     * The ID of the VPC where the load balancer will be located.
     * 
     */
    @Export(name="vpcUuid", type=String.class, parameters={})
    private Output<String> vpcUuid;

    /**
     * @return The ID of the VPC where the load balancer will be located.
     * 
     */
    public Output<String> vpcUuid() {
        return this.vpcUuid;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LoadBalancer(String name) {
        this(name, LoadBalancerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LoadBalancer(String name, LoadBalancerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LoadBalancer(String name, LoadBalancerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("digitalocean:index/loadBalancer:LoadBalancer", name, args == null ? LoadBalancerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LoadBalancer(String name, Output<String> id, @Nullable LoadBalancerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("digitalocean:index/loadBalancer:LoadBalancer", name, state, makeResourceOptions(options, id));
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
    public static LoadBalancer get(String name, Output<String> id, @Nullable LoadBalancerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LoadBalancer(name, id, state, options);
    }
}
