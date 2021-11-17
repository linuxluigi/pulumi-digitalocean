# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetRegionsResult',
    'AwaitableGetRegionsResult',
    'get_regions',
    'get_regions_output',
]

@pulumi.output_type
class GetRegionsResult:
    """
    A collection of values returned by getRegions.
    """
    def __init__(__self__, filters=None, id=None, regions=None, sorts=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if regions and not isinstance(regions, list):
            raise TypeError("Expected argument 'regions' to be a list")
        pulumi.set(__self__, "regions", regions)
        if sorts and not isinstance(sorts, list):
            raise TypeError("Expected argument 'sorts' to be a list")
        pulumi.set(__self__, "sorts", sorts)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetRegionsFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def regions(self) -> Sequence['outputs.GetRegionsRegionResult']:
        """
        A set of regions satisfying any `filter` and `sort` criteria. Each region has the following attributes:
        """
        return pulumi.get(self, "regions")

    @property
    @pulumi.getter
    def sorts(self) -> Optional[Sequence['outputs.GetRegionsSortResult']]:
        return pulumi.get(self, "sorts")


class AwaitableGetRegionsResult(GetRegionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegionsResult(
            filters=self.filters,
            id=self.id,
            regions=self.regions,
            sorts=self.sorts)


def get_regions(filters: Optional[Sequence[pulumi.InputType['GetRegionsFilterArgs']]] = None,
                sorts: Optional[Sequence[pulumi.InputType['GetRegionsSortArgs']]] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegionsResult:
    """
    Retrieve information about all supported DigitalOcean regions, with the ability to
    filter and sort the results. If no filters are specified, all regions will be returned.

    Note: You can use the `get_region` data source
    to obtain metadata about a single region if you already know the `slug` to retrieve.

    ## Example Usage

    Use the `filter` block with a `key` string and `values` list to filter regions.

    For example to find all available regions:

    ```python
    import pulumi
    import pulumi_digitalocean as digitalocean

    available = digitalocean.get_regions(filters=[digitalocean.GetRegionsFilterArgs(
        key="available",
        values=["true"],
    )])
    ```

    You can filter on multiple fields and sort the results as well:

    ```python
    import pulumi
    import pulumi_digitalocean as digitalocean

    available = digitalocean.get_regions(filters=[
            digitalocean.GetRegionsFilterArgs(
                key="available",
                values=["true"],
            ),
            digitalocean.GetRegionsFilterArgs(
                key="features",
                values=["private_networking"],
            ),
        ],
        sorts=[digitalocean.GetRegionsSortArgs(
            direction="desc",
            key="name",
        )])
    ```


    :param Sequence[pulumi.InputType['GetRegionsFilterArgs']] filters: Filter the results.
           The `filter` block is documented below.
    :param Sequence[pulumi.InputType['GetRegionsSortArgs']] sorts: Sort the results.
           The `sort` block is documented below.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['sorts'] = sorts
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('digitalocean:index/getRegions:getRegions', __args__, opts=opts, typ=GetRegionsResult).value

    return AwaitableGetRegionsResult(
        filters=__ret__.filters,
        id=__ret__.id,
        regions=__ret__.regions,
        sorts=__ret__.sorts)


@_utilities.lift_output_func(get_regions)
def get_regions_output(filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetRegionsFilterArgs']]]]] = None,
                       sorts: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetRegionsSortArgs']]]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRegionsResult]:
    """
    Retrieve information about all supported DigitalOcean regions, with the ability to
    filter and sort the results. If no filters are specified, all regions will be returned.

    Note: You can use the `get_region` data source
    to obtain metadata about a single region if you already know the `slug` to retrieve.

    ## Example Usage

    Use the `filter` block with a `key` string and `values` list to filter regions.

    For example to find all available regions:

    ```python
    import pulumi
    import pulumi_digitalocean as digitalocean

    available = digitalocean.get_regions(filters=[digitalocean.GetRegionsFilterArgs(
        key="available",
        values=["true"],
    )])
    ```

    You can filter on multiple fields and sort the results as well:

    ```python
    import pulumi
    import pulumi_digitalocean as digitalocean

    available = digitalocean.get_regions(filters=[
            digitalocean.GetRegionsFilterArgs(
                key="available",
                values=["true"],
            ),
            digitalocean.GetRegionsFilterArgs(
                key="features",
                values=["private_networking"],
            ),
        ],
        sorts=[digitalocean.GetRegionsSortArgs(
            direction="desc",
            key="name",
        )])
    ```


    :param Sequence[pulumi.InputType['GetRegionsFilterArgs']] filters: Filter the results.
           The `filter` block is documented below.
    :param Sequence[pulumi.InputType['GetRegionsSortArgs']] sorts: Sort the results.
           The `sort` block is documented below.
    """
    ...
