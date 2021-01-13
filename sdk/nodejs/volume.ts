// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

import {FilesystemType, Region} from "./index";

/**
 * Provides a DigitalOcean Block Storage volume which can be attached to a Droplet in order to provide expanded storage.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const foobarVolume = new digitalocean.Volume("foobarVolume", {
 *     region: "nyc1",
 *     size: 100,
 *     initialFilesystemType: "ext4",
 *     description: "an example volume",
 * });
 * const foobarDroplet = new digitalocean.Droplet("foobarDroplet", {
 *     size: "s-1vcpu-1gb",
 *     image: "ubuntu-18-04-x64",
 *     region: "nyc1",
 * });
 * const foobarVolumeAttachment = new digitalocean.VolumeAttachment("foobarVolumeAttachment", {
 *     dropletId: foobarDroplet.id,
 *     volumeId: foobarVolume.id,
 * });
 * ```
 *
 * You can also create a volume from an existing snapshot.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const foobarVolumeSnapshot = digitalocean.getVolumeSnapshot({
 *     name: "baz",
 * });
 * const foobarVolume = new digitalocean.Volume("foobarVolume", {
 *     region: "lon1",
 *     size: foobarVolumeSnapshot.then(foobarVolumeSnapshot => foobarVolumeSnapshot.minDiskSize),
 *     snapshotId: foobarVolumeSnapshot.then(foobarVolumeSnapshot => foobarVolumeSnapshot.id),
 * });
 * ```
 *
 * ## Import
 *
 * Volumes can be imported using the `volume id`, e.g.
 *
 * ```sh
 *  $ pulumi import digitalocean:index/volume:Volume volume 506f78a4-e098-11e5-ad9f-000f53306ae1
 * ```
 */
export class Volume extends pulumi.CustomResource {
    /**
     * Get an existing Volume resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VolumeState, opts?: pulumi.CustomResourceOptions): Volume {
        return new Volume(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean:index/volume:Volume';

    /**
     * Returns true if the given object is an instance of Volume.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Volume {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Volume.__pulumiType;
    }

    /**
     * A free-form text field up to a limit of 1024 bytes to describe a block storage volume.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A list of associated droplet ids.
     */
    public /*out*/ readonly dropletIds!: pulumi.Output<number[]>;
    /**
     * Filesystem label for the block storage volume.
     */
    public /*out*/ readonly filesystemLabel!: pulumi.Output<string>;
    /**
     * Filesystem type (`xfs` or `ext4`) for the block storage volume.
     *
     * @deprecated This fields functionality has been replaced by `initial_filesystem_type`. The property will still remain as a computed attribute representing the current volumes filesystem type.
     */
    public readonly filesystemType!: pulumi.Output<string>;
    /**
     * Initial filesystem label for the block storage volume.
     */
    public readonly initialFilesystemLabel!: pulumi.Output<string | undefined>;
    /**
     * Initial filesystem type (`xfs` or `ext4`) for the block storage volume.
     */
    public readonly initialFilesystemType!: pulumi.Output<FilesystemType | undefined>;
    /**
     * A name for the block storage volume. Must be lowercase and be composed only of numbers, letters and "-", up to a limit of 64 characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region that the block storage volume will be created in.
     */
    public readonly region!: pulumi.Output<Region>;
    /**
     * The size of the block storage volume in GiB. If updated, can only be expanded.
     */
    public readonly size!: pulumi.Output<number>;
    /**
     * The ID of an existing volume snapshot from which the new volume will be created. If supplied, the region and size will be limitied on creation to that of the referenced snapshot
     */
    public readonly snapshotId!: pulumi.Output<string | undefined>;
    /**
     * A list of the tags to be applied to this Volume.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The uniform resource name for the volume.
     */
    public /*out*/ readonly volumeUrn!: pulumi.Output<string>;

