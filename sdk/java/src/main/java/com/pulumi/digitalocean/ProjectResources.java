// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.digitalocean.ProjectResourcesArgs;
import com.pulumi.digitalocean.Utilities;
import com.pulumi.digitalocean.inputs.ProjectResourcesState;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Assign resources to a DigitalOcean Project. This is useful if you need to assign resources
 * managed via this provider to a DigitalOcean Project managed outside of the provider.
 * 
 * The following resource types can be associated with a project:
 * 
 * * App Platform Apps
 * * Database Clusters
 * * Domains
 * * Droplets
 * * Floating IPs
 * * Kubernetes Clusters
 * * Load Balancers
 * * Spaces Buckets
 * * Volumes
 * 
 * ## Example Usage
 * 
 * The following example assigns a droplet to a Project managed outside of the provider:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.digitalocean.DigitaloceanFunctions;
 * import com.pulumi.digitalocean.inputs.GetProjectArgs;
 * import com.pulumi.digitalocean.Droplet;
 * import com.pulumi.digitalocean.DropletArgs;
 * import com.pulumi.digitalocean.ProjectResources;
 * import com.pulumi.digitalocean.ProjectResourcesArgs;
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
 *         final var playground = DigitaloceanFunctions.getProject(GetProjectArgs.builder()
 *             .name(&#34;playground&#34;)
 *             .build());
 * 
 *         var foobar = new Droplet(&#34;foobar&#34;, DropletArgs.builder()        
 *             .size(&#34;s-1vcpu-1gb&#34;)
 *             .image(&#34;ubuntu-22-04-x64&#34;)
 *             .region(&#34;nyc3&#34;)
 *             .build());
 * 
 *         var barfoo = new ProjectResources(&#34;barfoo&#34;, ProjectResourcesArgs.builder()        
 *             .project(playground.applyValue(getProjectResult -&gt; getProjectResult.id()))
 *             .resources(foobar.dropletUrn())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Importing this resource is not supported.
 * 
 */
@ResourceType(type="digitalocean:index/projectResources:ProjectResources")
public class ProjectResources extends com.pulumi.resources.CustomResource {
    /**
     * the ID of the project
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return the ID of the project
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * a list of uniform resource names (URNs) for the resources associated with the project
     * 
     */
    @Export(name="resources", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> resources;

    /**
     * @return a list of uniform resource names (URNs) for the resources associated with the project
     * 
     */
    public Output<List<String>> resources() {
        return this.resources;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProjectResources(String name) {
        this(name, ProjectResourcesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProjectResources(String name, ProjectResourcesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProjectResources(String name, ProjectResourcesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("digitalocean:index/projectResources:ProjectResources", name, args == null ? ProjectResourcesArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProjectResources(String name, Output<String> id, @Nullable ProjectResourcesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("digitalocean:index/projectResources:ProjectResources", name, state, makeResourceOptions(options, id));
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
    public static ProjectResources get(String name, Output<String> id, @Nullable ProjectResourcesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProjectResources(name, id, state, options);
    }
}
