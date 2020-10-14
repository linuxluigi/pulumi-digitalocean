// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface AppSpec {
    databases?: pulumi.Input<pulumi.Input<inputs.AppSpecDatabase>[]>;
    /**
     * A list of hostnames where the application will be available.
     */
    domains?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the component
     */
    name: pulumi.Input<string>;
    /**
     * The slug for the DigitalOcean data center region hosting the app.
     */
    region?: pulumi.Input<string>;
    services?: pulumi.Input<pulumi.Input<inputs.AppSpecService>[]>;
    staticSites?: pulumi.Input<pulumi.Input<inputs.AppSpecStaticSite>[]>;
    workers?: pulumi.Input<pulumi.Input<inputs.AppSpecWorker>[]>;
}

export interface AppSpecDatabase {
    clusterName?: pulumi.Input<string>;
    dbName?: pulumi.Input<string>;
    dbUser?: pulumi.Input<string>;
    engine?: pulumi.Input<string>;
    /**
     * The name of the component
     */
    name?: pulumi.Input<string>;
    production?: pulumi.Input<boolean>;
    version?: pulumi.Input<string>;
}

export interface AppSpecService {
    /**
     * An optional build command to run while building this component from source.
     */
    buildCommand?: pulumi.Input<string>;
    /**
     * The path to a Dockerfile relative to the root of the repo. If set, overrides usage of buildpacks.
     */
    dockerfilePath?: pulumi.Input<string>;
    /**
     * An environment slug describing the type of this app.
     */
    environmentSlug?: pulumi.Input<string>;
    /**
     * Describes an environment variable made available to an app competent.
     */
    envs?: pulumi.Input<pulumi.Input<inputs.AppSpecServiceEnv>[]>;
    /**
     * A Git repo to use as component's source. Only one of `git` and `github` may be set.
     */
    git?: pulumi.Input<inputs.AppSpecServiceGit>;
    /**
     * A GitHub repo to use as component's source. Only one of `git` and `github` may be set.
     */
    github?: pulumi.Input<inputs.AppSpecServiceGithub>;
    /**
     * A health check to determine the availability of this component.
     */
    healthCheck?: pulumi.Input<inputs.AppSpecServiceHealthCheck>;
    /**
     * The internal port on which this service's run command will listen.
     */
    httpPort?: pulumi.Input<number>;
    /**
     * The amount of instances that this component should be scaled to.
     */
    instanceCount?: pulumi.Input<number>;
    /**
     * The instance size to use for this component.
     */
    instanceSizeSlug?: pulumi.Input<string>;
    /**
     * The name of the component
     */
    name: pulumi.Input<string>;
    routes?: pulumi.Input<inputs.AppSpecServiceRoutes>;
    /**
     * An optional run command to override the component's default.
     */
    runCommand?: pulumi.Input<string>;
    /**
     * An optional path to the working directory to use for the build.
     */
    sourceDir?: pulumi.Input<string>;
}

export interface AppSpecServiceEnv {
    /**
     * The name of the environment variable.
     */
    key?: pulumi.Input<string>;
    /**
     * The visibility scope of the environment variable. One of `RUN_TIME`, `BUILD_TIME`, or `RUN_AND_BUILD_TIME` (default).
     */
    scope?: pulumi.Input<string>;
    /**
     * The type of the environment variable, `GENERAL` or `SECRET`.
     */
    type?: pulumi.Input<string>;
    /**
     * The value of the environment variable.
     */
    value?: pulumi.Input<string>;
}

export interface AppSpecServiceGit {
    /**
     * The name of the branch to use.
     */
    branch?: pulumi.Input<string>;
    /**
     * The clone URL of the repo.
     */
    repoCloneUrl?: pulumi.Input<string>;
}

export interface AppSpecServiceGithub {
    /**
     * The name of the branch to use.
     */
    branch?: pulumi.Input<string>;
    /**
     * Whether to automatically deploy new commits made to the repo.
     */
    deployOnPush?: pulumi.Input<boolean>;
    /**
     * The name of the repo in the format `owner/repo`.
     */
    repo?: pulumi.Input<string>;
}

