# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class GetKubernetesClusterResult:
    """
    A collection of values returned by getKubernetesCluster.
    """
    def __init__(__self__, cluster_subnet=None, created_at=None, endpoint=None, ipv4_address=None, kube_configs=None, name=None, node_pools=None, region=None, service_subnet=None, status=None, tags=None, updated_at=None, version=None, id=None):
        if cluster_subnet and not isinstance(cluster_subnet, str):
            raise TypeError("Expected argument 'cluster_subnet' to be a str")
        __self__.cluster_subnet = cluster_subnet
        """
        The range of IP addresses in the overlay network of the Kubernetes cluster.
        """
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        __self__.created_at = created_at
        """
        The date and time when the Kubernetes cluster was created.
        """
        if endpoint and not isinstance(endpoint, str):
            raise TypeError("Expected argument 'endpoint' to be a str")
        __self__.endpoint = endpoint
        """
        The base URL of the API server on the Kubernetes master node.
        """
        if ipv4_address and not isinstance(ipv4_address, str):
            raise TypeError("Expected argument 'ipv4_address' to be a str")
        __self__.ipv4_address = ipv4_address
        """
        The public IPv4 address of the Kubernetes master node.
        """
        if kube_configs and not isinstance(kube_configs, list):
            raise TypeError("Expected argument 'kube_configs' to be a list")
        __self__.kube_configs = kube_configs
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if node_pools and not isinstance(node_pools, list):
            raise TypeError("Expected argument 'node_pools' to be a list")
        __self__.node_pools = node_pools
        """
        A list of node pools associated with the cluster. Each node pool exports the following attributes:
        - `id` -  The unique ID that can be used to identify and reference the node pool.
        - `name` - The name of the node pool.
        - `size` - The slug identifier for the type of Droplet used as workers in the node pool.
        - `node_count` - The number of Droplet instances in the node pool.
        - `tags` - A list of tag names applied to the node pool.
        - `nodes` - A list of nodes in the pool. Each node exports the following attributes:
        + `id` -  A unique ID that can be used to identify and reference the node.
        + `name` - The auto-generated name for the node.
        + `status` -  A string indicating the current status of the individual node.
        + `created_at` - The date and time when the node was created.
        + `updated_at` - The date and time when the node was last updated.
        """
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        """
        The slug identifier for the region where the Kubernetes cluster is located.
        """
        if service_subnet and not isinstance(service_subnet, str):
            raise TypeError("Expected argument 'service_subnet' to be a str")
        __self__.service_subnet = service_subnet
        """
        The range of assignable IP addresses for services running in the Kubernetes cluster.
        """
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        __self__.status = status
        """
        A string indicating the current status of the cluster. Potential values include running, provisioning, and errored.
        """
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        __self__.tags = tags
        """
        A list of tag names to be applied to the Kubernetes cluster.
        """
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        __self__.updated_at = updated_at
        """
        The date and time when the Kubernetes cluster was last updated.
        * `kube_config.0` - A representation of the Kubernetes cluster's kubeconfig with the following attributes:
        - `raw_config` - The full contents of the Kubernetes cluster's kubeconfig file.
        - `host` - The URL of the API server on the Kubernetes master node.
        - `client_key` - The base64 encoded private key used by clients to access the cluster.
        - `client_certificate` - The base64 encoded public certificate used by clients to access the cluster.
        - `cluster_ca_certificate` - The base64 encoded public certificate for the cluster's certificate authority.
        """
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        __self__.version = version
        """
        The slug identifier for the version of Kubernetes used for the cluster.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetKubernetesClusterResult(GetKubernetesClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKubernetesClusterResult(
            cluster_subnet=self.cluster_subnet,
            created_at=self.created_at,
            endpoint=self.endpoint,
            ipv4_address=self.ipv4_address,
            kube_configs=self.kube_configs,
            name=self.name,
            node_pools=self.node_pools,
            region=self.region,
            service_subnet=self.service_subnet,
            status=self.status,
            tags=self.tags,
            updated_at=self.updated_at,
            version=self.version,
            id=self.id)

def get_kubernetes_cluster(name=None,tags=None,opts=None):
    """
    > This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/d/kubernetes_cluster.html.markdown.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('digitalocean:index/getKubernetesCluster:getKubernetesCluster', __args__, opts=opts).value

    return AwaitableGetKubernetesClusterResult(
        cluster_subnet=__ret__.get('clusterSubnet'),
        created_at=__ret__.get('createdAt'),
        endpoint=__ret__.get('endpoint'),
        ipv4_address=__ret__.get('ipv4Address'),
        kube_configs=__ret__.get('kubeConfigs'),
        name=__ret__.get('name'),
        node_pools=__ret__.get('nodePools'),
        region=__ret__.get('region'),
        service_subnet=__ret__.get('serviceSubnet'),
        status=__ret__.get('status'),
        tags=__ret__.get('tags'),
        updated_at=__ret__.get('updatedAt'),
        version=__ret__.get('version'),
        id=__ret__.get('id'))
