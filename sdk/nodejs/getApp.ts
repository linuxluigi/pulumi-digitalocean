// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

/**
 * Get information on a DigitalOcean App.
 *
 * ## Example Usage
 *
 * Get the account:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const example = digitalocean.getApp({
 *     appId: "e665d18d-7b56-44a9-92ce-31979174d544",
 * });
 * export const defaultIngress = example.then(example => example.defaultIngress);
 * ```
 */
export function getApp(args: GetAppArgs, opts?: pulumi.InvokeOptions): Promise<GetAppResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("digitalocean:index/getApp:getApp", {
        "appId": args.appId,
    }, opts);
}

/**
 * A collection of arguments for invoking getApp.
 */
export interface GetAppArgs {
    /**
     * The ID of the app to retrieve information about.
     */
    appId: string;
}

/**
 * A collection of values returned by getApp.
 */
export interface GetAppResult {
    /**
     * The ID the app's currently active deployment.
     */
    readonly activeDeploymentId: string;
    readonly appId: string;
    /**
     * The date and time of when the app was created.
     */
    readonly createdAt: string;
    /**
     * The default URL to access the app.
     */
    readonly defaultIngress: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The live URL of the app.
     */
    readonly liveUrl: string;
    /**
     * A DigitalOcean App spec describing the app.
     */
    readonly specs: outputs.GetAppSpec[];
    /**
     * The date and time of when the app was last updated.
     */
    readonly updatedAt: string;
}

export function getAppOutput(args: GetAppOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppResult> {
    return pulumi.output(args).apply(a => getApp(a, opts))
}

/**
 * A collection of arguments for invoking getApp.
 */
export interface GetAppOutputArgs {
    /**
     * The ID of the app to retrieve information about.
     */
    appId: pulumi.Input<string>;
}
