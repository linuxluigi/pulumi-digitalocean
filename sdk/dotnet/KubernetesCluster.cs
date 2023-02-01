// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean
{
    /// <summary>
    /// ## Import
    /// 
    /// Before importing a Kubernetes cluster, the cluster's default node pool must be tagged with the `terraform:default-node-pool` tag. The provider will automatically add this tag if the cluster has a single node pool. Clusters with more than one node pool, however, will require that you manually add the `terraform:default-node-pool` tag to the node pool that you intend to be the default node pool. Then the Kubernetes cluster and all of its node pools can be imported using the cluster's `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import digitalocean:index/kubernetesCluster:KubernetesCluster mycluster 1b8b2100-0e9f-4e8f-ad78-9eb578c2a0af
    /// ```
    /// </summary>
    [DigitalOceanResourceType("digitalocean:index/kubernetesCluster:KubernetesCluster")]
    public partial class KubernetesCluster : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A boolean value indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window.
        /// </summary>
        [Output("autoUpgrade")]
        public Output<bool?> AutoUpgrade { get; private set; } = null!;

        /// <summary>
        /// The range of IP addresses in the overlay network of the Kubernetes cluster.
        /// </summary>
        [Output("clusterSubnet")]
        public Output<string> ClusterSubnet { get; private set; } = null!;

        /// <summary>
        /// The uniform resource name (URN) for the Kubernetes cluster.
        /// </summary>
        [Output("clusterUrn")]
        public Output<string> ClusterUrn { get; private set; } = null!;

        /// <summary>
        /// The date and time when the node was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The base URL of the API server on the Kubernetes master node.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// Enable/disable the high availability control plane for a cluster. High availability can only be set when creating a cluster. Any update will create a new cluster. Default: false
        /// </summary>
        [Output("ha")]
        public Output<bool?> Ha { get; private set; } = null!;

        /// <summary>
        /// The public IPv4 address of the Kubernetes master node. This will not be set if high availability is configured on the cluster (v1.21+)
        /// </summary>
        [Output("ipv4Address")]
        public Output<string> Ipv4Address { get; private set; } = null!;

        [Output("kubeConfigs")]
        public Output<ImmutableArray<Outputs.KubernetesClusterKubeConfig>> KubeConfigs { get; private set; } = null!;

        /// <summary>
        /// A block representing the cluster's maintenance window. Updates will be applied within this window. If not specified, a default maintenance window will be chosen. `auto_upgrade` must be set to `true` for this to have an effect.
        /// </summary>
        [Output("maintenancePolicy")]
        public Output<Outputs.KubernetesClusterMaintenancePolicy> MaintenancePolicy { get; private set; } = null!;

        /// <summary>
        /// A name for the node pool.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A block representing the cluster's default node pool. Additional node pools may be added to the cluster using the `digitalocean.KubernetesNodePool` resource. The following arguments may be specified:
        /// </summary>
        [Output("nodePool")]
        public Output<Outputs.KubernetesClusterNodePool> NodePool { get; private set; } = null!;

        /// <summary>
        /// The slug identifier for the region where the Kubernetes cluster will be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The range of assignable IP addresses for services running in the Kubernetes cluster.
        /// </summary>
        [Output("serviceSubnet")]
        public Output<string> ServiceSubnet { get; private set; } = null!;

        /// <summary>
        /// A string indicating the current status of the individual node.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Enable/disable surge upgrades for a cluster. Default: false
        /// </summary>
        [Output("surgeUpgrade")]
        public Output<bool?> SurgeUpgrade { get; private set; } = null!;

        /// <summary>
        /// A list of tag names to be applied to the Kubernetes cluster.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The date and time when the node was last updated.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The slug identifier for the version of Kubernetes used for the cluster. Use [doctl](https://github.com/digitalocean/doctl) to find the available versions `doctl kubernetes options versions`. (**Note:** A cluster may only be upgraded to newer versions in-place. If the version is decreased, a new resource will be created.)
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPC where the Kubernetes cluster will be located.
        /// </summary>
        [Output("vpcUuid")]
        public Output<string> VpcUuid { get; private set; } = null!;


        /// <summary>
        /// Create a KubernetesCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KubernetesCluster(string name, KubernetesClusterArgs args, CustomResourceOptions? options = null)
            : base("digitalocean:index/kubernetesCluster:KubernetesCluster", name, args ?? new KubernetesClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KubernetesCluster(string name, Input<string> id, KubernetesClusterState? state = null, CustomResourceOptions? options = null)
            : base("digitalocean:index/kubernetesCluster:KubernetesCluster", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "kubeConfigs",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing KubernetesCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KubernetesCluster Get(string name, Input<string> id, KubernetesClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new KubernetesCluster(name, id, state, options);
        }
    }

    public sealed class KubernetesClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A boolean value indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window.
        /// </summary>
        [Input("autoUpgrade")]
        public Input<bool>? AutoUpgrade { get; set; }

        /// <summary>
        /// Enable/disable the high availability control plane for a cluster. High availability can only be set when creating a cluster. Any update will create a new cluster. Default: false
        /// </summary>
        [Input("ha")]
        public Input<bool>? Ha { get; set; }

        /// <summary>
        /// A block representing the cluster's maintenance window. Updates will be applied within this window. If not specified, a default maintenance window will be chosen. `auto_upgrade` must be set to `true` for this to have an effect.
        /// </summary>
        [Input("maintenancePolicy")]
        public Input<Inputs.KubernetesClusterMaintenancePolicyArgs>? MaintenancePolicy { get; set; }

        /// <summary>
        /// A name for the node pool.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A block representing the cluster's default node pool. Additional node pools may be added to the cluster using the `digitalocean.KubernetesNodePool` resource. The following arguments may be specified:
        /// </summary>
        [Input("nodePool", required: true)]
        public Input<Inputs.KubernetesClusterNodePoolArgs> NodePool { get; set; } = null!;

        /// <summary>
        /// The slug identifier for the region where the Kubernetes cluster will be created.
        /// </summary>
        [Input("region", required: true)]
        public InputUnion<string, Pulumi.DigitalOcean.Region> Region { get; set; } = null!;

        /// <summary>
        /// Enable/disable surge upgrades for a cluster. Default: false
        /// </summary>
        [Input("surgeUpgrade")]
        public Input<bool>? SurgeUpgrade { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tag names to be applied to the Kubernetes cluster.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The slug identifier for the version of Kubernetes used for the cluster. Use [doctl](https://github.com/digitalocean/doctl) to find the available versions `doctl kubernetes options versions`. (**Note:** A cluster may only be upgraded to newer versions in-place. If the version is decreased, a new resource will be created.)
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        /// <summary>
        /// The ID of the VPC where the Kubernetes cluster will be located.
        /// </summary>
        [Input("vpcUuid")]
        public Input<string>? VpcUuid { get; set; }

        public KubernetesClusterArgs()
        {
        }
        public static new KubernetesClusterArgs Empty => new KubernetesClusterArgs();
    }

    public sealed class KubernetesClusterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A boolean value indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window.
        /// </summary>
        [Input("autoUpgrade")]
        public Input<bool>? AutoUpgrade { get; set; }

        /// <summary>
        /// The range of IP addresses in the overlay network of the Kubernetes cluster.
        /// </summary>
        [Input("clusterSubnet")]
        public Input<string>? ClusterSubnet { get; set; }

        /// <summary>
        /// The uniform resource name (URN) for the Kubernetes cluster.
        /// </summary>
        [Input("clusterUrn")]
        public Input<string>? ClusterUrn { get; set; }

        /// <summary>
        /// The date and time when the node was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The base URL of the API server on the Kubernetes master node.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// Enable/disable the high availability control plane for a cluster. High availability can only be set when creating a cluster. Any update will create a new cluster. Default: false
        /// </summary>
        [Input("ha")]
        public Input<bool>? Ha { get; set; }

        /// <summary>
        /// The public IPv4 address of the Kubernetes master node. This will not be set if high availability is configured on the cluster (v1.21+)
        /// </summary>
        [Input("ipv4Address")]
        public Input<string>? Ipv4Address { get; set; }

        [Input("kubeConfigs")]
        private InputList<Inputs.KubernetesClusterKubeConfigGetArgs>? _kubeConfigs;
        public InputList<Inputs.KubernetesClusterKubeConfigGetArgs> KubeConfigs
        {
            get => _kubeConfigs ?? (_kubeConfigs = new InputList<Inputs.KubernetesClusterKubeConfigGetArgs>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableArray.Create<Inputs.KubernetesClusterKubeConfigGetArgs>());
                _kubeConfigs = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// A block representing the cluster's maintenance window. Updates will be applied within this window. If not specified, a default maintenance window will be chosen. `auto_upgrade` must be set to `true` for this to have an effect.
        /// </summary>
        [Input("maintenancePolicy")]
        public Input<Inputs.KubernetesClusterMaintenancePolicyGetArgs>? MaintenancePolicy { get; set; }

        /// <summary>
        /// A name for the node pool.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A block representing the cluster's default node pool. Additional node pools may be added to the cluster using the `digitalocean.KubernetesNodePool` resource. The following arguments may be specified:
        /// </summary>
        [Input("nodePool")]
        public Input<Inputs.KubernetesClusterNodePoolGetArgs>? NodePool { get; set; }

        /// <summary>
        /// The slug identifier for the region where the Kubernetes cluster will be created.
        /// </summary>
        [Input("region")]
        public InputUnion<string, Pulumi.DigitalOcean.Region>? Region { get; set; }

        /// <summary>
        /// The range of assignable IP addresses for services running in the Kubernetes cluster.
        /// </summary>
        [Input("serviceSubnet")]
        public Input<string>? ServiceSubnet { get; set; }

        /// <summary>
        /// A string indicating the current status of the individual node.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Enable/disable surge upgrades for a cluster. Default: false
        /// </summary>
        [Input("surgeUpgrade")]
        public Input<bool>? SurgeUpgrade { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tag names to be applied to the Kubernetes cluster.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The date and time when the node was last updated.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// The slug identifier for the version of Kubernetes used for the cluster. Use [doctl](https://github.com/digitalocean/doctl) to find the available versions `doctl kubernetes options versions`. (**Note:** A cluster may only be upgraded to newer versions in-place. If the version is decreased, a new resource will be created.)
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The ID of the VPC where the Kubernetes cluster will be located.
        /// </summary>
        [Input("vpcUuid")]
        public Input<string>? VpcUuid { get; set; }

        public KubernetesClusterState()
        {
        }
        public static new KubernetesClusterState Empty => new KubernetesClusterState();
    }
}
