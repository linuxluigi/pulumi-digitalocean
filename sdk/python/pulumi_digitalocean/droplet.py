# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class Droplet(pulumi.CustomResource):
    backups: pulumi.Output[bool]
    """
    Boolean controlling if backups are made. Defaults to
    false.
    """
    disk: pulumi.Output[float]
    """
    The size of the instance's disk in GB
    """
    image: pulumi.Output[str]
    """
    The Droplet image ID or slug.
    """
    ipv4_address: pulumi.Output[str]
    """
    The IPv4 address
    """
    ipv4_address_private: pulumi.Output[str]
    """
    The private networking IPv4 address
    """
    ipv6: pulumi.Output[bool]
    """
    Boolean controlling if IPv6 is enabled. Defaults to false.
    """
    ipv6_address: pulumi.Output[str]
    """
    The IPv6 address
    """
    ipv6_address_private: pulumi.Output[str]
    locked: pulumi.Output[bool]
    """
    Is the Droplet locked
    """
    memory: pulumi.Output[float]
    monitoring: pulumi.Output[bool]
    """
    Boolean controlling whether monitoring agent is installed.
    Defaults to false.
    """
    name: pulumi.Output[str]
    """
    The Droplet name.
    """
    price_hourly: pulumi.Output[float]
    """
    Droplet hourly price
    """
    price_monthly: pulumi.Output[float]
    """
    Droplet monthly price
    """
    private_networking: pulumi.Output[bool]
    """
    Boolean controlling if private networks are
    enabled. Defaults to false.
    """
    region: pulumi.Output[str]
    """
    The region to start in.
    """
    resize_disk: pulumi.Output[bool]
    """
    Boolean controlling whether to increase the disk
    size when resizing a Droplet. It defaults to `true`. When set to `false`,
    only the Droplet's RAM and CPU will be resized. **Increasing a Droplet's disk
    size is a permanent change**. Increasing only RAM and CPU is reversible.
    """
    size: pulumi.Output[str]
    """
    The unique slug that indentifies the type of Droplet. You can find a list of available slugs on [DigitalOcean API documentation](https://developers.digitalocean.com/documentation/v2/#list-all-sizes).
    """
    ssh_keys: pulumi.Output[list]
    """
    A list of SSH IDs or fingerprints to enable in
    the format `[12345, 123456]`. To retrieve this info, use a tool such
    as `curl` with the [DigitalOcean API](https://developers.digitalocean.com/documentation/v2/#ssh-keys),
    to retrieve them.
    """
    status: pulumi.Output[str]
    """
    The status of the Droplet
    """
    tags: pulumi.Output[list]
    """
    A list of the tags to be applied to this Droplet.
    """
    urn: pulumi.Output[str]
    """
    The uniform resource name of the Droplet
    * `name`- The name of the Droplet
    """
    user_data: pulumi.Output[str]
    """
    A string of the desired User Data for the Droplet.
    """
    vcpus: pulumi.Output[float]
    """
    The number of the instance's virtual CPUs
    """
    volume_ids: pulumi.Output[list]
    """
    A list of the IDs of each [block storage volume](https://www.terraform.io/docs/providers/do/r/volume.html) to be attached to the Droplet.
    """
    def __init__(__self__, resource_name, opts=None, backups=None, image=None, ipv6=None, monitoring=None, name=None, private_networking=None, region=None, resize_disk=None, size=None, ssh_keys=None, tags=None, user_data=None, volume_ids=None, __name__=None, __opts__=None):
        """
        Provides a DigitalOcean Droplet resource. This can be used to create,
        modify, and delete Droplets. Droplets also support
        [provisioning](https://www.terraform.io/docs/provisioners/index.html).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] backups: Boolean controlling if backups are made. Defaults to
               false.
        :param pulumi.Input[str] image: The Droplet image ID or slug.
        :param pulumi.Input[bool] ipv6: Boolean controlling if IPv6 is enabled. Defaults to false.
        :param pulumi.Input[bool] monitoring: Boolean controlling whether monitoring agent is installed.
               Defaults to false.
        :param pulumi.Input[str] name: The Droplet name.
        :param pulumi.Input[bool] private_networking: Boolean controlling if private networks are
               enabled. Defaults to false.
        :param pulumi.Input[str] region: The region to start in.
        :param pulumi.Input[bool] resize_disk: Boolean controlling whether to increase the disk
               size when resizing a Droplet. It defaults to `true`. When set to `false`,
               only the Droplet's RAM and CPU will be resized. **Increasing a Droplet's disk
               size is a permanent change**. Increasing only RAM and CPU is reversible.
        :param pulumi.Input[str] size: The unique slug that indentifies the type of Droplet. You can find a list of available slugs on [DigitalOcean API documentation](https://developers.digitalocean.com/documentation/v2/#list-all-sizes).
        :param pulumi.Input[list] ssh_keys: A list of SSH IDs or fingerprints to enable in
               the format `[12345, 123456]`. To retrieve this info, use a tool such
               as `curl` with the [DigitalOcean API](https://developers.digitalocean.com/documentation/v2/#ssh-keys),
               to retrieve them.
        :param pulumi.Input[list] tags: A list of the tags to be applied to this Droplet.
        :param pulumi.Input[str] user_data: A string of the desired User Data for the Droplet.
        :param pulumi.Input[list] volume_ids: A list of the IDs of each [block storage volume](https://www.terraform.io/docs/providers/do/r/volume.html) to be attached to the Droplet.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['backups'] = backups

        if image is None:
            raise TypeError("Missing required property 'image'")
        __props__['image'] = image

        __props__['ipv6'] = ipv6

        __props__['monitoring'] = monitoring

        __props__['name'] = name

        __props__['private_networking'] = private_networking

        if region is None:
            raise TypeError("Missing required property 'region'")
        __props__['region'] = region

        __props__['resize_disk'] = resize_disk

        if size is None:
            raise TypeError("Missing required property 'size'")
        __props__['size'] = size

        __props__['ssh_keys'] = ssh_keys

        __props__['tags'] = tags

        __props__['user_data'] = user_data

        __props__['volume_ids'] = volume_ids

        __props__['disk'] = None
        __props__['ipv4_address'] = None
        __props__['ipv4_address_private'] = None
        __props__['ipv6_address'] = None
        __props__['ipv6_address_private'] = None
        __props__['locked'] = None
        __props__['memory'] = None
        __props__['price_hourly'] = None
        __props__['price_monthly'] = None
        __props__['status'] = None
        __props__['urn'] = None
        __props__['vcpus'] = None

        super(Droplet, __self__).__init__(
            'digitalocean:index/droplet:Droplet',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
