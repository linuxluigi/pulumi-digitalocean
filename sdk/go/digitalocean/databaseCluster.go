// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DigitalOcean database cluster resource.
//
// ## Example Usage
// ### Create a new PostgreSQL database cluster
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean/index"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := digitalocean.NewDatabaseCluster(ctx, "postgres_example", &digitalocean.DatabaseClusterArgs{
// 			Engine:    pulumi.String("pg"),
// 			NodeCount: pulumi.Int(1),
// 			Region:    pulumi.String("nyc1"),
// 			Size:      pulumi.String("db-s-1vcpu-1gb"),
// 			Version:   pulumi.String("11"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Create a new MySQL database cluster
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean/index"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := digitalocean.NewDatabaseCluster(ctx, "mysql_example", &digitalocean.DatabaseClusterArgs{
// 			Engine:    pulumi.String("mysql"),
// 			NodeCount: pulumi.Int(1),
// 			Region:    pulumi.String("nyc1"),
// 			Size:      pulumi.String("db-s-1vcpu-1gb"),
// 			Version:   pulumi.String("8"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Create a new Redis database cluster
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean/index"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := digitalocean.NewDatabaseCluster(ctx, "redis_example", &digitalocean.DatabaseClusterArgs{
// 			Engine:    pulumi.String("redis"),
// 			NodeCount: pulumi.Int(1),
// 			Region:    pulumi.String("nyc1"),
// 			Size:      pulumi.String("db-s-1vcpu-1gb"),
// 			Version:   pulumi.String("6"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Create a new MongoDB database cluster
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean/index"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := digitalocean.NewDatabaseCluster(ctx, "mongodb_example", &digitalocean.DatabaseClusterArgs{
// 			Engine:    pulumi.String("mongodb"),
// 			NodeCount: pulumi.Int(1),
// 			Region:    pulumi.String("nyc3"),
// 			Size:      pulumi.String("db-s-1vcpu-1gb"),
// 			Version:   pulumi.String("4"),
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
// Database clusters can be imported using the `id` returned from DigitalOcean, e.g.
//
// ```sh
//  $ pulumi import digitalocean:index/databaseCluster:DatabaseCluster mycluster 245bcfd0-7f31-4ce6-a2bc-475a116cca97
// ```
type DatabaseCluster struct {
	pulumi.CustomResourceState

	// The uniform resource name of the database cluster.
	ClusterUrn pulumi.StringOutput `pulumi:"clusterUrn"`
	// Name of the cluster's default database.
	Database pulumi.StringOutput `pulumi:"database"`
	// Database engine used by the cluster (ex. `pg` for PostreSQL, `mysql` for MySQL, `redis` for Redis, or `mongodb` for MongoDB).
	Engine pulumi.StringOutput `pulumi:"engine"`
	// A string specifying the eviction policy for a Redis cluster. Valid values are: `noeviction`, `allkeysLru`, `allkeysRandom`, `volatileLru`, `volatileRandom`, or `volatileTtl`.
	EvictionPolicy pulumi.StringPtrOutput `pulumi:"evictionPolicy"`
	// Database cluster's hostname.
	Host pulumi.StringOutput `pulumi:"host"`
	// Defines when the automatic maintenance should be performed for the database cluster.
	MaintenanceWindows DatabaseClusterMaintenanceWindowArrayOutput `pulumi:"maintenanceWindows"`
	// The name of the database cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of nodes that will be included in the cluster.
	NodeCount pulumi.IntOutput `pulumi:"nodeCount"`
	// Password for the cluster's default user.
	Password pulumi.StringOutput `pulumi:"password"`
	// Network port that the database cluster is listening on.
	Port pulumi.IntOutput `pulumi:"port"`
	// Same as `host`, but only accessible from resources within the account and in the same region.
	PrivateHost pulumi.StringOutput `pulumi:"privateHost"`
	// The ID of the VPC where the database cluster will be located.
	PrivateNetworkUuid pulumi.StringOutput `pulumi:"privateNetworkUuid"`
	// Same as `uri`, but only accessible from resources within the account and in the same region.
	PrivateUri pulumi.StringOutput `pulumi:"privateUri"`
	// DigitalOcean region where the cluster will reside.
	Region pulumi.StringOutput `pulumi:"region"`
	// Database Droplet size associated with the cluster (ex. `db-s-1vcpu-1gb`). See here for a [list of valid size slugs](https://docs.digitalocean.com/reference/api/api-reference/#tag/Databases).
	Size pulumi.StringOutput `pulumi:"size"`
	// A comma separated string specifying the  SQL modes for a MySQL cluster.
	SqlMode pulumi.StringPtrOutput `pulumi:"sqlMode"`
	// A list of tag names to be applied to the database cluster.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The full URI for connecting to the database cluster.
	Uri pulumi.StringOutput `pulumi:"uri"`
	// Username for the cluster's default user.
	User pulumi.StringOutput `pulumi:"user"`
	// Engine version used by the cluster (ex. `11` for PostgreSQL 11).
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewDatabaseCluster registers a new resource with the given unique name, arguments, and options.
func NewDatabaseCluster(ctx *pulumi.Context,
	name string, args *DatabaseClusterArgs, opts ...pulumi.ResourceOption) (*DatabaseCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Engine == nil {
		return nil, errors.New("invalid value for required argument 'Engine'")
	}
	if args.NodeCount == nil {
		return nil, errors.New("invalid value for required argument 'NodeCount'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	var resource DatabaseCluster
	err := ctx.RegisterResource("digitalocean:index/databaseCluster:DatabaseCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseCluster gets an existing DatabaseCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseClusterState, opts ...pulumi.ResourceOption) (*DatabaseCluster, error) {
	var resource DatabaseCluster
	err := ctx.ReadResource("digitalocean:index/databaseCluster:DatabaseCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseCluster resources.
type databaseClusterState struct {
	// The uniform resource name of the database cluster.
	ClusterUrn *string `pulumi:"clusterUrn"`
	// Name of the cluster's default database.
	Database *string `pulumi:"database"`
	// Database engine used by the cluster (ex. `pg` for PostreSQL, `mysql` for MySQL, `redis` for Redis, or `mongodb` for MongoDB).
	Engine *string `pulumi:"engine"`
	// A string specifying the eviction policy for a Redis cluster. Valid values are: `noeviction`, `allkeysLru`, `allkeysRandom`, `volatileLru`, `volatileRandom`, or `volatileTtl`.
	EvictionPolicy *string `pulumi:"evictionPolicy"`
	// Database cluster's hostname.
	Host *string `pulumi:"host"`
	// Defines when the automatic maintenance should be performed for the database cluster.
	MaintenanceWindows []DatabaseClusterMaintenanceWindow `pulumi:"maintenanceWindows"`
	// The name of the database cluster.
	Name *string `pulumi:"name"`
	// Number of nodes that will be included in the cluster.
	NodeCount *int `pulumi:"nodeCount"`
	// Password for the cluster's default user.
	Password *string `pulumi:"password"`
	// Network port that the database cluster is listening on.
	Port *int `pulumi:"port"`
	// Same as `host`, but only accessible from resources within the account and in the same region.
	PrivateHost *string `pulumi:"privateHost"`
	// The ID of the VPC where the database cluster will be located.
	PrivateNetworkUuid *string `pulumi:"privateNetworkUuid"`
	// Same as `uri`, but only accessible from resources within the account and in the same region.
	PrivateUri *string `pulumi:"privateUri"`
	// DigitalOcean region where the cluster will reside.
	Region *string `pulumi:"region"`
	// Database Droplet size associated with the cluster (ex. `db-s-1vcpu-1gb`). See here for a [list of valid size slugs](https://docs.digitalocean.com/reference/api/api-reference/#tag/Databases).
	Size *string `pulumi:"size"`
	// A comma separated string specifying the  SQL modes for a MySQL cluster.
	SqlMode *string `pulumi:"sqlMode"`
	// A list of tag names to be applied to the database cluster.
	Tags []string `pulumi:"tags"`
	// The full URI for connecting to the database cluster.
	Uri *string `pulumi:"uri"`
	// Username for the cluster's default user.
	User *string `pulumi:"user"`
	// Engine version used by the cluster (ex. `11` for PostgreSQL 11).
	Version *string `pulumi:"version"`
}

type DatabaseClusterState struct {
	// The uniform resource name of the database cluster.
	ClusterUrn pulumi.StringPtrInput
	// Name of the cluster's default database.
	Database pulumi.StringPtrInput
	// Database engine used by the cluster (ex. `pg` for PostreSQL, `mysql` for MySQL, `redis` for Redis, or `mongodb` for MongoDB).
	Engine pulumi.StringPtrInput
	// A string specifying the eviction policy for a Redis cluster. Valid values are: `noeviction`, `allkeysLru`, `allkeysRandom`, `volatileLru`, `volatileRandom`, or `volatileTtl`.
	EvictionPolicy pulumi.StringPtrInput
	// Database cluster's hostname.
	Host pulumi.StringPtrInput
	// Defines when the automatic maintenance should be performed for the database cluster.
	MaintenanceWindows DatabaseClusterMaintenanceWindowArrayInput
	// The name of the database cluster.
	Name pulumi.StringPtrInput
	// Number of nodes that will be included in the cluster.
	NodeCount pulumi.IntPtrInput
	// Password for the cluster's default user.
	Password pulumi.StringPtrInput
	// Network port that the database cluster is listening on.
	Port pulumi.IntPtrInput
	// Same as `host`, but only accessible from resources within the account and in the same region.
	PrivateHost pulumi.StringPtrInput
	// The ID of the VPC where the database cluster will be located.
	PrivateNetworkUuid pulumi.StringPtrInput
	// Same as `uri`, but only accessible from resources within the account and in the same region.
	PrivateUri pulumi.StringPtrInput
	// DigitalOcean region where the cluster will reside.
	Region pulumi.StringPtrInput
	// Database Droplet size associated with the cluster (ex. `db-s-1vcpu-1gb`). See here for a [list of valid size slugs](https://docs.digitalocean.com/reference/api/api-reference/#tag/Databases).
	Size pulumi.StringPtrInput
	// A comma separated string specifying the  SQL modes for a MySQL cluster.
	SqlMode pulumi.StringPtrInput
	// A list of tag names to be applied to the database cluster.
	Tags pulumi.StringArrayInput
	// The full URI for connecting to the database cluster.
	Uri pulumi.StringPtrInput
	// Username for the cluster's default user.
	User pulumi.StringPtrInput
	// Engine version used by the cluster (ex. `11` for PostgreSQL 11).
	Version pulumi.StringPtrInput
}

func (DatabaseClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseClusterState)(nil)).Elem()
}

type databaseClusterArgs struct {
	// Database engine used by the cluster (ex. `pg` for PostreSQL, `mysql` for MySQL, `redis` for Redis, or `mongodb` for MongoDB).
	Engine string `pulumi:"engine"`
	// A string specifying the eviction policy for a Redis cluster. Valid values are: `noeviction`, `allkeysLru`, `allkeysRandom`, `volatileLru`, `volatileRandom`, or `volatileTtl`.
	EvictionPolicy *string `pulumi:"evictionPolicy"`
	// Defines when the automatic maintenance should be performed for the database cluster.
	MaintenanceWindows []DatabaseClusterMaintenanceWindow `pulumi:"maintenanceWindows"`
	// The name of the database cluster.
	Name *string `pulumi:"name"`
	// Number of nodes that will be included in the cluster.
	NodeCount int `pulumi:"nodeCount"`
	// The ID of the VPC where the database cluster will be located.
	PrivateNetworkUuid *string `pulumi:"privateNetworkUuid"`
	// DigitalOcean region where the cluster will reside.
	Region string `pulumi:"region"`
	// Database Droplet size associated with the cluster (ex. `db-s-1vcpu-1gb`). See here for a [list of valid size slugs](https://docs.digitalocean.com/reference/api/api-reference/#tag/Databases).
	Size string `pulumi:"size"`
	// A comma separated string specifying the  SQL modes for a MySQL cluster.
	SqlMode *string `pulumi:"sqlMode"`
	// A list of tag names to be applied to the database cluster.
	Tags []string `pulumi:"tags"`
	// Engine version used by the cluster (ex. `11` for PostgreSQL 11).
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a DatabaseCluster resource.
type DatabaseClusterArgs struct {
	// Database engine used by the cluster (ex. `pg` for PostreSQL, `mysql` for MySQL, `redis` for Redis, or `mongodb` for MongoDB).
	Engine pulumi.StringInput
	// A string specifying the eviction policy for a Redis cluster. Valid values are: `noeviction`, `allkeysLru`, `allkeysRandom`, `volatileLru`, `volatileRandom`, or `volatileTtl`.
	EvictionPolicy pulumi.StringPtrInput
	// Defines when the automatic maintenance should be performed for the database cluster.
	MaintenanceWindows DatabaseClusterMaintenanceWindowArrayInput
	// The name of the database cluster.
	Name pulumi.StringPtrInput
	// Number of nodes that will be included in the cluster.
	NodeCount pulumi.IntInput
	// The ID of the VPC where the database cluster will be located.
	PrivateNetworkUuid pulumi.StringPtrInput
	// DigitalOcean region where the cluster will reside.
	Region pulumi.StringInput
	// Database Droplet size associated with the cluster (ex. `db-s-1vcpu-1gb`). See here for a [list of valid size slugs](https://docs.digitalocean.com/reference/api/api-reference/#tag/Databases).
	Size pulumi.StringInput
	// A comma separated string specifying the  SQL modes for a MySQL cluster.
	SqlMode pulumi.StringPtrInput
	// A list of tag names to be applied to the database cluster.
	Tags pulumi.StringArrayInput
	// Engine version used by the cluster (ex. `11` for PostgreSQL 11).
	Version pulumi.StringPtrInput
}

func (DatabaseClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseClusterArgs)(nil)).Elem()
}

type DatabaseClusterInput interface {
	pulumi.Input

	ToDatabaseClusterOutput() DatabaseClusterOutput
	ToDatabaseClusterOutputWithContext(ctx context.Context) DatabaseClusterOutput
}

func (*DatabaseCluster) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseCluster)(nil))
}

