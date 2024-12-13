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
    public sealed class DropletAutoscaleCurrentUtilization
    {
        /// <summary>
        /// Average CPU utilization
        /// </summary>
        public readonly double? Cpu;
        /// <summary>
        /// Average Memory utilization
        /// </summary>
        public readonly double? Memory;

        [OutputConstructor]
        private DropletAutoscaleCurrentUtilization(
            double? cpu,

            double? memory)
        {
            Cpu = cpu;
            Memory = memory;
        }
    }
}
