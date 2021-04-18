// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get Docker credentials for your DigitalOcean container registry.
//
// An error is triggered if the provided container registry name does not exist.
//
// ## Example Usage
// ### Basic Example
//
// Get the container registry:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := digitalocean.NewContainerRegistryDockerCredentials(ctx, "example", &digitalocean.ContainerRegistryDockerCredentialsArgs{
// 			RegistryName: pulumi.String("example"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Docker Provider Example
//
// Use the `endpoint` and `dockerCredentials` with the Docker provider:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := digitalocean.LookupContainerRegistry(ctx, &digitalocean.LookupContainerRegistryArgs{
// 			Name: "example",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = digitalocean.NewContainerRegistryDockerCredentials(ctx, "exampleContainerRegistryDockerCredentials", &digitalocean.ContainerRegistryDockerCredentialsArgs{
// 			RegistryName: pulumi.String("example"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ContainerRegistryDockerCredentials struct {
	pulumi.CustomResourceState

	CredentialExpirationTime pulumi.StringOutput `pulumi:"credentialExpirationTime"`
	DockerCredentials        pulumi.StringOutput `pulumi:"dockerCredentials"`
	// The amount of time to pass before the Docker credentials expire in seconds. Defaults to 1576800000, or roughly 50 years. Must be greater than 0 and less than 1576800000.
	ExpirySeconds pulumi.IntPtrOutput `pulumi:"expirySeconds"`
	// The name of the container registry.
	RegistryName pulumi.StringOutput `pulumi:"registryName"`
	// Allow for write access to the container registry. Defaults to false.
	Write pulumi.BoolPtrOutput `pulumi:"write"`
}

// NewContainerRegistryDockerCredentials registers a new resource with the given unique name, arguments, and options.
func NewContainerRegistryDockerCredentials(ctx *pulumi.Context,
	name string, args *ContainerRegistryDockerCredentialsArgs, opts ...pulumi.ResourceOption) (*ContainerRegistryDockerCredentials, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	var resource ContainerRegistryDockerCredentials
	err := ctx.RegisterResource("digitalocean:index/containerRegistryDockerCredentials:ContainerRegistryDockerCredentials", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerRegistryDockerCredentials gets an existing ContainerRegistryDockerCredentials resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerRegistryDockerCredentials(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerRegistryDockerCredentialsState, opts ...pulumi.ResourceOption) (*ContainerRegistryDockerCredentials, error) {
	var resource ContainerRegistryDockerCredentials
	err := ctx.ReadResource("digitalocean:index/containerRegistryDockerCredentials:ContainerRegistryDockerCredentials", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerRegistryDockerCredentials resources.
type containerRegistryDockerCredentialsState struct {
	CredentialExpirationTime *string `pulumi:"credentialExpirationTime"`
	DockerCredentials        *string `pulumi:"dockerCredentials"`
	// The amount of time to pass before the Docker credentials expire in seconds. Defaults to 1576800000, or roughly 50 years. Must be greater than 0 and less than 1576800000.
	ExpirySeconds *int `pulumi:"expirySeconds"`
	// The name of the container registry.
	RegistryName *string `pulumi:"registryName"`
	// Allow for write access to the container registry. Defaults to false.
	Write *bool `pulumi:"write"`
}

type ContainerRegistryDockerCredentialsState struct {
	CredentialExpirationTime pulumi.StringPtrInput
	DockerCredentials        pulumi.StringPtrInput
	// The amount of time to pass before the Docker credentials expire in seconds. Defaults to 1576800000, or roughly 50 years. Must be greater than 0 and less than 1576800000.
	ExpirySeconds pulumi.IntPtrInput
	// The name of the container registry.
	RegistryName pulumi.StringPtrInput
	// Allow for write access to the container registry. Defaults to false.
	Write pulumi.BoolPtrInput
}

func (ContainerRegistryDockerCredentialsState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryDockerCredentialsState)(nil)).Elem()
}

type containerRegistryDockerCredentialsArgs struct {
	// The amount of time to pass before the Docker credentials expire in seconds. Defaults to 1576800000, or roughly 50 years. Must be greater than 0 and less than 1576800000.
	ExpirySeconds *int `pulumi:"expirySeconds"`
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// Allow for write access to the container registry. Defaults to false.
	Write *bool `pulumi:"write"`
}

// The set of arguments for constructing a ContainerRegistryDockerCredentials resource.
type ContainerRegistryDockerCredentialsArgs struct {
	// The amount of time to pass before the Docker credentials expire in seconds. Defaults to 1576800000, or roughly 50 years. Must be greater than 0 and less than 1576800000.
	ExpirySeconds pulumi.IntPtrInput
	// The name of the container registry.
	RegistryName pulumi.StringInput
	// Allow for write access to the container registry. Defaults to false.
	Write pulumi.BoolPtrInput
}

func (ContainerRegistryDockerCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryDockerCredentialsArgs)(nil)).Elem()
}

