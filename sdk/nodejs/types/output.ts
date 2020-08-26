// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface DatabaseClusterMaintenanceWindow {
    /**
     * The day of the week on which to apply maintenance updates.
     */
    day: string;
    /**
     * The hour in UTC at which maintenance updates will be applied in 24 hour format.
     */
    hour: string;
}

export interface DatabaseFirewallRule {
    /**
     * The date and time when the firewall rule was created.
     */
    createdAt: string;
    /**
     * The type of resource that the firewall rule allows to access the database cluster. The possible values are: `droplet`, `k8s`, `ipAddr`, or `tag`.
     */
    type: string;
    /**
     * A unique identifier for the firewall rule.
     */
    uuid: string;
    /**
     * The ID of the specific resource, the name of a tag applied to a group of resources, or the IP address that the firewall rule allows to access the database cluster.
     */
    value: string;
}

export interface FirewallInboundRule {
    /**
     * The ports on which traffic will be allowed
     * specified as a string containing a single port, a range (e.g. "8000-9000"),
     * or "1-65535" to open all ports for a protocol. Required for when protocol is
     * `tcp` or `udp`.
     */
    portRange?: string;
    /**
     * The type of traffic to be allowed.
     * This may be one of "tcp", "udp", or "icmp".
     */
    protocol: string;
    /**
     * An array of strings containing the IPv4
     * addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs from which the
     * inbound traffic will be accepted.
     */
    sourceAddresses?: string[];
    /**
     * An array containing the IDs of
     * the Droplets from which the inbound traffic will be accepted.
     */
    sourceDropletIds?: number[];
    /**
     * An array containing the IDs
     * of the Load Balancers from which the inbound traffic will be accepted.
     */
    sourceLoadBalancerUids?: string[];
    /**
     * An array containing the names of Tags
     * corresponding to groups of Droplets from which the inbound traffic
     * will be accepted.
     */
    sourceTags?: string[];
}

export interface FirewallOutboundRule {
    /**
     * An array of strings containing the IPv4
     * addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs to which the
     * outbound traffic will be allowed.
     */
    destinationAddresses?: string[];
    /**
     * An array containing the IDs of
     * the Droplets to which the outbound traffic will be allowed.
     */
    destinationDropletIds?: number[];
    /**
     * An array containing the IDs
     * of the Load Balancers to which the outbound traffic will be allowed.
     */
    destinationLoadBalancerUids?: string[];
    /**
     * An array containing the names of Tags
     * corresponding to groups of Droplets to which the outbound traffic will
     * be allowed.
     * traffic.
     */
    destinationTags?: string[];
    /**
     * The ports on which traffic will be allowed
     * specified as a string containing a single port, a range (e.g. "8000-9000"),
     * or "1-65535" to open all ports for a protocol. Required for when protocol is
     * `tcp` or `udp`.
     */
    portRange?: string;
    /**
     * The type of traffic to be allowed.
     * This may be one of "tcp", "udp", or "icmp".
     */
    protocol: string;
}

export interface FirewallPendingChange {
    dropletId?: number;
    removing?: boolean;
    /**
     * A status string indicating the current state of the Firewall.
     * This can be "waiting", "succeeded", or "failed".
     */
    status?: string;
}

export interface GetDatabaseClusterMaintenanceWindow {
    /**
     * The day of the week on which to apply maintenance updates.
     */
    day: string;
    /**
     * The hour in UTC at which maintenance updates will be applied in 24 hour format.
     */
    hour: string;
}

export interface GetDropletsDroplet {
    /**
     * Whether backups are enabled.
     */
    backups: boolean;
    createdAt: string;
    /**
     * The size of the Droplet's disk in GB.
     */
    disk: number;
    /**
     * The ID of the Droplet.
     */
    id: number;
    /**
     * The Droplet image ID or slug.
     */
    image: string;
    /**
     * The Droplet's public IPv4 address
     */
    ipv4Address: string;
    /**
     * The Droplet's private IPv4 address
     */
    ipv4AddressPrivate: string;
    /**
     * Whether IPv6 is enabled.
     */
    ipv6: boolean;
    /**
     * The Droplet's public IPv6 address
     */
    ipv6Address: string;
    /**
     * The Droplet's private IPv6 address
     */
    ipv6AddressPrivate: string;
    /**
     * Whether the Droplet is locked.
     */
    locked: boolean;
    /**
     * The amount of the Droplet's memory in MB.
     */
    memory: number;
    /**
     * Whether monitoring agent is installed.
     */
    monitoring: boolean;
    name: string;
    /**
     * Droplet hourly price.
     */
    priceHourly: number;
    /**
     * Droplet monthly price.
     */
    priceMonthly: number;
    /**
     * Whether private networks are enabled.
     */
    privateNetworking: boolean;
    /**
     * The region the Droplet is running in.
     */
    region: string;
    /**
     * The unique slug that identifies the type of Droplet.
     */
    size: string;
    /**
     * The status of the Droplet.
     */
    status: string;
    /**
     * A list of the tags associated to the Droplet.
     */
    tags: string[];
    /**
     * The uniform resource name of the Droplet
     */
    urn: string;
    /**
     * The number of the Droplet's virtual CPUs.
     */
    vcpus: number;
    /**
     * List of the IDs of each volumes attached to the Droplet.
     */
    volumeIds: string[];
    /**
     * The ID of the VPC where the Droplet is located.
     */
    vpcUuid: string;
}

