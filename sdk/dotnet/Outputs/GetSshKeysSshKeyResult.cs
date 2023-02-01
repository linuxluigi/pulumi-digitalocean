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
    public sealed class GetSshKeysSshKeyResult
    {
        /// <summary>
        /// The fingerprint of the public key of the ssh key.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// The ID of the ssh key.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// The name of the ssh key.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The public key of the ssh key.
        /// </summary>
        public readonly string PublicKey;

        [OutputConstructor]
        private GetSshKeysSshKeyResult(
            string fingerprint,

            int id,

            string name,

            string publicKey)
        {
            Fingerprint = fingerprint;
            Id = id;
            Name = name;
            PublicKey = publicKey;
        }
    }
}
