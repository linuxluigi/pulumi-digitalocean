// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### Limiting access to specific IP addresses
//
// ```go
// package main
//
// import (
// 	"encoding/json"
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		foobarSpacesBucket, err := digitalocean.NewSpacesBucket(ctx, "foobarSpacesBucket", &digitalocean.SpacesBucketArgs{
// 			Region: pulumi.String("nyc3"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = digitalocean.NewSpacesBucketPolicy(ctx, "foobarSpacesBucketPolicy", &digitalocean.SpacesBucketPolicyArgs{
// 			Region: foobarSpacesBucket.Region,
// 			Bucket: foobarSpacesBucket.Name,
// 			Policy: pulumi.All(foobarSpacesBucket.Name, foobarSpacesBucket.Name).ApplyT(func(_args []interface{}) (string, error) {
// 				foobarSpacesBucketName := _args[0].(string)
// 				foobarSpacesBucketName1 := _args[1].(string)
// 				var _zero string
// 				tmpJSON0, err := json.Marshal(map[string]interface{}{
// 					"Version": "2012-10-17",
// 					"Statement": []map[string]interface{}{
// 						map[string]interface{}{
// 							"Sid":       "IPAllow",
// 							"Effect":    "Deny",
// 							"Principal": "*",
// 							"Action":    "s3:*",
// 							"Resource": []string{
// 								fmt.Sprintf("%v%v", "arn:aws:s3:::", foobarSpacesBucketName),
// 								fmt.Sprintf("%v%v%v", "arn:aws:s3:::", foobarSpacesBucketName1, "/*"),
// 							},
// 							"Condition": map[string]interface{}{
// 								"NotIpAddress": map[string]interface{}{
// 									"aws:SourceIp": "54.240.143.0/24",
// 								},
// 							},
// 						},
// 					},
// 				})
// 				if err != nil {
// 					return _zero, err
// 				}
// 				json0 := string(tmpJSON0)
// 				return json0, nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// !> **Warning:** Before using this policy, replace the 54.240.143.0/24 IP address range in this example with an appropriate value for your use case. Otherwise, you will lose the ability to access your bucket.
//
// ## Import
//
// Bucket policies can be imported using the `region` and `bucket` attributes (delimited by a comma)
//
// ```sh
//  $ pulumi import digitalocean:index/spacesBucketPolicy:SpacesBucketPolicy foobar `region`,`bucket`
// ```
type SpacesBucketPolicy struct {
	pulumi.CustomResourceState

	// The name of the bucket to which to apply the policy.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The text of the policy.
	Policy pulumi.StringOutput `pulumi:"policy"`
	// The region where the bucket resides.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewSpacesBucketPolicy registers a new resource with the given unique name, arguments, and options.
func NewSpacesBucketPolicy(ctx *pulumi.Context,
	name string, args *SpacesBucketPolicyArgs, opts ...pulumi.ResourceOption) (*SpacesBucketPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource SpacesBucketPolicy
	err := ctx.RegisterResource("digitalocean:index/spacesBucketPolicy:SpacesBucketPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpacesBucketPolicy gets an existing SpacesBucketPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpacesBucketPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpacesBucketPolicyState, opts ...pulumi.ResourceOption) (*SpacesBucketPolicy, error) {
	var resource SpacesBucketPolicy
	err := ctx.ReadResource("digitalocean:index/spacesBucketPolicy:SpacesBucketPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpacesBucketPolicy resources.
type spacesBucketPolicyState struct {
	// The name of the bucket to which to apply the policy.
	Bucket *string `pulumi:"bucket"`
	// The text of the policy.
	Policy *string `pulumi:"policy"`
	// The region where the bucket resides.
	Region *string `pulumi:"region"`
}

type SpacesBucketPolicyState struct {
	// The name of the bucket to which to apply the policy.
	Bucket pulumi.StringPtrInput
	// The text of the policy.
	Policy pulumi.StringPtrInput
	// The region where the bucket resides.
	Region pulumi.StringPtrInput
}

func (SpacesBucketPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*spacesBucketPolicyState)(nil)).Elem()
}

type spacesBucketPolicyArgs struct {
	// The name of the bucket to which to apply the policy.
	Bucket string `pulumi:"bucket"`
	// The text of the policy.
	Policy string `pulumi:"policy"`
	// The region where the bucket resides.
	Region string `pulumi:"region"`
}

// The set of arguments for constructing a SpacesBucketPolicy resource.
type SpacesBucketPolicyArgs struct {
	// The name of the bucket to which to apply the policy.
	Bucket pulumi.StringInput
	// The text of the policy.
	Policy pulumi.StringInput
	// The region where the bucket resides.
	Region pulumi.StringInput
}

func (SpacesBucketPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spacesBucketPolicyArgs)(nil)).Elem()
}

type SpacesBucketPolicyInput interface {
	pulumi.Input

	ToSpacesBucketPolicyOutput() SpacesBucketPolicyOutput
	ToSpacesBucketPolicyOutputWithContext(ctx context.Context) SpacesBucketPolicyOutput
}

func (*SpacesBucketPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SpacesBucketPolicy)(nil)).Elem()
}

func (i *SpacesBucketPolicy) ToSpacesBucketPolicyOutput() SpacesBucketPolicyOutput {
	return i.ToSpacesBucketPolicyOutputWithContext(context.Background())
}

func (i *SpacesBucketPolicy) ToSpacesBucketPolicyOutputWithContext(ctx context.Context) SpacesBucketPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpacesBucketPolicyOutput)
}