export interface GetDropletsFilter {
    /**
     * Filter the Droplets by this key. This may be one of '`backups`, `createdAt`, `disk`, `id`,
     * `image`, `ipv4Address`, `ipv4AddressPrivate`, `ipv6`, `ipv6Address`, `ipv6AddressPrivate`, `locked`,
     * `memory`, `monitoring`, `name`, `priceHourly`, `priceMonthly`, `privateNetworking`, `region`, `size`,
     * `status`, `tags`, `urn`, `vcpus`, `volumeIds`, or `vpcUuid`'.
     */
    key: string;
    /**
     * A list of values to match against the `key` field. Only retrieves Droplets
     * where the `key` field takes on one or more of the values provided here.
     */
    values: string[];
}

export interface GetDropletsSort {
    /**
     * The sort direction. This may be either `asc` or `desc`.
     */
    direction?: string;
    /**
     * Sort the Droplets by this key. This may be one of `backups`, `createdAt`, `disk`, `id`,
     * `image`, `ipv4Address`, `ipv4AddressPrivate`, `ipv6`, `ipv6Address`, `ipv6AddressPrivate`, `locked`,
     * `memory`, `monitoring`, `name`, `priceHourly`, `priceMonthly`, `privateNetworking`, `region`, `size`,
     * `status`, `urn`, `vcpus`, or `vpcUuid`.
     */
    key: string;
}

export interface GetImagesFilter {
    /**
     * Filter the images by this key. This may be one of `distribution`, `errorMessage`,
     * `id`, `image`, `minDiskSize`, `name`, `private`, `regions`, `sizeGigabytes`, `slug`, `status`,
     * `tags`, or `type`.
     */
    key: string;
    /**
     * A list of values to match against the `key` field. Only retrieves images
     * where the `key` field takes on one or more of the values provided here.
     */
    values: string[];
}

export interface GetImagesImage {
    created: string;
    /**
     * The name of the distribution of the OS of the image.
     * - `minDiskSize`: The minimum 'disk' required for the image.
     * - `sizeGigabytes`: The size of the image in GB.
     */
    distribution: string;
    errorMessage: string;
    id: number;
    /**
     * The id of the image (legacy parameter).
     */
    image: string;
    minDiskSize: number;
    name: string;
    /**
     * Is image a public image or not. Public images represent
     * Linux distributions or One-Click Applications, while non-public images represent
     * snapshots and backups and are only available within your account.
     * - `regions`: A set of the regions that the image is available in.
     * - `tags`: A set of tags applied to the image
     * - `created`: When the image was created
     * - `status`: Current status of the image
     * - `errorMessage`: Any applicable error message pertaining to the image
     */
    private: boolean;
    regions: string[];
    sizeGigabytes: number;
    slug: string;
    status: string;
    tags: string[];
    type: string;
}

export interface GetImagesSort {
    /**
     * The sort direction. This may be either `asc` or `desc`.
     */
    direction?: string;
    /**
     * Sort the images by this key. This may be one of `distribution`, `errorMessage`, `id`,
     * `image`, `minDiskSize`, `name`, `private`, `sizeGigabytes`, `slug`, `status`, or `type`.
     */
    key: string;
}

