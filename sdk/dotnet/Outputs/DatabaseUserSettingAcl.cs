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
    public sealed class DatabaseUserSettingAcl
    {
        /// <summary>
        /// An identifier for the ACL, this will be automatically assigned when you create an ACL entry
        /// </summary>
        public readonly string? Id;
        public readonly string Permission;
        public readonly string Topic;

        [OutputConstructor]
        private DatabaseUserSettingAcl(
            string? id,

            string permission,

            string topic)
        {
            Id = id;
            Permission = permission;
            Topic = topic;
        }
    }
}
