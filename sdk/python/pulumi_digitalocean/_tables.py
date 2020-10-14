# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

SNAKE_TO_CAMEL_CASE_TABLE = {
    "active_deployment_id": "activeDeploymentId",
    "actual_node_count": "actualNodeCount",
    "api_endpoint": "apiEndpoint",
    "auto_scale": "autoScale",
    "auto_upgrade": "autoUpgrade",
    "bucket_domain_name": "bucketDomainName",
    "bucket_urn": "bucketUrn",
    "cache_control": "cacheControl",
    "certificate_chain": "certificateChain",
    "certificate_id": "certificateId",
    "cluster_id": "clusterId",
    "cluster_subnet": "clusterSubnet",
    "cluster_urn": "clusterUrn",
    "content_base64": "contentBase64",
    "content_disposition": "contentDisposition",
    "content_encoding": "contentEncoding",
    "content_language": "contentLanguage",
    "content_type": "contentType",
    "cors_rules": "corsRules",
    "created_at": "createdAt",
    "credential_expiration_time": "credentialExpirationTime",
    "custom_domain": "customDomain",
    "databases_count": "databasesCount",
    "db_name": "dbName",
    "default_ingress": "defaultIngress",
    "docker_credentials": "dockerCredentials",
    "domain_urn": "domainUrn",
    "droplet_id": "dropletId",
    "droplet_ids": "dropletIds",
    "droplet_tag": "dropletTag",
    "droplet_urn": "dropletUrn",
    "droplets_count": "dropletsCount",
    "enable_backend_keepalive": "enableBackendKeepalive",
    "enable_proxy_protocol": "enableProxyProtocol",
    "eviction_policy": "evictionPolicy",
    "expiry_seconds": "expirySeconds",
    "filesystem_label": "filesystemLabel",
    "filesystem_type": "filesystemType",
    "floating_ip_urn": "floatingIpUrn",
    "force_destroy": "forceDestroy",
    "forwarding_rules": "forwardingRules",
    "images_count": "imagesCount",
    "inbound_rules": "inboundRules",
    "initial_filesystem_label": "initialFilesystemLabel",
    "initial_filesystem_type": "initialFilesystemType",
    "ip_address": "ipAddress",
    "ip_range": "ipRange",
    "ipv4_address": "ipv4Address",
    "ipv4_address_private": "ipv4AddressPrivate",
    "ipv6_address": "ipv6Address",
    "is_default": "isDefault",
    "kube_configs": "kubeConfigs",
    "leaf_certificate": "leafCertificate",
    "lifecycle_rules": "lifecycleRules",
    "live_url": "liveUrl",
    "load_balancer_urn": "loadBalancerUrn",
    "maintenance_windows": "maintenanceWindows",
    "max_nodes": "maxNodes",
    "min_disk_size": "minDiskSize",
    "min_nodes": "minNodes",
    "mysql_auth_plugin": "mysqlAuthPlugin",
    "node_count": "nodeCount",
    "node_pool": "nodePool",
    "not_after": "notAfter",
    "outbound_rules": "outboundRules",
    "owner_id": "ownerId",
    "owner_uuid": "ownerUuid",
    "pending_changes": "pendingChanges",
    "price_hourly": "priceHourly",
    "price_monthly": "priceMonthly",
    "private_host": "privateHost",
    "private_key": "privateKey",
    "private_network_uuid": "privateNetworkUuid",
    "private_networking": "privateNetworking",
    "private_uri": "privateUri",
    "public_key": "publicKey",
    "redirect_http_to_https": "redirectHttpToHttps",
    "registry_name": "registryName",
    "resize_disk": "resizeDisk",
    "server_url": "serverUrl",
    "service_subnet": "serviceSubnet",
    "sha1_fingerprint": "sha1Fingerprint",
    "snapshot_id": "snapshotId",
    "spaces_access_id": "spacesAccessId",
    "spaces_endpoint": "spacesEndpoint",
    "spaces_secret_key": "spacesSecretKey",
    "sql_mode": "sqlMode",
    "ssh_keys": "sshKeys",
    "sticky_sessions": "stickySessions",
    "surge_upgrade": "surgeUpgrade",
    "total_resource_count": "totalResourceCount",
    "updated_at": "updatedAt",
    "user_data": "userData",
    "version_id": "versionId",
    "volume_id": "volumeId",
    "volume_ids": "volumeIds",
    "volume_snapshots_count": "volumeSnapshotsCount",
    "volume_urn": "volumeUrn",
    "volumes_count": "volumesCount",
    "vpc_urn": "vpcUrn",
    "vpc_uuid": "vpcUuid",
    "website_redirect": "websiteRedirect",
}

