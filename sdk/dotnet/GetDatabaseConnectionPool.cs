// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean
{
    public static class GetDatabaseConnectionPool
    {
        /// <summary>
        /// Provides information on a DigitalOcean PostgreSQL database connection pool.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = DigitalOcean.GetDatabaseCluster.Invoke(new()
        ///     {
        ///         Name = "example-cluster",
        ///     });
        /// 
        ///     var read_only = DigitalOcean.GetDatabaseConnectionPool.Invoke(new()
        ///     {
        ///         ClusterId = example.Apply(getDatabaseClusterResult =&gt; getDatabaseClusterResult.Id),
        ///         Name = "pool-01",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["connectionPoolUriOutput"] = read_only.Apply(read_only =&gt; read_only.Apply(getDatabaseConnectionPoolResult =&gt; getDatabaseConnectionPoolResult.Uri)),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetDatabaseConnectionPoolResult> InvokeAsync(GetDatabaseConnectionPoolArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseConnectionPoolResult>("digitalocean:index/getDatabaseConnectionPool:getDatabaseConnectionPool", args ?? new GetDatabaseConnectionPoolArgs(), options.WithDefaults());

        /// <summary>
        /// Provides information on a DigitalOcean PostgreSQL database connection pool.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = DigitalOcean.GetDatabaseCluster.Invoke(new()
        ///     {
        ///         Name = "example-cluster",
        ///     });
        /// 
        ///     var read_only = DigitalOcean.GetDatabaseConnectionPool.Invoke(new()
        ///     {
        ///         ClusterId = example.Apply(getDatabaseClusterResult =&gt; getDatabaseClusterResult.Id),
        ///         Name = "pool-01",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["connectionPoolUriOutput"] = read_only.Apply(read_only =&gt; read_only.Apply(getDatabaseConnectionPoolResult =&gt; getDatabaseConnectionPoolResult.Uri)),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetDatabaseConnectionPoolResult> Invoke(GetDatabaseConnectionPoolInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabaseConnectionPoolResult>("digitalocean:index/getDatabaseConnectionPool:getDatabaseConnectionPool", args ?? new GetDatabaseConnectionPoolInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Provides information on a DigitalOcean PostgreSQL database connection pool.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = DigitalOcean.GetDatabaseCluster.Invoke(new()
        ///     {
        ///         Name = "example-cluster",
        ///     });
        /// 
        ///     var read_only = DigitalOcean.GetDatabaseConnectionPool.Invoke(new()
        ///     {
        ///         ClusterId = example.Apply(getDatabaseClusterResult =&gt; getDatabaseClusterResult.Id),
        ///         Name = "pool-01",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["connectionPoolUriOutput"] = read_only.Apply(read_only =&gt; read_only.Apply(getDatabaseConnectionPoolResult =&gt; getDatabaseConnectionPoolResult.Uri)),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetDatabaseConnectionPoolResult> Invoke(GetDatabaseConnectionPoolInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabaseConnectionPoolResult>("digitalocean:index/getDatabaseConnectionPool:getDatabaseConnectionPool", args ?? new GetDatabaseConnectionPoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabaseConnectionPoolArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the original source database cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// The name for the database connection pool.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetDatabaseConnectionPoolArgs()
        {
        }
        public static new GetDatabaseConnectionPoolArgs Empty => new GetDatabaseConnectionPoolArgs();
    }

    public sealed class GetDatabaseConnectionPoolInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the original source database cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The name for the database connection pool.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetDatabaseConnectionPoolInvokeArgs()
        {
        }
        public static new GetDatabaseConnectionPoolInvokeArgs Empty => new GetDatabaseConnectionPoolInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabaseConnectionPoolResult
    {
        public readonly string ClusterId;
        /// <summary>
        /// Name of the connection pool's default database.
        /// </summary>
        public readonly string DbName;
        /// <summary>
        /// Connection pool hostname.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The transaction mode for the connection pool.
        /// </summary>
        public readonly string Mode;
        public readonly string Name;
        /// <summary>
        /// Password for the connection pool's default user.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// Network port that the connection pool is listening on.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// Same as `host`, but only accessible from resources within the account and in the same region.
        /// </summary>
        public readonly string PrivateHost;
        /// <summary>
        /// Same as `uri`, but only accessible from resources within the account and in the same region.
        /// </summary>
        public readonly string PrivateUri;
        /// <summary>
        /// Size of the connection pool.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// The full URI for connecting to the database connection pool.
        /// </summary>
        public readonly string Uri;
        /// <summary>
        /// Username for the connection pool's default user.
        /// </summary>
        public readonly string User;

        [OutputConstructor]
        private GetDatabaseConnectionPoolResult(
            string clusterId,

            string dbName,

            string host,

            string id,

            string mode,

            string name,

            string password,

            int port,

            string privateHost,

            string privateUri,

            int size,

            string uri,

            string user)
        {
            ClusterId = clusterId;
            DbName = dbName;
            Host = host;
            Id = id;
            Mode = mode;
            Name = name;
            Password = password;
            Port = port;
            PrivateHost = privateHost;
            PrivateUri = privateUri;
            Size = size;
            Uri = uri;
            User = user;
        }
    }
}
