// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean
{
    public static class GetImages
    {
        public static Task<GetImagesResult> InvokeAsync(GetImagesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetImagesResult>("digitalocean:index/getImages:getImages", args ?? new GetImagesArgs(), options.WithVersion());
    }


    public sealed class GetImagesArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetImagesFilterArgs>? _filters;

        /// <summary>
        /// Filter the results.
        /// The `filter` block is documented below.
        /// </summary>
        public List<Inputs.GetImagesFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetImagesFilterArgs>());
            set => _filters = value;
        }

        [Input("sorts")]
        private List<Inputs.GetImagesSortArgs>? _sorts;

        /// <summary>
        /// Sort the results.
        /// The `sort` block is documented below.
        /// </summary>
        public List<Inputs.GetImagesSortArgs> Sorts
        {
            get => _sorts ?? (_sorts = new List<Inputs.GetImagesSortArgs>());
            set => _sorts = value;
        }

        public GetImagesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetImagesResult
    {
        public readonly ImmutableArray<Outputs.GetImagesFilterResult> Filters;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A set of images satisfying any `filter` and `sort` criteria. Each image has the following attributes:  
        /// - `slug`: Unique text identifier of the image.
        /// - `id`: The ID of the image.
        /// - `name`: The name of the image.
        /// - `type`: Type of the image.
        /// - `distribution` - The name of the distribution of the OS of the image.
        /// - `min_disk_size`: The minimum 'disk' required for the image.
        /// - `size_gigabytes`: The size of the image in GB.
        /// - `private` - Is image a public image or not. Public images represent
        /// Linux distributions or One-Click Applications, while non-public images represent
        /// snapshots and backups and are only available within your account.
        /// - `regions`: A set of the regions that the image is available in.
        /// - `tags`: A set of tags applied to the image
        /// - `created`: When the image was created
        /// - `status`: Current status of the image
        /// - `error_message`: Any applicable error message pertaining to the image
        /// - `image` - The id of the image (legacy parameter).
        /// </summary>
        public readonly ImmutableArray<Outputs.GetImagesImageResult> Images;
        public readonly ImmutableArray<Outputs.GetImagesSortResult> Sorts;

        [OutputConstructor]
        private GetImagesResult(
            ImmutableArray<Outputs.GetImagesFilterResult> filters,

            string id,

            ImmutableArray<Outputs.GetImagesImageResult> images,

            ImmutableArray<Outputs.GetImagesSortResult> sorts)
        {
            Filters = filters;
            Id = id;
            Images = images;
            Sorts = sorts;
        }
    }
}