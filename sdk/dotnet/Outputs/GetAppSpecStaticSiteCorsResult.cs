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
    public sealed class GetAppSpecStaticSiteCorsResult
    {
        public readonly bool? AllowCredentials;
        public readonly ImmutableArray<string> AllowHeaders;
        public readonly ImmutableArray<string> AllowMethods;
        public readonly Outputs.GetAppSpecStaticSiteCorsAllowOriginsResult? AllowOrigins;
        public readonly ImmutableArray<string> ExposeHeaders;
        public readonly string? MaxAge;

        [OutputConstructor]
        private GetAppSpecStaticSiteCorsResult(
            bool? allowCredentials,

            ImmutableArray<string> allowHeaders,

            ImmutableArray<string> allowMethods,

            Outputs.GetAppSpecStaticSiteCorsAllowOriginsResult? allowOrigins,

            ImmutableArray<string> exposeHeaders,

            string? maxAge)
        {
            AllowCredentials = allowCredentials;
            AllowHeaders = allowHeaders;
            AllowMethods = allowMethods;
            AllowOrigins = allowOrigins;
            ExposeHeaders = exposeHeaders;
            MaxAge = maxAge;
        }
    }
}