// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Get information on an images for use in other resources (e.g. creating a Droplet
// based on snapshot). This data source provides all of the image properties as
// configured on your DigitalOcean account. This is useful if the image in question
// is not managed by Terraform or you need to utilize any of the image's data.
// 
// An error is triggered if zero or more than one result is returned by the query.
func LookupImage(ctx *pulumi.Context, args *GetImageArgs) (*GetImageResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["slug"] = args.Slug
	}
	outputs, err := ctx.Invoke("digitalocean:index/getImage:getImage", inputs)
	if err != nil {
		return nil, err
	}
	return &GetImageResult{
		Distribution: outputs["distribution"],
		Image: outputs["image"],
		MinDiskSize: outputs["minDiskSize"],
		Name: outputs["name"],
		Private: outputs["private"],
		Regions: outputs["regions"],
		Slug: outputs["slug"],
		Type: outputs["type"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getImage.
type GetImageArgs struct {
	// The name of the private image.
	Name interface{}
	// The slug of the official image.
	Slug interface{}
}

// A collection of values returned by getImage.
type GetImageResult struct {
	// The name of the distribution of the OS of the image.
	// * `min_disk_size`: The minimum 'disk' required for the image.
	Distribution interface{}
	// The id of the image.
	Image interface{}
	MinDiskSize interface{}
	Name interface{}
	// Is image a public image or not. Public images represent
	// Linux distributions or One-Click Applications, while non-public images represent
	// snapshots and backups and are only available within your account.
	// * `regions`: The regions that the image is available in.
	// * `type`: Type of the image.
	Private interface{}
	Regions interface{}
	Slug interface{}
	Type interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}