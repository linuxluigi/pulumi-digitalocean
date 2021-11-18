// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean.Inputs
{

    public sealed class GetFirewallOutboundRuleInputArgs : Pulumi.ResourceArgs
    {
        [Input("destinationAddresses")]
        private InputList<string>? _destinationAddresses;

        /// <summary>
        /// An array of strings containing the IPv4
        /// addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs to which the
        /// outbound traffic will be allowed.
        /// </summary>
        public InputList<string> DestinationAddresses
        {
            get => _destinationAddresses ?? (_destinationAddresses = new InputList<string>());
            set => _destinationAddresses = value;
        }

        [Input("destinationDropletIds")]
        private InputList<int>? _destinationDropletIds;

        /// <summary>
        /// An array containing the IDs of
        /// the Droplets to which the outbound traffic will be allowed.
        /// </summary>
        public InputList<int> DestinationDropletIds
        {
            get => _destinationDropletIds ?? (_destinationDropletIds = new InputList<int>());
            set => _destinationDropletIds = value;
        }

        [Input("destinationLoadBalancerUids")]
        private InputList<string>? _destinationLoadBalancerUids;

        /// <summary>
        /// An array containing the IDs
        /// of the Load Balancers to which the outbound traffic will be allowed.
        /// </summary>
        public InputList<string> DestinationLoadBalancerUids
        {
            get => _destinationLoadBalancerUids ?? (_destinationLoadBalancerUids = new InputList<string>());
            set => _destinationLoadBalancerUids = value;
        }

        [Input("destinationTags")]
        private InputList<string>? _destinationTags;

        /// <summary>
        /// An array containing the names of Tags
        /// corresponding to groups of Droplets to which the outbound traffic will
        /// be allowed.
        /// traffic.
        /// </summary>
        public InputList<string> DestinationTags
        {
            get => _destinationTags ?? (_destinationTags = new InputList<string>());
            set => _destinationTags = value;
        }

        /// <summary>
        /// The ports on which traffic will be allowed
        /// specified as a string containing a single port, a range (e.g. "8000-9000"),
        /// or "1-65535" to open all ports for a protocol. Required for when protocol is
        /// `tcp` or `udp`.
        /// </summary>
        [Input("portRange")]
        public Input<string>? PortRange { get; set; }

        /// <summary>
        /// The type of traffic to be allowed.
        /// This may be one of "tcp", "udp", or "icmp".
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        public GetFirewallOutboundRuleInputArgs()
        {
        }
    }
}
