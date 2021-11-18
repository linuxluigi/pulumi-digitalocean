// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.DigitalOcean
{
    public static class GetSshKeys
    {
        /// <summary>
        /// Get information on SSH Keys for use in other resources.
        /// 
        /// This data source is useful if the SSH Keys in question are not managed by the provider or you need to
        /// utilize any of the SSH Keys' data.
        /// 
        /// Note: You can use the `digitalocean.SshKey` data source to obtain metadata
        /// about a single SSH Key if you already know the unique `name` to retrieve.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// For example to find all SSH Keys:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var keys = Output.Create(DigitalOcean.GetSshKeys.InvokeAsync(new DigitalOcean.GetSshKeysArgs
        ///         {
        ///             Sorts = 
        ///             {
        ///                 new DigitalOcean.Inputs.GetSshKeysSortArgs
        ///                 {
        ///                     Direction = "asc",
        ///                     Key = "name",
        ///                 },
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSshKeysResult> InvokeAsync(GetSshKeysArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSshKeysResult>("digitalocean:index/getSshKeys:getSshKeys", args ?? new GetSshKeysArgs(), options.WithVersion());

        /// <summary>
        /// Get information on SSH Keys for use in other resources.
        /// 
        /// This data source is useful if the SSH Keys in question are not managed by the provider or you need to
        /// utilize any of the SSH Keys' data.
        /// 
        /// Note: You can use the `digitalocean.SshKey` data source to obtain metadata
        /// about a single SSH Key if you already know the unique `name` to retrieve.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// For example to find all SSH Keys:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var keys = Output.Create(DigitalOcean.GetSshKeys.InvokeAsync(new DigitalOcean.GetSshKeysArgs
        ///         {
        ///             Sorts = 
        ///             {
        ///                 new DigitalOcean.Inputs.GetSshKeysSortArgs
        ///                 {
        ///                     Direction = "asc",
        ///                     Key = "name",
        ///                 },
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSshKeysResult> Invoke(GetSshKeysInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSshKeysResult>("digitalocean:index/getSshKeys:getSshKeys", args ?? new GetSshKeysInvokeArgs(), options.WithVersion());
    }


    public sealed class GetSshKeysArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetSshKeysFilterArgs>? _filters;

        /// <summary>
        /// Filter the results.
        /// The `filter` block is documented below.
        /// </summary>
        public List<Inputs.GetSshKeysFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetSshKeysFilterArgs>());
            set => _filters = value;
        }

        [Input("sorts")]
        private List<Inputs.GetSshKeysSortArgs>? _sorts;

        /// <summary>
        /// Sort the results.
        /// The `sort` block is documented below.
        /// </summary>
        public List<Inputs.GetSshKeysSortArgs> Sorts
        {
            get => _sorts ?? (_sorts = new List<Inputs.GetSshKeysSortArgs>());
            set => _sorts = value;
        }

        public GetSshKeysArgs()
        {
        }
    }

    public sealed class GetSshKeysInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private InputList<Inputs.GetSshKeysFilterInputArgs>? _filters;

        /// <summary>
        /// Filter the results.
        /// The `filter` block is documented below.
        /// </summary>
        public InputList<Inputs.GetSshKeysFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetSshKeysFilterInputArgs>());
            set => _filters = value;
        }

        [Input("sorts")]
        private InputList<Inputs.GetSshKeysSortInputArgs>? _sorts;

        /// <summary>
        /// Sort the results.
        /// The `sort` block is documented below.
        /// </summary>
        public InputList<Inputs.GetSshKeysSortInputArgs> Sorts
        {
            get => _sorts ?? (_sorts = new InputList<Inputs.GetSshKeysSortInputArgs>());
            set => _sorts = value;
        }

        public GetSshKeysInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSshKeysResult
    {
        public readonly ImmutableArray<Outputs.GetSshKeysFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetSshKeysSortResult> Sorts;
        /// <summary>
        /// A list of SSH Keys. Each SSH Key has the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSshKeysSshKeyResult> SshKeys;

        [OutputConstructor]
        private GetSshKeysResult(
            ImmutableArray<Outputs.GetSshKeysFilterResult> filters,

            string id,

            ImmutableArray<Outputs.GetSshKeysSortResult> sorts,

            ImmutableArray<Outputs.GetSshKeysSshKeyResult> sshKeys)
        {
            Filters = filters;
            Id = id;
            Sorts = sorts;
            SshKeys = sshKeys;
        }
    }
}
