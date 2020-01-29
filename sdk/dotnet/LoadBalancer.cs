// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean
{
    /// <summary>
    /// Provides a DigitalOcean Load Balancer resource. This can be used to create,
    /// modify, and delete Load Balancers.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/r/loadbalancer.html.markdown.
    /// </summary>
    public partial class LoadBalancer : Pulumi.CustomResource
    {
        /// <summary>
        /// The load balancing algorithm used to determine
        /// which backend Droplet will be selected by a client. It must be either `round_robin`
        /// or `least_connections`. The default value is `round_robin`.
        /// </summary>
        [Output("algorithm")]
        public Output<string?> Algorithm { get; private set; } = null!;

        /// <summary>
        /// A list of the IDs of each droplet to be attached to the Load Balancer.
        /// </summary>
        [Output("dropletIds")]
        public Output<ImmutableArray<int>> DropletIds { get; private set; } = null!;

        /// <summary>
        /// The name of a Droplet tag corresponding to Droplets to be assigned to the Load Balancer.
        /// </summary>
        [Output("dropletTag")]
        public Output<string?> DropletTag { get; private set; } = null!;

        /// <summary>
        /// A boolean value indicating whether PROXY
        /// Protocol should be used to pass information from connecting client requests to
        /// the backend service. Default value is `false`.
        /// </summary>
        [Output("enableProxyProtocol")]
        public Output<bool?> EnableProxyProtocol { get; private set; } = null!;

        /// <summary>
        /// A list of `forwarding_rule` to be assigned to the
        /// Load Balancer. The `forwarding_rule` block is documented below.
        /// </summary>
        [Output("forwardingRules")]
        public Output<ImmutableArray<Outputs.LoadBalancerForwardingRules>> ForwardingRules { get; private set; } = null!;

        /// <summary>
        /// A `healthcheck` block to be assigned to the
        /// Load Balancer. The `healthcheck` block is documented below. Only 1 healthcheck is allowed.
        /// </summary>
        [Output("healthcheck")]
        public Output<Outputs.LoadBalancerHealthcheck> Healthcheck { get; private set; } = null!;

        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// The Load Balancer name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A boolean value indicating whether
        /// HTTP requests to the Load Balancer on port 80 will be redirected to HTTPS on port 443.
        /// Default value is `false`.
        /// </summary>
        [Output("redirectHttpToHttps")]
        public Output<bool?> RedirectHttpToHttps { get; private set; } = null!;

        /// <summary>
        /// The region to start in
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A `sticky_sessions` block to be assigned to the
        /// Load Balancer. The `sticky_sessions` block is documented below. Only 1 sticky_sessions block is allowed.
        /// </summary>
        [Output("stickySessions")]
        public Output<Outputs.LoadBalancerStickySessions> StickySessions { get; private set; } = null!;

        /// <summary>
        /// The uniform resource name for the Load Balancer
        /// </summary>
        [Output("urn")]
        public Output<string> Urn { get; private set; } = null!;


        /// <summary>
        /// Create a LoadBalancer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoadBalancer(string name, LoadBalancerArgs args, CustomResourceOptions? options = null)
            : base("digitalocean:index/loadBalancer:LoadBalancer", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private LoadBalancer(string name, Input<string> id, LoadBalancerState? state = null, CustomResourceOptions? options = null)
            : base("digitalocean:index/loadBalancer:LoadBalancer", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LoadBalancer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoadBalancer Get(string name, Input<string> id, LoadBalancerState? state = null, CustomResourceOptions? options = null)
        {
            return new LoadBalancer(name, id, state, options);
        }
    }

    public sealed class LoadBalancerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The load balancing algorithm used to determine
        /// which backend Droplet will be selected by a client. It must be either `round_robin`
        /// or `least_connections`. The default value is `round_robin`.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        [Input("dropletIds")]
        private InputList<int>? _dropletIds;

        /// <summary>
        /// A list of the IDs of each droplet to be attached to the Load Balancer.
        /// </summary>
        public InputList<int> DropletIds
        {
            get => _dropletIds ?? (_dropletIds = new InputList<int>());
            set => _dropletIds = value;
        }

        /// <summary>
        /// The name of a Droplet tag corresponding to Droplets to be assigned to the Load Balancer.
        /// </summary>
        [Input("dropletTag")]
        public Input<string>? DropletTag { get; set; }

        /// <summary>
        /// A boolean value indicating whether PROXY
        /// Protocol should be used to pass information from connecting client requests to
        /// the backend service. Default value is `false`.
        /// </summary>
        [Input("enableProxyProtocol")]
        public Input<bool>? EnableProxyProtocol { get; set; }

        [Input("forwardingRules", required: true)]
        private InputList<Inputs.LoadBalancerForwardingRulesArgs>? _forwardingRules;

        /// <summary>
        /// A list of `forwarding_rule` to be assigned to the
        /// Load Balancer. The `forwarding_rule` block is documented below.
        /// </summary>
        public InputList<Inputs.LoadBalancerForwardingRulesArgs> ForwardingRules
        {
            get => _forwardingRules ?? (_forwardingRules = new InputList<Inputs.LoadBalancerForwardingRulesArgs>());
            set => _forwardingRules = value;
        }

        /// <summary>
        /// A `healthcheck` block to be assigned to the
        /// Load Balancer. The `healthcheck` block is documented below. Only 1 healthcheck is allowed.
        /// </summary>
        [Input("healthcheck")]
        public Input<Inputs.LoadBalancerHealthcheckArgs>? Healthcheck { get; set; }

        /// <summary>
        /// The Load Balancer name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A boolean value indicating whether
        /// HTTP requests to the Load Balancer on port 80 will be redirected to HTTPS on port 443.
        /// Default value is `false`.
        /// </summary>
        [Input("redirectHttpToHttps")]
        public Input<bool>? RedirectHttpToHttps { get; set; }

        /// <summary>
        /// The region to start in
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// A `sticky_sessions` block to be assigned to the
        /// Load Balancer. The `sticky_sessions` block is documented below. Only 1 sticky_sessions block is allowed.
        /// </summary>
        [Input("stickySessions")]
        public Input<Inputs.LoadBalancerStickySessionsArgs>? StickySessions { get; set; }

        public LoadBalancerArgs()
        {
        }
    }

    public sealed class LoadBalancerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The load balancing algorithm used to determine
        /// which backend Droplet will be selected by a client. It must be either `round_robin`
        /// or `least_connections`. The default value is `round_robin`.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        [Input("dropletIds")]
        private InputList<int>? _dropletIds;

        /// <summary>
        /// A list of the IDs of each droplet to be attached to the Load Balancer.
        /// </summary>
        public InputList<int> DropletIds
        {
            get => _dropletIds ?? (_dropletIds = new InputList<int>());
            set => _dropletIds = value;
        }

        /// <summary>
        /// The name of a Droplet tag corresponding to Droplets to be assigned to the Load Balancer.
        /// </summary>
        [Input("dropletTag")]
        public Input<string>? DropletTag { get; set; }

        /// <summary>
        /// A boolean value indicating whether PROXY
        /// Protocol should be used to pass information from connecting client requests to
        /// the backend service. Default value is `false`.
        /// </summary>
        [Input("enableProxyProtocol")]
        public Input<bool>? EnableProxyProtocol { get; set; }

        [Input("forwardingRules")]
        private InputList<Inputs.LoadBalancerForwardingRulesGetArgs>? _forwardingRules;

        /// <summary>
        /// A list of `forwarding_rule` to be assigned to the
        /// Load Balancer. The `forwarding_rule` block is documented below.
        /// </summary>
        public InputList<Inputs.LoadBalancerForwardingRulesGetArgs> ForwardingRules
        {
            get => _forwardingRules ?? (_forwardingRules = new InputList<Inputs.LoadBalancerForwardingRulesGetArgs>());
            set => _forwardingRules = value;
        }

        /// <summary>
        /// A `healthcheck` block to be assigned to the
        /// Load Balancer. The `healthcheck` block is documented below. Only 1 healthcheck is allowed.
        /// </summary>
        [Input("healthcheck")]
        public Input<Inputs.LoadBalancerHealthcheckGetArgs>? Healthcheck { get; set; }

        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The Load Balancer name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A boolean value indicating whether
        /// HTTP requests to the Load Balancer on port 80 will be redirected to HTTPS on port 443.
        /// Default value is `false`.
        /// </summary>
        [Input("redirectHttpToHttps")]
        public Input<bool>? RedirectHttpToHttps { get; set; }

        /// <summary>
        /// The region to start in
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// A `sticky_sessions` block to be assigned to the
        /// Load Balancer. The `sticky_sessions` block is documented below. Only 1 sticky_sessions block is allowed.
        /// </summary>
        [Input("stickySessions")]
        public Input<Inputs.LoadBalancerStickySessionsGetArgs>? StickySessions { get; set; }

        /// <summary>
        /// The uniform resource name for the Load Balancer
        /// </summary>
        [Input("urn")]
        public Input<string>? Urn { get; set; }

        public LoadBalancerState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class LoadBalancerForwardingRulesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the TLS certificate to be used for SSL termination.
        /// </summary>
        [Input("certificateId")]
        public Input<string>? CertificateId { get; set; }

        /// <summary>
        /// An integer representing the port on which the Load Balancer instance will listen.
        /// </summary>
        [Input("entryPort", required: true)]
        public Input<int> EntryPort { get; set; } = null!;

        /// <summary>
        /// The protocol used for traffic to the Load Balancer. The possible values are: `http`, `https`, `http2` or `tcp`.
        /// </summary>
        [Input("entryProtocol", required: true)]
        public Input<string> EntryProtocol { get; set; } = null!;

        /// <summary>
        /// An integer representing the port on the backend Droplets to which the Load Balancer will send traffic.
        /// </summary>
        [Input("targetPort", required: true)]
        public Input<int> TargetPort { get; set; } = null!;

        /// <summary>
        /// The protocol used for traffic from the Load Balancer to the backend Droplets. The possible values are: `http`, `https`, `http2` or `tcp`.
        /// </summary>
        [Input("targetProtocol", required: true)]
        public Input<string> TargetProtocol { get; set; } = null!;

        /// <summary>
        /// A boolean value indicating whether SSL encrypted traffic will be passed through to the backend Droplets. The default value is `false`.
        /// </summary>
        [Input("tlsPassthrough")]
        public Input<bool>? TlsPassthrough { get; set; }

        public LoadBalancerForwardingRulesArgs()
        {
        }
    }

    public sealed class LoadBalancerForwardingRulesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the TLS certificate to be used for SSL termination.
        /// </summary>
        [Input("certificateId")]
        public Input<string>? CertificateId { get; set; }

        /// <summary>
        /// An integer representing the port on which the Load Balancer instance will listen.
        /// </summary>
        [Input("entryPort", required: true)]
        public Input<int> EntryPort { get; set; } = null!;

        /// <summary>
        /// The protocol used for traffic to the Load Balancer. The possible values are: `http`, `https`, `http2` or `tcp`.
        /// </summary>
        [Input("entryProtocol", required: true)]
        public Input<string> EntryProtocol { get; set; } = null!;

        /// <summary>
        /// An integer representing the port on the backend Droplets to which the Load Balancer will send traffic.
        /// </summary>
        [Input("targetPort", required: true)]
        public Input<int> TargetPort { get; set; } = null!;

        /// <summary>
        /// The protocol used for traffic from the Load Balancer to the backend Droplets. The possible values are: `http`, `https`, `http2` or `tcp`.
        /// </summary>
        [Input("targetProtocol", required: true)]
        public Input<string> TargetProtocol { get; set; } = null!;

        /// <summary>
        /// A boolean value indicating whether SSL encrypted traffic will be passed through to the backend Droplets. The default value is `false`.
        /// </summary>
        [Input("tlsPassthrough")]
        public Input<bool>? TlsPassthrough { get; set; }

        public LoadBalancerForwardingRulesGetArgs()
        {
        }
    }

    public sealed class LoadBalancerHealthcheckArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of seconds between between two consecutive health checks. If not specified, the default value is `10`.
        /// </summary>
        [Input("checkIntervalSeconds")]
        public Input<int>? CheckIntervalSeconds { get; set; }

        /// <summary>
        /// The number of times a health check must pass for a backend Droplet to be marked "healthy" and be re-added to the pool. If not specified, the default value is `5`.
        /// </summary>
        [Input("healthyThreshold")]
        public Input<int>? HealthyThreshold { get; set; }

        /// <summary>
        /// The path on the backend Droplets to which the Load Balancer instance will send a request.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// An integer representing the port on the backend Droplets on which the health check will attempt a connection.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// The protocol used for health checks sent to the backend Droplets. The possible values are `http` or `tcp`.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// The number of seconds the Load Balancer instance will wait for a response until marking a health check as failed. If not specified, the default value is `5`.
        /// </summary>
        [Input("responseTimeoutSeconds")]
        public Input<int>? ResponseTimeoutSeconds { get; set; }

        /// <summary>
        /// The number of times a health check must fail for a backend Droplet to be marked "unhealthy" and be removed from the pool. If not specified, the default value is `3`.
        /// </summary>
        [Input("unhealthyThreshold")]
        public Input<int>? UnhealthyThreshold { get; set; }

        public LoadBalancerHealthcheckArgs()
        {
        }
    }

    public sealed class LoadBalancerHealthcheckGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of seconds between between two consecutive health checks. If not specified, the default value is `10`.
        /// </summary>
        [Input("checkIntervalSeconds")]
        public Input<int>? CheckIntervalSeconds { get; set; }

        /// <summary>
        /// The number of times a health check must pass for a backend Droplet to be marked "healthy" and be re-added to the pool. If not specified, the default value is `5`.
        /// </summary>
        [Input("healthyThreshold")]
        public Input<int>? HealthyThreshold { get; set; }

        /// <summary>
        /// The path on the backend Droplets to which the Load Balancer instance will send a request.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// An integer representing the port on the backend Droplets on which the health check will attempt a connection.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// The protocol used for health checks sent to the backend Droplets. The possible values are `http` or `tcp`.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// The number of seconds the Load Balancer instance will wait for a response until marking a health check as failed. If not specified, the default value is `5`.
        /// </summary>
        [Input("responseTimeoutSeconds")]
        public Input<int>? ResponseTimeoutSeconds { get; set; }

        /// <summary>
        /// The number of times a health check must fail for a backend Droplet to be marked "unhealthy" and be removed from the pool. If not specified, the default value is `3`.
        /// </summary>
        [Input("unhealthyThreshold")]
        public Input<int>? UnhealthyThreshold { get; set; }

        public LoadBalancerHealthcheckGetArgs()
        {
        }
    }

    public sealed class LoadBalancerStickySessionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name to be used for the cookie sent to the client. This attribute is required when using `cookies` for the sticky sessions type.
        /// </summary>
        [Input("cookieName")]
        public Input<string>? CookieName { get; set; }

        /// <summary>
        /// The number of seconds until the cookie set by the Load Balancer expires. This attribute is required when using `cookies` for the sticky sessions type.
        /// </summary>
        [Input("cookieTtlSeconds")]
        public Input<int>? CookieTtlSeconds { get; set; }

        /// <summary>
        /// An attribute indicating how and if requests from a client will be persistently served by the same backend Droplet. The possible values are `cookies` or `none`. If not specified, the default value is `none`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public LoadBalancerStickySessionsArgs()
        {
        }
    }

    public sealed class LoadBalancerStickySessionsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name to be used for the cookie sent to the client. This attribute is required when using `cookies` for the sticky sessions type.
        /// </summary>
        [Input("cookieName")]
        public Input<string>? CookieName { get; set; }

        /// <summary>
        /// The number of seconds until the cookie set by the Load Balancer expires. This attribute is required when using `cookies` for the sticky sessions type.
        /// </summary>
        [Input("cookieTtlSeconds")]
        public Input<int>? CookieTtlSeconds { get; set; }

        /// <summary>
        /// An attribute indicating how and if requests from a client will be persistently served by the same backend Droplet. The possible values are `cookies` or `none`. If not specified, the default value is `none`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public LoadBalancerStickySessionsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class LoadBalancerForwardingRules
    {
        /// <summary>
        /// The ID of the TLS certificate to be used for SSL termination.
        /// </summary>
        public readonly string? CertificateId;
        /// <summary>
        /// An integer representing the port on which the Load Balancer instance will listen.
        /// </summary>
        public readonly int EntryPort;
        /// <summary>
        /// The protocol used for traffic to the Load Balancer. The possible values are: `http`, `https`, `http2` or `tcp`.
        /// </summary>
        public readonly string EntryProtocol;
        /// <summary>
        /// An integer representing the port on the backend Droplets to which the Load Balancer will send traffic.
        /// </summary>
        public readonly int TargetPort;
        /// <summary>
        /// The protocol used for traffic from the Load Balancer to the backend Droplets. The possible values are: `http`, `https`, `http2` or `tcp`.
        /// </summary>
        public readonly string TargetProtocol;
        /// <summary>
        /// A boolean value indicating whether SSL encrypted traffic will be passed through to the backend Droplets. The default value is `false`.
        /// </summary>
        public readonly bool? TlsPassthrough;

        [OutputConstructor]
        private LoadBalancerForwardingRules(
            string? certificateId,
            int entryPort,
            string entryProtocol,
            int targetPort,
            string targetProtocol,
            bool? tlsPassthrough)
        {
            CertificateId = certificateId;
            EntryPort = entryPort;
            EntryProtocol = entryProtocol;
            TargetPort = targetPort;
            TargetProtocol = targetProtocol;
            TlsPassthrough = tlsPassthrough;
        }
    }

    [OutputType]
    public sealed class LoadBalancerHealthcheck
    {
        /// <summary>
        /// The number of seconds between between two consecutive health checks. If not specified, the default value is `10`.
        /// </summary>
        public readonly int? CheckIntervalSeconds;
        /// <summary>
        /// The number of times a health check must pass for a backend Droplet to be marked "healthy" and be re-added to the pool. If not specified, the default value is `5`.
        /// </summary>
        public readonly int? HealthyThreshold;
        /// <summary>
        /// The path on the backend Droplets to which the Load Balancer instance will send a request.
        /// </summary>
        public readonly string? Path;
        /// <summary>
        /// An integer representing the port on the backend Droplets on which the health check will attempt a connection.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The protocol used for health checks sent to the backend Droplets. The possible values are `http` or `tcp`.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// The number of seconds the Load Balancer instance will wait for a response until marking a health check as failed. If not specified, the default value is `5`.
        /// </summary>
        public readonly int? ResponseTimeoutSeconds;
        /// <summary>
        /// The number of times a health check must fail for a backend Droplet to be marked "unhealthy" and be removed from the pool. If not specified, the default value is `3`.
        /// </summary>
        public readonly int? UnhealthyThreshold;

        [OutputConstructor]
        private LoadBalancerHealthcheck(
            int? checkIntervalSeconds,
            int? healthyThreshold,
            string? path,
            int port,
            string protocol,
            int? responseTimeoutSeconds,
            int? unhealthyThreshold)
        {
            CheckIntervalSeconds = checkIntervalSeconds;
            HealthyThreshold = healthyThreshold;
            Path = path;
            Port = port;
            Protocol = protocol;
            ResponseTimeoutSeconds = responseTimeoutSeconds;
            UnhealthyThreshold = unhealthyThreshold;
        }
    }

    [OutputType]
    public sealed class LoadBalancerStickySessions
    {
        /// <summary>
        /// The name to be used for the cookie sent to the client. This attribute is required when using `cookies` for the sticky sessions type.
        /// </summary>
        public readonly string? CookieName;
        /// <summary>
        /// The number of seconds until the cookie set by the Load Balancer expires. This attribute is required when using `cookies` for the sticky sessions type.
        /// </summary>
        public readonly int? CookieTtlSeconds;
        /// <summary>
        /// An attribute indicating how and if requests from a client will be persistently served by the same backend Droplet. The possible values are `cookies` or `none`. If not specified, the default value is `none`.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private LoadBalancerStickySessions(
            string? cookieName,
            int? cookieTtlSeconds,
            string? type)
        {
            CookieName = cookieName;
            CookieTtlSeconds = cookieTtlSeconds;
            Type = type;
        }
    }
    }
}