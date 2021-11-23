// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean.Inputs
{

    public sealed class GetSpacesBucketsFilterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Set to `true` to require that a field match all of the `values` instead of just one or more of
        /// them. This is useful when matching against multi-valued fields such as lists or sets where you want to ensure
        /// that all of the `values` are present in the list or set.
        /// </summary>
        [Input("all")]
        public bool? All { get; set; }

        /// <summary>
        /// Filter the images by this key. This may be one of `bucket_domain_name`, `name`, `region`, or `urn`.
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        /// <summary>
        /// One of `exact` (default), `re`, or `substring`. For string-typed fields, specify `re` to
        /// match by using the `values` as regular expressions, or specify `substring` to match by treating the `values` as
        /// substrings to find within the string field.
        /// </summary>
        [Input("matchBy")]
        public string? MatchBy { get; set; }

        [Input("values", required: true)]
        private List<string>? _values;

        /// <summary>
        /// A list of values to match against the `key` field. Only retrieves Spaces buckets
        /// where the `key` field takes on one or more of the values provided here.
        /// </summary>
        public List<string> Values
        {
            get => _values ?? (_values = new List<string>());
            set => _values = value;
        }

        public GetSpacesBucketsFilterArgs()
        {
        }
    }
}