export interface AppSpecServiceHealthCheck {
    /**
     * The number of failed health checks before considered unhealthy.
     */
    failureThreshold?: pulumi.Input<number>;
    /**
     * The route path used for the HTTP health check ping.
     */
    httpPath?: pulumi.Input<string>;
    /**
     * The number of seconds to wait before beginning health checks.
     */
    initialDelaySeconds?: pulumi.Input<number>;
    /**
     * The number of seconds to wait between health checks.
     */
    periodSeconds?: pulumi.Input<number>;
    /**
     * The number of successful health checks before considered healthy.
     */
    successThreshold?: pulumi.Input<number>;
    /**
     * The number of seconds after which the check times out.
     */
    timeoutSeconds?: pulumi.Input<number>;
}

export interface AppSpecServiceRoutes {
    /**
     * Paths must start with `/` and must be unique within the app.
     */
    path?: pulumi.Input<string>;
}

export interface AppSpecStaticSite {
    /**
     * An optional build command to run while building this component from source.
     */
    buildCommand?: pulumi.Input<string>;
    /**
     * The path to a Dockerfile relative to the root of the repo. If set, overrides usage of buildpacks.
     */
    dockerfilePath?: pulumi.Input<string>;
    /**
     * An environment slug describing the type of this app.
     */
    environmentSlug?: pulumi.Input<string>;
    /**
     * Describes an environment variable made available to an app competent.
     */
    envs?: pulumi.Input<pulumi.Input<inputs.AppSpecStaticSiteEnv>[]>;
    /**
     * The name of the error document to use when serving this static site*
     */
    errorDocument?: pulumi.Input<string>;
    /**
     * A Git repo to use as component's source. Only one of `git` and `github` may be set.
     */
    git?: pulumi.Input<inputs.AppSpecStaticSiteGit>;
    /**
     * A GitHub repo to use as component's source. Only one of `git` and `github` may be set.
     */
    github?: pulumi.Input<inputs.AppSpecStaticSiteGithub>;
    /**
     * The name of the index document to use when serving this static site.
     */
    indexDocument?: pulumi.Input<string>;
    /**
     * The name of the component
     */
    name: pulumi.Input<string>;
    /**
     * An optional path to where the built assets will be located, relative to the build context. If not set, App Platform will automatically scan for these directory names: `_static`, `dist`, `public`.
     */
    outputDir?: pulumi.Input<string>;
    routes?: pulumi.Input<inputs.AppSpecStaticSiteRoutes>;
    /**
     * An optional path to the working directory to use for the build.
     */
    sourceDir?: pulumi.Input<string>;
}

export interface AppSpecStaticSiteEnv {
    /**
     * The name of the environment variable.
     */
    key?: pulumi.Input<string>;
    /**
     * The visibility scope of the environment variable. One of `RUN_TIME`, `BUILD_TIME`, or `RUN_AND_BUILD_TIME` (default).
     */
    scope?: pulumi.Input<string>;
    /**
     * The type of the environment variable, `GENERAL` or `SECRET`.
     */
    type?: pulumi.Input<string>;
    /**
     * The value of the environment variable.
     */
    value?: pulumi.Input<string>;
}

export interface AppSpecStaticSiteGit {
    /**
     * The name of the branch to use.
     */
    branch?: pulumi.Input<string>;
    /**
     * The clone URL of the repo.
     */
    repoCloneUrl?: pulumi.Input<string>;
}

export interface AppSpecStaticSiteGithub {
    /**
     * The name of the branch to use.
     */
    branch?: pulumi.Input<string>;
    /**
     * Whether to automatically deploy new commits made to the repo.
     */
    deployOnPush?: pulumi.Input<boolean>;
    /**
     * The name of the repo in the format `owner/repo`.
     */
    repo?: pulumi.Input<string>;
}

export interface AppSpecStaticSiteRoutes {
    /**
     * Paths must start with `/` and must be unique within the app.
     */
    path?: pulumi.Input<string>;
}

