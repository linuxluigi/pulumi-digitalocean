// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.digitalocean.DomainArgs;
import com.pulumi.digitalocean.Utilities;
import com.pulumi.digitalocean.inputs.DomainState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a DigitalOcean domain resource.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.digitalocean.Domain;
 * import com.pulumi.digitalocean.DomainArgs;
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
 *         // Create a new domain
 *         var default_ = new Domain(&#34;default&#34;, DomainArgs.builder()        
 *             .name(&#34;example.com&#34;)
 *             .ipAddress(digitalocean_droplet.foo().ipv4_address())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Domains can be imported using the `domain name`, e.g.
 * 
 * ```sh
 * $ pulumi import digitalocean:index/domain:Domain mydomain mytestdomain.com
 * ```
 * 
 */
@ResourceType(type="digitalocean:index/domain:Domain")
public class Domain extends com.pulumi.resources.CustomResource {
    /**
     * The uniform resource name of the domain
     * 
     */
    @Export(name="domainUrn", refs={String.class}, tree="[0]")
    private Output<String> domainUrn;

    /**
     * @return The uniform resource name of the domain
     * 
     */
    public Output<String> domainUrn() {
        return this.domainUrn;
    }
    /**
     * The IP address of the domain. If specified, this IP
     * is used to created an initial A record for the domain.
     * 
     */
    @Export(name="ipAddress", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ipAddress;

    /**
     * @return The IP address of the domain. If specified, this IP
     * is used to created an initial A record for the domain.
     * 
     */
    public Output<Optional<String>> ipAddress() {
        return Codegen.optional(this.ipAddress);
    }
    /**
     * The name of the domain
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the domain
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The TTL value of the domain
     * 
     */
    @Export(name="ttl", refs={Integer.class}, tree="[0]")
    private Output<Integer> ttl;

    /**
     * @return The TTL value of the domain
     * 
     */
    public Output<Integer> ttl() {
        return this.ttl;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Domain(String name) {
        this(name, DomainArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Domain(String name, DomainArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Domain(String name, DomainArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("digitalocean:index/domain:Domain", name, args == null ? DomainArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Domain(String name, Output<String> id, @Nullable DomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("digitalocean:index/domain:Domain", name, state, makeResourceOptions(options, id));
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
    public static Domain get(String name, Output<String> id, @Nullable DomainState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Domain(name, id, state, options);
    }
}
