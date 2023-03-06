// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a DigitalOcean database connection pool resource.
//
// ## Example Usage
// ### Create a new PostgreSQL database connection pool
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := digitalocean.NewDatabaseCluster(ctx, "postgres-example", &digitalocean.DatabaseClusterArgs{
//				Engine:    pulumi.String("pg"),
//				Version:   pulumi.String("11"),
//				Size:      pulumi.String("db-s-1vcpu-1gb"),
//				Region:    pulumi.String("nyc1"),
//				NodeCount: pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = digitalocean.NewDatabaseConnectionPool(ctx, "pool-01", &digitalocean.DatabaseConnectionPoolArgs{
//				ClusterId: postgres_example.ID(),
//				Mode:      pulumi.String("transaction"),
//				Size:      pulumi.Int(20),
//				DbName:    pulumi.String("defaultdb"),
//				User:      pulumi.String("doadmin"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Database connection pools can be imported using the `id` of the source database cluster and the `name` of the connection pool joined with a comma. For example
//
// ```sh
//
//	$ pulumi import digitalocean:index/databaseConnectionPool:DatabaseConnectionPool pool-01 245bcfd0-7f31-4ce6-a2bc-475a116cca97,pool-01
//
// ```
type DatabaseConnectionPool struct {
	pulumi.CustomResourceState

	// The ID of the source database cluster. Note: This must be a PostgreSQL cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The database for use with the connection pool.
	DbName pulumi.StringOutput `pulumi:"dbName"`
	// The hostname used to connect to the database connection pool.
	Host pulumi.StringOutput `pulumi:"host"`
	// The PGBouncer transaction mode for the connection pool. The allowed values are session, transaction, and statement.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// The name for the database connection pool.
	Name pulumi.StringOutput `pulumi:"name"`
	// Password for the connection pool's user.
	Password pulumi.StringOutput `pulumi:"password"`
	// Network port that the database connection pool is listening on.
	Port pulumi.IntOutput `pulumi:"port"`
	// Same as `host`, but only accessible from resources within the account and in the same region.
	PrivateHost pulumi.StringOutput `pulumi:"privateHost"`
	// Same as `uri`, but only accessible from resources within the account and in the same region.
	PrivateUri pulumi.StringOutput `pulumi:"privateUri"`
	// The desired size of the PGBouncer connection pool.
	Size pulumi.IntOutput `pulumi:"size"`
	// The full URI for connecting to the database connection pool.
	Uri pulumi.StringOutput `pulumi:"uri"`
	// The name of the database user for use with the connection pool. When excluded, all sessions connect to the database as the inbound user.
	User pulumi.StringPtrOutput `pulumi:"user"`
}

// NewDatabaseConnectionPool registers a new resource with the given unique name, arguments, and options.
func NewDatabaseConnectionPool(ctx *pulumi.Context,
	name string, args *DatabaseConnectionPoolArgs, opts ...pulumi.ResourceOption) (*DatabaseConnectionPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.DbName == nil {
		return nil, errors.New("invalid value for required argument 'DbName'")
	}
	if args.Mode == nil {
		return nil, errors.New("invalid value for required argument 'Mode'")
	}
	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
		"privateUri",
		"uri",
	})
	opts = append(opts, secrets)
	var resource DatabaseConnectionPool
	err := ctx.RegisterResource("digitalocean:index/databaseConnectionPool:DatabaseConnectionPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseConnectionPool gets an existing DatabaseConnectionPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseConnectionPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseConnectionPoolState, opts ...pulumi.ResourceOption) (*DatabaseConnectionPool, error) {
	var resource DatabaseConnectionPool
	err := ctx.ReadResource("digitalocean:index/databaseConnectionPool:DatabaseConnectionPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseConnectionPool resources.
type databaseConnectionPoolState struct {
	// The ID of the source database cluster. Note: This must be a PostgreSQL cluster.
	ClusterId *string `pulumi:"clusterId"`
	// The database for use with the connection pool.
	DbName *string `pulumi:"dbName"`
	// The hostname used to connect to the database connection pool.
	Host *string `pulumi:"host"`
	// The PGBouncer transaction mode for the connection pool. The allowed values are session, transaction, and statement.
	Mode *string `pulumi:"mode"`
	// The name for the database connection pool.
	Name *string `pulumi:"name"`
	// Password for the connection pool's user.
	Password *string `pulumi:"password"`
	// Network port that the database connection pool is listening on.
	Port *int `pulumi:"port"`
	// Same as `host`, but only accessible from resources within the account and in the same region.
	PrivateHost *string `pulumi:"privateHost"`
	// Same as `uri`, but only accessible from resources within the account and in the same region.
	PrivateUri *string `pulumi:"privateUri"`
	// The desired size of the PGBouncer connection pool.
	Size *int `pulumi:"size"`
	// The full URI for connecting to the database connection pool.
	Uri *string `pulumi:"uri"`
	// The name of the database user for use with the connection pool. When excluded, all sessions connect to the database as the inbound user.
	User *string `pulumi:"user"`
}

type DatabaseConnectionPoolState struct {
	// The ID of the source database cluster. Note: This must be a PostgreSQL cluster.
	ClusterId pulumi.StringPtrInput
	// The database for use with the connection pool.
	DbName pulumi.StringPtrInput
	// The hostname used to connect to the database connection pool.
	Host pulumi.StringPtrInput
	// The PGBouncer transaction mode for the connection pool. The allowed values are session, transaction, and statement.
	Mode pulumi.StringPtrInput
	// The name for the database connection pool.
	Name pulumi.StringPtrInput
	// Password for the connection pool's user.
	Password pulumi.StringPtrInput
	// Network port that the database connection pool is listening on.
	Port pulumi.IntPtrInput
	// Same as `host`, but only accessible from resources within the account and in the same region.
	PrivateHost pulumi.StringPtrInput
	// Same as `uri`, but only accessible from resources within the account and in the same region.
	PrivateUri pulumi.StringPtrInput
	// The desired size of the PGBouncer connection pool.
	Size pulumi.IntPtrInput
	// The full URI for connecting to the database connection pool.
	Uri pulumi.StringPtrInput
	// The name of the database user for use with the connection pool. When excluded, all sessions connect to the database as the inbound user.
	User pulumi.StringPtrInput
}

func (DatabaseConnectionPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseConnectionPoolState)(nil)).Elem()
}

type databaseConnectionPoolArgs struct {
	// The ID of the source database cluster. Note: This must be a PostgreSQL cluster.
	ClusterId string `pulumi:"clusterId"`
	// The database for use with the connection pool.
	DbName string `pulumi:"dbName"`
	// The PGBouncer transaction mode for the connection pool. The allowed values are session, transaction, and statement.
	Mode string `pulumi:"mode"`
	// The name for the database connection pool.
	Name *string `pulumi:"name"`
	// The desired size of the PGBouncer connection pool.
	Size int `pulumi:"size"`
	// The name of the database user for use with the connection pool. When excluded, all sessions connect to the database as the inbound user.
	User *string `pulumi:"user"`
}

// The set of arguments for constructing a DatabaseConnectionPool resource.
type DatabaseConnectionPoolArgs struct {
	// The ID of the source database cluster. Note: This must be a PostgreSQL cluster.
	ClusterId pulumi.StringInput
	// The database for use with the connection pool.
	DbName pulumi.StringInput
	// The PGBouncer transaction mode for the connection pool. The allowed values are session, transaction, and statement.
	Mode pulumi.StringInput
	// The name for the database connection pool.
	Name pulumi.StringPtrInput
	// The desired size of the PGBouncer connection pool.
	Size pulumi.IntInput
	// The name of the database user for use with the connection pool. When excluded, all sessions connect to the database as the inbound user.
	User pulumi.StringPtrInput
}

func (DatabaseConnectionPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseConnectionPoolArgs)(nil)).Elem()
}

