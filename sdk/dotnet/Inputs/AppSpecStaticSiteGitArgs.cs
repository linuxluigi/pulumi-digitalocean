// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean.Inputs
{

    public sealed class AppSpecStaticSiteGitArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the branch to use.
        /// </summary>
        [Input("branch")]
        public Input<string>? Branch { get; set; }

        /// <summary>
        /// The clone URL of the repo.
        /// </summary>
        [Input("repoCloneUrl")]
        public Input<string>? RepoCloneUrl { get; set; }

        public AppSpecStaticSiteGitArgs()
        {
        }
    }
}
