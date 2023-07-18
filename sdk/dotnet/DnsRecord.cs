// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean
{
    /// <summary>
    /// Provides a DigitalOcean DNS record resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using DigitalOcean = Pulumi.DigitalOcean;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new DigitalOcean.Domain("default", new()
    ///     {
    ///         Name = "example.com",
    ///     });
    /// 
    ///     // Add an A record to the domain for www.example.com.
    ///     var www = new DigitalOcean.DnsRecord("www", new()
    ///     {
    ///         Domain = @default.Id,
    ///         Type = "A",
    ///         Value = "192.168.0.11",
    ///     });
    /// 
    ///     // Add a MX record for the example.com domain itself.
    ///     var mx = new DigitalOcean.DnsRecord("mx", new()
    ///     {
    ///         Domain = @default.Id,
    ///         Type = "MX",
    ///         Priority = 10,
    ///         Value = "mail.example.com.",
    ///     });
    /// 
    ///     return new Dictionary&lt;string, object?&gt;
    ///     {
    ///         ["wwwFqdn"] = www.Fqdn,
    ///         ["mxFqdn"] = mx.Fqdn,
    ///     };
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Records can be imported using the domain name and record `id` when joined with a comma. See the following example
    /// 
    /// ```sh
    ///  $ pulumi import digitalocean:index/dnsRecord:DnsRecord example_record example.com,12345678
    /// ```
    /// 
    ///  ~&gt;
    /// 
    /// You find the `id` of the records [using the DigitalOcean API](https://docs.digitalocean.com/reference/api/api-reference/#operation/domains_list_records) or CLI. Run the follow command to list the IDs for all DNS records on a domain`doctl compute domain records list &lt;domain.name&gt;`
    /// </summary>
    [DigitalOceanResourceType("digitalocean:index/dnsRecord:DnsRecord")]
    public partial class DnsRecord : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The domain to add the record to.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// The flags of the record. Only valid when type is `CAA`. Must be between 0 and 255.
        /// </summary>
        [Output("flags")]
        public Output<int?> Flags { get; private set; } = null!;

        /// <summary>
        /// The FQDN of the record
        /// </summary>
        [Output("fqdn")]
        public Output<string> Fqdn { get; private set; } = null!;

        /// <summary>
        /// The hostname of the record. Use `@` for records on domain's name itself.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The port of the record. Only valid when type is `SRV`.  Must be between 1 and 65535.
        /// </summary>
        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        /// <summary>
        /// The priority of the record. Only valid when type is `MX` or `SRV`. Must be between 0 and 65535.
        /// </summary>
        [Output("priority")]
        public Output<int?> Priority { get; private set; } = null!;

        /// <summary>
        /// The tag of the record. Only valid when type is `CAA`. Must be one of `issue`, `issuewild`, or `iodef`.
        /// </summary>
        [Output("tag")]
        public Output<string?> Tag { get; private set; } = null!;

        /// <summary>
        /// The time to live for the record, in seconds. Must be at least 0. Defaults to 1800.
        /// </summary>
        [Output("ttl")]
        public Output<int> Ttl { get; private set; } = null!;

        /// <summary>
        /// The type of record. Must be one of `A`, `AAAA`, `CAA`, `CNAME`, `MX`, `NS`, `TXT`, or `SRV`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The value of the record.
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;

        /// <summary>
        /// The weight of the record. Only valid when type is `SRV`.  Must be between 0 and 65535.
        /// </summary>
        [Output("weight")]
        public Output<int?> Weight { get; private set; } = null!;


        /// <summary>
        /// Create a DnsRecord resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DnsRecord(string name, DnsRecordArgs args, CustomResourceOptions? options = null)
            : base("digitalocean:index/dnsRecord:DnsRecord", name, args ?? new DnsRecordArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DnsRecord(string name, Input<string> id, DnsRecordState? state = null, CustomResourceOptions? options = null)
            : base("digitalocean:index/dnsRecord:DnsRecord", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DnsRecord resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DnsRecord Get(string name, Input<string> id, DnsRecordState? state = null, CustomResourceOptions? options = null)
        {
            return new DnsRecord(name, id, state, options);
        }
    }

    public sealed class DnsRecordArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain to add the record to.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// The flags of the record. Only valid when type is `CAA`. Must be between 0 and 255.
        /// </summary>
        [Input("flags")]
        public Input<int>? Flags { get; set; }

        /// <summary>
        /// The hostname of the record. Use `@` for records on domain's name itself.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The port of the record. Only valid when type is `SRV`.  Must be between 1 and 65535.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The priority of the record. Only valid when type is `MX` or `SRV`. Must be between 0 and 65535.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The tag of the record. Only valid when type is `CAA`. Must be one of `issue`, `issuewild`, or `iodef`.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// The time to live for the record, in seconds. Must be at least 0. Defaults to 1800.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// The type of record. Must be one of `A`, `AAAA`, `CAA`, `CNAME`, `MX`, `NS`, `TXT`, or `SRV`.
        /// </summary>
        [Input("type", required: true)]
        public InputUnion<string, Pulumi.DigitalOcean.RecordType> Type { get; set; } = null!;

        /// <summary>
        /// The value of the record.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        /// <summary>
        /// The weight of the record. Only valid when type is `SRV`.  Must be between 0 and 65535.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public DnsRecordArgs()
        {
        }
        public static new DnsRecordArgs Empty => new DnsRecordArgs();
    }

    public sealed class DnsRecordState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain to add the record to.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The flags of the record. Only valid when type is `CAA`. Must be between 0 and 255.
        /// </summary>
        [Input("flags")]
        public Input<int>? Flags { get; set; }

        /// <summary>
        /// The FQDN of the record
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        /// <summary>
        /// The hostname of the record. Use `@` for records on domain's name itself.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The port of the record. Only valid when type is `SRV`.  Must be between 1 and 65535.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The priority of the record. Only valid when type is `MX` or `SRV`. Must be between 0 and 65535.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The tag of the record. Only valid when type is `CAA`. Must be one of `issue`, `issuewild`, or `iodef`.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// The time to live for the record, in seconds. Must be at least 0. Defaults to 1800.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// The type of record. Must be one of `A`, `AAAA`, `CAA`, `CNAME`, `MX`, `NS`, `TXT`, or `SRV`.
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.DigitalOcean.RecordType>? Type { get; set; }

        /// <summary>
        /// The value of the record.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        /// <summary>
        /// The weight of the record. Only valid when type is `SRV`.  Must be between 0 and 65535.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public DnsRecordState()
        {
        }
        public static new DnsRecordState Empty => new DnsRecordState();
    }
}
