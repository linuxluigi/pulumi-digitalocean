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
    public static class GetDroplet
    {
        /// <summary>
        /// Get information on a Droplet for use in other resources. This data source provides
        /// all of the Droplet's properties as configured on your DigitalOcean account. This
        /// is useful if the Droplet in question is not managed by this provider or you need to
        /// utilize any of the Droplet's data.
        /// 
        /// **Note:** This data source returns a single Droplet. When specifying a `tag`, an
        /// error is triggered if more than one Droplet is found.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Get the Droplet by name:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(DigitalOcean.GetDroplet.InvokeAsync(new DigitalOcean.GetDropletArgs
        ///         {
        ///             Name = "web",
        ///         }));
        ///         this.DropletOutput = example.Apply(example =&gt; example.Ipv4Address);
        ///     }
        /// 
        ///     [Output("dropletOutput")]
        ///     public Output&lt;string&gt; DropletOutput { get; set; }
        /// }
        /// ```
        /// 
        /// Get the Droplet by tag:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(DigitalOcean.GetDroplet.InvokeAsync(new DigitalOcean.GetDropletArgs
        ///         {
        ///             Tag = "web",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// Get the Droplet by ID:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(DigitalOcean.GetDroplet.InvokeAsync(new DigitalOcean.GetDropletArgs
        ///         {
        ///             Id = digitalocean_kubernetes_cluster.Example.Node_pool[0].Nodes[0].Droplet_id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDropletResult> InvokeAsync(GetDropletArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDropletResult>("digitalocean:index/getDroplet:getDroplet", args ?? new GetDropletArgs(), options.WithVersion());

        /// <summary>
        /// Get information on a Droplet for use in other resources. This data source provides
        /// all of the Droplet's properties as configured on your DigitalOcean account. This
        /// is useful if the Droplet in question is not managed by this provider or you need to
        /// utilize any of the Droplet's data.
        /// 
        /// **Note:** This data source returns a single Droplet. When specifying a `tag`, an
        /// error is triggered if more than one Droplet is found.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Get the Droplet by name:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(DigitalOcean.GetDroplet.InvokeAsync(new DigitalOcean.GetDropletArgs
        ///         {
        ///             Name = "web",
        ///         }));
        ///         this.DropletOutput = example.Apply(example =&gt; example.Ipv4Address);
        ///     }
        /// 
        ///     [Output("dropletOutput")]
        ///     public Output&lt;string&gt; DropletOutput { get; set; }
        /// }
        /// ```
        /// 
        /// Get the Droplet by tag:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(DigitalOcean.GetDroplet.InvokeAsync(new DigitalOcean.GetDropletArgs
        ///         {
        ///             Tag = "web",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// Get the Droplet by ID:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(DigitalOcean.GetDroplet.InvokeAsync(new DigitalOcean.GetDropletArgs
        ///         {
        ///             Id = digitalocean_kubernetes_cluster.Example.Node_pool[0].Nodes[0].Droplet_id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDropletResult> Invoke(GetDropletInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDropletResult>("digitalocean:index/getDroplet:getDroplet", args ?? new GetDropletInvokeArgs(), options.WithVersion());
    }


    public sealed class GetDropletArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Droplet
        /// </summary>
        [Input("id")]
        public int? Id { get; set; }

        /// <summary>
        /// The name of the Droplet.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// A tag applied to the Droplet.
        /// </summary>
        [Input("tag")]
        public string? Tag { get; set; }

        public GetDropletArgs()
        {
        }
    }

    public sealed class GetDropletInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Droplet
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// The name of the Droplet.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A tag applied to the Droplet.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        public GetDropletInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDropletResult
    {
        /// <summary>
        /// Whether backups are enabled.
        /// </summary>
        public readonly bool Backups;
        public readonly string CreatedAt;
        /// <summary>
        /// The size of the Droplets disk in GB.
        /// </summary>
        public readonly int Disk;
        public readonly int Id;
        /// <summary>
        /// The Droplet image ID or slug.
        /// </summary>
        public readonly string Image;
        /// <summary>
        /// The Droplets public IPv4 address
        /// </summary>
        public readonly string Ipv4Address;
        /// <summary>
        /// The Droplets private IPv4 address
        /// </summary>
        public readonly string Ipv4AddressPrivate;
        /// <summary>
        /// Whether IPv6 is enabled.
        /// </summary>
        public readonly bool Ipv6;
        /// <summary>
        /// The Droplets public IPv6 address
        /// </summary>
        public readonly string Ipv6Address;
        /// <summary>
        /// The Droplets private IPv6 address
        /// </summary>
        public readonly string Ipv6AddressPrivate;
        /// <summary>
        /// Whether the Droplet is locked.
        /// </summary>
        public readonly bool Locked;
        /// <summary>
        /// The amount of the Droplets memory in MB.
        /// </summary>
        public readonly int Memory;
        /// <summary>
        /// Whether monitoring agent is installed.
        /// </summary>
        public readonly bool Monitoring;
        public readonly string Name;
        /// <summary>
        /// Droplet hourly price.
        /// </summary>
        public readonly double PriceHourly;
        /// <summary>
        /// Droplet monthly price.
        /// </summary>
        public readonly double PriceMonthly;
        /// <summary>
        /// Whether private networks are enabled.
        /// </summary>
        public readonly bool PrivateNetworking;
        /// <summary>
        /// The region the Droplet is running in.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The unique slug that indentifies the type of Droplet.
        /// </summary>
        public readonly string Size;
        /// <summary>
        /// The status of the Droplet.
        /// </summary>
        public readonly string Status;
        public readonly string? Tag;
        /// <summary>
        /// A list of the tags associated to the Droplet.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The uniform resource name of the Droplet
        /// </summary>
        public readonly string Urn;
        /// <summary>
        /// The number of the Droplets virtual CPUs.
        /// </summary>
        public readonly int Vcpus;
        /// <summary>
        /// List of the IDs of each volumes attached to the Droplet.
        /// </summary>
        public readonly ImmutableArray<string> VolumeIds;
        /// <summary>
        /// The ID of the VPC where the Droplet is located.
        /// </summary>
        public readonly string VpcUuid;

        [OutputConstructor]
        private GetDropletResult(
            bool backups,

            string createdAt,

            int disk,

            int id,

            string image,

            string ipv4Address,

            string ipv4AddressPrivate,

            bool ipv6,

            string ipv6Address,

            string ipv6AddressPrivate,

            bool locked,

            int memory,

            bool monitoring,

            string name,

            double priceHourly,

            double priceMonthly,

            bool privateNetworking,

            string region,

            string size,

            string status,

            string? tag,

            ImmutableArray<string> tags,

            string urn,

            int vcpus,

            ImmutableArray<string> volumeIds,

            string vpcUuid)
        {
            Backups = backups;
            CreatedAt = createdAt;
            Disk = disk;
            Id = id;
            Image = image;
            Ipv4Address = ipv4Address;
            Ipv4AddressPrivate = ipv4AddressPrivate;
            Ipv6 = ipv6;
            Ipv6Address = ipv6Address;
            Ipv6AddressPrivate = ipv6AddressPrivate;
            Locked = locked;
            Memory = memory;
            Monitoring = monitoring;
            Name = name;
            PriceHourly = priceHourly;
            PriceMonthly = priceMonthly;
            PrivateNetworking = privateNetworking;
            Region = region;
            Size = size;
            Status = status;
            Tag = tag;
            Tags = tags;
            Urn = urn;
            Vcpus = vcpus;
            VolumeIds = volumeIds;
            VpcUuid = vpcUuid;
        }
    }
}
