// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean
{
    public static class GetTags
    {
        /// <summary>
        /// Returns a list of tags in your DigitalOcean account, with the ability to
        /// filter and sort the results. If no filters are specified, all tags will be
        /// returned.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var list = DigitalOcean.GetTags.Invoke(new()
        ///     {
        ///         Sorts = new[]
        ///         {
        ///             new DigitalOcean.Inputs.GetTagsSortInputArgs
        ///             {
        ///                 Key = "total_resource_count",
        ///                 Direction = "asc",
        ///             },
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["sortedTags"] = list.Apply(getTagsResult =&gt; getTagsResult.Tags),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTagsResult> InvokeAsync(GetTagsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTagsResult>("digitalocean:index/getTags:getTags", args ?? new GetTagsArgs(), options.WithDefaults());

        /// <summary>
        /// Returns a list of tags in your DigitalOcean account, with the ability to
        /// filter and sort the results. If no filters are specified, all tags will be
        /// returned.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var list = DigitalOcean.GetTags.Invoke(new()
        ///     {
        ///         Sorts = new[]
        ///         {
        ///             new DigitalOcean.Inputs.GetTagsSortInputArgs
        ///             {
        ///                 Key = "total_resource_count",
        ///                 Direction = "asc",
        ///             },
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["sortedTags"] = list.Apply(getTagsResult =&gt; getTagsResult.Tags),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetTagsResult> Invoke(GetTagsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTagsResult>("digitalocean:index/getTags:getTags", args ?? new GetTagsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTagsArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetTagsFilterArgs>? _filters;

        /// <summary>
        /// Filter the results.
        /// The `filter` block is documented below.
        /// </summary>
        public List<Inputs.GetTagsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetTagsFilterArgs>());
            set => _filters = value;
        }

        [Input("sorts")]
        private List<Inputs.GetTagsSortArgs>? _sorts;

        /// <summary>
        /// Sort the results.
        /// The `sort` block is documented below.
        /// </summary>
        public List<Inputs.GetTagsSortArgs> Sorts
        {
            get => _sorts ?? (_sorts = new List<Inputs.GetTagsSortArgs>());
            set => _sorts = value;
        }

        public GetTagsArgs()
        {
        }
        public static new GetTagsArgs Empty => new GetTagsArgs();
    }

    public sealed class GetTagsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetTagsFilterInputArgs>? _filters;

        /// <summary>
        /// Filter the results.
        /// The `filter` block is documented below.
        /// </summary>
        public InputList<Inputs.GetTagsFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetTagsFilterInputArgs>());
            set => _filters = value;
        }

        [Input("sorts")]
        private InputList<Inputs.GetTagsSortInputArgs>? _sorts;

        /// <summary>
        /// Sort the results.
        /// The `sort` block is documented below.
        /// </summary>
        public InputList<Inputs.GetTagsSortInputArgs> Sorts
        {
            get => _sorts ?? (_sorts = new InputList<Inputs.GetTagsSortInputArgs>());
            set => _sorts = value;
        }

        public GetTagsInvokeArgs()
        {
        }
        public static new GetTagsInvokeArgs Empty => new GetTagsInvokeArgs();
    }


    [OutputType]
    public sealed class GetTagsResult
    {
        public readonly ImmutableArray<Outputs.GetTagsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetTagsSortResult> Sorts;
        public readonly ImmutableArray<Outputs.GetTagsTagResult> Tags;

        [OutputConstructor]
        private GetTagsResult(
            ImmutableArray<Outputs.GetTagsFilterResult> filters,

            string id,

            ImmutableArray<Outputs.GetTagsSortResult> sorts,

            ImmutableArray<Outputs.GetTagsTagResult> tags)
        {
            Filters = filters;
            Id = id;
            Sorts = sorts;
            Tags = tags;
        }
    }
}
