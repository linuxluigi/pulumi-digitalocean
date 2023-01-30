// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean.Inputs
{

    public sealed class LoadBalancerForwardingRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// **Deprecated** The ID of the TLS certificate to be used for SSL termination.
        /// </summary>
        [Input("certificateId")]
        public Input<string>? CertificateId { get; set; }

        /// <summary>
        /// The unique name of the TLS certificate to be used for SSL termination.
        /// </summary>
        [Input("certificateName")]
        public Input<string>? CertificateName { get; set; }

        /// <summary>
        /// An integer representing the port on which the Load Balancer instance will listen.
        /// </summary>
        [Input("entryPort", required: true)]
        public Input<int> EntryPort { get; set; } = null!;

        /// <summary>
        /// The protocol used for traffic to the Load Balancer. The possible values are: `http`, `https`, `http2`, `http3`, `tcp`, or `udp`.
        /// </summary>
        [Input("entryProtocol", required: true)]
        public Input<string> EntryProtocol { get; set; } = null!;

        /// <summary>
        /// An integer representing the port on the backend Droplets to which the Load Balancer will send traffic.
        /// </summary>
        [Input("targetPort", required: true)]
        public Input<int> TargetPort { get; set; } = null!;

        /// <summary>
        /// The protocol used for traffic from the Load Balancer to the backend Droplets. The possible values are: `http`, `https`, `http2`, `tcp`, or `udp`.
        /// </summary>
        [Input("targetProtocol", required: true)]
        public Input<string> TargetProtocol { get; set; } = null!;

        /// <summary>
        /// A boolean value indicating whether SSL encrypted traffic will be passed through to the backend Droplets. The default value is `false`.
        /// </summary>
        [Input("tlsPassthrough")]
        public Input<bool>? TlsPassthrough { get; set; }

        public LoadBalancerForwardingRuleArgs()
        {
        }
        public static new LoadBalancerForwardingRuleArgs Empty => new LoadBalancerForwardingRuleArgs();
    }
}