func (i *DatabaseCluster) ToDatabaseClusterOutput() DatabaseClusterOutput {
	return i.ToDatabaseClusterOutputWithContext(context.Background())
}

func (i *DatabaseCluster) ToDatabaseClusterOutputWithContext(ctx context.Context) DatabaseClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseClusterOutput)
}

func (i *DatabaseCluster) ToDatabaseClusterPtrOutput() DatabaseClusterPtrOutput {
	return i.ToDatabaseClusterPtrOutputWithContext(context.Background())
}

func (i *DatabaseCluster) ToDatabaseClusterPtrOutputWithContext(ctx context.Context) DatabaseClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseClusterPtrOutput)
}

type DatabaseClusterPtrInput interface {
	pulumi.Input

	ToDatabaseClusterPtrOutput() DatabaseClusterPtrOutput
	ToDatabaseClusterPtrOutputWithContext(ctx context.Context) DatabaseClusterPtrOutput
}

type databaseClusterPtrType DatabaseClusterArgs

func (*databaseClusterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseCluster)(nil))
}

func (i *databaseClusterPtrType) ToDatabaseClusterPtrOutput() DatabaseClusterPtrOutput {
	return i.ToDatabaseClusterPtrOutputWithContext(context.Background())
}

func (i *databaseClusterPtrType) ToDatabaseClusterPtrOutputWithContext(ctx context.Context) DatabaseClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseClusterPtrOutput)
}

