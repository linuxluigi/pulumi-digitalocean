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
    public sealed class DatabaseUserSetting
    {
        /// <summary>
        /// A set of ACLs (Access Control Lists) specifying permission on topics with a Kafka cluster. The properties of an individual ACL are described below:
        /// 
        /// An individual ACL includes the following:
        /// </summary>
        public readonly ImmutableArray<Outputs.DatabaseUserSettingAcl> Acls;

        [OutputConstructor]
        private DatabaseUserSetting(ImmutableArray<Outputs.DatabaseUserSettingAcl> acls)
        {
            Acls = acls;
        }
    }
}