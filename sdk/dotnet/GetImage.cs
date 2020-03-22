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
        public static Task<GetImageResult> GetImage(GetImageArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetImageResult>("digitalocean:index/getImage:getImage", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetImageArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the image
        /// </summary>
        [Input("id")]
        public int? Id { get; set; }

        /// <summary>
        /// The name of the image.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The slug of the official image.
        /// </summary>
        [Input("slug")]
        public string? Slug { get; set; }

        /// <summary>
        /// Restrict the search to one of the following categories of images:
        /// - `all` - All images (whether public or private)
        /// - `applications` - One-click applications
        /// - `distributions` - Distributions
        /// - `user` - (Default) User (private) images
        /// </summary>
        [Input("source")]
        public string? Source { get; set; }

        public GetImageArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetImageResult
    {
        public readonly string Created;
        /// <summary>
        /// The name of the distribution of the OS of the image.
        /// * `min_disk_size`: The minimum 'disk' required for the image.
        /// * `size_gigabytes`: The size of the image in GB.
        /// </summary>
        public readonly string Distribution;
        public readonly string ErrorMessage;
        public readonly int Id;
        /// <summary>
        /// The id of the image (legacy parameter).
        /// </summary>
        public readonly string Image;
        public readonly int MinDiskSize;
        public readonly string Name;
        /// <summary>
        /// Is image a public image or not. Public images represent
        /// Linux distributions or One-Click Applications, while non-public images represent
        /// snapshots and backups and are only available within your account.
        /// * `regions`: A set of the regions that the image is available in.
        /// * `tags`: A set of tags applied to the image
        /// * `created`: When the image was created
        /// * `status`: Current status of the image
        /// * `error_message`: Any applicable error message pertaining to the image
        /// </summary>
        public readonly bool Private;
        public readonly ImmutableArray<string> Regions;
        public readonly double SizeGigabytes;
        public readonly string Slug;
        public readonly string? Source;
        public readonly string Status;
        public readonly ImmutableArray<string> Tags;
        public readonly string Type;

        [OutputConstructor]
        private GetImageResult(
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
            string? source,
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
            Source = source;
            Status = status;
            Tags = tags;
            Type = type;
        }
    }
}
