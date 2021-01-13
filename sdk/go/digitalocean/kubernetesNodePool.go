// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a DigitalOcean Kubernetes node pool resource. While the default node pool must be defined in the `KubernetesCluster` resource, this resource can be used to add additional ones to a cluster.
//
// ## Example Usage
// ### Basic Example
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v3/go/digitalocean"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		foo, err := digitalocean.NewKubernetesCluster(ctx, "foo", &digitalocean.KubernetesClusterArgs{
// 			Region:  pulumi.String("nyc1"),
// 			Version: pulumi.String("1.15.5-do.1"),
// 			NodePool: &digitalocean.KubernetesClusterNodePoolArgs{
// 				Name:      pulumi.String("front-end-pool"),
// 				Size:      pulumi.String("s-2vcpu-2gb"),
// 				NodeCount: pulumi.Int(3),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = digitalocean.NewKubernetesNodePool(ctx, "bar", &digitalocean.KubernetesNodePoolArgs{
// 			ClusterId: foo.ID(),
// 			Size:      pulumi.String("c-2"),
// 			NodeCount: pulumi.Int(2),
// 			Tags: pulumi.StringArray{
// 				pulumi.String("backend"),
// 			},
// 			Labels: pulumi.StringMap{
// 				"service":  pulumi.String("backend"),
// 				"priority": pulumi.String("high"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Autoscaling Example
//
// Node pools may also be configured to [autoscale](https://www.digitalocean.com/docs/kubernetes/how-to/autoscale/).
// For example:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v3/go/digitalocean"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := digitalocean.NewKubernetesNodePool(ctx, "autoscale_pool_01", &digitalocean.KubernetesNodePoolArgs{
// 			ClusterId: pulumi.Any(digitalocean_kubernetes_cluster.Foo.Id),
// 			Size:      pulumi.String("s-1vcpu-2gb"),
// 			AutoScale: pulumi.Bool(true),
// 			MinNodes:  pulumi.Int(1),
// 			MaxNodes:  pulumi.Int(5),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// If you are importing an existing Kubernetes cluster, just import the cluster. Importing a cluster also imports all of its associated node pools. If you still need to import a single node pool, then import it by using its `id`, e.g.
//
// ```sh
//  $ pulumi import digitalocean:index/kubernetesNodePool:KubernetesNodePool mynodepool 9d76f410-9284-4436-9633-4066852442c8
// ```
//
//  NoteIf the node pool has the `terraform:default-node-pool` tag, then it is a default node pool for an existing cluster. The provider will refuse to import the node pool in that case because the node pool is managed by the `digitalocean_kubernetes_cluster` resource and not by this `digitalocean_kubernetes_node_pool` resource.
type KubernetesNodePool struct {
	pulumi.CustomResourceState

	// A computed field representing the actual number of nodes in the node pool, which is especially useful when auto-scaling is enabled.
	ActualNodeCount pulumi.IntOutput `pulumi:"actualNodeCount"`
	// Enable auto-scaling of the number of nodes in the node pool within the given min/max range.
	AutoScale pulumi.BoolPtrOutput `pulumi:"autoScale"`
	// The ID of the Kubernetes cluster to which the node pool is associated.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// A map of key/value pairs to apply to nodes in the pool. The labels are exposed in the Kubernetes API as labels in the metadata of the corresponding [Node resources](https://kubernetes.io/docs/concepts/architecture/nodes/).
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to.
	MaxNodes pulumi.IntPtrOutput `pulumi:"maxNodes"`
	// If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to.
	MinNodes pulumi.IntPtrOutput `pulumi:"minNodes"`
	// A name for the node pool.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of Droplet instances in the node pool. If auto-scaling is enabled, this should only be set if the desired result is to explicitly reset the number of nodes to this value. If auto-scaling is enabled, and the node count is outside of the given min/max range, it will use the min nodes value.
	NodeCount pulumi.IntPtrOutput `pulumi:"nodeCount"`
	// A list of nodes in the pool. Each node exports the following attributes:
	Nodes KubernetesNodePoolNodeArrayOutput `pulumi:"nodes"`
	// The slug identifier for the type of Droplet to be used as workers in the node pool.
	Size pulumi.StringOutput `pulumi:"size"`
	// A list of tag names to be applied to the Kubernetes cluster.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewKubernetesNodePool registers a new resource with the given unique name, arguments, and options.
func NewKubernetesNodePool(ctx *pulumi.Context,
	name string, args *KubernetesNodePoolArgs, opts ...pulumi.ResourceOption) (*KubernetesNodePool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	var resource KubernetesNodePool
	err := ctx.RegisterResource("digitalocean:index/kubernetesNodePool:KubernetesNodePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKubernetesNodePool gets an existing KubernetesNodePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKubernetesNodePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubernetesNodePoolState, opts ...pulumi.ResourceOption) (*KubernetesNodePool, error) {
	var resource KubernetesNodePool
	err := ctx.ReadResource("digitalocean:index/kubernetesNodePool:KubernetesNodePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KubernetesNodePool resources.
type kubernetesNodePoolState struct {
	// A computed field representing the actual number of nodes in the node pool, which is especially useful when auto-scaling is enabled.
	ActualNodeCount *int `pulumi:"actualNodeCount"`
	// Enable auto-scaling of the number of nodes in the node pool within the given min/max range.
	AutoScale *bool `pulumi:"autoScale"`
	// The ID of the Kubernetes cluster to which the node pool is associated.
	ClusterId *string `pulumi:"clusterId"`
	// A map of key/value pairs to apply to nodes in the pool. The labels are exposed in the Kubernetes API as labels in the metadata of the corresponding [Node resources](https://kubernetes.io/docs/concepts/architecture/nodes/).
	Labels map[string]string `pulumi:"labels"`
	// If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to.
	MaxNodes *int `pulumi:"maxNodes"`
	// If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to.
	MinNodes *int `pulumi:"minNodes"`
	// A name for the node pool.
	Name *string `pulumi:"name"`
	// The number of Droplet instances in the node pool. If auto-scaling is enabled, this should only be set if the desired result is to explicitly reset the number of nodes to this value. If auto-scaling is enabled, and the node count is outside of the given min/max range, it will use the min nodes value.
	NodeCount *int `pulumi:"nodeCount"`
	// A list of nodes in the pool. Each node exports the following attributes:
	Nodes []KubernetesNodePoolNode `pulumi:"nodes"`
	// The slug identifier for the type of Droplet to be used as workers in the node pool.
	Size *string `pulumi:"size"`
	// A list of tag names to be applied to the Kubernetes cluster.
	Tags []string `pulumi:"tags"`
}

type KubernetesNodePoolState struct {
	// A computed field representing the actual number of nodes in the node pool, which is especially useful when auto-scaling is enabled.
	ActualNodeCount pulumi.IntPtrInput
	// Enable auto-scaling of the number of nodes in the node pool within the given min/max range.
	AutoScale pulumi.BoolPtrInput
	// The ID of the Kubernetes cluster to which the node pool is associated.
	ClusterId pulumi.StringPtrInput
	// A map of key/value pairs to apply to nodes in the pool. The labels are exposed in the Kubernetes API as labels in the metadata of the corresponding [Node resources](https://kubernetes.io/docs/concepts/architecture/nodes/).
	Labels pulumi.StringMapInput
	// If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to.
	MaxNodes pulumi.IntPtrInput
	// If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to.
	MinNodes pulumi.IntPtrInput
	// A name for the node pool.
	Name pulumi.StringPtrInput
	// The number of Droplet instances in the node pool. If auto-scaling is enabled, this should only be set if the desired result is to explicitly reset the number of nodes to this value. If auto-scaling is enabled, and the node count is outside of the given min/max range, it will use the min nodes value.
	NodeCount pulumi.IntPtrInput
	// A list of nodes in the pool. Each node exports the following attributes:
	Nodes KubernetesNodePoolNodeArrayInput
	// The slug identifier for the type of Droplet to be used as workers in the node pool.
	Size pulumi.StringPtrInput
	// A list of tag names to be applied to the Kubernetes cluster.
	Tags pulumi.StringArrayInput
}

func (KubernetesNodePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesNodePoolState)(nil)).Elem()
}

type kubernetesNodePoolArgs struct {
	// Enable auto-scaling of the number of nodes in the node pool within the given min/max range.
	AutoScale *bool `pulumi:"autoScale"`
	// The ID of the Kubernetes cluster to which the node pool is associated.
	ClusterId string `pulumi:"clusterId"`
	// A map of key/value pairs to apply to nodes in the pool. The labels are exposed in the Kubernetes API as labels in the metadata of the corresponding [Node resources](https://kubernetes.io/docs/concepts/architecture/nodes/).
	Labels map[string]string `pulumi:"labels"`
	// If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to.
	MaxNodes *int `pulumi:"maxNodes"`
	// If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to.
	MinNodes *int `pulumi:"minNodes"`
	// A name for the node pool.
	Name *string `pulumi:"name"`
	// The number of Droplet instances in the node pool. If auto-scaling is enabled, this should only be set if the desired result is to explicitly reset the number of nodes to this value. If auto-scaling is enabled, and the node count is outside of the given min/max range, it will use the min nodes value.
	NodeCount *int `pulumi:"nodeCount"`
	// The slug identifier for the type of Droplet to be used as workers in the node pool.
	Size string `pulumi:"size"`
	// A list of tag names to be applied to the Kubernetes cluster.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a KubernetesNodePool resource.
type KubernetesNodePoolArgs struct {
	// Enable auto-scaling of the number of nodes in the node pool within the given min/max range.
	AutoScale pulumi.BoolPtrInput
	// The ID of the Kubernetes cluster to which the node pool is associated.
	ClusterId pulumi.StringInput
	// A map of key/value pairs to apply to nodes in the pool. The labels are exposed in the Kubernetes API as labels in the metadata of the corresponding [Node resources](https://kubernetes.io/docs/concepts/architecture/nodes/).
	Labels pulumi.StringMapInput
	// If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to.
	MaxNodes pulumi.IntPtrInput
	// If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to.
	MinNodes pulumi.IntPtrInput
	// A name for the node pool.
	Name pulumi.StringPtrInput
	// The number of Droplet instances in the node pool. If auto-scaling is enabled, this should only be set if the desired result is to explicitly reset the number of nodes to this value. If auto-scaling is enabled, and the node count is outside of the given min/max range, it will use the min nodes value.
	NodeCount pulumi.IntPtrInput
	// The slug identifier for the type of Droplet to be used as workers in the node pool.
	Size pulumi.StringInput
	// A list of tag names to be applied to the Kubernetes cluster.
	Tags pulumi.StringArrayInput
}

func (KubernetesNodePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesNodePoolArgs)(nil)).Elem()
}

type KubernetesNodePoolInput interface {
	pulumi.Input

	ToKubernetesNodePoolOutput() KubernetesNodePoolOutput
	ToKubernetesNodePoolOutputWithContext(ctx context.Context) KubernetesNodePoolOutput
}

func (KubernetesNodePool) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesNodePool)(nil)).Elem()
}

func (i KubernetesNodePool) ToKubernetesNodePoolOutput() KubernetesNodePoolOutput {
	return i.ToKubernetesNodePoolOutputWithContext(context.Background())
}

func (i KubernetesNodePool) ToKubernetesNodePoolOutputWithContext(ctx context.Context) KubernetesNodePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesNodePoolOutput)
}

type KubernetesNodePoolOutput struct {
	*pulumi.OutputState
}

func (KubernetesNodePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesNodePoolOutput)(nil)).Elem()
}

func (o KubernetesNodePoolOutput) ToKubernetesNodePoolOutput() KubernetesNodePoolOutput {
	return o
}

func (o KubernetesNodePoolOutput) ToKubernetesNodePoolOutputWithContext(ctx context.Context) KubernetesNodePoolOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(KubernetesNodePoolOutput{})
}