export interface GetKubernetesClusterKubeConfig {
    /**
     * The base64 encoded public certificate used by clients to access the cluster. Only available if token authentication is not supported on your cluster.
     */
    clientCertificate: string;
    /**
     * The base64 encoded private key used by clients to access the cluster. Only available if token authentication is not supported on your cluster.
     */
    clientKey: string;
    /**
     * The base64 encoded public certificate for the cluster's certificate authority.
     */
    clusterCaCertificate: string;
    /**
     * The date and time when the credentials will expire and need to be regenerated.
     */
    expiresAt: string;
    /**
     * The URL of the API server on the Kubernetes master node.
     */
    host: string;
    /**
     * The full contents of the Kubernetes cluster's kubeconfig file.
     */
    rawConfig: string;
    /**
     * The DigitalOcean API access token used by clients to access the cluster.
     */
    token: string;
}

export interface GetKubernetesClusterNodePool {
    /**
     * The actual number of nodes in the node pool, which is especially useful when auto-scaling is enabled.
     */
    actualNodeCount: number;
    /**
     * A boolean indicating whether auto-scaling is enabled on the node pool.
     */
    autoScale: boolean;
    /**
     * A unique ID that can be used to identify and reference the node.
     */
    id: string;
    /**
     * A map of key/value pairs applied to nodes in the pool. The labels are exposed in the Kubernetes API as labels in the metadata of the corresponding [Node resources](https://kubernetes.io/docs/concepts/architecture/nodes/).
     */
    labels: {[key: string]: string};
    /**
     * If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to.
     */
    maxNodes: number;
    /**
     * If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to.
     */
    minNodes: number;
    /**
     * The name of Kubernetes cluster.
     */
    name: string;
    /**
     * The number of Droplet instances in the node pool.
     */
    nodeCount: number;
    /**
     * A list of nodes in the pool. Each node exports the following attributes:
     */
    nodes: outputs.GetKubernetesClusterNodePoolNode[];
    /**
     * The slug identifier for the type of Droplet used as workers in the node pool.
     */
    size: string;
    /**
     * A list of tag names applied to the node pool.
     */
    tags?: string[];
}

export interface GetKubernetesClusterNodePoolNode {
    /**
     * The date and time when the node was created.
     */
    createdAt: string;
    dropletId: string;
    /**
     * A unique ID that can be used to identify and reference the node.
     */
    id: string;
    /**
     * The name of Kubernetes cluster.
     */
    name: string;
    /**
     * A string indicating the current status of the individual node.
     */
    status: string;
    /**
     * The date and time when the node was last updated.
     */
    updatedAt: string;
}

export interface GetLoadBalancerForwardingRule {
    certificateId: string;
    entryPort: number;
    entryProtocol: string;
    targetPort: number;
    targetProtocol: string;
    tlsPassthrough: boolean;
}

export interface GetLoadBalancerHealthcheck {
    checkIntervalSeconds: number;
    healthyThreshold: number;
    path: string;
    port: number;
    protocol: string;
    responseTimeoutSeconds: number;
    unhealthyThreshold: number;
}

export interface GetLoadBalancerStickySessions {
    cookieName: string;
    cookieTtlSeconds: number;
    type: string;
}

export interface GetProjectsFilter {
    /**
     * Filter the projects by this key. This may be one of `name`,
     * `purpose`, `description`, `environment`, or `isDefault`.
     */
    key: string;
    /**
     * A list of values to match against the `key` field. Only retrieves projects
     * where the `key` field takes on one or more of the values provided here.
     */
    values: string[];
}

export interface GetProjectsProject {
    /**
     * The date and time when the project was created, (ISO8601)
     */
    createdAt: string;
    /**
     * The description of the project
     */
    description: string;
    /**
     * The environment of the project's resources. The possible values are: `Development`, `Staging`, `Production`.
     */
    environment: string;
    /**
     * The ID of the project
     */
    id: string;
    isDefault: boolean;
    /**
     * The name of the project
     */
    name: string;
    /**
     * The ID of the project owner
     */
    ownerId: number;
    /**
     * The unique universal identifier of the project owner
     */
    ownerUuid: string;
    /**
     * The purpose of the project (Default: "Web Application")
     */
    purpose: string;
    /**
     * A set of uniform resource names (URNs) for the resources associated with the project
     */
    resources: string[];
    /**
     * The date and time when the project was last updated, (ISO8601)
     */
    updatedAt: string;
}

export interface GetProjectsSort {
    /**
     * The sort direction. This may be either `asc` or `desc`.
     */
    direction?: string;
    /**
     * Sort the projects by this key. This may be one of `name`,
     * `purpose`, `description`, or `environment`.
     */
    key: string;
}

export interface GetRegionsFilter {
    /**
     * Filter the regions by this key. This may be one of `slug`,
     * `name`, `available`, `features`, or `sizes`.
     */
    key: string;
    /**
     * A list of values to match against the `key` field. Only retrieves regions
     * where the `key` field takes on one or more of the values provided here.
     */
    values: string[];
}

