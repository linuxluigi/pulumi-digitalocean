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
    public static class GetDropletSnapshot
    {
        /// <summary>
        /// Droplet snapshots are saved instances of a Droplet. Use this data
        /// source to retrieve the ID of a DigitalOcean Droplet snapshot for use in other
        /// resources.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Get the Droplet snapshot:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var web_snapshot = Output.Create(DigitalOcean.GetDropletSnapshot.InvokeAsync(new DigitalOcean.GetDropletSnapshotArgs
        ///         {
        ///             MostRecent = true,
        ///             NameRegex = "^web",
        ///             Region = "nyc3",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDropletSnapshotResult> InvokeAsync(GetDropletSnapshotArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDropletSnapshotResult>("digitalocean:index/getDropletSnapshot:getDropletSnapshot", args ?? new GetDropletSnapshotArgs(), options.WithVersion());

        /// <summary>
        /// Droplet snapshots are saved instances of a Droplet. Use this data
        /// source to retrieve the ID of a DigitalOcean Droplet snapshot for use in other
        /// resources.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Get the Droplet snapshot:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var web_snapshot = Output.Create(DigitalOcean.GetDropletSnapshot.InvokeAsync(new DigitalOcean.GetDropletSnapshotArgs
        ///         {
        ///             MostRecent = true,
        ///             NameRegex = "^web",
        ///             Region = "nyc3",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDropletSnapshotResult> Invoke(GetDropletSnapshotInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDropletSnapshotResult>("digitalocean:index/getDropletSnapshot:getDropletSnapshot", args ?? new GetDropletSnapshotInvokeArgs(), options.WithVersion());
    }


    public sealed class GetDropletSnapshotArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// If more than one result is returned, use the most recent Droplet snapshot.
        /// </summary>
        [Input("mostRecent")]
        public bool? MostRecent { get; set; }

        /// <summary>
        /// The name of the Droplet snapshot.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// A regex string to apply to the Droplet snapshot list returned by DigitalOcean. This allows more advanced filtering not supported from the DigitalOcean API. This filtering is done locally on what DigitalOcean returns.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// A "slug" representing a DigitalOcean region (e.g. `nyc1`). If set, only Droplet snapshots available in the region will be returned.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetDropletSnapshotArgs()
        {
        }
    }

    public sealed class GetDropletSnapshotInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// If more than one result is returned, use the most recent Droplet snapshot.
        /// </summary>
        [Input("mostRecent")]
        public Input<bool>? MostRecent { get; set; }

        /// <summary>
        /// The name of the Droplet snapshot.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A regex string to apply to the Droplet snapshot list returned by DigitalOcean. This allows more advanced filtering not supported from the DigitalOcean API. This filtering is done locally on what DigitalOcean returns.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// A "slug" representing a DigitalOcean region (e.g. `nyc1`). If set, only Droplet snapshots available in the region will be returned.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetDropletSnapshotInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDropletSnapshotResult
    {
        /// <summary>
        /// The date and time the Droplet snapshot was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The ID of the Droplet from which the Droplet snapshot originated.
        /// </summary>
        public readonly string DropletId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The minimum size in gigabytes required for a Droplet to be created based on this Droplet snapshot.
        /// </summary>
        public readonly int MinDiskSize;
        public readonly bool? MostRecent;
        public readonly string? Name;
        public readonly string? NameRegex;
        public readonly string? Region;
        /// <summary>
        /// A list of DigitalOcean region "slugs" indicating where the Droplet snapshot is available.
        /// </summary>
        public readonly ImmutableArray<string> Regions;
        /// <summary>
        /// The billable size of the Droplet snapshot in gigabytes.
        /// </summary>
        public readonly double Size;

        [OutputConstructor]
        private GetDropletSnapshotResult(
            string createdAt,

            string dropletId,

            string id,

            int minDiskSize,

            bool? mostRecent,

            string? name,

            string? nameRegex,

            string? region,

            ImmutableArray<string> regions,

            double size)
        {
            CreatedAt = createdAt;
            DropletId = dropletId;
            Id = id;
            MinDiskSize = minDiskSize;
            MostRecent = mostRecent;
            Name = name;
            NameRegex = nameRegex;
            Region = region;
            Regions = regions;
            Size = size;
        }
    }
}
