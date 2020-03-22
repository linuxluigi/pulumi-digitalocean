// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Retrieve information about all DigitalOcean projects associated with an account, with
 * the ability to filter and sort the results. If no filters are specified, all projects
 * will be returned.
 * 
 * Note: You can use the [`digitalocean..Project`](https://www.terraform.io/docs/providers/do/d/project.html) data source to
 * obtain metadata about a single project if you already know the `id` to retrieve or the unique
 * `name` of the project.
 * 
 * ## Example Usage
 * 
 * Use the `filter` block with a `key` string and `values` list to filter projects.
 * 
 * For example to find all staging environment projects:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 * 
 * const staging = digitalocean.getProjects({
 *     filters: [{
 *         key: "environment",
 *         values: ["Staging"],
 *     }],
 * });
 * ```
 * 
 * You can filter on multiple fields and sort the results as well:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 * 
 * const nonDefaultProduction = digitalocean.getProjects({
 *     filters: [
 *         {
 *             key: "environment",
 *             values: ["Production"],
 *         },
 *         {
 *             key: "isDefault",
 *             values: ["false"],
 *         },
 *     ],
 *     sorts: [{
 *         direction: "asc",
 *         key: "name",
 *     }],
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/d/projects.html.md.
 */
export function getProjects(args?: GetProjectsArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectsResult> & GetProjectsResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetProjectsResult> = pulumi.runtime.invoke("digitalocean:index/getProjects:getProjects", {
        "filters": args.filters,
        "sorts": args.sorts,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getProjects.
 */
export interface GetProjectsArgs {
    /**
     * Filter the results.
     * The `filter` block is documented below.
     */
    readonly filters?: inputs.GetProjectsFilter[];
    /**
     * Sort the results.
     * The `sort` block is documented below.
     */
    readonly sorts?: inputs.GetProjectsSort[];
}

/**
 * A collection of values returned by getProjects.
 */
export interface GetProjectsResult {
    readonly filters?: outputs.GetProjectsFilter[];
    /**
     * A set of projects satisfying any `filter` and `sort` criteria. Each project has
     * the following attributes:
     * - `id` - The ID of the project
     * - `name` - The name of the project
     * - `description` - The description of the project
     * - `purpose` -  The purpose of the project (Default: "Web Application")
     * - `environment` - The environment of the project's resources. The possible values are: `Development`, `Staging`, `Production`.
     * - `resources` - A set of uniform resource names (URNs) for the resources associated with the project
     * - `ownerUuid` - The unique universal identifier of the project owner
     * - `ownerId` - The ID of the project owner
     * - `createdAt` - The date and time when the project was created, (ISO8601)
     * - `updatedAt` - The date and time when the project was last updated, (ISO8601)
     */
    readonly projects: outputs.GetProjectsProject[];
    readonly sorts?: outputs.GetProjectsSort[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
