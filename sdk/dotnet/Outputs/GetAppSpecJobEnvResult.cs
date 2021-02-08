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
    public sealed class GetAppSpecJobEnvResult
    {
        /// <summary>
        /// The name of the environment variable.
        /// </summary>
        public readonly string? Key;
        /// <summary>
        /// The visibility scope of the environment variable. One of `RUN_TIME`, `BUILD_TIME`, or `RUN_AND_BUILD_TIME` (default).
        /// </summary>
        public readonly string? Scope;
        /// <summary>
        /// The type of the environment variable, `GENERAL` or `SECRET`.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The value of the environment variable.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private GetAppSpecJobEnvResult(
            string? key,

            string? scope,

            string type,

            string? value)
        {
            Key = key;
            Scope = scope;
            Type = type;
            Value = value;
        }
    }
}
