// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean.Inputs
{

    public sealed class AppSpecServiceCorsAllowOriginsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Exact string match.
        /// </summary>
        [Input("exact")]
        public Input<string>? Exact { get; set; }

        /// <summary>
        /// Prefix-based match.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// RE2 style regex-based match.
        /// </summary>
        [Input("regex")]
        public Input<string>? Regex { get; set; }

        public AppSpecServiceCorsAllowOriginsGetArgs()
        {
        }
        public static new AppSpecServiceCorsAllowOriginsGetArgs Empty => new AppSpecServiceCorsAllowOriginsGetArgs();
    }
}