// DatabaseClusterArrayInput is an input type that accepts DatabaseClusterArray and DatabaseClusterArrayOutput values.
// You can construct a concrete instance of `DatabaseClusterArrayInput` via:
//
//          DatabaseClusterArray{ DatabaseClusterArgs{...} }
type DatabaseClusterArrayInput interface {
	pulumi.Input

	ToDatabaseClusterArrayOutput() DatabaseClusterArrayOutput
	ToDatabaseClusterArrayOutputWithContext(context.Context) DatabaseClusterArrayOutput
}

type DatabaseClusterArray []DatabaseClusterInput

func (DatabaseClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseCluster)(nil)).Elem()
}

func (i DatabaseClusterArray) ToDatabaseClusterArrayOutput() DatabaseClusterArrayOutput {
	return i.ToDatabaseClusterArrayOutputWithContext(context.Background())
}

func (i DatabaseClusterArray) ToDatabaseClusterArrayOutputWithContext(ctx context.Context) DatabaseClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseClusterArrayOutput)
}

// DatabaseClusterMapInput is an input type that accepts DatabaseClusterMap and DatabaseClusterMapOutput values.
// You can construct a concrete instance of `DatabaseClusterMapInput` via:
//
//          DatabaseClusterMap{ "key": DatabaseClusterArgs{...} }
type DatabaseClusterMapInput interface {
	pulumi.Input

	ToDatabaseClusterMapOutput() DatabaseClusterMapOutput
	ToDatabaseClusterMapOutputWithContext(context.Context) DatabaseClusterMapOutput
}

