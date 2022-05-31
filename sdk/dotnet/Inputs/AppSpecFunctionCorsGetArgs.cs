// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean.Inputs
{

    public sealed class AppSpecFunctionCorsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether browsers should expose the response to the client-side JavaScript code when the request's credentials mode is `include`. This configures the `Access-Control-Allow-Credentials` header.
        /// </summary>
        [Input("allowCredentials")]
        public Input<bool>? AllowCredentials { get; set; }

        [Input("allowHeaders")]
        private InputList<string>? _allowHeaders;

        /// <summary>
        /// The set of allowed HTTP request headers. This configures the `Access-Control-Allow-Headers` header.
        /// </summary>
        public InputList<string> AllowHeaders
        {
            get => _allowHeaders ?? (_allowHeaders = new InputList<string>());
            set => _allowHeaders = value;
        }

        [Input("allowMethods")]
        private InputList<string>? _allowMethods;

        /// <summary>
        /// The set of allowed HTTP methods. This configures the `Access-Control-Allow-Methods` header.
        /// </summary>
        public InputList<string> AllowMethods
        {
            get => _allowMethods ?? (_allowMethods = new InputList<string>());
            set => _allowMethods = value;
        }

        /// <summary>
        /// The `Access-Control-Allow-Origin` can be
        /// </summary>
        [Input("allowOrigins")]
        public Input<Inputs.AppSpecFunctionCorsAllowOriginsGetArgs>? AllowOrigins { get; set; }

        [Input("exposeHeaders")]
        private InputList<string>? _exposeHeaders;

        /// <summary>
        /// The set of HTTP response headers that browsers are allowed to access. This configures the `Access-Control-Expose-Headers` header.
        /// </summary>
        public InputList<string> ExposeHeaders
        {
            get => _exposeHeaders ?? (_exposeHeaders = new InputList<string>());
            set => _exposeHeaders = value;
        }

        /// <summary>
        /// An optional duration specifying how long browsers can cache the results of a preflight request. This configures the Access-Control-Max-Age header. Example: `5h30m`.
        /// </summary>
        [Input("maxAge")]
        public Input<string>? MaxAge { get; set; }

        public AppSpecFunctionCorsGetArgs()
        {
        }
    }
}