type DatabaseConnectionPoolInput interface {
	pulumi.Input

	ToDatabaseConnectionPoolOutput() DatabaseConnectionPoolOutput
	ToDatabaseConnectionPoolOutputWithContext(ctx context.Context) DatabaseConnectionPoolOutput
}

func (*DatabaseConnectionPool) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseConnectionPool)(nil)).Elem()
}

func (i *DatabaseConnectionPool) ToDatabaseConnectionPoolOutput() DatabaseConnectionPoolOutput {
	return i.ToDatabaseConnectionPoolOutputWithContext(context.Background())
}

func (i *DatabaseConnectionPool) ToDatabaseConnectionPoolOutputWithContext(ctx context.Context) DatabaseConnectionPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseConnectionPoolOutput)
}

// DatabaseConnectionPoolArrayInput is an input type that accepts DatabaseConnectionPoolArray and DatabaseConnectionPoolArrayOutput values.
// You can construct a concrete instance of `DatabaseConnectionPoolArrayInput` via:
//
//	DatabaseConnectionPoolArray{ DatabaseConnectionPoolArgs{...} }
type DatabaseConnectionPoolArrayInput interface {
	pulumi.Input

	ToDatabaseConnectionPoolArrayOutput() DatabaseConnectionPoolArrayOutput
	ToDatabaseConnectionPoolArrayOutputWithContext(context.Context) DatabaseConnectionPoolArrayOutput
}

