// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean.Inputs
{

    public sealed class AppSpecServiceImageGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The registry name. Must be left empty for the `DOCR` registry type. Required for the `DOCKER_HUB` registry type.
        /// </summary>
        [Input("registry")]
        public Input<string>? Registry { get; set; }

        /// <summary>
        /// The registry type. One of `DOCR` (DigitalOcean container registry) or `DOCKER_HUB`.
        /// </summary>
        [Input("registryType", required: true)]
        public Input<string> RegistryType { get; set; } = null!;

        /// <summary>
        /// The repository name.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        /// <summary>
        /// The repository tag. Defaults to `latest` if not provided.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        public AppSpecServiceImageGetArgs()
        {
        }
    }
}
