# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetAppResult',
    'AwaitableGetAppResult',
    'get_app',
    'get_app_output',
]

@pulumi.output_type
class GetAppResult:
    """
    A collection of values returned by getApp.
    """
    def __init__(__self__, active_deployment_id=None, app_id=None, created_at=None, default_ingress=None, id=None, live_url=None, specs=None, updated_at=None):
        if active_deployment_id and not isinstance(active_deployment_id, str):
            raise TypeError("Expected argument 'active_deployment_id' to be a str")
        pulumi.set(__self__, "active_deployment_id", active_deployment_id)
        if app_id and not isinstance(app_id, str):
            raise TypeError("Expected argument 'app_id' to be a str")
        pulumi.set(__self__, "app_id", app_id)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if default_ingress and not isinstance(default_ingress, str):
            raise TypeError("Expected argument 'default_ingress' to be a str")
        pulumi.set(__self__, "default_ingress", default_ingress)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if live_url and not isinstance(live_url, str):
            raise TypeError("Expected argument 'live_url' to be a str")
        pulumi.set(__self__, "live_url", live_url)
        if specs and not isinstance(specs, list):
            raise TypeError("Expected argument 'specs' to be a list")
        pulumi.set(__self__, "specs", specs)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="activeDeploymentId")
    def active_deployment_id(self) -> str:
        """
        The ID the app's currently active deployment.
        """
        return pulumi.get(self, "active_deployment_id")

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> str:
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The date and time of when the app was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="defaultIngress")
    def default_ingress(self) -> str:
        """
        The default URL to access the app.
        """
        return pulumi.get(self, "default_ingress")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="liveUrl")
    def live_url(self) -> str:
        """
        The live URL of the app.
        """
        return pulumi.get(self, "live_url")

    @property
    @pulumi.getter
    def specs(self) -> Sequence['outputs.GetAppSpecResult']:
        """
        A DigitalOcean App spec describing the app.
        """
        return pulumi.get(self, "specs")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        The date and time of when the app was last updated.
        """
        return pulumi.get(self, "updated_at")


class AwaitableGetAppResult(GetAppResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppResult(
            active_deployment_id=self.active_deployment_id,
            app_id=self.app_id,
            created_at=self.created_at,
            default_ingress=self.default_ingress,
            id=self.id,
            live_url=self.live_url,
            specs=self.specs,
            updated_at=self.updated_at)


def get_app(app_id: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppResult:
    """
    Get information on a DigitalOcean App.

    ## Example Usage

    Get the account:

    ```python
    import pulumi
    import pulumi_digitalocean as digitalocean

    example = digitalocean.get_app(app_id="e665d18d-7b56-44a9-92ce-31979174d544")
    pulumi.export("defaultIngress", example.default_ingress)
    ```


    :param str app_id: The ID of the app to retrieve information about.
    """
    __args__ = dict()
    __args__['appId'] = app_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('digitalocean:index/getApp:getApp', __args__, opts=opts, typ=GetAppResult).value

    return AwaitableGetAppResult(
        active_deployment_id=__ret__.active_deployment_id,
        app_id=__ret__.app_id,
        created_at=__ret__.created_at,
        default_ingress=__ret__.default_ingress,
        id=__ret__.id,
        live_url=__ret__.live_url,
        specs=__ret__.specs,
        updated_at=__ret__.updated_at)


@_utilities.lift_output_func(get_app)
def get_app_output(app_id: Optional[pulumi.Input[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAppResult]:
    """
    Get information on a DigitalOcean App.

    ## Example Usage

    Get the account:

    ```python
    import pulumi
    import pulumi_digitalocean as digitalocean

    example = digitalocean.get_app(app_id="e665d18d-7b56-44a9-92ce-31979174d544")
    pulumi.export("defaultIngress", example.default_ingress)
    ```


    :param str app_id: The ID of the app to retrieve information about.
    """
    ...
