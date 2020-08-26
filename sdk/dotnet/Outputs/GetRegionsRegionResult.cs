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
    public sealed class GetRegionsRegionResult
    {
        /// <summary>
        /// A boolean value that represents whether new Droplets can be created in this region.
        /// </summary>
        public readonly bool Available;
        /// <summary>
        /// A set of features available in this region.
        /// </summary>
        public readonly ImmutableArray<string> Features;
        /// <summary>
        /// The display name of the region.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A set of identifying slugs for the Droplet sizes available in this region.
        /// </summary>
        public readonly ImmutableArray<string> Sizes;
        /// <summary>
        /// A human-readable string that is used as a unique identifier for each region.
        /// </summary>
        public readonly string Slug;

        [OutputConstructor]
        private GetRegionsRegionResult(
            bool available,

            ImmutableArray<string> features,

            string name,

            ImmutableArray<string> sizes,

            string slug)
        {
            Available = available;
            Features = features;
            Name = name;
            Sizes = sizes;
            Slug = slug;
        }
    }
}