type DatabaseConnectionPoolArray []DatabaseConnectionPoolInput

func (DatabaseConnectionPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseConnectionPool)(nil)).Elem()
}

func (i DatabaseConnectionPoolArray) ToDatabaseConnectionPoolArrayOutput() DatabaseConnectionPoolArrayOutput {
	return i.ToDatabaseConnectionPoolArrayOutputWithContext(context.Background())
}

func (i DatabaseConnectionPoolArray) ToDatabaseConnectionPoolArrayOutputWithContext(ctx context.Context) DatabaseConnectionPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseConnectionPoolArrayOutput)
}

// DatabaseConnectionPoolMapInput is an input type that accepts DatabaseConnectionPoolMap and DatabaseConnectionPoolMapOutput values.
// You can construct a concrete instance of `DatabaseConnectionPoolMapInput` via:
//
//	DatabaseConnectionPoolMap{ "key": DatabaseConnectionPoolArgs{...} }
type DatabaseConnectionPoolMapInput interface {
	pulumi.Input

	ToDatabaseConnectionPoolMapOutput() DatabaseConnectionPoolMapOutput
	ToDatabaseConnectionPoolMapOutputWithContext(context.Context) DatabaseConnectionPoolMapOutput
}

type DatabaseConnectionPoolMap map[string]DatabaseConnectionPoolInput

func (DatabaseConnectionPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseConnectionPool)(nil)).Elem()
}

func (i DatabaseConnectionPoolMap) ToDatabaseConnectionPoolMapOutput() DatabaseConnectionPoolMapOutput {
	return i.ToDatabaseConnectionPoolMapOutputWithContext(context.Background())
}

func (i DatabaseConnectionPoolMap) ToDatabaseConnectionPoolMapOutputWithContext(ctx context.Context) DatabaseConnectionPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseConnectionPoolMapOutput)
}

type DatabaseConnectionPoolOutput struct{ *pulumi.OutputState }

func (DatabaseConnectionPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseConnectionPool)(nil)).Elem()
}

func (o DatabaseConnectionPoolOutput) ToDatabaseConnectionPoolOutput() DatabaseConnectionPoolOutput {
	return o
}

func (o DatabaseConnectionPoolOutput) ToDatabaseConnectionPoolOutputWithContext(ctx context.Context) DatabaseConnectionPoolOutput {
	return o
}

// The ID of the source database cluster. Note: This must be a PostgreSQL cluster.
func (o DatabaseConnectionPoolOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseConnectionPool) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The database for use with the connection pool.
func (o DatabaseConnectionPoolOutput) DbName() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseConnectionPool) pulumi.StringOutput { return v.DbName }).(pulumi.StringOutput)
}

