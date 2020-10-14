// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean.Inputs
{

    public sealed class AppSpecWorkerGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional build command to run while building this component from source.
        /// </summary>
        [Input("buildCommand")]
        public Input<string>? BuildCommand { get; set; }

        /// <summary>
        /// The path to a Dockerfile relative to the root of the repo. If set, overrides usage of buildpacks.
        /// </summary>
        [Input("dockerfilePath")]
        public Input<string>? DockerfilePath { get; set; }

        /// <summary>
        /// An environment slug describing the type of this app.
        /// </summary>
        [Input("environmentSlug")]
        public Input<string>? EnvironmentSlug { get; set; }

        [Input("envs")]
        private InputList<Inputs.AppSpecWorkerEnvGetArgs>? _envs;

        /// <summary>
        /// Describes an environment variable made available to an app competent.
        /// </summary>
        public InputList<Inputs.AppSpecWorkerEnvGetArgs> Envs
        {
            get => _envs ?? (_envs = new InputList<Inputs.AppSpecWorkerEnvGetArgs>());
            set => _envs = value;
        }

        /// <summary>
        /// A Git repo to use as component's source. Only one of `git` and `github` may be set.
        /// </summary>
        [Input("git")]
        public Input<Inputs.AppSpecWorkerGitGetArgs>? Git { get; set; }

        /// <summary>
        /// A GitHub repo to use as component's source. Only one of `git` and `github` may be set.
        /// </summary>
        [Input("github")]
        public Input<Inputs.AppSpecWorkerGithubGetArgs>? Github { get; set; }

        /// <summary>
        /// The amount of instances that this component should be scaled to.
        /// </summary>
        [Input("instanceCount")]
        public Input<int>? InstanceCount { get; set; }

        /// <summary>
        /// The instance size to use for this component.
        /// </summary>
        [Input("instanceSizeSlug")]
        public Input<string>? InstanceSizeSlug { get; set; }

        /// <summary>
        /// The name of the component
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("routes")]
        public Input<Inputs.AppSpecWorkerRoutesGetArgs>? Routes { get; set; }

        /// <summary>
        /// An optional run command to override the component's default.
        /// </summary>
        [Input("runCommand")]
        public Input<string>? RunCommand { get; set; }

        /// <summary>
        /// An optional path to the working directory to use for the build.
        /// </summary>
        [Input("sourceDir")]
        public Input<string>? SourceDir { get; set; }

        public AppSpecWorkerGetArgs()
        {
        }
    }
}
