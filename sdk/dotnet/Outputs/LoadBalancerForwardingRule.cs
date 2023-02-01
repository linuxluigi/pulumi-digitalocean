// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean.Outputs
{

    [OutputType]
    public sealed class LoadBalancerForwardingRule
    {
        /// <summary>
        /// **Deprecated** The ID of the TLS certificate to be used for SSL termination.
        /// </summary>
        public readonly string? CertificateId;
        /// <summary>
        /// The unique name of the TLS certificate to be used for SSL termination.
        /// </summary>
        public readonly string? CertificateName;
        /// <summary>
        /// An integer representing the port on which the Load Balancer instance will listen.
        /// </summary>
        public readonly int EntryPort;
        /// <summary>
        /// The protocol used for traffic to the Load Balancer. The possible values are: `http`, `https`, `http2`, `http3`, `tcp`, or `udp`.
        /// </summary>
        public readonly string EntryProtocol;
        /// <summary>
        /// An integer representing the port on the backend Droplets to which the Load Balancer will send traffic.
        /// </summary>
        public readonly int TargetPort;
        /// <summary>
        /// The protocol used for traffic from the Load Balancer to the backend Droplets. The possible values are: `http`, `https`, `http2`, `tcp`, or `udp`.
        /// </summary>
        public readonly string TargetProtocol;
        /// <summary>
        /// A boolean value indicating whether SSL encrypted traffic will be passed through to the backend Droplets. The default value is `false`.
        /// </summary>
        public readonly bool? TlsPassthrough;

        [OutputConstructor]
        private LoadBalancerForwardingRule(
            string? certificateId,

            string? certificateName,

            int entryPort,

            string entryProtocol,

            int targetPort,

            string targetProtocol,

            bool? tlsPassthrough)
        {
            CertificateId = certificateId;
            CertificateName = certificateName;
            EntryPort = entryPort;
            EntryProtocol = entryProtocol;
            TargetPort = targetPort;
            TargetProtocol = targetProtocol;
            TlsPassthrough = tlsPassthrough;
        }
    }
}