export interface GetRegionsRegion {
    /**
     * A boolean value that represents whether new Droplets can be created in this region.
     */
    available: boolean;
    /**
     * A set of features available in this region.
     */
    features: string[];
    /**
     * The display name of the region.
     */
    name: string;
    /**
     * A set of identifying slugs for the Droplet sizes available in this region.
     */
    sizes: string[];
    /**
     * A human-readable string that is used as a unique identifier for each region.
     */
    slug: string;
}

export interface GetRegionsSort {
    /**
     * The sort direction. This may be either `asc` or `desc`.
     */
    direction?: string;
    /**
     * Sort the regions by this key. This may be one of `slug`,
     * `name`, or `available`.
     */
    key: string;
}

export interface GetSizesFilter {
    /**
     * Filter the sizes by this key. This may be one of `slug`,
     * `regions`, `memory`, `vcpus`, `disk`, `transfer`, `priceMonthly`,
     * `priceHourly`, or `available`.
     */
    key: string;
    /**
     * Only retrieves images which keys has value that matches
     * one of the values provided here.
     */
    values: string[];
}

export interface GetSizesSize {
    /**
     * This represents whether new Droplets can be created with this size.
     */
    available: boolean;
    /**
     * The amount of disk space set aside for Droplets of this size. The value is measured in gigabytes.
     */
    disk: number;
    /**
     * The amount of RAM allocated to Droplets created of this size. The value is measured in megabytes.
     */
    memory: number;
    /**
     * The hourly cost of Droplets created in this size as measured hourly. The value is measured in US dollars.
     */
    priceHourly: number;
    /**
     * The monthly cost of Droplets created in this size if they are kept for an entire month. The value is measured in US dollars.
     */
    priceMonthly: number;
    /**
     * List of region slugs where Droplets can be created in this size.
     */
    regions: string[];
    /**
     * A human-readable string that is used to uniquely identify each size.
     */
    slug: string;
    /**
     * The amount of transfer bandwidth that is available for Droplets created in this size. This only counts traffic on the public interface. The value is given in terabytes.
     */
    transfer: number;
    /**
     * The number of CPUs allocated to Droplets of this size.
     */
    vcpus: number;
}

export interface GetSizesSort {
    /**
     * The sort direction. This may be either `asc` or `desc`.
     */
    direction?: string;
    /**
     * Sort the sizes by this key. This may be one of `slug`,
     * `memory`, `vcpus`, `disk`, `transfer`, `priceMonthly`, or `priceHourly`.
     */
    key: string;
}

export interface GetSpacesBucketsBucket {
    /**
     * The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com)
     */
    bucketDomainName: string;
    /**
     * The name of the Spaces bucket
     */
    name: string;
    /**
     * The slug of the region where the bucket is stored.
     */
    region: string;
    /**
     * The uniform resource name of the bucket
     */
    urn: string;
}

export interface GetSpacesBucketsFilter {
    /**
     * Filter the images by this key. This may be one of `bucketDomainName`, `name`, `region`, or `urn`.
     */
    key: string;
    /**
     * A list of values to match against the `key` field. Only retrieves images
     * where the `key` field takes on one or more of the values provided here.
     */
    values: string[];
}

export interface GetSpacesBucketsSort {
    /**
     * The sort direction. This may be either `asc` or `desc`.
     */
    direction?: string;
    /**
     * Sort the images by this key. This may be one of `bucketDomainName`, `name`, `region`, or `urn`.
     */
    key: string;
}

export interface GetTagsFilter {
    /**
     * Filter the tags by this key. This may be one of `name`, `totalResourceCount`,  `dropletsCount`, `imagesCount`, `volumesCount`, `volumeSnapshotsCount`, or `databasesCount`.
     */
    key: string;
    /**
     * Only retrieves tags which keys has value that matches
     * one of the values provided here.
     */
    values: string[];
}

export interface GetTagsSort {
    /**
     * The sort direction. This may be either `asc` or `desc`.
     */
    direction?: string;
    /**
     * Sort the tags by this key. This may be one of `name`, `totalResourceCount`,  `dropletsCount`, `imagesCount`, `volumesCount`, `volumeSnapshotsCount`, or `databasesCount`.
     */
    key: string;
}

