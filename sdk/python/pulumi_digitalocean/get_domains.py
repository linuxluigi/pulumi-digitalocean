# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = [
    'GetDomainsResult',
    'AwaitableGetDomainsResult',
    'get_domains',
]

@pulumi.output_type
class GetDomainsResult:
    """
    A collection of values returned by getDomains.
    """
    def __init__(__self__, domains=None, filters=None, id=None, sorts=None):
        if domains and not isinstance(domains, list):
            raise TypeError("Expected argument 'domains' to be a list")
        pulumi.set(__self__, "domains", domains)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if sorts and not isinstance(sorts, list):
            raise TypeError("Expected argument 'sorts' to be a list")
        pulumi.set(__self__, "sorts", sorts)

    @property
    @pulumi.getter
    def domains(self) -> Sequence['outputs.GetDomainsDomainResult']:
        """
        A list of domains satisfying any `filter` and `sort` criteria. Each domain has the following attributes:
        """
        return pulumi.get(self, "domains")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetDomainsFilterResult']]:
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
    def sorts(self) -> Optional[Sequence['outputs.GetDomainsSortResult']]:
        return pulumi.get(self, "sorts")


class AwaitableGetDomainsResult(GetDomainsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDomainsResult(
            domains=self.domains,
            filters=self.filters,
            id=self.id,
            sorts=self.sorts)


def get_domains(filters: Optional[Sequence[pulumi.InputType['GetDomainsFilterArgs']]] = None,
                sorts: Optional[Sequence[pulumi.InputType['GetDomainsSortArgs']]] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDomainsResult:
    """
    Use this data source to access information about an existing resource.

    :param Sequence[pulumi.InputType['GetDomainsFilterArgs']] filters: Filter the results.
           The `filter` block is documented below.
    :param Sequence[pulumi.InputType['GetDomainsSortArgs']] sorts: Sort the results.
           The `sort` block is documented below.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['sorts'] = sorts
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('digitalocean:index/getDomains:getDomains', __args__, opts=opts, typ=GetDomainsResult).value

    return AwaitableGetDomainsResult(
        domains=__ret__.domains,
        filters=__ret__.filters,
        id=__ret__.id,
        sorts=__ret__.sorts)