type DatabaseClusterMap map[string]DatabaseClusterInput

func (DatabaseClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseCluster)(nil)).Elem()
}

func (i DatabaseClusterMap) ToDatabaseClusterMapOutput() DatabaseClusterMapOutput {
	return i.ToDatabaseClusterMapOutputWithContext(context.Background())
}

func (i DatabaseClusterMap) ToDatabaseClusterMapOutputWithContext(ctx context.Context) DatabaseClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseClusterMapOutput)
}

type DatabaseClusterOutput struct{ *pulumi.OutputState }

func (DatabaseClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseCluster)(nil))
}

func (o DatabaseClusterOutput) ToDatabaseClusterOutput() DatabaseClusterOutput {
	return o
}

func (o DatabaseClusterOutput) ToDatabaseClusterOutputWithContext(ctx context.Context) DatabaseClusterOutput {
	return o
}

func (o DatabaseClusterOutput) ToDatabaseClusterPtrOutput() DatabaseClusterPtrOutput {
	return o.ToDatabaseClusterPtrOutputWithContext(context.Background())
}

func (o DatabaseClusterOutput) ToDatabaseClusterPtrOutputWithContext(ctx context.Context) DatabaseClusterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatabaseCluster) *DatabaseCluster {
		return &v
	}).(DatabaseClusterPtrOutput)
}