// The hostname used to connect to the database connection pool.
func (o DatabaseConnectionPoolOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseConnectionPool) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// The PGBouncer transaction mode for the connection pool. The allowed values are session, transaction, and statement.
func (o DatabaseConnectionPoolOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseConnectionPool) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

// The name for the database connection pool.
func (o DatabaseConnectionPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseConnectionPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Password for the connection pool's user.
func (o DatabaseConnectionPoolOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseConnectionPool) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// Network port that the database connection pool is listening on.
func (o DatabaseConnectionPoolOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *DatabaseConnectionPool) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// Same as `host`, but only accessible from resources within the account and in the same region.
func (o DatabaseConnectionPoolOutput) PrivateHost() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseConnectionPool) pulumi.StringOutput { return v.PrivateHost }).(pulumi.StringOutput)
}

// Same as `uri`, but only accessible from resources within the account and in the same region.
func (o DatabaseConnectionPoolOutput) PrivateUri() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseConnectionPool) pulumi.StringOutput { return v.PrivateUri }).(pulumi.StringOutput)
}

// The desired size of the PGBouncer connection pool.
func (o DatabaseConnectionPoolOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *DatabaseConnectionPool) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// The full URI for connecting to the database connection pool.
func (o DatabaseConnectionPoolOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseConnectionPool) pulumi.StringOutput { return v.Uri }).(pulumi.StringOutput)
}

// The name of the database user for use with the connection pool. When excluded, all sessions connect to the database as the inbound user.
func (o DatabaseConnectionPoolOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DatabaseConnectionPool) pulumi.StringPtrOutput { return v.User }).(pulumi.StringPtrOutput)
}

type DatabaseConnectionPoolArrayOutput struct{ *pulumi.OutputState }

func (DatabaseConnectionPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseConnectionPool)(nil)).Elem()
}

func (o DatabaseConnectionPoolArrayOutput) ToDatabaseConnectionPoolArrayOutput() DatabaseConnectionPoolArrayOutput {
	return o
}

func (o DatabaseConnectionPoolArrayOutput) ToDatabaseConnectionPoolArrayOutputWithContext(ctx context.Context) DatabaseConnectionPoolArrayOutput {
	return o
}

func (o DatabaseConnectionPoolArrayOutput) Index(i pulumi.IntInput) DatabaseConnectionPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatabaseConnectionPool {
		return vs[0].([]*DatabaseConnectionPool)[vs[1].(int)]
	}).(DatabaseConnectionPoolOutput)
}

type DatabaseConnectionPoolMapOutput struct{ *pulumi.OutputState }

func (DatabaseConnectionPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseConnectionPool)(nil)).Elem()
}

func (o DatabaseConnectionPoolMapOutput) ToDatabaseConnectionPoolMapOutput() DatabaseConnectionPoolMapOutput {
	return o
}

func (o DatabaseConnectionPoolMapOutput) ToDatabaseConnectionPoolMapOutputWithContext(ctx context.Context) DatabaseConnectionPoolMapOutput {
	return o
}

func (o DatabaseConnectionPoolMapOutput) MapIndex(k pulumi.StringInput) DatabaseConnectionPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatabaseConnectionPool {
		return vs[0].(map[string]*DatabaseConnectionPool)[vs[1].(string)]
	}).(DatabaseConnectionPoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseConnectionPoolInput)(nil)).Elem(), &DatabaseConnectionPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseConnectionPoolArrayInput)(nil)).Elem(), DatabaseConnectionPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseConnectionPoolMapInput)(nil)).Elem(), DatabaseConnectionPoolMap{})
	pulumi.RegisterOutputType(DatabaseConnectionPoolOutput{})
	pulumi.RegisterOutputType(DatabaseConnectionPoolArrayOutput{})
	pulumi.RegisterOutputType(DatabaseConnectionPoolMapOutput{})
}