export interface GetTagsTag {
    /**
     * A count of the database clusters that the tag is applied to.
     */
    databasesCount: number;
    /**
     * A count of the Droplets the tag is applied to.
     */
    dropletsCount: number;
    /**
     * A count of the images that the tag is applied to.
     */
    imagesCount: number;
    /**
     * The name of the tag.
     */
    name: string;
    /**
     * A count of the total number of resources that the tag is applied to.
     */
    totalResourceCount: number;
    /**
     * A count of the volume snapshots that the tag is applied to.
     */
    volumeSnapshotsCount: number;
    /**
     * A count of the volumes that the tag is applied to.
     */
    volumesCount: number;
}

export interface KubernetesClusterKubeConfig {
    /**
     * The base64 encoded public certificate used by clients to access the cluster. Only available if token authentication is not supported on your cluster.
     */
    clientCertificate: string;
    /**
     * The base64 encoded private key used by clients to access the cluster. Only available if token authentication is not supported on your cluster.
     */
    clientKey: string;
    /**
     * The base64 encoded public certificate for the cluster's certificate authority.
     */
    clusterCaCertificate: string;
    /**
     * The date and time when the credentials will expire and need to be regenerated.
     */
    expiresAt: string;
    /**
     * The URL of the API server on the Kubernetes master node.
     */
    host: string;
    /**
     * The full contents of the Kubernetes cluster's kubeconfig file.
     */
    rawConfig: string;
    /**
     * The DigitalOcean API access token used by clients to access the cluster.
     */
    token: string;
}

export interface KubernetesClusterNodePool {
    /**
     * A computed field representing the actual number of nodes in the node pool, which is especially useful when auto-scaling is enabled.
     */
    actualNodeCount: number;
    /**
     * Enable auto-scaling of the number of nodes in the node pool within the given min/max range.
     */
    autoScale?: boolean;
    /**
     * A unique ID that can be used to identify and reference the node.
     */
    id: string;
    /**
     * A map of key/value pairs to apply to nodes in the pool. The labels are exposed in the Kubernetes API as labels in the metadata of the corresponding [Node resources](https://kubernetes.io/docs/concepts/architecture/nodes/).
     */
    labels?: {[key: string]: string};
    /**
     * If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to.
     */
    maxNodes?: number;
    /**
     * If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to.
     */
    minNodes?: number;
    /**
     * A name for the node pool.
     */
    name: string;
    /**
     * The number of Droplet instances in the node pool. If auto-scaling is enabled, this should only be set if the desired result is to explicitly reset the number of nodes to this value. If auto-scaling is enabled, and the node count is outside of the given min/max range, it will use the min nodes value.
     */
    nodeCount?: number;
    /**
     * A list of nodes in the pool. Each node exports the following attributes:
     */
    nodes: outputs.KubernetesClusterNodePoolNode[];
    /**
     * The slug identifier for the type of Droplet to be used as workers in the node pool.
     */
    size: string;
    /**
     * A list of tag names to be applied to the Kubernetes cluster.
     */
    tags?: string[];
}

export interface KubernetesClusterNodePoolNode {
    /**
     * The date and time when the node was created.
     */
    createdAt: string;
    /**
     * The id of the node's droplet
     */
    dropletId: string;
    /**
     * A unique ID that can be used to identify and reference the node.
     */
    id: string;
    /**
     * A name for the node pool.
     */
    name: string;
    /**
     * A string indicating the current status of the individual node.
     */
    status: string;
    /**
     * The date and time when the node was last updated.
     */
    updatedAt: string;
}

export interface KubernetesNodePoolNode {
    /**
     * The date and time when the node was created.
     */
    createdAt: string;
    /**
     * The id of the node's droplet
     */
    dropletId: string;
    /**
     * A unique ID that can be used to identify and reference the node.
     */
    id: string;
    /**
     * A name for the node pool.
     */
    name: string;
    /**
     * A string indicating the current status of the individual node.
     */
    status: string;
    /**
     * The date and time when the node was last updated.
     */
    updatedAt: string;
}

export interface LoadBalancerForwardingRule {
    /**
     * The ID of the TLS certificate to be used for SSL termination.
     */
    certificateId?: string;
    /**
     * An integer representing the port on which the Load Balancer instance will listen.
     */
    entryPort: number;
    /**
     * The protocol used for traffic to the Load Balancer. The possible values are: `http`, `https`, `http2` or `tcp`.
     */
    entryProtocol: string;
    /**
     * An integer representing the port on the backend Droplets to which the Load Balancer will send traffic.
     */
    targetPort: number;
    /**
     * The protocol used for traffic from the Load Balancer to the backend Droplets. The possible values are: `http`, `https`, `http2` or `tcp`.
     */
    targetProtocol: string;
    /**
     * A boolean value indicating whether SSL encrypted traffic will be passed through to the backend Droplets. The default value is `false`.
     */
    tlsPassthrough?: boolean;
}

