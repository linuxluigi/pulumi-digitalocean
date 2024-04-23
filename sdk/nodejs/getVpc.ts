// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Retrieve information about a VPC for use in other resources.
 *
 * This data source provides all of the VPC's properties as configured on your
 * DigitalOcean account. This is useful if the VPC in question is not managed by
 * the provider or you need to utilize any of the VPC's data.
 *
 * VPCs may be looked up by `id` or `name`. Specifying a `region` will
 * return that that region's default VPC.
 *
 * ## Example Usage
 *
 * ### VPC By Name
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const example = digitalocean.getVpc({
 *     name: "example-network",
 * });
 * ```
 *
 * Reuse the data about a VPC to assign a Droplet to it:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const exampleVpc = digitalocean.getVpc({
 *     name: "example-network",
 * });
 * const exampleDroplet = new digitalocean.Droplet("exampleDroplet", {
 *     size: digitalocean.DropletSlug.DropletS1VCPU1GB,
 *     image: "ubuntu-18-04-x64",
 *     region: digitalocean.Region.NYC3,
 *     vpcUuid: exampleVpc.then(exampleVpc => exampleVpc.id),
 * });
 * ```
 */
export function getVpc(args?: GetVpcArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean:index/getVpc:getVpc", {
        "id": args.id,
        "name": args.name,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpc.
 */
export interface GetVpcArgs {
    /**
     * The unique identifier of an existing VPC.
     */
    id?: string;
    /**
     * The name of an existing VPC.
     */
    name?: string;
    /**
     * The DigitalOcean region slug for the VPC's location.
     */
    region?: string;
}

/**
 * A collection of values returned by getVpc.
 */
export interface GetVpcResult {
    /**
     * The date and time of when the VPC was created.
     */
    readonly createdAt: string;
    /**
     * A boolean indicating whether or not the VPC is the default one for the region.
     */
    readonly default: boolean;
    /**
     * A free-form text field describing the VPC.
     */
    readonly description: string;
    /**
     * The unique identifier for the VPC.
     */
    readonly id: string;
    /**
     * The range of IP addresses for the VPC in CIDR notation.
     */
    readonly ipRange: string;
    /**
     * The name of the VPC.
     */
    readonly name: string;
    /**
     * The DigitalOcean region slug for the VPC's location.
     */
    readonly region: string;
    /**
     * The uniform resource name (URN) for the VPC.
     */
    readonly urn: string;
}
/**
 * Retrieve information about a VPC for use in other resources.
 *
 * This data source provides all of the VPC's properties as configured on your
 * DigitalOcean account. This is useful if the VPC in question is not managed by
 * the provider or you need to utilize any of the VPC's data.
 *
 * VPCs may be looked up by `id` or `name`. Specifying a `region` will
 * return that that region's default VPC.
 *
 * ## Example Usage
 *
 * ### VPC By Name
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const example = digitalocean.getVpc({
 *     name: "example-network",
 * });
 * ```
 *
 * Reuse the data about a VPC to assign a Droplet to it:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const exampleVpc = digitalocean.getVpc({
 *     name: "example-network",
 * });
 * const exampleDroplet = new digitalocean.Droplet("exampleDroplet", {
 *     size: digitalocean.DropletSlug.DropletS1VCPU1GB,
 *     image: "ubuntu-18-04-x64",
 *     region: digitalocean.Region.NYC3,
 *     vpcUuid: exampleVpc.then(exampleVpc => exampleVpc.id),
 * });
 * ```
 */
export function getVpcOutput(args?: GetVpcOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVpcResult> {
    return pulumi.output(args).apply((a: any) => getVpc(a, opts))
}

/**
 * A collection of arguments for invoking getVpc.
 */
export interface GetVpcOutputArgs {
    /**
     * The unique identifier of an existing VPC.
     */
    id?: pulumi.Input<string>;
    /**
     * The name of an existing VPC.
     */
    name?: pulumi.Input<string>;
    /**
     * The DigitalOcean region slug for the VPC's location.
     */
    region?: pulumi.Input<string>;
}