export interface AppSpecWorker {
    /**
     * An optional build command to run while building this component from source.
     */
    buildCommand?: pulumi.Input<string>;
    /**
     * The path to a Dockerfile relative to the root of the repo. If set, overrides usage of buildpacks.
     */
    dockerfilePath?: pulumi.Input<string>;
    /**
     * An environment slug describing the type of this app.
     */
    environmentSlug?: pulumi.Input<string>;
    /**
     * Describes an environment variable made available to an app competent.
     */
    envs?: pulumi.Input<pulumi.Input<inputs.AppSpecWorkerEnv>[]>;
    /**
     * A Git repo to use as component's source. Only one of `git` and `github` may be set.
     */
    git?: pulumi.Input<inputs.AppSpecWorkerGit>;
    /**
     * A GitHub repo to use as component's source. Only one of `git` and `github` may be set.
     */
    github?: pulumi.Input<inputs.AppSpecWorkerGithub>;
    /**
     * The amount of instances that this component should be scaled to.
     */
    instanceCount?: pulumi.Input<number>;
    /**
     * The instance size to use for this component.
     */
    instanceSizeSlug?: pulumi.Input<string>;
    /**
     * The name of the component
     */
    name: pulumi.Input<string>;
    routes?: pulumi.Input<inputs.AppSpecWorkerRoutes>;
    /**
     * An optional run command to override the component's default.
     */
    runCommand?: pulumi.Input<string>;
    /**
     * An optional path to the working directory to use for the build.
     */
    sourceDir?: pulumi.Input<string>;
}

export interface AppSpecWorkerEnv {
    /**
     * The name of the environment variable.
     */
    key?: pulumi.Input<string>;
    /**
     * The visibility scope of the environment variable. One of `RUN_TIME`, `BUILD_TIME`, or `RUN_AND_BUILD_TIME` (default).
     */
    scope?: pulumi.Input<string>;
    /**
     * The type of the environment variable, `GENERAL` or `SECRET`.
     */
    type?: pulumi.Input<string>;
    /**
     * The value of the environment variable.
     */
    value?: pulumi.Input<string>;
}

export interface AppSpecWorkerGit {
    /**
     * The name of the branch to use.
     */
    branch?: pulumi.Input<string>;
    /**
     * The clone URL of the repo.
     */
    repoCloneUrl?: pulumi.Input<string>;
}

export interface AppSpecWorkerGithub {
    /**
     * The name of the branch to use.
     */
    branch?: pulumi.Input<string>;
    /**
     * Whether to automatically deploy new commits made to the repo.
     */
    deployOnPush?: pulumi.Input<boolean>;
    /**
     * The name of the repo in the format `owner/repo`.
     */
    repo?: pulumi.Input<string>;
}

export interface AppSpecWorkerRoutes {
    /**
     * Paths must start with `/` and must be unique within the app.
     */
    path?: pulumi.Input<string>;
}

export interface DatabaseClusterMaintenanceWindow {
    /**
     * The day of the week on which to apply maintenance updates.
     */
    day: pulumi.Input<string>;
    /**
     * The hour in UTC at which maintenance updates will be applied in 24 hour format.
     */
    hour: pulumi.Input<string>;
}

export interface DatabaseFirewallRule {
    /**
     * The date and time when the firewall rule was created.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The type of resource that the firewall rule allows to access the database cluster. The possible values are: `droplet`, `k8s`, `ipAddr`, or `tag`.
     */
    type: pulumi.Input<string>;
    /**
     * A unique identifier for the firewall rule.
     */
    uuid?: pulumi.Input<string>;
    /**
     * The ID of the specific resource, the name of a tag applied to a group of resources, or the IP address that the firewall rule allows to access the database cluster.
     */
    value: pulumi.Input<string>;
}