export interface LoadBalancerHealthcheck {
    /**
     * The number of seconds between between two consecutive health checks. If not specified, the default value is `10`.
     */
    checkIntervalSeconds?: number;
    /**
     * The number of times a health check must pass for a backend Droplet to be marked "healthy" and be re-added to the pool. If not specified, the default value is `5`.
     */
    healthyThreshold?: number;
    /**
     * The path on the backend Droplets to which the Load Balancer instance will send a request.
     */
    path?: string;
    /**
     * An integer representing the port on the backend Droplets on which the health check will attempt a connection.
     */
    port: number;
    /**
     * The protocol used for health checks sent to the backend Droplets. The possible values are `http` or `tcp`.
     */
    protocol: string;
    /**
     * The number of seconds the Load Balancer instance will wait for a response until marking a health check as failed. If not specified, the default value is `5`.
     */
    responseTimeoutSeconds?: number;
    /**
     * The number of times a health check must fail for a backend Droplet to be marked "unhealthy" and be removed from the pool. If not specified, the default value is `3`.
     */
    unhealthyThreshold?: number;
}

export interface LoadBalancerStickySessions {
    /**
     * The name to be used for the cookie sent to the client. This attribute is required when using `cookies` for the sticky sessions type.
     */
    cookieName?: string;
    /**
     * The number of seconds until the cookie set by the Load Balancer expires. This attribute is required when using `cookies` for the sticky sessions type.
     */
    cookieTtlSeconds?: number;
    /**
     * An attribute indicating how and if requests from a client will be persistently served by the same backend Droplet. The possible values are `cookies` or `none`. If not specified, the default value is `none`.
     */
    type?: string;
}

export interface SpacesBucketCorsRule {
    /**
     * A list of headers that will be included in the CORS preflight request's `Access-Control-Request-Headers`. A header may contain one wildcard (e.g. `x-amz-*`).
     */
    allowedHeaders?: string[];
    /**
     * A list of HTTP methods (e.g. `GET`) which are allowed from the specified origin.
     */
    allowedMethods: string[];
    /**
     * A list of hosts from which requests using the specified methods are allowed. A host may contain one wildcard (e.g. http://*.example.com).
     */
    allowedOrigins: string[];
    /**
     * The time in seconds that browser can cache the response for a preflight request.
     */
    maxAgeSeconds?: number;
}

export interface SpacesBucketLifecycleRule {
    /**
     * Specifies the number of days after initiating a multipart
     * upload when the multipart upload must be completed or else Spaces will abort the upload.
     */
    abortIncompleteMultipartUploadDays?: number;
    /**
     * Specifies lifecycle rule status.
     */
    enabled: boolean;
    /**
     * Specifies a time period after which applicable objects expire (documented below).
     */
    expiration?: outputs.SpacesBucketLifecycleRuleExpiration;
    /**
     * Unique identifier for the rule.
     */
    id: string;
    /**
     * Specifies when non-current object versions expire (documented below).
     */
    noncurrentVersionExpiration?: outputs.SpacesBucketLifecycleRuleNoncurrentVersionExpiration;
    /**
     * Object key prefix identifying one or more objects to which the rule applies.
     */
    prefix?: string;
}

export interface SpacesBucketLifecycleRuleExpiration {
    /**
     * Specifies the date/time after which you want applicable objects to expire. The argument uses
     * RFC3339 format, e.g. "2020-03-22T15:03:55Z" or parts thereof e.g. "2019-02-28".
     */
    date?: string;
    /**
     * Specifies the number of days after object creation when the applicable objects will expire.
     */
    days?: number;
    /**
     * On a versioned bucket (versioning-enabled or versioning-suspended
     * bucket), setting this to true directs Spaces to delete expired object delete markers.
     */
    expiredObjectDeleteMarker?: boolean;
}

export interface SpacesBucketLifecycleRuleNoncurrentVersionExpiration {
    /**
     * Specifies the number of days after which an object's non-current versions expire.
     */
    days?: number;
}

export interface SpacesBucketVersioning {
    /**
     * Enable versioning. Once you version-enable a bucket, it can never return to an unversioned
     * state. You can, however, suspend versioning on that bucket.
     */
    enabled?: boolean;
}
