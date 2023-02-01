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
        /// <summary>
        /// Get information on images for use in other resources (e.g. creating a Droplet
        /// based on a snapshot), with the ability to filter and sort the results. If no filters are specified,
        /// all images will be returned.
        /// 
        /// This data source is useful if the image in question is not managed by the provider or you need to utilize any
        /// of the image's data.
        /// 
        /// Note: You can use the `digitalocean.getImage` data source to obtain metadata
        /// about a single image if you already know the `slug`, unique `name`, or `id` to retrieve.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Use the `filter` block with a `key` string and `values` list to filter images.
        /// 
        /// For example to find all Ubuntu images:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ubuntu = DigitalOcean.GetImages.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new DigitalOcean.Inputs.GetImagesFilterInputArgs
        ///             {
        ///                 Key = "distribution",
        ///                 Values = new[]
        ///                 {
        ///                     "Ubuntu",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// You can filter on multiple fields and sort the results as well:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var available = DigitalOcean.GetImages.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new DigitalOcean.Inputs.GetImagesFilterInputArgs
        ///             {
        ///                 Key = "distribution",
        ///                 Values = new[]
        ///                 {
        ///                     "Ubuntu",
        ///                 },
        ///             },
        ///             new DigitalOcean.Inputs.GetImagesFilterInputArgs
        ///             {
        ///                 Key = "regions",
        ///                 Values = new[]
        ///                 {
        ///                     "nyc3",
        ///                 },
        ///             },
        ///         },
        ///         Sorts = new[]
        ///         {
        ///             new DigitalOcean.Inputs.GetImagesSortInputArgs
        ///             {
        ///                 Direction = "desc",
        ///                 Key = "created",
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetImagesResult> InvokeAsync(GetImagesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetImagesResult>("digitalocean:index/getImages:getImages", args ?? new GetImagesArgs(), options.WithDefaults());

        /// <summary>
        /// Get information on images for use in other resources (e.g. creating a Droplet
        /// based on a snapshot), with the ability to filter and sort the results. If no filters are specified,
        /// all images will be returned.
        /// 
        /// This data source is useful if the image in question is not managed by the provider or you need to utilize any
        /// of the image's data.
        /// 
        /// Note: You can use the `digitalocean.getImage` data source to obtain metadata
        /// about a single image if you already know the `slug`, unique `name`, or `id` to retrieve.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Use the `filter` block with a `key` string and `values` list to filter images.
        /// 
        /// For example to find all Ubuntu images:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ubuntu = DigitalOcean.GetImages.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new DigitalOcean.Inputs.GetImagesFilterInputArgs
        ///             {
        ///                 Key = "distribution",
        ///                 Values = new[]
        ///                 {
        ///                     "Ubuntu",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// You can filter on multiple fields and sort the results as well:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var available = DigitalOcean.GetImages.Invoke(new()
        ///     {
        ///         Filters = new[]
        ///         {
        ///             new DigitalOcean.Inputs.GetImagesFilterInputArgs
        ///             {
        ///                 Key = "distribution",
        ///                 Values = new[]
        ///                 {
        ///                     "Ubuntu",
        ///                 },
        ///             },
        ///             new DigitalOcean.Inputs.GetImagesFilterInputArgs
        ///             {
        ///                 Key = "regions",
        ///                 Values = new[]
        ///                 {
        ///                     "nyc3",
        ///                 },
        ///             },
        ///         },
        ///         Sorts = new[]
        ///         {
        ///             new DigitalOcean.Inputs.GetImagesSortInputArgs
        ///             {
        ///                 Direction = "desc",
        ///                 Key = "created",
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetImagesResult> Invoke(GetImagesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetImagesResult>("digitalocean:index/getImages:getImages", args ?? new GetImagesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetImagesArgs : global::Pulumi.InvokeArgs
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
        public static new GetImagesArgs Empty => new GetImagesArgs();
    }

    public sealed class GetImagesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetImagesFilterInputArgs>? _filters;

        /// <summary>
        /// Filter the results.
        /// The `filter` block is documented below.
        /// </summary>
        public InputList<Inputs.GetImagesFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetImagesFilterInputArgs>());
            set => _filters = value;
        }

        [Input("sorts")]
        private InputList<Inputs.GetImagesSortInputArgs>? _sorts;

        /// <summary>
        /// Sort the results.
        /// The `sort` block is documented below.
        /// </summary>
        public InputList<Inputs.GetImagesSortInputArgs> Sorts
        {
            get => _sorts ?? (_sorts = new InputList<Inputs.GetImagesSortInputArgs>());
            set => _sorts = value;
        }

        public GetImagesInvokeArgs()
        {
        }
        public static new GetImagesInvokeArgs Empty => new GetImagesInvokeArgs();
    }


    [OutputType]
    public sealed class GetImagesResult
    {
        public readonly ImmutableArray<Outputs.GetImagesFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A set of images satisfying any `filter` and `sort` criteria. Each image has the following attributes:
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
