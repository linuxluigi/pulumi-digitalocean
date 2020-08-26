# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Firewall']


class Firewall(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 droplet_ids: Optional[pulumi.Input[List[pulumi.Input[float]]]] = None,
                 inbound_rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['FirewallInboundRuleArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 outbound_rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['FirewallOutboundRuleArgs']]]]] = None,
                 tags: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a DigitalOcean Cloud Firewall resource. This can be used to create,
        modify, and delete Firewalls.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_digitalocean as digitalocean

        web_droplet = digitalocean.Droplet("webDroplet",
            size="s-1vcpu-1gb",
            image="ubuntu-18-04-x64",
            region="nyc3")
        web_firewall = digitalocean.Firewall("webFirewall",
            droplet_ids=[web_droplet.id],
            inbound_rules=[
                digitalocean.FirewallInboundRuleArgs(
                    protocol="tcp",
                    port_range="22",
                    source_addresses=[
                        "192.168.1.0/24",
                        "2002:1:2::/48",
                    ],
                ),
                digitalocean.FirewallInboundRuleArgs(
                    protocol="tcp",
                    port_range="80",
                    source_addresses=[
                        "0.0.0.0/0",
                        "::/0",
                    ],
                ),
                digitalocean.FirewallInboundRuleArgs(
                    protocol="tcp",
                    port_range="443",
                    source_addresses=[
                        "0.0.0.0/0",
                        "::/0",
                    ],
                ),
                digitalocean.FirewallInboundRuleArgs(
                    protocol="icmp",
                    source_addresses=[
                        "0.0.0.0/0",
                        "::/0",
                    ],
                ),
            ],
            outbound_rules=[
                digitalocean.FirewallOutboundRuleArgs(
                    protocol="tcp",
                    port_range="53",
                    destination_addresses=[
                        "0.0.0.0/0",
                        "::/0",
                    ],
                ),
                digitalocean.FirewallOutboundRuleArgs(
                    protocol="udp",
                    port_range="53",
                    destination_addresses=[
                        "0.0.0.0/0",
                        "::/0",
                    ],
                ),
                digitalocean.FirewallOutboundRuleArgs(
                    protocol="icmp",
                    destination_addresses=[
                        "0.0.0.0/0",
                        "::/0",
                    ],
                ),
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[float]]] droplet_ids: The list of the IDs of the Droplets assigned
               to the Firewall.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['FirewallInboundRuleArgs']]]] inbound_rules: The inbound access rule block for the Firewall.
               The `inbound_rule` block is documented below.
        :param pulumi.Input[str] name: The Firewall name
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['FirewallOutboundRuleArgs']]]] outbound_rules: The outbound access rule block for the Firewall.
               The `outbound_rule` block is documented below.
        :param pulumi.Input[List[pulumi.Input[str]]] tags: The names of the Tags assigned to the Firewall.
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

            __props__['droplet_ids'] = droplet_ids
            __props__['inbound_rules'] = inbound_rules
            __props__['name'] = name
            __props__['outbound_rules'] = outbound_rules
            __props__['tags'] = tags
            __props__['created_at'] = None
            __props__['pending_changes'] = None
            __props__['status'] = None
        super(Firewall, __self__).__init__(
            'digitalocean:index/firewall:Firewall',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            droplet_ids: Optional[pulumi.Input[List[pulumi.Input[float]]]] = None,
            inbound_rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['FirewallInboundRuleArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            outbound_rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['FirewallOutboundRuleArgs']]]]] = None,
            pending_changes: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['FirewallPendingChangeArgs']]]]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> 'Firewall':
        """
        Get an existing Firewall resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: A time value given in ISO8601 combined date and time format
               that represents when the Firewall was created.
        :param pulumi.Input[List[pulumi.Input[float]]] droplet_ids: The list of the IDs of the Droplets assigned
               to the Firewall.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['FirewallInboundRuleArgs']]]] inbound_rules: The inbound access rule block for the Firewall.
               The `inbound_rule` block is documented below.
        :param pulumi.Input[str] name: The Firewall name
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['FirewallOutboundRuleArgs']]]] outbound_rules: The outbound access rule block for the Firewall.
               The `outbound_rule` block is documented below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['FirewallPendingChangeArgs']]]] pending_changes: An list of object containing the fields, "droplet_id",
               "removing", and "status".  It is provided to detail exactly which Droplets
               are having their security policies updated.  When empty, all changes
               have been successfully applied.
        :param pulumi.Input[str] status: A status string indicating the current state of the Firewall.
               This can be "waiting", "succeeded", or "failed".
        :param pulumi.Input[List[pulumi.Input[str]]] tags: The names of the Tags assigned to the Firewall.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["created_at"] = created_at
        __props__["droplet_ids"] = droplet_ids
        __props__["inbound_rules"] = inbound_rules
        __props__["name"] = name
        __props__["outbound_rules"] = outbound_rules
        __props__["pending_changes"] = pending_changes
        __props__["status"] = status
        __props__["tags"] = tags
        return Firewall(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        A time value given in ISO8601 combined date and time format
        that represents when the Firewall was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="dropletIds")
    def droplet_ids(self) -> Optional[List[float]]:
        """
        The list of the IDs of the Droplets assigned
        to the Firewall.
        """
        return pulumi.get(self, "droplet_ids")

    @property
    @pulumi.getter(name="inboundRules")
    def inbound_rules(self) -> Optional[List['outputs.FirewallInboundRule']]:
        """
        The inbound access rule block for the Firewall.
        The `inbound_rule` block is documented below.
        """
        return pulumi.get(self, "inbound_rules")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The Firewall name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outboundRules")
    def outbound_rules(self) -> Optional[List['outputs.FirewallOutboundRule']]:
        """
        The outbound access rule block for the Firewall.
        The `outbound_rule` block is documented below.
        """
        return pulumi.get(self, "outbound_rules")

    @property
    @pulumi.getter(name="pendingChanges")
    def pending_changes(self) -> List['outputs.FirewallPendingChange']:
        """
        An list of object containing the fields, "droplet_id",
        "removing", and "status".  It is provided to detail exactly which Droplets
        are having their security policies updated.  When empty, all changes
        have been successfully applied.
        """
        return pulumi.get(self, "pending_changes")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        A status string indicating the current state of the Firewall.
        This can be "waiting", "succeeded", or "failed".
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[List[str]]:
        """
        The names of the Tags assigned to the Firewall.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

