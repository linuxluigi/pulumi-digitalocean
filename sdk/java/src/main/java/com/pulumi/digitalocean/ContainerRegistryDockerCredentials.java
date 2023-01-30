// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.digitalocean.ContainerRegistryDockerCredentialsArgs;
import com.pulumi.digitalocean.Utilities;
import com.pulumi.digitalocean.inputs.ContainerRegistryDockerCredentialsState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Get Docker credentials for your DigitalOcean container registry.
 * 
 * An error is triggered if the provided container registry name does not exist.
 * 
 * ## Example Usage
 * ### Basic Example
 * 
 * Get the container registry:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.digitalocean.ContainerRegistryDockerCredentials;
 * import com.pulumi.digitalocean.ContainerRegistryDockerCredentialsArgs;
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
 *         var example = new ContainerRegistryDockerCredentials(&#34;example&#34;, ContainerRegistryDockerCredentialsArgs.builder()        
 *             .registryName(&#34;example&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Docker Provider Example
 * 
 * Use the `endpoint` and `docker_credentials` with the Docker provider:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.digitalocean.DigitaloceanFunctions;
 * import com.pulumi.digitalocean.inputs.GetContainerRegistryArgs;
 * import com.pulumi.digitalocean.ContainerRegistryDockerCredentials;
 * import com.pulumi.digitalocean.ContainerRegistryDockerCredentialsArgs;
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
 *         final var exampleContainerRegistry = DigitaloceanFunctions.getContainerRegistry(GetContainerRegistryArgs.builder()
 *             .name(&#34;example&#34;)
 *             .build());
 * 
 *         var exampleContainerRegistryDockerCredentials = new ContainerRegistryDockerCredentials(&#34;exampleContainerRegistryDockerCredentials&#34;, ContainerRegistryDockerCredentialsArgs.builder()        
 *             .registryName(&#34;example&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="digitalocean:index/containerRegistryDockerCredentials:ContainerRegistryDockerCredentials")
public class ContainerRegistryDockerCredentials extends com.pulumi.resources.CustomResource {
    /**
     * The date and time the registry access token will expire.
     * 
     */
    @Export(name="credentialExpirationTime", type=String.class, parameters={})
    private Output<String> credentialExpirationTime;

    /**
     * @return The date and time the registry access token will expire.
     * 
     */
    public Output<String> credentialExpirationTime() {
        return this.credentialExpirationTime;
    }
    /**
     * Credentials for the container registry.
     * 
     */
    @Export(name="dockerCredentials", type=String.class, parameters={})
    private Output<String> dockerCredentials;

    /**
     * @return Credentials for the container registry.
     * 
     */
    public Output<String> dockerCredentials() {
        return this.dockerCredentials;
    }
    /**
     * The amount of time to pass before the Docker credentials expire in seconds. Defaults to 1576800000, or roughly 50 years. Must be greater than 0 and less than 1576800000.
     * 
     */
    @Export(name="expirySeconds", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> expirySeconds;

    /**
     * @return The amount of time to pass before the Docker credentials expire in seconds. Defaults to 1576800000, or roughly 50 years. Must be greater than 0 and less than 1576800000.
     * 
     */
    public Output<Optional<Integer>> expirySeconds() {
        return Codegen.optional(this.expirySeconds);
    }
    /**
     * The name of the container registry.
     * 
     */
    @Export(name="registryName", type=String.class, parameters={})
    private Output<String> registryName;

    /**
     * @return The name of the container registry.
     * 
     */
    public Output<String> registryName() {
        return this.registryName;
    }
    /**
     * Allow for write access to the container registry. Defaults to false.
     * 
     */
    @Export(name="write", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> write;

    /**
     * @return Allow for write access to the container registry. Defaults to false.
     * 
     */
    public Output<Optional<Boolean>> write() {
        return Codegen.optional(this.write);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ContainerRegistryDockerCredentials(String name) {
        this(name, ContainerRegistryDockerCredentialsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ContainerRegistryDockerCredentials(String name, ContainerRegistryDockerCredentialsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ContainerRegistryDockerCredentials(String name, ContainerRegistryDockerCredentialsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("digitalocean:index/containerRegistryDockerCredentials:ContainerRegistryDockerCredentials", name, args == null ? ContainerRegistryDockerCredentialsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ContainerRegistryDockerCredentials(String name, Output<String> id, @Nullable ContainerRegistryDockerCredentialsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("digitalocean:index/containerRegistryDockerCredentials:ContainerRegistryDockerCredentials", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "dockerCredentials"
            ))
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
    public static ContainerRegistryDockerCredentials get(String name, Output<String> id, @Nullable ContainerRegistryDockerCredentialsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ContainerRegistryDockerCredentials(name, id, state, options);
    }
}
