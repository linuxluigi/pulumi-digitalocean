# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class VolumeSnapshot(pulumi.CustomResource):
    created_at: pulumi.Output[str]
    """
    The date and time the volume snapshot was created.
    """
    min_disk_size: pulumi.Output[float]
    """
    The minimum size in gigabytes required for a volume to be created based on this volume snapshot.
    """
    name: pulumi.Output[str]
    """
    A name for the volume snapshot.
    """
    regions: pulumi.Output[list]
    """
    A list of DigitalOcean region "slugs" indicating where the volume snapshot is available.
    """
    size: pulumi.Output[float]
    """
    The billable size of the volume snapshot in gigabytes.
    """
    volume_id: pulumi.Output[str]
    """
    The ID of the volume from which the volume snapshot originated.
    """
    def __init__(__self__, resource_name, opts=None, name=None, volume_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a DigitalOcean Volume Snapshot which can be used to create a snapshot from an existing volume.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: A name for the volume snapshot.
        :param pulumi.Input[str] volume_id: The ID of the volume from which the volume snapshot originated.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/r/volume_snapshot.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['name'] = name
            if volume_id is None:
                raise TypeError("Missing required property 'volume_id'")
            __props__['volume_id'] = volume_id
            __props__['created_at'] = None
            __props__['min_disk_size'] = None
            __props__['regions'] = None
            __props__['size'] = None
        super(VolumeSnapshot, __self__).__init__(
            'digitalocean:index/volumeSnapshot:VolumeSnapshot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, created_at=None, min_disk_size=None, name=None, regions=None, size=None, volume_id=None):
        """
        Get an existing VolumeSnapshot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The date and time the volume snapshot was created.
        :param pulumi.Input[float] min_disk_size: The minimum size in gigabytes required for a volume to be created based on this volume snapshot.
        :param pulumi.Input[str] name: A name for the volume snapshot.
        :param pulumi.Input[list] regions: A list of DigitalOcean region "slugs" indicating where the volume snapshot is available.
        :param pulumi.Input[float] size: The billable size of the volume snapshot in gigabytes.
        :param pulumi.Input[str] volume_id: The ID of the volume from which the volume snapshot originated.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/r/volume_snapshot.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["created_at"] = created_at
        __props__["min_disk_size"] = min_disk_size
        __props__["name"] = name
        __props__["regions"] = regions
        __props__["size"] = size
        __props__["volume_id"] = volume_id
        return VolumeSnapshot(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

