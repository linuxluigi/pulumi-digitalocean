// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * Provides information on a DigitalOcean database cluster resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const example = digitalocean.getDatabaseCluster({
 *     name: "example-cluster",
 * });
 * export const databaseOutput = example.then(example => example.uri);
 * ```
 */
export function getDatabaseCluster(args: GetDatabaseClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabaseClusterResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("digitalocean:index/getDatabaseCluster:getDatabaseCluster", {
        "name": args.name,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabaseCluster.
 */
export interface GetDatabaseClusterArgs {
    /**
     * The name of the database cluster.
     */
    name: string;
    tags?: string[];
}

/**
 * A collection of values returned by getDatabaseCluster.
 */
export interface GetDatabaseClusterResult {
    /**
     * Name of the cluster's default database.
     */
    readonly database: string;
    /**
     * Database engine used by the cluster (ex. `pg` for PostreSQL).
     */
    readonly engine: string;
    /**
     * Database cluster's hostname.
     */
    readonly host: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Defines when the automatic maintenance should be performed for the database cluster.
     */
    readonly maintenanceWindows: outputs.GetDatabaseClusterMaintenanceWindow[];
    readonly name: string;
    /**
     * Number of nodes that will be included in the cluster.
     */
    readonly nodeCount: number;
    /**
     * Password for the cluster's default user.
     */
    readonly password: string;
    /**
     * Network port that the database cluster is listening on.
     */
    readonly port: number;
    /**
     * Same as `host`, but only accessible from resources within the account and in the same region.
     */
    readonly privateHost: string;
    /**
     * The ID of the VPC where the database cluster is located.
     */
    readonly privateNetworkUuid: string;
    /**
     * Same as `uri`, but only accessible from resources within the account and in the same region.
     */
    readonly privateUri: string;
    /**
     * DigitalOcean region where the cluster will reside.
     */
    readonly region: string;
    /**
     * Database droplet size associated with the cluster (ex. `db-s-1vcpu-1gb`).
     */
    readonly size: string;
    readonly tags?: string[];
    /**
     * The full URI for connecting to the database cluster.
     */
    readonly uri: string;
    /**
     * The uniform resource name of the database cluster.
     */
    readonly urn: string;
    /**
     * Username for the cluster's default user.
     */
    readonly user: string;
    /**
     * Engine version used by the cluster (ex. `11` for PostgreSQL 11).
     */
    readonly version: string;
}
/**
 * Provides information on a DigitalOcean database cluster resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const example = digitalocean.getDatabaseCluster({
 *     name: "example-cluster",
 * });
 * export const databaseOutput = example.then(example => example.uri);
 * ```
 */
export function getDatabaseClusterOutput(args: GetDatabaseClusterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatabaseClusterResult> {
    return pulumi.output(args).apply((a: any) => getDatabaseCluster(a, opts))
}

/**
 * A collection of arguments for invoking getDatabaseCluster.
 */
export interface GetDatabaseClusterOutputArgs {
    /**
     * The name of the database cluster.
     */
    name: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
