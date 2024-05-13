// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.digitalocean.SpacesBucketPolicyArgs;
import com.pulumi.digitalocean.Utilities;
import com.pulumi.digitalocean.inputs.SpacesBucketPolicyState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * ### Limiting access to specific IP addresses
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.digitalocean.SpacesBucket;
 * import com.pulumi.digitalocean.SpacesBucketArgs;
 * import com.pulumi.digitalocean.SpacesBucketPolicy;
 * import com.pulumi.digitalocean.SpacesBucketPolicyArgs;
 * import static com.pulumi.codegen.internal.Serialization.*;
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
 *         var foobarSpacesBucket = new SpacesBucket("foobarSpacesBucket", SpacesBucketArgs.builder()        
 *             .region("nyc3")
 *             .build());
 * 
 *         var foobarSpacesBucketPolicy = new SpacesBucketPolicy("foobarSpacesBucketPolicy", SpacesBucketPolicyArgs.builder()        
 *             .region(foobarSpacesBucket.region())
 *             .bucket(foobarSpacesBucket.name())
 *             .policy(Output.tuple(foobarSpacesBucket.name(), foobarSpacesBucket.name()).applyValue(values -> {
 *                 var foobarSpacesBucketName = values.t1;
 *                 var foobarSpacesBucketName1 = values.t2;
 *                 return serializeJson(
 *                     jsonObject(
 *                         jsonProperty("Version", "2012-10-17"),
 *                         jsonProperty("Statement", jsonArray(jsonObject(
 *                             jsonProperty("Sid", "IPAllow"),
 *                             jsonProperty("Effect", "Deny"),
 *                             jsonProperty("Principal", "*"),
 *                             jsonProperty("Action", "s3:*"),
 *                             jsonProperty("Resource", jsonArray(
 *                                 String.format("arn:aws:s3:::%s", foobarSpacesBucketName), 
 *                                 String.format("arn:aws:s3:::%s/*", foobarSpacesBucketName1)
 *                             )),
 *                             jsonProperty("Condition", jsonObject(
 *                                 jsonProperty("NotIpAddress", jsonObject(
 *                                     jsonProperty("aws:SourceIp", "54.240.143.0/24")
 *                                 ))
 *                             ))
 *                         )))
 *                     ));
 *             }))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * !&gt; **Warning:** Before using this policy, replace the 54.240.143.0/24 IP address range in this example with an appropriate value for your use case. Otherwise, you will lose the ability to access your bucket.
 * 
 * ## Import
 * 
 * Bucket policies can be imported using the `region` and `bucket` attributes (delimited by a comma):
 * 
 * ```sh
 * $ pulumi import digitalocean:index/spacesBucketPolicy:SpacesBucketPolicy foobar `region`,`bucket`
 * ```
 * 
 */
@ResourceType(type="digitalocean:index/spacesBucketPolicy:SpacesBucketPolicy")
public class SpacesBucketPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The name of the bucket to which to apply the policy.
     * 
     */
    @Export(name="bucket", refs={String.class}, tree="[0]")
    private Output<String> bucket;

    /**
     * @return The name of the bucket to which to apply the policy.
     * 
     */
    public Output<String> bucket() {
        return this.bucket;
    }
    /**
     * The text of the policy.
     * 
     */
    @Export(name="policy", refs={String.class}, tree="[0]")
    private Output<String> policy;

    /**
     * @return The text of the policy.
     * 
     */
    public Output<String> policy() {
        return this.policy;
    }
    /**
     * The region where the bucket resides.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region where the bucket resides.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SpacesBucketPolicy(String name) {
        this(name, SpacesBucketPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SpacesBucketPolicy(String name, SpacesBucketPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SpacesBucketPolicy(String name, SpacesBucketPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("digitalocean:index/spacesBucketPolicy:SpacesBucketPolicy", name, args == null ? SpacesBucketPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SpacesBucketPolicy(String name, Output<String> id, @Nullable SpacesBucketPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("digitalocean:index/spacesBucketPolicy:SpacesBucketPolicy", name, state, makeResourceOptions(options, id));
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
    public static SpacesBucketPolicy get(String name, Output<String> id, @Nullable SpacesBucketPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SpacesBucketPolicy(name, id, state, options);
    }
}