export interface FirewallInboundRule {
    /**
     * The ports on which traffic will be allowed
     * specified as a string containing a single port, a range (e.g. "8000-9000"),
     * or "1-65535" to open all ports for a protocol. Required for when protocol is
     * `tcp` or `udp`.
     */
    portRange?: pulumi.Input<string>;
    /**
     * The type of traffic to be allowed.
     * This may be one of "tcp", "udp", or "icmp".
     */
    protocol: pulumi.Input<string>;
    /**
     * An array of strings containing the IPv4
     * addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs from which the
     * inbound traffic will be accepted.
     */
    sourceAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An array containing the IDs of
     * the Droplets from which the inbound traffic will be accepted.
     */
    sourceDropletIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * An array containing the IDs
     * of the Load Balancers from which the inbound traffic will be accepted.
     */
    sourceLoadBalancerUids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An array containing the names of Tags
     * corresponding to groups of Droplets from which the inbound traffic
     * will be accepted.
     */
    sourceTags?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface FirewallOutboundRule {
    /**
     * An array of strings containing the IPv4
     * addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs to which the
     * outbound traffic will be allowed.
     */
    destinationAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An array containing the IDs of
     * the Droplets to which the outbound traffic will be allowed.
     */
    destinationDropletIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * An array containing the IDs
     * of the Load Balancers to which the outbound traffic will be allowed.
     */
    destinationLoadBalancerUids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An array containing the names of Tags
     * corresponding to groups of Droplets to which the outbound traffic will
     * be allowed.
     * traffic.
     */
    destinationTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ports on which traffic will be allowed
     * specified as a string containing a single port, a range (e.g. "8000-9000"),
     * or "1-65535" to open all ports for a protocol. Required for when protocol is
     * `tcp` or `udp`.
     */
    portRange?: pulumi.Input<string>;
    /**
     * The type of traffic to be allowed.
     * This may be one of "tcp", "udp", or "icmp".
     */
    protocol: pulumi.Input<string>;
}

export interface FirewallPendingChange {
    dropletId?: pulumi.Input<number>;
    removing?: pulumi.Input<boolean>;
    /**
     * A status string indicating the current state of the Firewall.
     * This can be "waiting", "succeeded", or "failed".
     */
    status?: pulumi.Input<string>;
}

export interface GetDomainsFilter {
    /**
     * Set to `true` to require that a field match all of the `values` instead of just one or more of
     * them. This is useful when matching against multi-valued fields such as lists or sets where you want to ensure
     * that all of the `values` are present in the list or set.
     */
    all?: boolean;
    /**
     * Filter the domains by this key. This may be one of `name`, `urn`, and `ttl`.
     */
    key: string;
    /**
     * One of `exact` (default), `re`, or `substring`. For string-typed fields, specify `re` to
     * match by using the `values` as regular expressions, or specify `substring` to match by treating the `values` as
     * substrings to find within the string field.
     */
    matchBy?: string;
    /**
     * A list of values to match against the `key` field. Only retrieves domains
     * where the `key` field takes on one or more of the values provided here.
     */
    values: string[];
}

export interface GetDomainsSort {
    /**
     * The sort direction. This may be either `asc` or `desc`.
     */
    direction?: string;
    /**
     * Sort the domains by this key. This may be one of `name`, `urn`, and `ttl`.
     */
    key: string;
}

export interface GetDropletsFilter {
    /**
     * Set to `true` to require that a field match all of the `values` instead of just one or more of
     * them. This is useful when matching against multi-valued fields such as lists or sets where you want to ensure
     * that all of the `values` are present in the list or set.
     */
    all?: boolean;
    /**
     * Filter the Droplets by this key. This may be one of `backups`, `createdAt`, `disk`, `id`,
     * `image`, `ipv4Address`, `ipv4AddressPrivate`, `ipv6`, `ipv6Address`, `ipv6AddressPrivate`, `locked`,
     * `memory`, `monitoring`, `name`, `priceHourly`, `priceMonthly`, `privateNetworking`, `region`, `size`,
     * `status`, `tags`, `urn`, `vcpus`, `volumeIds`, or `vpcUuid`.
     */
    key: string;
    /**
     * One of `exact` (default), `re`, or `substring`. For string-typed fields, specify `re` to
     * match by using the `values` as regular expressions, or specify `substring` to match by treating the `values` as
     * substrings to find within the string field.
     */
    matchBy?: string;
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
     * Set to `true` to require that a field match all of the `values` instead of just one or more of
     * them. This is useful when matching against multi-valued fields such as lists or sets where you want to ensure
     * that all of the `values` are present in the list or set.
     */
    all?: boolean;
    /**
     * Filter the images by this key. This may be one of `distribution`, `errorMessage`,
     * `id`, `image`, `minDiskSize`, `name`, `private`, `regions`, `sizeGigabytes`, `slug`, `status`,
     * `tags`, or `type`.
     */
    key: string;
    /**
     * One of `exact` (default), `re`, or `substring`. For string-typed fields, specify `re` to
     * match by using the `values` as regular expressions, or specify `substring` to match by treating the `values` as
     * substrings to find within the string field.
     */
    matchBy?: string;
    /**
     * A list of values to match against the `key` field. Only retrieves images
     * where the `key` field takes on one or more of the values provided here.
     */
    values: string[];
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

export interface GetProjectsFilter {
    /**
     * Set to `true` to require that a field match all of the `values` instead of just one or more of
     * them. This is useful when matching against multi-valued fields such as lists or sets where you want to ensure
     * that all of the `values` are present in the list or set.
     */
    all?: boolean;
    /**
     * Filter the projects by this key. This may be one of `name`,
     * `purpose`, `description`, `environment`, or `isDefault`.
     */
    key: string;
    /**
     * One of `exact` (default), `re`, or `substring`. For string-typed fields, specify `re` to
     * match by using the `values` as regular expressions, or specify `substring` to match by treating the `values` as
     * substrings to find within the string field.
     */
    matchBy?: string;
    /**
     * A list of values to match against the `key` field. Only retrieves projects
     * where the `key` field takes on one or more of the values provided here.
     */
    values: string[];
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
     * Set to `true` to require that a field match all of the `values` instead of just one or more of
     * them. This is useful when matching against multi-valued fields such as lists or sets where you want to ensure
     * that all of the `values` are present in the list or set.
     */
    all?: boolean;
    /**
     * Filter the regions by this key. This may be one of `slug`,
     * `name`, `available`, `features`, or `sizes`.
     */
    key: string;
    /**
     * One of `exact` (default), `re`, or `substring`. For string-typed fields, specify `re` to
     * match by using the `values` as regular expressions, or specify `substring` to match by treating the `values` as
     * substrings to find within the string field.
     */
    matchBy?: string;
    /**
     * A list of values to match against the `key` field. Only retrieves regions
     * where the `key` field takes on one or more of the values provided here.
     */
    values: string[];
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
     * Set to `true` to require that a field match all of the `values` instead of just one or more of
     * them. This is useful when matching against multi-valued fields such as lists or sets where you want to ensure
     * that all of the `values` are present in the list or set.
     */
    all?: boolean;
    /**
     * Filter the sizes by this key. This may be one of `slug`,
     * `regions`, `memory`, `vcpus`, `disk`, `transfer`, `priceMonthly`,
     * `priceHourly`, or `available`.
     */
    key: string;
    /**
     * One of `exact` (default), `re`, or `substring`. For string-typed fields, specify `re` to
     * match by using the `values` as regular expressions, or specify `substring` to match by treating the `values` as
     * substrings to find within the string field.
     */
    matchBy?: string;
    /**
     * Only retrieves sizes which keys has value that matches
     * one of the values provided here.
     */
    values: string[];
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

export interface GetSpacesBucketsFilter {
    /**
     * Set to `true` to require that a field match all of the `values` instead of just one or more of
     * them. This is useful when matching against multi-valued fields such as lists or sets where you want to ensure
     * that all of the `values` are present in the list or set.
     */
    all?: boolean;
    /**
     * Filter the images by this key. This may be one of `bucketDomainName`, `name`, `region`, or `urn`.
     */
    key: string;
    /**
     * One of `exact` (default), `re`, or `substring`. For string-typed fields, specify `re` to
     * match by using the `values` as regular expressions, or specify `substring` to match by treating the `values` as
     * substrings to find within the string field.
     */
    matchBy?: string;
    /**
     * A list of values to match against the `key` field. Only retrieves Spaces buckets
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
     * Set to `true` to require that a field match all of the `values` instead of just one or more of
     * them. This is useful when matching against multi-valued fields such as lists or sets where you want to ensure
     * that all of the `values` are present in the list or set.
     */
    all?: boolean;
    /**
     * Filter the tags by this key. This may be one of `name`, `totalResourceCount`,  `dropletsCount`, `imagesCount`, `volumesCount`, `volumeSnapshotsCount`, or `databasesCount`.
     */
    key: string;
    /**
     * One of `exact` (default), `re`, or `substring`. For string-typed fields, specify `re` to
     * match by using the `values` as regular expressions, or specify `substring` to match by treating the `values` as
     * substrings to find within the string field.
     */
    matchBy?: string;
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

export interface KubernetesClusterKubeConfig {
    /**
     * The base64 encoded public certificate used by clients to access the cluster. Only available if token authentication is not supported on your cluster.
     */
    clientCertificate?: pulumi.Input<string>;
    /**
     * The base64 encoded private key used by clients to access the cluster. Only available if token authentication is not supported on your cluster.
     */
    clientKey?: pulumi.Input<string>;
    /**
     * The base64 encoded public certificate for the cluster's certificate authority.
     */
    clusterCaCertificate?: pulumi.Input<string>;
    /**
     * The date and time when the credentials will expire and need to be regenerated.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * The URL of the API server on the Kubernetes master node.
     */
    host?: pulumi.Input<string>;
    /**
     * The full contents of the Kubernetes cluster's kubeconfig file.
     */
    rawConfig?: pulumi.Input<string>;
    /**
     * The DigitalOcean API access token used by clients to access the cluster.
     */
    token?: pulumi.Input<string>;
}

export interface KubernetesClusterNodePool {
    /**
     * A computed field representing the actual number of nodes in the node pool, which is especially useful when auto-scaling is enabled.
     */
    actualNodeCount?: pulumi.Input<number>;
    /**
     * Enable auto-scaling of the number of nodes in the node pool within the given min/max range.
     */
    autoScale?: pulumi.Input<boolean>;
    /**
     * A unique ID that can be used to identify and reference the node.
     */
    id?: pulumi.Input<string>;
    /**
     * A map of key/value pairs to apply to nodes in the pool. The labels are exposed in the Kubernetes API as labels in the metadata of the corresponding [Node resources](https://kubernetes.io/docs/concepts/architecture/nodes/).
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to.
     */
    maxNodes?: pulumi.Input<number>;
    /**
     * If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to.
     */
    minNodes?: pulumi.Input<number>;
    /**
     * A name for the node pool.
     */
    name: pulumi.Input<string>;
    /**
     * The number of Droplet instances in the node pool. If auto-scaling is enabled, this should only be set if the desired result is to explicitly reset the number of nodes to this value. If auto-scaling is enabled, and the node count is outside of the given min/max range, it will use the min nodes value.
     */
    nodeCount?: pulumi.Input<number>;
    /**
     * A list of nodes in the pool. Each node exports the following attributes:
     */
    nodes?: pulumi.Input<pulumi.Input<inputs.KubernetesClusterNodePoolNode>[]>;
    /**
     * The slug identifier for the type of Droplet to be used as workers in the node pool.
     */
    size: pulumi.Input<string>;
    /**
     * A list of tag names to be applied to the Kubernetes cluster.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface KubernetesClusterNodePoolNode {
    /**
     * The date and time when the node was created.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The id of the node's droplet
     */
    dropletId?: pulumi.Input<string>;
    /**
     * A unique ID that can be used to identify and reference the node.
     */
    id?: pulumi.Input<string>;
    /**
     * A name for the node pool.
     */
    name?: pulumi.Input<string>;
    /**
     * A string indicating the current status of the individual node.
     */
    status?: pulumi.Input<string>;
    /**
     * The date and time when the node was last updated.
     */
    updatedAt?: pulumi.Input<string>;
}

export interface KubernetesNodePoolNode {
    /**
     * The date and time when the node was created.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The id of the node's droplet
     */
    dropletId?: pulumi.Input<string>;
    /**
     * A unique ID that can be used to identify and reference the node.
     */
    id?: pulumi.Input<string>;
    /**
     * A name for the node pool.
     */
    name?: pulumi.Input<string>;
    /**
     * A string indicating the current status of the individual node.
     */
    status?: pulumi.Input<string>;
    /**
     * The date and time when the node was last updated.
     */
    updatedAt?: pulumi.Input<string>;
}

export interface LoadBalancerForwardingRule {
    /**
     * The ID of the TLS certificate to be used for SSL termination.
     */
    certificateId?: pulumi.Input<string>;
    /**
     * An integer representing the port on which the Load Balancer instance will listen.
     */
    entryPort: pulumi.Input<number>;
    /**
     * The protocol used for traffic to the Load Balancer. The possible values are: `http`, `https`, `http2` or `tcp`.
     */
    entryProtocol: pulumi.Input<string>;
    /**
     * An integer representing the port on the backend Droplets to which the Load Balancer will send traffic.
     */
    targetPort: pulumi.Input<number>;
    /**
     * The protocol used for traffic from the Load Balancer to the backend Droplets. The possible values are: `http`, `https`, `http2` or `tcp`.
     */
    targetProtocol: pulumi.Input<string>;
    /**
     * A boolean value indicating whether SSL encrypted traffic will be passed through to the backend Droplets. The default value is `false`.
     */
    tlsPassthrough?: pulumi.Input<boolean>;
}

export interface LoadBalancerHealthcheck {
    /**
     * The number of seconds between between two consecutive health checks. If not specified, the default value is `10`.
     */
    checkIntervalSeconds?: pulumi.Input<number>;
    /**
     * The number of times a health check must pass for a backend Droplet to be marked "healthy" and be re-added to the pool. If not specified, the default value is `5`.
     */
    healthyThreshold?: pulumi.Input<number>;
    /**
     * The path on the backend Droplets to which the Load Balancer instance will send a request.
     */
    path?: pulumi.Input<string>;
    /**
     * An integer representing the port on the backend Droplets on which the health check will attempt a connection.
     */
    port: pulumi.Input<number>;
    /**
     * The protocol used for health checks sent to the backend Droplets. The possible values are `http`, `https` or `tcp`.
     */
    protocol: pulumi.Input<string>;
    /**
     * The number of seconds the Load Balancer instance will wait for a response until marking a health check as failed. If not specified, the default value is `5`.
     */
    responseTimeoutSeconds?: pulumi.Input<number>;
    /**
     * The number of times a health check must fail for a backend Droplet to be marked "unhealthy" and be removed from the pool. If not specified, the default value is `3`.
     */
    unhealthyThreshold?: pulumi.Input<number>;
}

export interface LoadBalancerStickySessions {
    /**
     * The name to be used for the cookie sent to the client. This attribute is required when using `cookies` for the sticky sessions type.
     */
    cookieName?: pulumi.Input<string>;
    /**
     * The number of seconds until the cookie set by the Load Balancer expires. This attribute is required when using `cookies` for the sticky sessions type.
     */
    cookieTtlSeconds?: pulumi.Input<number>;
    /**
     * An attribute indicating how and if requests from a client will be persistently served by the same backend Droplet. The possible values are `cookies` or `none`. If not specified, the default value is `none`.
     */
    type?: pulumi.Input<string>;
}

export interface SpacesBucketCorsRule {
    /**
     * A list of headers that will be included in the CORS preflight request's `Access-Control-Request-Headers`. A header may contain one wildcard (e.g. `x-amz-*`).
     */
    allowedHeaders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of HTTP methods (e.g. `GET`) which are allowed from the specified origin.
     */
    allowedMethods: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of hosts from which requests using the specified methods are allowed. A host may contain one wildcard (e.g. http://*.example.com).
     */
    allowedOrigins: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The time in seconds that browser can cache the response for a preflight request.
     */
    maxAgeSeconds?: pulumi.Input<number>;
}

export interface SpacesBucketLifecycleRule {
    /**
     * Specifies the number of days after initiating a multipart
     * upload when the multipart upload must be completed or else Spaces will abort the upload.
     */
    abortIncompleteMultipartUploadDays?: pulumi.Input<number>;
    /**
     * Specifies lifecycle rule status.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Specifies a time period after which applicable objects expire (documented below).
     */
    expiration?: pulumi.Input<inputs.SpacesBucketLifecycleRuleExpiration>;
    /**
     * Unique identifier for the rule.
     */
    id?: pulumi.Input<string>;
    /**
     * Specifies when non-current object versions expire (documented below).
     */
    noncurrentVersionExpiration?: pulumi.Input<inputs.SpacesBucketLifecycleRuleNoncurrentVersionExpiration>;
    /**
     * Object key prefix identifying one or more objects to which the rule applies.
     */
    prefix?: pulumi.Input<string>;
}

export interface SpacesBucketLifecycleRuleExpiration {
    /**
     * Specifies the date/time after which you want applicable objects to expire. The argument uses
     * RFC3339 format, e.g. "2020-03-22T15:03:55Z" or parts thereof e.g. "2019-02-28".
     */
    date?: pulumi.Input<string>;
    /**
     * Specifies the number of days after object creation when the applicable objects will expire.
     */
    days?: pulumi.Input<number>;
    /**
     * On a versioned bucket (versioning-enabled or versioning-suspended
     * bucket), setting this to true directs Spaces to delete expired object delete markers.
     */
    expiredObjectDeleteMarker?: pulumi.Input<boolean>;
}

export interface SpacesBucketLifecycleRuleNoncurrentVersionExpiration {
    /**
     * Specifies the number of days after which an object's non-current versions expire.
     */
    days?: pulumi.Input<number>;
}

export interface SpacesBucketVersioning {
    /**
     * Enable versioning. Once you version-enable a bucket, it can never return to an unversioned
     * state. You can, however, suspend versioning on that bucket.
     */
    enabled?: pulumi.Input<boolean>;
}
