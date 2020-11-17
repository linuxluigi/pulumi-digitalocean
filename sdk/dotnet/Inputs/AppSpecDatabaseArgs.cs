// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean.Inputs
{

    public sealed class AppSpecDatabaseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the underlying DigitalOcean DBaaS cluster. This is required for production databases. For dev databases, if `cluster_name` is not set, a new cluster will be provisioned.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// The name of the MySQL or PostgreSQL database to configure.
        /// </summary>
        [Input("dbName")]
        public Input<string>? DbName { get; set; }

        /// <summary>
        /// The name of the MySQL or PostgreSQL user to configure.
        /// </summary>
        [Input("dbUser")]
        public Input<string>? DbUser { get; set; }

        /// <summary>
        /// The database engine to use (`MYSQL`, `PG`, or `REDIS`).
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// The name of the component.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether this is a production or dev database.
        /// </summary>
        [Input("production")]
        public Input<bool>? Production { get; set; }

        /// <summary>
        /// The version of the database engine.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public AppSpecDatabaseArgs()
        {
        }
    }
}
