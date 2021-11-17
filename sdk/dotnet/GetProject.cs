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
    public static class GetProject
    {
        /// <summary>
        /// Get information on a single DigitalOcean project. If neither the `id` nor `name` attributes are provided,
        /// then this data source returns the default project.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var @default = Output.Create(DigitalOcean.GetProject.InvokeAsync());
        ///         var staging = Output.Create(DigitalOcean.GetProject.InvokeAsync(new DigitalOcean.GetProjectArgs
        ///         {
        ///             Name = "My Staging Project",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetProjectResult> InvokeAsync(GetProjectArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProjectResult>("digitalocean:index/getProject:getProject", args ?? new GetProjectArgs(), options.WithVersion());

        /// <summary>
        /// Get information on a single DigitalOcean project. If neither the `id` nor `name` attributes are provided,
        /// then this data source returns the default project.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var @default = Output.Create(DigitalOcean.GetProject.InvokeAsync());
        ///         var staging = Output.Create(DigitalOcean.GetProject.InvokeAsync(new DigitalOcean.GetProjectArgs
        ///         {
        ///             Name = "My Staging Project",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetProjectResult> Invoke(GetProjectInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetProjectResult>("digitalocean:index/getProject:getProject", args ?? new GetProjectInvokeArgs(), options.WithVersion());
    }


    public sealed class GetProjectArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// the ID of the project to retrieve
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// the name of the project to retrieve. The data source will raise an error if more than
        /// one project has the provided name or if no project has that name.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetProjectArgs()
        {
        }
    }

    public sealed class GetProjectInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// the ID of the project to retrieve
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// the name of the project to retrieve. The data source will raise an error if more than
        /// one project has the provided name or if no project has that name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetProjectInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProjectResult
    {
        /// <summary>
        /// The date and time when the project was created, (ISO8601)
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The description of the project
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The environment of the project's resources. The possible values are: `Development`, `Staging`, `Production`.
        /// </summary>
        public readonly string Environment;
        public readonly string Id;
        public readonly bool IsDefault;
        public readonly string Name;
        /// <summary>
        /// The ID of the project owner.
        /// </summary>
        public readonly int OwnerId;
        /// <summary>
        /// The unique universal identifier of the project owner.
        /// </summary>
        public readonly string OwnerUuid;
        /// <summary>
        /// The purpose of the project, (Default: "Web Application")
        /// </summary>
        public readonly string Purpose;
        /// <summary>
        /// A set of uniform resource names (URNs) for the resources associated with the project
        /// </summary>
        public readonly ImmutableArray<string> Resources;
        /// <summary>
        /// The date and time when the project was last updated, (ISO8601)
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetProjectResult(
            string createdAt,

            string description,

            string environment,

            string id,

            bool isDefault,

            string name,

            int ownerId,

            string ownerUuid,

            string purpose,

            ImmutableArray<string> resources,

            string updatedAt)
        {
            CreatedAt = createdAt;
            Description = description;
            Environment = environment;
            Id = id;
            IsDefault = isDefault;
            Name = name;
            OwnerId = ownerId;
            OwnerUuid = ownerUuid;
            Purpose = purpose;
            Resources = resources;
            UpdatedAt = updatedAt;
        }
    }
}