// SpacesBucketPolicyArrayInput is an input type that accepts SpacesBucketPolicyArray and SpacesBucketPolicyArrayOutput values.
// You can construct a concrete instance of `SpacesBucketPolicyArrayInput` via:
//
//          SpacesBucketPolicyArray{ SpacesBucketPolicyArgs{...} }
type SpacesBucketPolicyArrayInput interface {
	pulumi.Input

	ToSpacesBucketPolicyArrayOutput() SpacesBucketPolicyArrayOutput
	ToSpacesBucketPolicyArrayOutputWithContext(context.Context) SpacesBucketPolicyArrayOutput
}

type SpacesBucketPolicyArray []SpacesBucketPolicyInput

func (SpacesBucketPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SpacesBucketPolicy)(nil)).Elem()
}

func (i SpacesBucketPolicyArray) ToSpacesBucketPolicyArrayOutput() SpacesBucketPolicyArrayOutput {
	return i.ToSpacesBucketPolicyArrayOutputWithContext(context.Background())
}

func (i SpacesBucketPolicyArray) ToSpacesBucketPolicyArrayOutputWithContext(ctx context.Context) SpacesBucketPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpacesBucketPolicyArrayOutput)
}

// SpacesBucketPolicyMapInput is an input type that accepts SpacesBucketPolicyMap and SpacesBucketPolicyMapOutput values.
// You can construct a concrete instance of `SpacesBucketPolicyMapInput` via:
//
//          SpacesBucketPolicyMap{ "key": SpacesBucketPolicyArgs{...} }
type SpacesBucketPolicyMapInput interface {
	pulumi.Input

	ToSpacesBucketPolicyMapOutput() SpacesBucketPolicyMapOutput
	ToSpacesBucketPolicyMapOutputWithContext(context.Context) SpacesBucketPolicyMapOutput
}

type SpacesBucketPolicyMap map[string]SpacesBucketPolicyInput

func (SpacesBucketPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SpacesBucketPolicy)(nil)).Elem()
}

func (i SpacesBucketPolicyMap) ToSpacesBucketPolicyMapOutput() SpacesBucketPolicyMapOutput {
	return i.ToSpacesBucketPolicyMapOutputWithContext(context.Background())
}

func (i SpacesBucketPolicyMap) ToSpacesBucketPolicyMapOutputWithContext(ctx context.Context) SpacesBucketPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpacesBucketPolicyMapOutput)
}

type SpacesBucketPolicyOutput struct{ *pulumi.OutputState }

func (SpacesBucketPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SpacesBucketPolicy)(nil)).Elem()
}

func (o SpacesBucketPolicyOutput) ToSpacesBucketPolicyOutput() SpacesBucketPolicyOutput {
	return o
}

func (o SpacesBucketPolicyOutput) ToSpacesBucketPolicyOutputWithContext(ctx context.Context) SpacesBucketPolicyOutput {
	return o
}

// The name of the bucket to which to apply the policy.
func (o SpacesBucketPolicyOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *SpacesBucketPolicy) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// The text of the policy.
func (o SpacesBucketPolicyOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *SpacesBucketPolicy) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

// The region where the bucket resides.
func (o SpacesBucketPolicyOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *SpacesBucketPolicy) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type SpacesBucketPolicyArrayOutput struct{ *pulumi.OutputState }

func (SpacesBucketPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SpacesBucketPolicy)(nil)).Elem()
}

func (o SpacesBucketPolicyArrayOutput) ToSpacesBucketPolicyArrayOutput() SpacesBucketPolicyArrayOutput {
	return o
}

func (o SpacesBucketPolicyArrayOutput) ToSpacesBucketPolicyArrayOutputWithContext(ctx context.Context) SpacesBucketPolicyArrayOutput {
	return o
}

func (o SpacesBucketPolicyArrayOutput) Index(i pulumi.IntInput) SpacesBucketPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SpacesBucketPolicy {
		return vs[0].([]*SpacesBucketPolicy)[vs[1].(int)]
	}).(SpacesBucketPolicyOutput)
}

type SpacesBucketPolicyMapOutput struct{ *pulumi.OutputState }

func (SpacesBucketPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SpacesBucketPolicy)(nil)).Elem()
}

func (o SpacesBucketPolicyMapOutput) ToSpacesBucketPolicyMapOutput() SpacesBucketPolicyMapOutput {
	return o
}

func (o SpacesBucketPolicyMapOutput) ToSpacesBucketPolicyMapOutputWithContext(ctx context.Context) SpacesBucketPolicyMapOutput {
	return o
}

func (o SpacesBucketPolicyMapOutput) MapIndex(k pulumi.StringInput) SpacesBucketPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SpacesBucketPolicy {
		return vs[0].(map[string]*SpacesBucketPolicy)[vs[1].(string)]
	}).(SpacesBucketPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SpacesBucketPolicyInput)(nil)).Elem(), &SpacesBucketPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpacesBucketPolicyArrayInput)(nil)).Elem(), SpacesBucketPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SpacesBucketPolicyMapInput)(nil)).Elem(), SpacesBucketPolicyMap{})
	pulumi.RegisterOutputType(SpacesBucketPolicyOutput{})
	pulumi.RegisterOutputType(SpacesBucketPolicyArrayOutput{})
	pulumi.RegisterOutputType(SpacesBucketPolicyMapOutput{})
}
