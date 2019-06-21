// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > **NOTE:** DigitalOcean Kubernetes is currently in [Limited Availability](https://www.digitalocean.com/docs/platform/product-lifecycle/). In order to access its API, you must first enable Kubernetes on your account by opting-in via the [cloud control panel](https://cloud.digitalocean.com/kubernetes/clusters). While the Kubernetes Cluster functionality is currently in limited availability the structure of this resource may change over time. Please share any feedback you may have by [opening an issue on GitHub](https://github.com/terraform-providers/terraform-provider-digitalocean/issues).
// 
// Provides a DigitalOcean Kubernetes node pool resource. While the default node pool must be defined in the `digitalocean_kubernetes_cluster` resource, this resource can be used to add additional ones to a cluster.
type KubernetesNodePool struct {
	s *pulumi.ResourceState
}

// NewKubernetesNodePool registers a new resource with the given unique name, arguments, and options.
func NewKubernetesNodePool(ctx *pulumi.Context,
	name string, args *KubernetesNodePoolArgs, opts ...pulumi.ResourceOpt) (*KubernetesNodePool, error) {
	if args == nil || args.ClusterId == nil {
		return nil, errors.New("missing required argument 'ClusterId'")
	}
	if args == nil || args.NodeCount == nil {
		return nil, errors.New("missing required argument 'NodeCount'")
	}
	if args == nil || args.Size == nil {
		return nil, errors.New("missing required argument 'Size'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["clusterId"] = nil
		inputs["name"] = nil
		inputs["nodeCount"] = nil
		inputs["size"] = nil
		inputs["tags"] = nil
	} else {
		inputs["clusterId"] = args.ClusterId
		inputs["name"] = args.Name
		inputs["nodeCount"] = args.NodeCount
		inputs["size"] = args.Size
		inputs["tags"] = args.Tags
	}
	inputs["nodes"] = nil
	s, err := ctx.RegisterResource("digitalocean:index/kubernetesNodePool:KubernetesNodePool", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &KubernetesNodePool{s: s}, nil
}

// GetKubernetesNodePool gets an existing KubernetesNodePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKubernetesNodePool(ctx *pulumi.Context,
	name string, id pulumi.ID, state *KubernetesNodePoolState, opts ...pulumi.ResourceOpt) (*KubernetesNodePool, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["clusterId"] = state.ClusterId
		inputs["name"] = state.Name
		inputs["nodeCount"] = state.NodeCount
		inputs["nodes"] = state.Nodes
		inputs["size"] = state.Size
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("digitalocean:index/kubernetesNodePool:KubernetesNodePool", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &KubernetesNodePool{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *KubernetesNodePool) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *KubernetesNodePool) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ID of the Kubernetes cluster to which the node pool is associated.
func (r *KubernetesNodePool) ClusterId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["clusterId"])
}

// A name for the node pool.
func (r *KubernetesNodePool) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The number of Droplet instances in the node pool.
func (r *KubernetesNodePool) NodeCount() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["nodeCount"])
}

// A list of nodes in the pool. Each node exports the following attributes:
// - `id` -  A unique ID that can be used to identify and reference the node.
// - `name` - The auto-generated name for the node.
// - `status` -  A string indicating the current status of the individual node.
// - `created_at` - The date and time when the node was created.
// - `updated_at` - The date and time when the node was last updated.
func (r *KubernetesNodePool) Nodes() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["nodes"])
}

// The slug identifier for the type of Droplet to be used as workers in the node pool.
func (r *KubernetesNodePool) Size() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["size"])
}

// A list of tag names to be applied to the Kubernetes cluster.
func (r *KubernetesNodePool) Tags() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering KubernetesNodePool resources.
type KubernetesNodePoolState struct {
	// The ID of the Kubernetes cluster to which the node pool is associated.
	ClusterId interface{}
	// A name for the node pool.
	Name interface{}
	// The number of Droplet instances in the node pool.
	NodeCount interface{}
	// A list of nodes in the pool. Each node exports the following attributes:
	// - `id` -  A unique ID that can be used to identify and reference the node.
	// - `name` - The auto-generated name for the node.
	// - `status` -  A string indicating the current status of the individual node.
	// - `created_at` - The date and time when the node was created.
	// - `updated_at` - The date and time when the node was last updated.
	Nodes interface{}
	// The slug identifier for the type of Droplet to be used as workers in the node pool.
	Size interface{}
	// A list of tag names to be applied to the Kubernetes cluster.
	Tags interface{}
}

// The set of arguments for constructing a KubernetesNodePool resource.
type KubernetesNodePoolArgs struct {
	// The ID of the Kubernetes cluster to which the node pool is associated.
	ClusterId interface{}
	// A name for the node pool.
	Name interface{}
	// The number of Droplet instances in the node pool.
	NodeCount interface{}
	// The slug identifier for the type of Droplet to be used as workers in the node pool.
	Size interface{}
	// A list of tag names to be applied to the Kubernetes cluster.
	Tags interface{}
}