type ContainerRegistryDockerCredentialsInput interface {
	pulumi.Input

	ToContainerRegistryDockerCredentialsOutput() ContainerRegistryDockerCredentialsOutput
	ToContainerRegistryDockerCredentialsOutputWithContext(ctx context.Context) ContainerRegistryDockerCredentialsOutput
}

func (*ContainerRegistryDockerCredentials) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerRegistryDockerCredentials)(nil))
}

func (i *ContainerRegistryDockerCredentials) ToContainerRegistryDockerCredentialsOutput() ContainerRegistryDockerCredentialsOutput {
	return i.ToContainerRegistryDockerCredentialsOutputWithContext(context.Background())
}

func (i *ContainerRegistryDockerCredentials) ToContainerRegistryDockerCredentialsOutputWithContext(ctx context.Context) ContainerRegistryDockerCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryDockerCredentialsOutput)
}

func (i *ContainerRegistryDockerCredentials) ToContainerRegistryDockerCredentialsPtrOutput() ContainerRegistryDockerCredentialsPtrOutput {
	return i.ToContainerRegistryDockerCredentialsPtrOutputWithContext(context.Background())
}

func (i *ContainerRegistryDockerCredentials) ToContainerRegistryDockerCredentialsPtrOutputWithContext(ctx context.Context) ContainerRegistryDockerCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryDockerCredentialsPtrOutput)
}

type ContainerRegistryDockerCredentialsPtrInput interface {
	pulumi.Input

	ToContainerRegistryDockerCredentialsPtrOutput() ContainerRegistryDockerCredentialsPtrOutput
	ToContainerRegistryDockerCredentialsPtrOutputWithContext(ctx context.Context) ContainerRegistryDockerCredentialsPtrOutput
}

type containerRegistryDockerCredentialsPtrType ContainerRegistryDockerCredentialsArgs

func (*containerRegistryDockerCredentialsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistryDockerCredentials)(nil))
}

func (i *containerRegistryDockerCredentialsPtrType) ToContainerRegistryDockerCredentialsPtrOutput() ContainerRegistryDockerCredentialsPtrOutput {
	return i.ToContainerRegistryDockerCredentialsPtrOutputWithContext(context.Background())
}

func (i *containerRegistryDockerCredentialsPtrType) ToContainerRegistryDockerCredentialsPtrOutputWithContext(ctx context.Context) ContainerRegistryDockerCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryDockerCredentialsPtrOutput)
}

// ContainerRegistryDockerCredentialsArrayInput is an input type that accepts ContainerRegistryDockerCredentialsArray and ContainerRegistryDockerCredentialsArrayOutput values.
// You can construct a concrete instance of `ContainerRegistryDockerCredentialsArrayInput` via:
//
//          ContainerRegistryDockerCredentialsArray{ ContainerRegistryDockerCredentialsArgs{...} }
type ContainerRegistryDockerCredentialsArrayInput interface {
	pulumi.Input

	ToContainerRegistryDockerCredentialsArrayOutput() ContainerRegistryDockerCredentialsArrayOutput
	ToContainerRegistryDockerCredentialsArrayOutputWithContext(context.Context) ContainerRegistryDockerCredentialsArrayOutput
}

type ContainerRegistryDockerCredentialsArray []ContainerRegistryDockerCredentialsInput

func (ContainerRegistryDockerCredentialsArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ContainerRegistryDockerCredentials)(nil))
}

func (i ContainerRegistryDockerCredentialsArray) ToContainerRegistryDockerCredentialsArrayOutput() ContainerRegistryDockerCredentialsArrayOutput {
	return i.ToContainerRegistryDockerCredentialsArrayOutputWithContext(context.Background())
}

func (i ContainerRegistryDockerCredentialsArray) ToContainerRegistryDockerCredentialsArrayOutputWithContext(ctx context.Context) ContainerRegistryDockerCredentialsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryDockerCredentialsArrayOutput)
}

// ContainerRegistryDockerCredentialsMapInput is an input type that accepts ContainerRegistryDockerCredentialsMap and ContainerRegistryDockerCredentialsMapOutput values.
// You can construct a concrete instance of `ContainerRegistryDockerCredentialsMapInput` via:
//
//          ContainerRegistryDockerCredentialsMap{ "key": ContainerRegistryDockerCredentialsArgs{...} }
type ContainerRegistryDockerCredentialsMapInput interface {
	pulumi.Input

	ToContainerRegistryDockerCredentialsMapOutput() ContainerRegistryDockerCredentialsMapOutput
	ToContainerRegistryDockerCredentialsMapOutputWithContext(context.Context) ContainerRegistryDockerCredentialsMapOutput
}

