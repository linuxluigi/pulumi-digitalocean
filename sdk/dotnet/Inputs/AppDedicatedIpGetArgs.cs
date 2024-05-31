// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean.Inputs
{

    public sealed class AppDedicatedIpGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the app.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The IP address of the dedicated egress IP.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The status of the dedicated egress IP: 'UNKNOWN', 'ASSIGNING', 'ASSIGNED', or 'REMOVED'
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public AppDedicatedIpGetArgs()
        {
        }
        public static new AppDedicatedIpGetArgs Empty => new AppDedicatedIpGetArgs();
    }
}