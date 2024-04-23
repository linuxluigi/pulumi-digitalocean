// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Get information on a Spaces bucket for use in other resources. This is useful if the Spaces bucket in question
 * is not managed by the provider or you need to utilize any of the bucket's data.
 *
 * ## Example Usage
 *
 * Get the bucket by name:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const example = digitalocean.getSpacesBucket({
 *     name: "my-spaces-bucket",
 *     region: "nyc3",
 * });
 * export const bucketDomainName = example.then(example => example.bucketDomainName);
 * ```
 */
export function getSpacesBucket(args: GetSpacesBucketArgs, opts?: pulumi.InvokeOptions): Promise<GetSpacesBucketResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean:index/getSpacesBucket:getSpacesBucket", {
        "name": args.name,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getSpacesBucket.
 */
export interface GetSpacesBucketArgs {
    /**
     * The name of the Spaces bucket.
     */
    name: string;
    /**
     * The slug of the region where the bucket is stored.
     */
    region: string;
}

/**
 * A collection of values returned by getSpacesBucket.
 */
export interface GetSpacesBucketResult {
    /**
     * The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com)
     */
    readonly bucketDomainName: string;
    /**
     * The FQDN of the bucket without the bucket name (e.g. nyc3.digitaloceanspaces.com)
     */
    readonly endpoint: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the Spaces bucket
     */
    readonly name: string;
    /**
     * The slug of the region where the bucket is stored.
     */
    readonly region: string;
    /**
     * The uniform resource name of the bucket
     */
    readonly urn: string;
}
/**
 * Get information on a Spaces bucket for use in other resources. This is useful if the Spaces bucket in question
 * is not managed by the provider or you need to utilize any of the bucket's data.
 *
 * ## Example Usage
 *
 * Get the bucket by name:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const example = digitalocean.getSpacesBucket({
 *     name: "my-spaces-bucket",
 *     region: "nyc3",
 * });
 * export const bucketDomainName = example.then(example => example.bucketDomainName);
 * ```
 */
export function getSpacesBucketOutput(args: GetSpacesBucketOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSpacesBucketResult> {
    return pulumi.output(args).apply((a: any) => getSpacesBucket(a, opts))
}

/**
 * A collection of arguments for invoking getSpacesBucket.
 */
export interface GetSpacesBucketOutputArgs {
    /**
     * The name of the Spaces bucket.
     */
    name: pulumi.Input<string>;
    /**
     * The slug of the region where the bucket is stored.
     */
    region: pulumi.Input<string>;
}
