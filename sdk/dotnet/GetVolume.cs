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
    public static class GetVolume
    {
        /// <summary>
        /// Get information on a volume for use in other resources. This data source provides
        /// all of the volumes properties as configured on your DigitalOcean account. This is
        /// useful if the volume in question is not managed by the provider or you need to utilize
        /// any of the volumes data.
        /// 
        /// An error is triggered if the provided volume name does not exist.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Get the volume:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(DigitalOcean.GetVolume.InvokeAsync(new DigitalOcean.GetVolumeArgs
        ///         {
        ///             Name = "app-data",
        ///             Region = "nyc3",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// Reuse the data about a volume to attach it to a Droplet:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var exampleVolume = Output.Create(DigitalOcean.GetVolume.InvokeAsync(new DigitalOcean.GetVolumeArgs
        ///         {
        ///             Name = "app-data",
        ///             Region = "nyc3",
        ///         }));
        ///         var exampleDroplet = new DigitalOcean.Droplet("exampleDroplet", new DigitalOcean.DropletArgs
        ///         {
        ///             Size = "s-1vcpu-1gb",
        ///             Image = "ubuntu-18-04-x64",
        ///             Region = "nyc3",
        ///         });
        ///         var foobar = new DigitalOcean.VolumeAttachment("foobar", new DigitalOcean.VolumeAttachmentArgs
        ///         {
        ///             DropletId = exampleDroplet.Id,
        ///             VolumeId = exampleVolume.Apply(exampleVolume =&gt; exampleVolume.Id),
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVolumeResult> InvokeAsync(GetVolumeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVolumeResult>("digitalocean:index/getVolume:getVolume", args ?? new GetVolumeArgs(), options.WithVersion());

        /// <summary>
        /// Get information on a volume for use in other resources. This data source provides
        /// all of the volumes properties as configured on your DigitalOcean account. This is
        /// useful if the volume in question is not managed by the provider or you need to utilize
        /// any of the volumes data.
        /// 
        /// An error is triggered if the provided volume name does not exist.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Get the volume:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(DigitalOcean.GetVolume.InvokeAsync(new DigitalOcean.GetVolumeArgs
        ///         {
        ///             Name = "app-data",
        ///             Region = "nyc3",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// Reuse the data about a volume to attach it to a Droplet:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var exampleVolume = Output.Create(DigitalOcean.GetVolume.InvokeAsync(new DigitalOcean.GetVolumeArgs
        ///         {
        ///             Name = "app-data",
        ///             Region = "nyc3",
        ///         }));
        ///         var exampleDroplet = new DigitalOcean.Droplet("exampleDroplet", new DigitalOcean.DropletArgs
        ///         {
        ///             Size = "s-1vcpu-1gb",
        ///             Image = "ubuntu-18-04-x64",
        ///             Region = "nyc3",
        ///         });
        ///         var foobar = new DigitalOcean.VolumeAttachment("foobar", new DigitalOcean.VolumeAttachmentArgs
        ///         {
        ///             DropletId = exampleDroplet.Id,
        ///             VolumeId = exampleVolume.Apply(exampleVolume =&gt; exampleVolume.Id),
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetVolumeResult> Invoke(GetVolumeInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetVolumeResult>("digitalocean:index/getVolume:getVolume", args ?? new GetVolumeInvokeArgs(), options.WithVersion());
    }


    public sealed class GetVolumeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Text describing a block storage volume.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The name of block storage volume.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The region the block storage volume is provisioned in.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetVolumeArgs()
        {
        }
    }

    public sealed class GetVolumeInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Text describing a block storage volume.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of block storage volume.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The region the block storage volume is provisioned in.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetVolumeInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVolumeResult
    {
        /// <summary>
        /// Text describing a block storage volume.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// A list of associated Droplet ids.
        /// </summary>
        public readonly ImmutableArray<int> DropletIds;
        /// <summary>
        /// Filesystem label currently in-use on the block storage volume.
        /// </summary>
        public readonly string FilesystemLabel;
        /// <summary>
        /// Filesystem type currently in-use on the block storage volume.
        /// </summary>
        public readonly string FilesystemType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? Region;
        /// <summary>
        /// The size of the block storage volume in GiB.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// A list of the tags associated to the Volume.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        public readonly string Urn;

        [OutputConstructor]
        private GetVolumeResult(
            string? description,

            ImmutableArray<int> dropletIds,

            string filesystemLabel,

            string filesystemType,

            string id,

            string name,

            string? region,

            int size,

            ImmutableArray<string> tags,

            string urn)
        {
            Description = description;
            DropletIds = dropletIds;
            FilesystemLabel = filesystemLabel;
            FilesystemType = filesystemType;
            Id = id;
            Name = name;
            Region = region;
            Size = size;
            Tags = tags;
            Urn = urn;
        }
    }
}
