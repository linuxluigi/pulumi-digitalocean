# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables

__all__ = [
    'GetCertificateResult',
    'AwaitableGetCertificateResult',
    'get_certificate',
]

@pulumi.output_type
class GetCertificateResult:
    """
    A collection of values returned by getCertificate.
    """
    def __init__(__self__, domains=None, id=None, name=None, not_after=None, sha1_fingerprint=None, state=None, type=None):
        if domains and not isinstance(domains, list):
            raise TypeError("Expected argument 'domains' to be a list")
        pulumi.set(__self__, "domains", domains)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if not_after and not isinstance(not_after, str):
            raise TypeError("Expected argument 'not_after' to be a str")
        pulumi.set(__self__, "not_after", not_after)
        if sha1_fingerprint and not isinstance(sha1_fingerprint, str):
            raise TypeError("Expected argument 'sha1_fingerprint' to be a str")
        pulumi.set(__self__, "sha1_fingerprint", sha1_fingerprint)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def domains(self) -> List[str]:
        return pulumi.get(self, "domains")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notAfter")
    def not_after(self) -> str:
        return pulumi.get(self, "not_after")

    @property
    @pulumi.getter(name="sha1Fingerprint")
    def sha1_fingerprint(self) -> str:
        return pulumi.get(self, "sha1_fingerprint")

    @property
    @pulumi.getter
    def state(self) -> str:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")


class AwaitableGetCertificateResult(GetCertificateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCertificateResult(
            domains=self.domains,
            id=self.id,
            name=self.name,
            not_after=self.not_after,
            sha1_fingerprint=self.sha1_fingerprint,
            state=self.state,
            type=self.type)


def get_certificate(name: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCertificateResult:
    """
    Get information on a certificate. This data source provides the name, type, state,
    domains, expiry date, and the sha1 fingerprint as configured on your DigitalOcean account.
    This is useful if the certificate in question is not managed by this provider or you need to utilize
    any of the certificates data.

    An error is triggered if the provided certificate name does not exist.

    ## Example Usage

    Get the certificate:

    ```python
    import pulumi
    import pulumi_digitalocean as digitalocean

    example = digitalocean.get_certificate(name="example")
    ```


    :param str name: The name of certificate.
    """
    __args__ = dict()
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('digitalocean:index/getCertificate:getCertificate', __args__, opts=opts, typ=GetCertificateResult).value

    return AwaitableGetCertificateResult(
        domains=__ret__.domains,
        id=__ret__.id,
        name=__ret__.name,
        not_after=__ret__.not_after,
        sha1_fingerprint=__ret__.sha1_fingerprint,
        state=__ret__.state,
        type=__ret__.type)
