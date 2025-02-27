// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getReservedIpv6(args: GetReservedIpv6Args, opts?: pulumi.InvokeOptions): Promise<GetReservedIpv6Result> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean:index/getReservedIpv6:getReservedIpv6", {
        "ip": args.ip,
    }, opts);
}

/**
 * A collection of arguments for invoking getReservedIpv6.
 */
export interface GetReservedIpv6Args {
    ip: string;
}

/**
 * A collection of values returned by getReservedIpv6.
 */
export interface GetReservedIpv6Result {
    readonly dropletId: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ip: string;
    readonly regionSlug: string;
    readonly urn: string;
}
export function getReservedIpv6Output(args: GetReservedIpv6OutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetReservedIpv6Result> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("digitalocean:index/getReservedIpv6:getReservedIpv6", {
        "ip": args.ip,
    }, opts);
}

/**
 * A collection of arguments for invoking getReservedIpv6.
 */
export interface GetReservedIpv6OutputArgs {
    ip: pulumi.Input<string>;
}