type ContainerRegistryDockerCredentialsMap map[string]ContainerRegistryDockerCredentialsInput

func (ContainerRegistryDockerCredentialsMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ContainerRegistryDockerCredentials)(nil))
}

func (i ContainerRegistryDockerCredentialsMap) ToContainerRegistryDockerCredentialsMapOutput() ContainerRegistryDockerCredentialsMapOutput {
	return i.ToContainerRegistryDockerCredentialsMapOutputWithContext(context.Background())
}

func (i ContainerRegistryDockerCredentialsMap) ToContainerRegistryDockerCredentialsMapOutputWithContext(ctx context.Context) ContainerRegistryDockerCredentialsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryDockerCredentialsMapOutput)
}

type ContainerRegistryDockerCredentialsOutput struct {
	*pulumi.OutputState
}

func (ContainerRegistryDockerCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerRegistryDockerCredentials)(nil))
}

func (o ContainerRegistryDockerCredentialsOutput) ToContainerRegistryDockerCredentialsOutput() ContainerRegistryDockerCredentialsOutput {
	return o
}

func (o ContainerRegistryDockerCredentialsOutput) ToContainerRegistryDockerCredentialsOutputWithContext(ctx context.Context) ContainerRegistryDockerCredentialsOutput {
	return o
}

func (o ContainerRegistryDockerCredentialsOutput) ToContainerRegistryDockerCredentialsPtrOutput() ContainerRegistryDockerCredentialsPtrOutput {
	return o.ToContainerRegistryDockerCredentialsPtrOutputWithContext(context.Background())
}

func (o ContainerRegistryDockerCredentialsOutput) ToContainerRegistryDockerCredentialsPtrOutputWithContext(ctx context.Context) ContainerRegistryDockerCredentialsPtrOutput {
	return o.ApplyT(func(v ContainerRegistryDockerCredentials) *ContainerRegistryDockerCredentials {
		return &v
	}).(ContainerRegistryDockerCredentialsPtrOutput)
}

type ContainerRegistryDockerCredentialsPtrOutput struct {
	*pulumi.OutputState
}

func (ContainerRegistryDockerCredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistryDockerCredentials)(nil))
}

func (o ContainerRegistryDockerCredentialsPtrOutput) ToContainerRegistryDockerCredentialsPtrOutput() ContainerRegistryDockerCredentialsPtrOutput {
	return o
}

func (o ContainerRegistryDockerCredentialsPtrOutput) ToContainerRegistryDockerCredentialsPtrOutputWithContext(ctx context.Context) ContainerRegistryDockerCredentialsPtrOutput {
	return o
}

type ContainerRegistryDockerCredentialsArrayOutput struct{ *pulumi.OutputState }

func (ContainerRegistryDockerCredentialsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerRegistryDockerCredentials)(nil))
}

func (o ContainerRegistryDockerCredentialsArrayOutput) ToContainerRegistryDockerCredentialsArrayOutput() ContainerRegistryDockerCredentialsArrayOutput {
	return o
}

func (o ContainerRegistryDockerCredentialsArrayOutput) ToContainerRegistryDockerCredentialsArrayOutputWithContext(ctx context.Context) ContainerRegistryDockerCredentialsArrayOutput {
	return o
}

func (o ContainerRegistryDockerCredentialsArrayOutput) Index(i pulumi.IntInput) ContainerRegistryDockerCredentialsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerRegistryDockerCredentials {
		return vs[0].([]ContainerRegistryDockerCredentials)[vs[1].(int)]
	}).(ContainerRegistryDockerCredentialsOutput)
}

type ContainerRegistryDockerCredentialsMapOutput struct{ *pulumi.OutputState }

func (ContainerRegistryDockerCredentialsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ContainerRegistryDockerCredentials)(nil))
}

func (o ContainerRegistryDockerCredentialsMapOutput) ToContainerRegistryDockerCredentialsMapOutput() ContainerRegistryDockerCredentialsMapOutput {
	return o
}

func (o ContainerRegistryDockerCredentialsMapOutput) ToContainerRegistryDockerCredentialsMapOutputWithContext(ctx context.Context) ContainerRegistryDockerCredentialsMapOutput {
	return o
}

func (o ContainerRegistryDockerCredentialsMapOutput) MapIndex(k pulumi.StringInput) ContainerRegistryDockerCredentialsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ContainerRegistryDockerCredentials {
		return vs[0].(map[string]ContainerRegistryDockerCredentials)[vs[1].(string)]
	}).(ContainerRegistryDockerCredentialsOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerRegistryDockerCredentialsOutput{})
	pulumi.RegisterOutputType(ContainerRegistryDockerCredentialsPtrOutput{})
	pulumi.RegisterOutputType(ContainerRegistryDockerCredentialsArrayOutput{})
	pulumi.RegisterOutputType(ContainerRegistryDockerCredentialsMapOutput{})
}
