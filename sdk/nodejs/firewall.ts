// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * Provides a DigitalOcean Cloud Firewall resource. This can be used to create,
 * modify, and delete Firewalls.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const webDroplet = new digitalocean.Droplet("webDroplet", {
 *     size: "s-1vcpu-1gb",
 *     image: "ubuntu-18-04-x64",
 *     region: "nyc3",
 * });
 * const webFirewall = new digitalocean.Firewall("webFirewall", {
 *     dropletIds: [webDroplet.id],
 *     inboundRules: [
 *         {
 *             protocol: "tcp",
 *             portRange: "22",
 *             sourceAddresses: [
 *                 "192.168.1.0/24",
 *                 "2002:1:2::/48",
 *             ],
 *         },
 *         {
 *             protocol: "tcp",
 *             portRange: "80",
 *             sourceAddresses: [
 *                 "0.0.0.0/0",
 *                 "::/0",
 *             ],
 *         },
 *         {
 *             protocol: "tcp",
 *             portRange: "443",
 *             sourceAddresses: [
 *                 "0.0.0.0/0",
 *                 "::/0",
 *             ],
 *         },
 *         {
 *             protocol: "icmp",
 *             sourceAddresses: [
 *                 "0.0.0.0/0",
 *                 "::/0",
 *             ],
 *         },
 *     ],
 *     outboundRules: [
 *         {
 *             protocol: "tcp",
 *             portRange: "53",
 *             destinationAddresses: [
 *                 "0.0.0.0/0",
 *                 "::/0",
 *             ],
 *         },
 *         {
 *             protocol: "udp",
 *             portRange: "53",
 *             destinationAddresses: [
 *                 "0.0.0.0/0",
 *                 "::/0",
 *             ],
 *         },
 *         {
 *             protocol: "icmp",
 *             destinationAddresses: [
 *                 "0.0.0.0/0",
 *                 "::/0",
 *             ],
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Firewalls can be imported using the firewall `id`, e.g.
 *
 * ```sh
 *  $ pulumi import digitalocean:index/firewall:Firewall myfirewall b8ecd2ab-2267-4a5e-8692-cbf1d32583e3
 * ```
 */
export class Firewall extends pulumi.CustomResource {
    /**
     * Get an existing Firewall resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallState, opts?: pulumi.CustomResourceOptions): Firewall {
        return new Firewall(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean:index/firewall:Firewall';

    /**
     * Returns true if the given object is an instance of Firewall.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Firewall {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Firewall.__pulumiType;
    }

    /**
     * A time value given in ISO8601 combined date and time format
     * that represents when the Firewall was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The list of the IDs of the Droplets assigned
     * to the Firewall.
     */
    public readonly dropletIds!: pulumi.Output<number[] | undefined>;
    /**
     * The inbound access rule block for the Firewall.
     * The `inboundRule` block is documented below.
     */
    public readonly inboundRules!: pulumi.Output<outputs.FirewallInboundRule[] | undefined>;
    /**
     * The Firewall name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The outbound access rule block for the Firewall.
     * The `outboundRule` block is documented below.
     */
    public readonly outboundRules!: pulumi.Output<outputs.FirewallOutboundRule[] | undefined>;
    /**
     * An list of object containing the fields, "dropletId",
     * "removing", and "status".  It is provided to detail exactly which Droplets
     * are having their security policies updated.  When empty, all changes
     * have been successfully applied.
     */
    public /*out*/ readonly pendingChanges!: pulumi.Output<outputs.FirewallPendingChange[]>;
    /**
     * A status string indicating the current state of the Firewall.
     * This can be "waiting", "succeeded", or "failed".
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The names of the Tags assigned to the Firewall.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Firewall resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallArgs | FirewallState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["dropletIds"] = state ? state.dropletIds : undefined;
            resourceInputs["inboundRules"] = state ? state.inboundRules : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["outboundRules"] = state ? state.outboundRules : undefined;
            resourceInputs["pendingChanges"] = state ? state.pendingChanges : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as FirewallArgs | undefined;
            resourceInputs["dropletIds"] = args ? args.dropletIds : undefined;
            resourceInputs["inboundRules"] = args ? args.inboundRules : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["outboundRules"] = args ? args.outboundRules : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["pendingChanges"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Firewall.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Firewall resources.
 */
export interface FirewallState {
    /**
     * A time value given in ISO8601 combined date and time format
     * that represents when the Firewall was created.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The list of the IDs of the Droplets assigned
     * to the Firewall.
     */
    dropletIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * The inbound access rule block for the Firewall.
     * The `inboundRule` block is documented below.
     */
    inboundRules?: pulumi.Input<pulumi.Input<inputs.FirewallInboundRule>[]>;
    /**
     * The Firewall name
     */
    name?: pulumi.Input<string>;
    /**
     * The outbound access rule block for the Firewall.
     * The `outboundRule` block is documented below.
     */
    outboundRules?: pulumi.Input<pulumi.Input<inputs.FirewallOutboundRule>[]>;
    /**
     * An list of object containing the fields, "dropletId",
     * "removing", and "status".  It is provided to detail exactly which Droplets
     * are having their security policies updated.  When empty, all changes
     * have been successfully applied.
     */
    pendingChanges?: pulumi.Input<pulumi.Input<inputs.FirewallPendingChange>[]>;
    /**
     * A status string indicating the current state of the Firewall.
     * This can be "waiting", "succeeded", or "failed".
     */
    status?: pulumi.Input<string>;
    /**
     * The names of the Tags assigned to the Firewall.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Firewall resource.
 */
export interface FirewallArgs {
    /**
     * The list of the IDs of the Droplets assigned
     * to the Firewall.
     */
    dropletIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * The inbound access rule block for the Firewall.
     * The `inboundRule` block is documented below.
     */
    inboundRules?: pulumi.Input<pulumi.Input<inputs.FirewallInboundRule>[]>;
    /**
     * The Firewall name
     */
    name?: pulumi.Input<string>;
    /**
     * The outbound access rule block for the Firewall.
     * The `outboundRule` block is documented below.
     */
    outboundRules?: pulumi.Input<pulumi.Input<inputs.FirewallOutboundRule>[]>;
    /**
     * The names of the Tags assigned to the Firewall.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
