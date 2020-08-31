# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables

__all__ = ['Project']


class Project(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 purpose: Optional[pulumi.Input[str]] = None,
                 resources: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a DigitalOcean Project resource.

        Projects allow you to organize your resources into groups that fit the way you work.
        You can group resources (like Droplets, Spaces, Load Balancers, domains, and Floating IPs)
        in ways that align with the applications you host on DigitalOcean.

        The following resource types can be associated with a project:

        * Database Clusters
        * Domains
        * Droplets
        * Floating IP
        * Load Balancers
        * Spaces Bucket
        * Volume

        **Note:** A managed project cannot be set as a default project.

        ## Example Usage

        The following example demonstrates the creation of an empty project:

        ```python
        import pulumi
        import pulumi_digitalocean as digitalocean

        playground = digitalocean.Project("playground",
            description="A project to represent development resources.",
            environment="Development",
            purpose="Web Application")
        ```

        The following example demonstrates the creation of a project with a Droplet resource:

        ```python
        import pulumi
        import pulumi_digitalocean as digitalocean

        foobar = digitalocean.Droplet("foobar",
            size="512mb",
            image="centos-7-x64",
            region="nyc3")
        playground = digitalocean.Project("playground",
            description="A project to represent development resources.",
            purpose="Web Application",
            environment="Development",
            resources=[foobar.droplet_urn])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: the description of the project
        :param pulumi.Input[str] environment: the environment of the project's resources. The possible values are: `Development`, `Staging`, `Production`)
        :param pulumi.Input[str] name: The name of the Project
        :param pulumi.Input[str] purpose: the purpose of the project, (Default: "Web Application")
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

            __props__['description'] = description
            __props__['environment'] = environment
            __props__['name'] = name
            __props__['purpose'] = purpose
            __props__['resources'] = resources
            __props__['created_at'] = None
            __props__['is_default'] = None
            __props__['owner_id'] = None
            __props__['owner_uuid'] = None
            __props__['updated_at'] = None
        super(Project, __self__).__init__(
            'digitalocean:index/project:Project',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            environment: Optional[pulumi.Input[str]] = None,
            is_default: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            owner_id: Optional[pulumi.Input[float]] = None,
            owner_uuid: Optional[pulumi.Input[str]] = None,
            purpose: Optional[pulumi.Input[str]] = None,
            resources: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'Project':
        """
        Get an existing Project resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: the date and time when the project was created, (ISO8601)
        :param pulumi.Input[str] description: the description of the project
        :param pulumi.Input[str] environment: the environment of the project's resources. The possible values are: `Development`, `Staging`, `Production`)
        :param pulumi.Input[str] name: The name of the Project
        :param pulumi.Input[float] owner_id: the id of the project owner.
        :param pulumi.Input[str] owner_uuid: the unique universal identifier of the project owner.
        :param pulumi.Input[str] purpose: the purpose of the project, (Default: "Web Application")
        :param pulumi.Input[List[pulumi.Input[str]]] resources: a list of uniform resource names (URNs) for the resources associated with the project
        :param pulumi.Input[str] updated_at: the date and time when the project was last updated, (ISO8601)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["created_at"] = created_at
        __props__["description"] = description
        __props__["environment"] = environment
        __props__["is_default"] = is_default
        __props__["name"] = name
        __props__["owner_id"] = owner_id
        __props__["owner_uuid"] = owner_uuid
        __props__["purpose"] = purpose
        __props__["resources"] = resources
        __props__["updated_at"] = updated_at
        return Project(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        the date and time when the project was created, (ISO8601)
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        the description of the project
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output[Optional[str]]:
        """
        the environment of the project's resources. The possible values are: `Development`, `Staging`, `Production`)
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "is_default")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Project
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[float]:
        """
        the id of the project owner.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="ownerUuid")
    def owner_uuid(self) -> pulumi.Output[str]:
        """
        the unique universal identifier of the project owner.
        """
        return pulumi.get(self, "owner_uuid")

    @property
    @pulumi.getter
    def purpose(self) -> pulumi.Output[Optional[str]]:
        """
        the purpose of the project, (Default: "Web Application")
        """
        return pulumi.get(self, "purpose")

    @property
    @pulumi.getter
    def resources(self) -> pulumi.Output[List[str]]:
        """
        a list of uniform resource names (URNs) for the resources associated with the project
        """
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        the date and time when the project was last updated, (ISO8601)
        """
        return pulumi.get(self, "updated_at")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

