// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Get information on an images for use in other resources (e.g. creating a Droplet
 * based on snapshot). This data source provides all of the image properties as
 * configured on your DigitalOcean account. This is useful if the image in question
 * is not managed by Terraform or you need to utilize any of the image's data.
 * 
 * An error is triggered if zero or more than one result is returned by the query.
 * 
 * ## Example Usage
 * 
 * Get the data about a snapshot:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 * 
 * const example1 = pulumi.output(digitalocean.getImage({
 *     name: "example-1.0.0",
 * }));
 * ```
 * 
 * Reuse the data about a snapshot to create a Droplet:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 * 
 * const example1Image = pulumi.output(digitalocean.getImage({
 *     name: "example-1.0.0",
 * }));
 * const example1Droplet = new digitalocean.Droplet("example1", {
 *     image: example1Image.image,
 *     region: "nyc2",
 *     size: "s-1vcpu-1gb",
 * });
 * ```
 * 
 * Get the data about an official image:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 * 
 * const example2 = pulumi.output(digitalocean.getImage({
 *     slug: "ubuntu-18-04-x64",
 * }));
 * ```
 */
export function getImage(args?: GetImageArgs, opts?: pulumi.InvokeOptions): Promise<GetImageResult> {
    args = args || {};
    return pulumi.runtime.invoke("digitalocean:index/getImage:getImage", {
        "name": args.name,
        "slug": args.slug,
    }, opts);
}

/**
 * A collection of arguments for invoking getImage.
 */
export interface GetImageArgs {
    /**
     * The name of the private image.
     */
    readonly name?: string;
    /**
     * The slug of the official image.
     */
    readonly slug?: string;
}

/**
 * A collection of values returned by getImage.
 */
export interface GetImageResult {
    /**
     * The name of the distribution of the OS of the image.
     * * `min_disk_size`: The minimum 'disk' required for the image.
     */
    readonly distribution: string;
    /**
     * The id of the image.
     */
    readonly image: string;
    readonly minDiskSize: number;
    readonly name?: string;
    /**
     * Is image a public image or not. Public images represent
     * Linux distributions or One-Click Applications, while non-public images represent
     * snapshots and backups and are only available within your account.
     * * `regions`: The regions that the image is available in.
     * * `type`: Type of the image.
     */
    readonly private: boolean;
    readonly regions: string[];
    readonly slug?: string;
    readonly type: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}