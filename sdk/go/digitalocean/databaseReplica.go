// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package digitalocean

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a DigitalOcean database replica resource.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/r/database_replica.html.markdown.
type DatabaseReplica struct {
	pulumi.CustomResourceState

	// The ID of the original source database cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Name of the replica's default database.
	Database pulumi.StringOutput `pulumi:"database"`
	// Database replica's hostname.
	Host pulumi.StringOutput `pulumi:"host"`
	// The name for the database replica.
	Name pulumi.StringOutput `pulumi:"name"`
	// Password for the replica's default user.
	Password pulumi.StringOutput `pulumi:"password"`
	// Network port that the database replica is listening on.
	Port pulumi.IntOutput `pulumi:"port"`
	// Same as `host`, but only accessible from resources within the account and in the same region.
	PrivateHost pulumi.StringOutput `pulumi:"privateHost"`
	// Same as `uri`, but only accessible from resources within the account and in the same region.
	PrivateUri pulumi.StringOutput `pulumi:"privateUri"`
	// DigitalOcean region where the replica will reside.
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// Database Droplet size associated with the replica (ex. `db-s-1vcpu-1gb`).
	Size pulumi.StringPtrOutput `pulumi:"size"`
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The full URI for connecting to the database replica.
	Uri pulumi.StringOutput `pulumi:"uri"`
	// Username for the replica's default user.
	User pulumi.StringOutput `pulumi:"user"`
}

// NewDatabaseReplica registers a new resource with the given unique name, arguments, and options.
func NewDatabaseReplica(ctx *pulumi.Context,
	name string, args *DatabaseReplicaArgs, opts ...pulumi.ResourceOption) (*DatabaseReplica, error) {
	if args == nil || args.ClusterId == nil {
		return nil, errors.New("missing required argument 'ClusterId'")
	}
	if args == nil {
		args = &DatabaseReplicaArgs{}
	}
	var resource DatabaseReplica
	err := ctx.RegisterResource("digitalocean:index/databaseReplica:DatabaseReplica", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseReplica gets an existing DatabaseReplica resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseReplica(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseReplicaState, opts ...pulumi.ResourceOption) (*DatabaseReplica, error) {
	var resource DatabaseReplica
	err := ctx.ReadResource("digitalocean:index/databaseReplica:DatabaseReplica", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseReplica resources.
type databaseReplicaState struct {
	// The ID of the original source database cluster.
	ClusterId *string `pulumi:"clusterId"`
	// Name of the replica's default database.
	Database *string `pulumi:"database"`
	// Database replica's hostname.
	Host *string `pulumi:"host"`
	// The name for the database replica.
	Name *string `pulumi:"name"`
	// Password for the replica's default user.
	Password *string `pulumi:"password"`
	// Network port that the database replica is listening on.
	Port *int `pulumi:"port"`
	// Same as `host`, but only accessible from resources within the account and in the same region.
	PrivateHost *string `pulumi:"privateHost"`
	// Same as `uri`, but only accessible from resources within the account and in the same region.
	PrivateUri *string `pulumi:"privateUri"`
	// DigitalOcean region where the replica will reside.
	Region *string `pulumi:"region"`
	// Database Droplet size associated with the replica (ex. `db-s-1vcpu-1gb`).
	Size *string `pulumi:"size"`
	Tags []string `pulumi:"tags"`
	// The full URI for connecting to the database replica.
	Uri *string `pulumi:"uri"`
	// Username for the replica's default user.
	User *string `pulumi:"user"`
}

type DatabaseReplicaState struct {
	// The ID of the original source database cluster.
	ClusterId pulumi.StringPtrInput
	// Name of the replica's default database.
	Database pulumi.StringPtrInput
	// Database replica's hostname.
	Host pulumi.StringPtrInput
	// The name for the database replica.
	Name pulumi.StringPtrInput
	// Password for the replica's default user.
	Password pulumi.StringPtrInput
	// Network port that the database replica is listening on.
	Port pulumi.IntPtrInput
	// Same as `host`, but only accessible from resources within the account and in the same region.
	PrivateHost pulumi.StringPtrInput
	// Same as `uri`, but only accessible from resources within the account and in the same region.
	PrivateUri pulumi.StringPtrInput
	// DigitalOcean region where the replica will reside.
	Region pulumi.StringPtrInput
	// Database Droplet size associated with the replica (ex. `db-s-1vcpu-1gb`).
	Size pulumi.StringPtrInput
	Tags pulumi.StringArrayInput
	// The full URI for connecting to the database replica.
	Uri pulumi.StringPtrInput
	// Username for the replica's default user.
	User pulumi.StringPtrInput
}

func (DatabaseReplicaState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseReplicaState)(nil)).Elem()
}

type databaseReplicaArgs struct {
	// The ID of the original source database cluster.
	ClusterId string `pulumi:"clusterId"`
	// The name for the database replica.
	Name *string `pulumi:"name"`
	// DigitalOcean region where the replica will reside.
	Region *string `pulumi:"region"`
	// Database Droplet size associated with the replica (ex. `db-s-1vcpu-1gb`).
	Size *string `pulumi:"size"`
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a DatabaseReplica resource.
type DatabaseReplicaArgs struct {
	// The ID of the original source database cluster.
	ClusterId pulumi.StringInput
	// The name for the database replica.
	Name pulumi.StringPtrInput
	// DigitalOcean region where the replica will reside.
	Region pulumi.StringPtrInput
	// Database Droplet size associated with the replica (ex. `db-s-1vcpu-1gb`).
	Size pulumi.StringPtrInput
	Tags pulumi.StringArrayInput
}

func (DatabaseReplicaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseReplicaArgs)(nil)).Elem()
}