type DatabaseClusterPtrOutput struct{ *pulumi.OutputState }

func (DatabaseClusterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseCluster)(nil))
}

func (o DatabaseClusterPtrOutput) ToDatabaseClusterPtrOutput() DatabaseClusterPtrOutput {
	return o
}

func (o DatabaseClusterPtrOutput) ToDatabaseClusterPtrOutputWithContext(ctx context.Context) DatabaseClusterPtrOutput {
	return o
}

func (o DatabaseClusterPtrOutput) Elem() DatabaseClusterOutput {
	return o.ApplyT(func(v *DatabaseCluster) DatabaseCluster {
		if v != nil {
			return *v
		}
		var ret DatabaseCluster
		return ret
	}).(DatabaseClusterOutput)
}

type DatabaseClusterArrayOutput struct{ *pulumi.OutputState }

func (DatabaseClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseCluster)(nil))
}

func (o DatabaseClusterArrayOutput) ToDatabaseClusterArrayOutput() DatabaseClusterArrayOutput {
	return o
}

func (o DatabaseClusterArrayOutput) ToDatabaseClusterArrayOutputWithContext(ctx context.Context) DatabaseClusterArrayOutput {
	return o
}

func (o DatabaseClusterArrayOutput) Index(i pulumi.IntInput) DatabaseClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabaseCluster {
		return vs[0].([]DatabaseCluster)[vs[1].(int)]
	}).(DatabaseClusterOutput)
}

type DatabaseClusterMapOutput struct{ *pulumi.OutputState }

func (DatabaseClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DatabaseCluster)(nil))
}

func (o DatabaseClusterMapOutput) ToDatabaseClusterMapOutput() DatabaseClusterMapOutput {
	return o
}

func (o DatabaseClusterMapOutput) ToDatabaseClusterMapOutputWithContext(ctx context.Context) DatabaseClusterMapOutput {
	return o
}

func (o DatabaseClusterMapOutput) MapIndex(k pulumi.StringInput) DatabaseClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DatabaseCluster {
		return vs[0].(map[string]DatabaseCluster)[vs[1].(string)]
	}).(DatabaseClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseClusterInput)(nil)).Elem(), &DatabaseCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseClusterPtrInput)(nil)).Elem(), &DatabaseCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseClusterArrayInput)(nil)).Elem(), DatabaseClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseClusterMapInput)(nil)).Elem(), DatabaseClusterMap{})
	pulumi.RegisterOutputType(DatabaseClusterOutput{})
	pulumi.RegisterOutputType(DatabaseClusterPtrOutput{})
	pulumi.RegisterOutputType(DatabaseClusterArrayOutput{})
	pulumi.RegisterOutputType(DatabaseClusterMapOutput{})
}