    /**
     * Create a Volume resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VolumeArgs | VolumeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VolumeState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["dropletIds"] = state ? state.dropletIds : undefined;
            inputs["filesystemLabel"] = state ? state.filesystemLabel : undefined;
            inputs["filesystemType"] = state ? state.filesystemType : undefined;
            inputs["initialFilesystemLabel"] = state ? state.initialFilesystemLabel : undefined;
            inputs["initialFilesystemType"] = state ? state.initialFilesystemType : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["size"] = state ? state.size : undefined;
            inputs["snapshotId"] = state ? state.snapshotId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["volumeUrn"] = state ? state.volumeUrn : undefined;
        } else {
            const args = argsOrState as VolumeArgs | undefined;
            if ((!args || args.region === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'region'");
            }
            if ((!args || args.size === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'size'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["filesystemType"] = args ? args.filesystemType : undefined;
            inputs["initialFilesystemLabel"] = args ? args.initialFilesystemLabel : undefined;
            inputs["initialFilesystemType"] = args ? args.initialFilesystemType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["size"] = args ? args.size : undefined;
            inputs["snapshotId"] = args ? args.snapshotId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["dropletIds"] = undefined /*out*/;
            inputs["filesystemLabel"] = undefined /*out*/;
            inputs["volumeUrn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Volume.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Volume resources.
 */
export interface VolumeState {
    /**
     * A free-form text field up to a limit of 1024 bytes to describe a block storage volume.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A list of associated droplet ids.
     */
    readonly dropletIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Filesystem label for the block storage volume.
     */
    readonly filesystemLabel?: pulumi.Input<string>;
    /**
     * Filesystem type (`xfs` or `ext4`) for the block storage volume.
     *
     * @deprecated This fields functionality has been replaced by `initial_filesystem_type`. The property will still remain as a computed attribute representing the current volumes filesystem type.
     */
    readonly filesystemType?: pulumi.Input<string>;
    /**
     * Initial filesystem label for the block storage volume.
     */
    readonly initialFilesystemLabel?: pulumi.Input<string>;
    /**
     * Initial filesystem type (`xfs` or `ext4`) for the block storage volume.
     */
    readonly initialFilesystemType?: pulumi.Input<FilesystemType>;
    /**
     * A name for the block storage volume. Must be lowercase and be composed only of numbers, letters and "-", up to a limit of 64 characters.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region that the block storage volume will be created in.
     */
    readonly region?: pulumi.Input<Region>;
    /**
     * The size of the block storage volume in GiB. If updated, can only be expanded.
     */
    readonly size?: pulumi.Input<number>;
    /**
     * The ID of an existing volume snapshot from which the new volume will be created. If supplied, the region and size will be limitied on creation to that of the referenced snapshot
     */
    readonly snapshotId?: pulumi.Input<string>;
    /**
     * A list of the tags to be applied to this Volume.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The uniform resource name for the volume.
     */
    readonly volumeUrn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Volume resource.
 */
export interface VolumeArgs {
    /**
     * A free-form text field up to a limit of 1024 bytes to describe a block storage volume.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Filesystem type (`xfs` or `ext4`) for the block storage volume.
     *
     * @deprecated This fields functionality has been replaced by `initial_filesystem_type`. The property will still remain as a computed attribute representing the current volumes filesystem type.
     */
    readonly filesystemType?: pulumi.Input<string>;
    /**
     * Initial filesystem label for the block storage volume.
     */
    readonly initialFilesystemLabel?: pulumi.Input<string>;
    /**
     * Initial filesystem type (`xfs` or `ext4`) for the block storage volume.
     */
    readonly initialFilesystemType?: pulumi.Input<FilesystemType>;
    /**
     * A name for the block storage volume. Must be lowercase and be composed only of numbers, letters and "-", up to a limit of 64 characters.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region that the block storage volume will be created in.
     */
    readonly region: pulumi.Input<Region>;
    /**
     * The size of the block storage volume in GiB. If updated, can only be expanded.
     */
    readonly size: pulumi.Input<number>;
    /**
     * The ID of an existing volume snapshot from which the new volume will be created. If supplied, the region and size will be limitied on creation to that of the referenced snapshot
     */
    readonly snapshotId?: pulumi.Input<string>;
    /**
     * A list of the tags to be applied to this Volume.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
}
