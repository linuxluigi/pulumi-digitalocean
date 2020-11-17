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
    public sealed class GetSshKeysFilterResult
    {
        public readonly bool? All;
        /// <summary>
        /// Filter the SSH Keys by this key. This may be one of `name`, `public_key`, or `fingerprint`.
        /// </summary>
        public readonly string Key;
        public readonly string? MatchBy;
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetSshKeysFilterResult(
            bool? all,

            string key,

            string? matchBy,

            ImmutableArray<string> values)
        {
            All = all;
            Key = key;
            MatchBy = matchBy;
            Values = values;
        }
    }
}
