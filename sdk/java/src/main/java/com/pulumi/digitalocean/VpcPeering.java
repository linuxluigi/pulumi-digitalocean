// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.digitalocean.Utilities;
import com.pulumi.digitalocean.VpcPeeringArgs;
import com.pulumi.digitalocean.inputs.VpcPeeringState;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides a DigitalOcean VPC Peering resource.
 * 
 * VPC Peerings are used to connect two VPC networks allowing resources in each
 * VPC to communicate with each other privately.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.digitalocean.VpcPeering;
 * import com.pulumi.digitalocean.VpcPeeringArgs;
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
 *         var example = new VpcPeering("example", VpcPeeringArgs.builder()
 *             .name("example-peering")
 *             .vpcIds(            
 *                 vpc1.id(),
 *                 vpc2.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Resource Assignement
 * 
 * You can use the VPC Peering resource to allow communication between resources
 * in different VPCs. For example:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.digitalocean.Vpc;
 * import com.pulumi.digitalocean.VpcArgs;
 * import com.pulumi.digitalocean.VpcPeering;
 * import com.pulumi.digitalocean.VpcPeeringArgs;
 * import com.pulumi.digitalocean.Droplet;
 * import com.pulumi.digitalocean.DropletArgs;
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
 *         var vpc1 = new Vpc("vpc1", VpcArgs.builder()
 *             .name("vpc1")
 *             .region("nyc3")
 *             .build());
 * 
 *         var vpc2 = new Vpc("vpc2", VpcArgs.builder()
 *             .name("vpc2")
 *             .region("nyc3")
 *             .build());
 * 
 *         var example = new VpcPeering("example", VpcPeeringArgs.builder()
 *             .name("example-peering")
 *             .vpcIds(            
 *                 vpc1.id(),
 *                 vpc2.id())
 *             .build());
 * 
 *         var example1 = new Droplet("example1", DropletArgs.builder()
 *             .name("example1")
 *             .size("s-1vcpu-1gb")
 *             .image("ubuntu-18-04-x64")
 *             .region("nyc3")
 *             .vpcUuid(vpc1.id())
 *             .build());
 * 
 *         var example2 = new Droplet("example2", DropletArgs.builder()
 *             .name("example2")
 *             .size("s-1vcpu-1gb")
 *             .image("ubuntu-18-04-x64")
 *             .region("nyc3")
 *             .vpcUuid(vpc2.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * A VPC Peering can be imported using its `id`, e.g.
 * 
 * ```sh
 * $ pulumi import digitalocean:index/vpcPeering:VpcPeering example 771ad360-c017-4b4e-a34e-73934f5f0190
 * ```
 * 
 */
@ResourceType(type="digitalocean:index/vpcPeering:VpcPeering")
public class VpcPeering extends com.pulumi.resources.CustomResource {
    /**
     * The date and time of when the VPC Peering was created.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return The date and time of when the VPC Peering was created.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * A name for the VPC Peering. Must be unique and contain alphanumeric characters, dashes, and periods only.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A name for the VPC Peering. Must be unique and contain alphanumeric characters, dashes, and periods only.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The status of the VPC Peering.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the VPC Peering.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * A set of two VPC IDs to be peered.
     * 
     */
    @Export(name="vpcIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> vpcIds;

    /**
     * @return A set of two VPC IDs to be peered.
     * 
     */
    public Output<List<String>> vpcIds() {
        return this.vpcIds;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VpcPeering(String name) {
        this(name, VpcPeeringArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VpcPeering(String name, VpcPeeringArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VpcPeering(String name, VpcPeeringArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("digitalocean:index/vpcPeering:VpcPeering", name, args == null ? VpcPeeringArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VpcPeering(String name, Output<String> id, @Nullable VpcPeeringState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("digitalocean:index/vpcPeering:VpcPeering", name, state, makeResourceOptions(options, id));
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
    public static VpcPeering get(String name, Output<String> id, @Nullable VpcPeeringState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VpcPeering(name, id, state, options);
    }
}
