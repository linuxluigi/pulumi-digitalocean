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
    public sealed class GetRegionsFilterResult
    {
        /// <summary>
        /// Filter the regions by this key. This may be one of `slug`,
        /// `name`, `available`, `features`, or `sizes`.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// A list of values to match against the `key` field. Only retrieves regions
        /// where the `key` field takes on one or more of the values provided here.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetRegionsFilterResult(
            string key,

            ImmutableArray<string> values)
        {
            Key = key;
            Values = values;
        }
    }
}