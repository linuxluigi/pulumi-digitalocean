# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['ContainerRegistryDockerCredentialsArgs', 'ContainerRegistryDockerCredentials']

@pulumi.input_type
class ContainerRegistryDockerCredentialsArgs:
    def __init__(__self__, *,
                 registry_name: pulumi.Input[str],
                 expiry_seconds: Optional[pulumi.Input[int]] = None,
                 write: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a ContainerRegistryDockerCredentials resource.
        :param pulumi.Input[str] registry_name: The name of the container registry.
        :param pulumi.Input[int] expiry_seconds: The amount of time to pass before the Docker credentials expire in seconds. Defaults to 1576800000, or roughly 50 years. Must be greater than 0 and less than 1576800000.
        :param pulumi.Input[bool] write: Allow for write access to the container registry. Defaults to false.
        """
        pulumi.set(__self__, "registry_name", registry_name)
        if expiry_seconds is not None:
            pulumi.set(__self__, "expiry_seconds", expiry_seconds)
        if write is not None:
            pulumi.set(__self__, "write", write)

    @property
    @pulumi.getter(name="registryName")
    def registry_name(self) -> pulumi.Input[str]:
        """
        The name of the container registry.
        """
        return pulumi.get(self, "registry_name")

    @registry_name.setter
    def registry_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "registry_name", value)

    @property
    @pulumi.getter(name="expirySeconds")
    def expiry_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The amount of time to pass before the Docker credentials expire in seconds. Defaults to 1576800000, or roughly 50 years. Must be greater than 0 and less than 1576800000.
        """
        return pulumi.get(self, "expiry_seconds")

    @expiry_seconds.setter
    def expiry_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "expiry_seconds", value)

    @property
    @pulumi.getter
    def write(self) -> Optional[pulumi.Input[bool]]:
        """
        Allow for write access to the container registry. Defaults to false.
        """
        return pulumi.get(self, "write")

    @write.setter
    def write(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "write", value)


@pulumi.input_type
class _ContainerRegistryDockerCredentialsState:
    def __init__(__self__, *,
                 credential_expiration_time: Optional[pulumi.Input[str]] = None,
                 docker_credentials: Optional[pulumi.Input[str]] = None,
                 expiry_seconds: Optional[pulumi.Input[int]] = None,
                 registry_name: Optional[pulumi.Input[str]] = None,
                 write: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering ContainerRegistryDockerCredentials resources.
        :param pulumi.Input[str] credential_expiration_time: The date and time the registry access token will expire.
        :param pulumi.Input[str] docker_credentials: Credentials for the container registry.
        :param pulumi.Input[int] expiry_seconds: The amount of time to pass before the Docker credentials expire in seconds. Defaults to 1576800000, or roughly 50 years. Must be greater than 0 and less than 1576800000.
        :param pulumi.Input[str] registry_name: The name of the container registry.
        :param pulumi.Input[bool] write: Allow for write access to the container registry. Defaults to false.
        """
        if credential_expiration_time is not None:
            pulumi.set(__self__, "credential_expiration_time", credential_expiration_time)
        if docker_credentials is not None:
            pulumi.set(__self__, "docker_credentials", docker_credentials)
        if expiry_seconds is not None:
            pulumi.set(__self__, "expiry_seconds", expiry_seconds)
        if registry_name is not None:
            pulumi.set(__self__, "registry_name", registry_name)
        if write is not None:
            pulumi.set(__self__, "write", write)

    @property
    @pulumi.getter(name="credentialExpirationTime")
    def credential_expiration_time(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time the registry access token will expire.
        """
        return pulumi.get(self, "credential_expiration_time")

    @credential_expiration_time.setter
    def credential_expiration_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "credential_expiration_time", value)

    @property
    @pulumi.getter(name="dockerCredentials")
    def docker_credentials(self) -> Optional[pulumi.Input[str]]:
        """
        Credentials for the container registry.
        """
        return pulumi.get(self, "docker_credentials")

    @docker_credentials.setter
    def docker_credentials(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "docker_credentials", value)

    @property
    @pulumi.getter(name="expirySeconds")
    def expiry_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The amount of time to pass before the Docker credentials expire in seconds. Defaults to 1576800000, or roughly 50 years. Must be greater than 0 and less than 1576800000.
        """
        return pulumi.get(self, "expiry_seconds")

    @expiry_seconds.setter
    def expiry_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "expiry_seconds", value)

    @property
    @pulumi.getter(name="registryName")
    def registry_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the container registry.
        """
        return pulumi.get(self, "registry_name")

    @registry_name.setter
    def registry_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "registry_name", value)

    @property
    @pulumi.getter
    def write(self) -> Optional[pulumi.Input[bool]]:
        """
        Allow for write access to the container registry. Defaults to false.
        """
        return pulumi.get(self, "write")

    @write.setter
    def write(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "write", value)


