// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean
{
    public static partial class Invokes
    {
        public static Task<GetImagesResult> GetImages(GetImagesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetImagesResult>("digitalocean:index/getImages:getImages", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetImagesArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetImagesFiltersArgs>? _filters;

        /// <summary>
        /// Filter the results.
        /// The `filter` block is documented below.
        /// </summary>
        public List<Inputs.GetImagesFiltersArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetImagesFiltersArgs>());
            set => _filters = value;
        }

        [Input("sorts")]
        private List<Inputs.GetImagesSortsArgs>? _sorts;

        /// <summary>
        /// Sort the results.
        /// The `sort` block is documented below.
        /// </summary>
        public List<Inputs.GetImagesSortsArgs> Sorts
        {
            get => _sorts ?? (_sorts = new List<Inputs.GetImagesSortsArgs>());
            set => _sorts = value;
        }

        public GetImagesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetImagesResult
    {
        public readonly ImmutableArray<Outputs.GetImagesFiltersResult> Filters;
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
        public readonly ImmutableArray<Outputs.GetImagesImagesResult> Images;
        public readonly ImmutableArray<Outputs.GetImagesSortsResult> Sorts;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetImagesResult(
            ImmutableArray<Outputs.GetImagesFiltersResult> filters,
            ImmutableArray<Outputs.GetImagesImagesResult> images,
            ImmutableArray<Outputs.GetImagesSortsResult> sorts,
            string id)
        {
            Filters = filters;
            Images = images;
            Sorts = sorts;
            Id = id;
        }
    }

    namespace Inputs
    {

    public sealed class GetImagesFiltersArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Sort the images by this key. This may be one of `distribution`, `error_message`, `id`,
        /// `image`, `min_disk_size`, `name`, `private`, `size_gigabytes`, `slug`, `status`, or `type`.
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        [Input("values", required: true)]
        private List<string>? _values;

        /// <summary>
        /// A list of values to match against the `key` field. Only retrieves images
        /// where the `key` field takes on one or more of the values provided here.
        /// </summary>
        public List<string> Values
        {
            get => _values ?? (_values = new List<string>());
            set => _values = value;
        }

        public GetImagesFiltersArgs()
        {
        }
    }

    public sealed class GetImagesSortsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The sort direction. This may be either `asc` or `desc`.
        /// </summary>
        [Input("direction")]
        public string? Direction { get; set; }

        /// <summary>
        /// Sort the images by this key. This may be one of `distribution`, `error_message`, `id`,
        /// `image`, `min_disk_size`, `name`, `private`, `size_gigabytes`, `slug`, `status`, or `type`.
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        public GetImagesSortsArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetImagesFiltersResult
    {
        /// <summary>
        /// Sort the images by this key. This may be one of `distribution`, `error_message`, `id`,
        /// `image`, `min_disk_size`, `name`, `private`, `size_gigabytes`, `slug`, `status`, or `type`.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// A list of values to match against the `key` field. Only retrieves images
        /// where the `key` field takes on one or more of the values provided here.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetImagesFiltersResult(
            string key,
            ImmutableArray<string> values)
        {
            Key = key;
            Values = values;
        }
    }

    [OutputType]
    public sealed class GetImagesImagesResult
    {
        public readonly string Created;
        public readonly string Distribution;
        public readonly string ErrorMessage;
        public readonly int Id;
        public readonly string Image;
        public readonly int MinDiskSize;
        public readonly string Name;
        public readonly bool Private;
        public readonly ImmutableArray<string> Regions;
        public readonly double SizeGigabytes;
        public readonly string Slug;
        public readonly string Status;
        public readonly ImmutableArray<string> Tags;
        public readonly string Type;

        [OutputConstructor]
        private GetImagesImagesResult(
            string created,
            string distribution,
            string errorMessage,
            int id,
            string image,
            int minDiskSize,
            string name,
            bool @private,
            ImmutableArray<string> regions,
            double sizeGigabytes,
            string slug,
            string status,
            ImmutableArray<string> tags,
            string type)
        {
            Created = created;
            Distribution = distribution;
            ErrorMessage = errorMessage;
            Id = id;
            Image = image;
            MinDiskSize = minDiskSize;
            Name = name;
            Private = @private;
            Regions = regions;
            SizeGigabytes = sizeGigabytes;
            Slug = slug;
            Status = status;
            Tags = tags;
            Type = type;
        }
    }

    [OutputType]
    public sealed class GetImagesSortsResult
    {
        /// <summary>
        /// The sort direction. This may be either `asc` or `desc`.
        /// </summary>
        public readonly string? Direction;
        /// <summary>
        /// Sort the images by this key. This may be one of `distribution`, `error_message`, `id`,
        /// `image`, `min_disk_size`, `name`, `private`, `size_gigabytes`, `slug`, `status`, or `type`.
        /// </summary>
        public readonly string Key;

        [OutputConstructor]
        private GetImagesSortsResult(
            string? direction,
            string key)
        {
            Direction = direction;
            Key = key;
        }
    }
    }
}
