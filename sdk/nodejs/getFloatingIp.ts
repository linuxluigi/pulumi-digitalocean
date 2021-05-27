// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

/**
 * Get information on a floating ip. This data source provides the region and Droplet id
 * as configured on your DigitalOcean account. This is useful if the floating IP
 * in question is not managed by the provider or you need to find the Droplet the IP is
 * attached to.
 *
 * An error is triggered if the provided floating IP does not exist.
 *
 * ## Example Usage
 *
 * Get the floating IP:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const config = new pulumi.Config();
 * const publicIp = config.requireObject("publicIp");
 * const example = digitalocean.getFloatingIp({
 *     ipAddress: publicIp,
 * });
 * export const fipOutput = example.then(example => example.dropletId);
 * ```
 */
export function getFloatingIp(args: GetFloatingIpArgs, opts?: pulumi.InvokeOptions): Promise<GetFloatingIpResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("digitalocean:index/getFloatingIp:getFloatingIp", {
        "ipAddress": args.ipAddress,
    }, opts);
}

/**
 * A collection of arguments for invoking getFloatingIp.
 */
export interface GetFloatingIpArgs {
    /**
     * The allocated IP address of the specific floating IP to retrieve.
     */
    ipAddress: string;
}

/**
 * A collection of values returned by getFloatingIp.
 */
export interface GetFloatingIpResult {
    readonly dropletId: number;
    readonly floatingIpUrn: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipAddress: string;
    readonly region: string;
}