CAMEL_TO_SNAKE_CASE_TABLE = {
    "activeDeploymentId": "active_deployment_id",
    "actualNodeCount": "actual_node_count",
    "apiEndpoint": "api_endpoint",
    "autoScale": "auto_scale",
    "autoUpgrade": "auto_upgrade",
    "bucketDomainName": "bucket_domain_name",
    "bucketUrn": "bucket_urn",
    "cacheControl": "cache_control",
    "certificateChain": "certificate_chain",
    "certificateId": "certificate_id",
    "clusterId": "cluster_id",
    "clusterSubnet": "cluster_subnet",
    "clusterUrn": "cluster_urn",
    "contentBase64": "content_base64",
    "contentDisposition": "content_disposition",
    "contentEncoding": "content_encoding",
    "contentLanguage": "content_language",
    "contentType": "content_type",
    "corsRules": "cors_rules",
    "createdAt": "created_at",
    "credentialExpirationTime": "credential_expiration_time",
    "customDomain": "custom_domain",
    "databasesCount": "databases_count",
    "dbName": "db_name",
    "defaultIngress": "default_ingress",
    "dockerCredentials": "docker_credentials",
    "domainUrn": "domain_urn",
    "dropletId": "droplet_id",
    "dropletIds": "droplet_ids",
    "dropletTag": "droplet_tag",
    "dropletUrn": "droplet_urn",
    "dropletsCount": "droplets_count",
    "enableBackendKeepalive": "enable_backend_keepalive",
    "enableProxyProtocol": "enable_proxy_protocol",
    "evictionPolicy": "eviction_policy",
    "expirySeconds": "expiry_seconds",
    "filesystemLabel": "filesystem_label",
    "filesystemType": "filesystem_type",
    "floatingIpUrn": "floating_ip_urn",
    "forceDestroy": "force_destroy",
    "forwardingRules": "forwarding_rules",
    "imagesCount": "images_count",
    "inboundRules": "inbound_rules",
    "initialFilesystemLabel": "initial_filesystem_label",
    "initialFilesystemType": "initial_filesystem_type",
    "ipAddress": "ip_address",
    "ipRange": "ip_range",
    "ipv4Address": "ipv4_address",
    "ipv4AddressPrivate": "ipv4_address_private",
    "ipv6Address": "ipv6_address",
    "isDefault": "is_default",
    "kubeConfigs": "kube_configs",
    "leafCertificate": "leaf_certificate",
    "lifecycleRules": "lifecycle_rules",
    "liveUrl": "live_url",
    "loadBalancerUrn": "load_balancer_urn",
    "maintenanceWindows": "maintenance_windows",
    "maxNodes": "max_nodes",
    "minDiskSize": "min_disk_size",
    "minNodes": "min_nodes",
    "mysqlAuthPlugin": "mysql_auth_plugin",
    "nodeCount": "node_count",
    "nodePool": "node_pool",
    "notAfter": "not_after",
    "outboundRules": "outbound_rules",
    "ownerId": "owner_id",
    "ownerUuid": "owner_uuid",
    "pendingChanges": "pending_changes",
    "priceHourly": "price_hourly",
    "priceMonthly": "price_monthly",
    "privateHost": "private_host",
    "privateKey": "private_key",
    "privateNetworkUuid": "private_network_uuid",
    "privateNetworking": "private_networking",
    "privateUri": "private_uri",
    "publicKey": "public_key",
    "redirectHttpToHttps": "redirect_http_to_https",
    "registryName": "registry_name",
    "resizeDisk": "resize_disk",
    "serverUrl": "server_url",
    "serviceSubnet": "service_subnet",
    "sha1Fingerprint": "sha1_fingerprint",
    "snapshotId": "snapshot_id",
    "spacesAccessId": "spaces_access_id",
    "spacesEndpoint": "spaces_endpoint",
    "spacesSecretKey": "spaces_secret_key",
    "sqlMode": "sql_mode",
    "sshKeys": "ssh_keys",
    "stickySessions": "sticky_sessions",
    "surgeUpgrade": "surge_upgrade",
    "totalResourceCount": "total_resource_count",
    "updatedAt": "updated_at",
    "userData": "user_data",
    "versionId": "version_id",
    "volumeId": "volume_id",
    "volumeIds": "volume_ids",
    "volumeSnapshotsCount": "volume_snapshots_count",
    "volumeUrn": "volume_urn",
    "volumesCount": "volumes_count",
    "vpcUrn": "vpc_urn",
    "vpcUuid": "vpc_uuid",
    "websiteRedirect": "website_redirect",
}
