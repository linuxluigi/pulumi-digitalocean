# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables

__all__ = ['ProjectResources']


class ProjectResources(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 resources: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Assign resources to a DigitalOcean Project. This is useful if you need to assign resources
        managed this provider to a DigitalOcean Project that is unmanaged by the provider.

        The following resource types can be associated with a project:

        * Database Clusters
        * Domains
        * Droplets
        * Floating IP
        * Load Balancers
        * Spaces Bucket
        * Volume

        ## Example Usage

        The following example assigns a droplet to a Project managed outside of this provider:

        ```python
        import pulumi
        import pulumi_digitalocean as digitalocean

        playground = digitalocean.get_project(name="playground")
        foobar = digitalocean.Droplet("foobar",
            size="512mb",
            image="centos-7-x64",
            region="nyc3")
        barfoo = digitalocean.ProjectResources("barfoo",
            project=data["digitalocean_project"]["foo"]["id"],
            resources=[foobar.droplet_urn])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: the ID of the project
        :param pulumi.Input[List[pulumi.Input[str]]] resources: a list of uniform resource names (URNs) for the resources associated with the project
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

            if project is None:
                raise TypeError("Missing required property 'project'")
            __props__['project'] = project
            if resources is None:
                raise TypeError("Missing required property 'resources'")
            __props__['resources'] = resources
        super(ProjectResources, __self__).__init__(
            'digitalocean:index/projectResources:ProjectResources',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            project: Optional[pulumi.Input[str]] = None,
            resources: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> 'ProjectResources':
        """
        Get an existing ProjectResources resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: the ID of the project
        :param pulumi.Input[List[pulumi.Input[str]]] resources: a list of uniform resource names (URNs) for the resources associated with the project
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["project"] = project
        __props__["resources"] = resources
        return ProjectResources(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        the ID of the project
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def resources(self) -> pulumi.Output[List[str]]:
        """
        a list of uniform resource names (URNs) for the resources associated with the project
        """
        return pulumi.get(self, "resources")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