class ContainerRegistryDockerCredentials(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expiry_seconds: Optional[pulumi.Input[int]] = None,
                 registry_name: Optional[pulumi.Input[str]] = None,
                 write: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Get Docker credentials for your DigitalOcean container registry.

        An error is triggered if the provided container registry name does not exist.

        ## Example Usage

        ### Basic Example

        Get the container registry:

        ```python
        import pulumi
        import pulumi_digitalocean as digitalocean

        example = digitalocean.ContainerRegistryDockerCredentials("example", registry_name="example")
        ```

        ### Docker Provider Example

        Use the `endpoint` and `docker_credentials` with the Docker provider:

        ```python
        import pulumi
        import pulumi_digitalocean as digitalocean

        example = digitalocean.get_container_registry(name="example")
        example_container_registry_docker_credentials = digitalocean.ContainerRegistryDockerCredentials("example", registry_name="example")
        ```

        ### Kubernetes Example

        Combined with the Kubernetes Provider's `kubernetes_secret` resource, you can
        access the registry from inside your cluster:

        ```python
        import pulumi
        import pulumi_digitalocean as digitalocean
        import pulumi_kubernetes as kubernetes

        example_container_registry_docker_credentials = digitalocean.ContainerRegistryDockerCredentials("example", registry_name="example")
        example = digitalocean.get_kubernetes_cluster(name="prod-cluster-01")
        example_secret = kubernetes.core.v1.Secret("example",
            metadata={
                "name": "docker-cfg",
            },
            data={
                ".dockerconfigjson": example_container_registry_docker_credentials.docker_credentials,
            },
            type="kubernetes.io/dockerconfigjson")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] expiry_seconds: The amount of time to pass before the Docker credentials expire in seconds. Defaults to 1576800000, or roughly 50 years. Must be greater than 0 and less than 1576800000.
        :param pulumi.Input[str] registry_name: The name of the container registry.
        :param pulumi.Input[bool] write: Allow for write access to the container registry. Defaults to false.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ContainerRegistryDockerCredentialsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Get Docker credentials for your DigitalOcean container registry.

        An error is triggered if the provided container registry name does not exist.

        ## Example Usage

        ### Basic Example

        Get the container registry:

        ```python
        import pulumi
        import pulumi_digitalocean as digitalocean

        example = digitalocean.ContainerRegistryDockerCredentials("example", registry_name="example")
        ```

        ### Docker Provider Example

        Use the `endpoint` and `docker_credentials` with the Docker provider:

        ```python
        import pulumi
        import pulumi_digitalocean as digitalocean

        example = digitalocean.get_container_registry(name="example")
        example_container_registry_docker_credentials = digitalocean.ContainerRegistryDockerCredentials("example", registry_name="example")
        ```

        ### Kubernetes Example

        Combined with the Kubernetes Provider's `kubernetes_secret` resource, you can
        access the registry from inside your cluster:

        ```python
        import pulumi
        import pulumi_digitalocean as digitalocean
        import pulumi_kubernetes as kubernetes

        example_container_registry_docker_credentials = digitalocean.ContainerRegistryDockerCredentials("example", registry_name="example")
        example = digitalocean.get_kubernetes_cluster(name="prod-cluster-01")
        example_secret = kubernetes.core.v1.Secret("example",
            metadata={
                "name": "docker-cfg",
            },
            data={
                ".dockerconfigjson": example_container_registry_docker_credentials.docker_credentials,
            },
            type="kubernetes.io/dockerconfigjson")
        ```

        :param str resource_name: The name of the resource.
        :param ContainerRegistryDockerCredentialsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ContainerRegistryDockerCredentialsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expiry_seconds: Optional[pulumi.Input[int]] = None,
                 registry_name: Optional[pulumi.Input[str]] = None,
                 write: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ContainerRegistryDockerCredentialsArgs.__new__(ContainerRegistryDockerCredentialsArgs)

            __props__.__dict__["expiry_seconds"] = expiry_seconds
            if registry_name is None and not opts.urn:
                raise TypeError("Missing required property 'registry_name'")
            __props__.__dict__["registry_name"] = registry_name
            __props__.__dict__["write"] = write
            __props__.__dict__["credential_expiration_time"] = None
            __props__.__dict__["docker_credentials"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["dockerCredentials"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ContainerRegistryDockerCredentials, __self__).__init__(
            'digitalocean:index/containerRegistryDockerCredentials:ContainerRegistryDockerCredentials',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            credential_expiration_time: Optional[pulumi.Input[str]] = None,
            docker_credentials: Optional[pulumi.Input[str]] = None,
            expiry_seconds: Optional[pulumi.Input[int]] = None,
            registry_name: Optional[pulumi.Input[str]] = None,
            write: Optional[pulumi.Input[bool]] = None) -> 'ContainerRegistryDockerCredentials':
        """
        Get an existing ContainerRegistryDockerCredentials resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] credential_expiration_time: The date and time the registry access token will expire.
        :param pulumi.Input[str] docker_credentials: Credentials for the container registry.
        :param pulumi.Input[int] expiry_seconds: The amount of time to pass before the Docker credentials expire in seconds. Defaults to 1576800000, or roughly 50 years. Must be greater than 0 and less than 1576800000.
        :param pulumi.Input[str] registry_name: The name of the container registry.
        :param pulumi.Input[bool] write: Allow for write access to the container registry. Defaults to false.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ContainerRegistryDockerCredentialsState.__new__(_ContainerRegistryDockerCredentialsState)

        __props__.__dict__["credential_expiration_time"] = credential_expiration_time
        __props__.__dict__["docker_credentials"] = docker_credentials
        __props__.__dict__["expiry_seconds"] = expiry_seconds
        __props__.__dict__["registry_name"] = registry_name
        __props__.__dict__["write"] = write
        return ContainerRegistryDockerCredentials(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="credentialExpirationTime")
    def credential_expiration_time(self) -> pulumi.Output[str]:
        """
        The date and time the registry access token will expire.
        """
        return pulumi.get(self, "credential_expiration_time")

    @property
    @pulumi.getter(name="dockerCredentials")
    def docker_credentials(self) -> pulumi.Output[str]:
        """
        Credentials for the container registry.
        """
        return pulumi.get(self, "docker_credentials")

    @property
    @pulumi.getter(name="expirySeconds")
    def expiry_seconds(self) -> pulumi.Output[Optional[int]]:
        """
        The amount of time to pass before the Docker credentials expire in seconds. Defaults to 1576800000, or roughly 50 years. Must be greater than 0 and less than 1576800000.
        """
        return pulumi.get(self, "expiry_seconds")

    @property
    @pulumi.getter(name="registryName")
    def registry_name(self) -> pulumi.Output[str]:
        """
        The name of the container registry.
        """
        return pulumi.get(self, "registry_name")

    @property
    @pulumi.getter
    def write(self) -> pulumi.Output[Optional[bool]]:
        """
        Allow for write access to the container registry. Defaults to false.
        """
        return pulumi.get(self, "write")

