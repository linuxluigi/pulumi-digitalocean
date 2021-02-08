// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Get information on an image for use in other resources (e.g. creating a Droplet
 * based on snapshot). This data source provides all of the image properties as
 * configured on your DigitalOcean account. This is useful if the image in question
 * is not managed by the provider or you need to utilize any of the image's data.
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
 * }, { async: true }));
 * ```
 *
 * Reuse the data about a snapshot to create a Droplet:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const exampleImage = digitalocean.getImage({
 *     name: "example-1.0.0",
 * });
 * const exampleDroplet = new digitalocean.Droplet("exampleDroplet", {
 *     image: exampleImage.then(exampleImage => exampleImage.id),
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
 * }, { async: true }));
 * ```
 */
export function getImage(args?: GetImageArgs, opts?: pulumi.InvokeOptions): Promise<GetImageResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("digitalocean:index/getImage:getImage", {
        "id": args.id,
        "name": args.name,
        "slug": args.slug,
        "source": args.source,
    }, opts);
}

/**
 * A collection of arguments for invoking getImage.
 */
export interface GetImageArgs {
    /**
     * The id of the image
     */
    readonly id?: number;
    /**
     * The name of the image.
     */
    readonly name?: string;
    /**
     * The slug of the official image.
     */
    readonly slug?: string;
    /**
     * Restrict the search to one of the following categories of images:
     */
    readonly source?: string;
}

/**
 * A collection of values returned by getImage.
 */
export interface GetImageResult {
    readonly created: string;
    readonly description: string;
    /**
     * The name of the distribution of the OS of the image.
     * * `minDiskSize`: The minimum 'disk' required for the image.
     * * `sizeGigabytes`: The size of the image in GB.
     */
    readonly distribution: string;
    readonly errorMessage: string;
    readonly id: number;
    /**
     * The id of the image (legacy parameter).
     */
    readonly image: string;
    readonly minDiskSize: number;
    readonly name: string;
    /**
     * Is image a public image or not. Public images represent
     * Linux distributions or One-Click Applications, while non-public images represent
     * snapshots and backups and are only available within your account.
     * * `regions`: A set of the regions that the image is available in.
     * * `tags`: A set of tags applied to the image
     * * `created`: When the image was created
     * * `status`: Current status of the image
     * * `errorMessage`: Any applicable error message pertaining to the image
     */
    readonly private: boolean;
    readonly regions: string[];
    readonly sizeGigabytes: number;
    readonly slug: string;
    readonly source?: string;
    readonly status: string;
    readonly tags: string[];
    readonly type: string;
}
