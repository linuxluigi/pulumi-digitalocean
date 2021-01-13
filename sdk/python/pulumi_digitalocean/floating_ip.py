# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['FloatingIp']


class FloatingIp(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 droplet_id: Optional[pulumi.Input[int]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a DigitalOcean Floating IP to represent a publicly-accessible static IP addresses that can be mapped to one of your Droplets.

        > **NOTE:** Floating IPs can be assigned to a Droplet either directly on the `FloatingIp` resource by setting a `droplet_id` or using the `FloatingIpAssignment` resource, but the two cannot be used together.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_digitalocean as digitalocean

        foobar_droplet = digitalocean.Droplet("foobarDroplet",
            size="s-1vcpu-1gb",
            image="ubuntu-18-04-x64",
            region="sgp1",
            ipv6=True,
            private_networking=True)
        foobar_floating_ip = digitalocean.FloatingIp("foobarFloatingIp",
            droplet_id=foobar_droplet.id,
            region=foobar_droplet.region)
        ```

        ## Import

        Floating IPs can be imported using the `ip`, e.g.

        ```sh
         $ pulumi import digitalocean:index/floatingIp:FloatingIp myip 192.168.0.1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] droplet_id: The ID of Droplet that the Floating IP will be assigned to.
        :param pulumi.Input[str] ip_address: The IP Address of the resource
        :param pulumi.Input[str] region: The region that the Floating IP is reserved to.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['droplet_id'] = droplet_id
            __props__['ip_address'] = ip_address
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__['region'] = region
            __props__['floating_ip_urn'] = None
        super(FloatingIp, __self__).__init__(
            'digitalocean:index/floatingIp:FloatingIp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            droplet_id: Optional[pulumi.Input[int]] = None,
            floating_ip_urn: Optional[pulumi.Input[str]] = None,
            ip_address: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'FloatingIp':
        """
        Get an existing FloatingIp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] droplet_id: The ID of Droplet that the Floating IP will be assigned to.
        :param pulumi.Input[str] floating_ip_urn: The uniform resource name of the floating ip
        :param pulumi.Input[str] ip_address: The IP Address of the resource
        :param pulumi.Input[str] region: The region that the Floating IP is reserved to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["droplet_id"] = droplet_id
        __props__["floating_ip_urn"] = floating_ip_urn
        __props__["ip_address"] = ip_address
        __props__["region"] = region
        return FloatingIp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dropletId")
    def droplet_id(self) -> pulumi.Output[Optional[int]]:
        """
        The ID of Droplet that the Floating IP will be assigned to.
        """
        return pulumi.get(self, "droplet_id")

    @property
    @pulumi.getter(name="floatingIpUrn")
    def floating_ip_urn(self) -> pulumi.Output[str]:
        """
        The uniform resource name of the floating ip
        """
        return pulumi.get(self, "floating_ip_urn")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Output[str]:
        """
        The IP Address of the resource
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region that the Floating IP is reserved to.
        """
        return pulumi.get(self, "region")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

