// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.DigitalOcean
{
    public static class GetRecord
    {
        /// <summary>
        /// Get information on a DNS record. This data source provides the name, TTL, and zone
        /// file as configured on your DigitalOcean account. This is useful if the record
        /// in question is not managed by the provider.
        /// 
        /// An error is triggered if the provided domain name or record are not managed with
        /// your DigitalOcean account.
        /// </summary>
        public static Task<GetRecordResult> InvokeAsync(GetRecordArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRecordResult>("digitalocean:index/getRecord:getRecord", args ?? new GetRecordArgs(), options.WithVersion());

        /// <summary>
        /// Get information on a DNS record. This data source provides the name, TTL, and zone
        /// file as configured on your DigitalOcean account. This is useful if the record
        /// in question is not managed by the provider.
        /// 
        /// An error is triggered if the provided domain name or record are not managed with
        /// your DigitalOcean account.
        /// </summary>
        public static Output<GetRecordResult> Invoke(GetRecordInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRecordResult>("digitalocean:index/getRecord:getRecord", args ?? new GetRecordInvokeArgs(), options.WithVersion());
    }


    public sealed class GetRecordArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The domain name of the record.
        /// </summary>
        [Input("domain", required: true)]
        public string Domain { get; set; } = null!;

        /// <summary>
        /// The name of the record.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetRecordArgs()
        {
        }
    }

    public sealed class GetRecordInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The domain name of the record.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// The name of the record.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetRecordInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRecordResult
    {
        public readonly string Data;
        public readonly string Domain;
        public readonly int Flags;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly int Port;
        public readonly int Priority;
        public readonly string Tag;
        public readonly int Ttl;
        public readonly string Type;
        public readonly int Weight;

        [OutputConstructor]
        private GetRecordResult(
            string data,

            string domain,

            int flags,

            string id,

            string name,

            int port,

            int priority,

            string tag,

            int ttl,

            string type,

            int weight)
        {
            Data = data;
            Domain = domain;
            Flags = flags;
            Id = id;
            Name = name;
            Port = port;
            Priority = priority;
            Tag = tag;
            Ttl = ttl;
            Type = type;
            Weight = weight;
        }
    }
}
