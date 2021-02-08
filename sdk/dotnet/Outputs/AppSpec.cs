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
    public sealed class AppSpec
    {
        public readonly ImmutableArray<Outputs.AppSpecDatabase> Databases;
        public readonly ImmutableArray<string> Domains;
        /// <summary>
        /// Describes an environment variable made available to an app competent.
        /// </summary>
        public readonly ImmutableArray<Outputs.AppSpecEnv> Envs;
        public readonly ImmutableArray<Outputs.AppSpecJob> Jobs;
        /// <summary>
        /// The name of the component.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The slug for the DigitalOcean data center region hosting the app.
        /// </summary>
        public readonly string? Region;
        public readonly ImmutableArray<Outputs.AppSpecService> Services;
        public readonly ImmutableArray<Outputs.AppSpecStaticSite> StaticSites;
        public readonly ImmutableArray<Outputs.AppSpecWorker> Workers;

        [OutputConstructor]
        private AppSpec(
            ImmutableArray<Outputs.AppSpecDatabase> databases,

            ImmutableArray<string> domains,

            ImmutableArray<Outputs.AppSpecEnv> envs,

            ImmutableArray<Outputs.AppSpecJob> jobs,

            string name,

            string? region,

            ImmutableArray<Outputs.AppSpecService> services,

            ImmutableArray<Outputs.AppSpecStaticSite> staticSites,

            ImmutableArray<Outputs.AppSpecWorker> workers)
        {
            Databases = databases;
            Domains = domains;
            Envs = envs;
            Jobs = jobs;
            Name = name;
            Region = region;
            Services = services;
            StaticSites = staticSites;
            Workers = workers;
        }
    }
}
