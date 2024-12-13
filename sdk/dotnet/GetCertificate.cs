// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean
{
    public static class GetCertificate
    {
        /// <summary>
        /// Get information on a certificate. This data source provides the name, type, state,
        /// domains, expiry date, and the sha1 fingerprint as configured on your DigitalOcean account.
        /// This is useful if the certificate in question is not managed by this provider or you need to utilize
        /// any of the certificates data.
        /// 
        /// An error is triggered if the provided certificate name does not exist.
        /// 
        /// ## Example Usage
        /// 
        /// Get the certificate:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = DigitalOcean.GetCertificate.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetCertificateResult> InvokeAsync(GetCertificateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCertificateResult>("digitalocean:index/getCertificate:getCertificate", args ?? new GetCertificateArgs(), options.WithDefaults());

        /// <summary>
        /// Get information on a certificate. This data source provides the name, type, state,
        /// domains, expiry date, and the sha1 fingerprint as configured on your DigitalOcean account.
        /// This is useful if the certificate in question is not managed by this provider or you need to utilize
        /// any of the certificates data.
        /// 
        /// An error is triggered if the provided certificate name does not exist.
        /// 
        /// ## Example Usage
        /// 
        /// Get the certificate:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = DigitalOcean.GetCertificate.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetCertificateResult> Invoke(GetCertificateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCertificateResult>("digitalocean:index/getCertificate:getCertificate", args ?? new GetCertificateInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get information on a certificate. This data source provides the name, type, state,
        /// domains, expiry date, and the sha1 fingerprint as configured on your DigitalOcean account.
        /// This is useful if the certificate in question is not managed by this provider or you need to utilize
        /// any of the certificates data.
        /// 
        /// An error is triggered if the provided certificate name does not exist.
        /// 
        /// ## Example Usage
        /// 
        /// Get the certificate:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using DigitalOcean = Pulumi.DigitalOcean;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = DigitalOcean.GetCertificate.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetCertificateResult> Invoke(GetCertificateInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCertificateResult>("digitalocean:index/getCertificate:getCertificate", args ?? new GetCertificateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCertificateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of certificate.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetCertificateArgs()
        {
        }
        public static new GetCertificateArgs Empty => new GetCertificateArgs();
    }

    public sealed class GetCertificateInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of certificate.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetCertificateInvokeArgs()
        {
        }
        public static new GetCertificateInvokeArgs Empty => new GetCertificateInvokeArgs();
    }


    [OutputType]
    public sealed class GetCertificateResult
    {
        /// <summary>
        /// Domains for which the certificate was issued.
        /// </summary>
        public readonly ImmutableArray<string> Domains;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// The expiration date and time of the certificate.
        /// </summary>
        public readonly string NotAfter;
        /// <summary>
        /// The SHA1 fingerprint of the certificate.
        /// </summary>
        public readonly string Sha1Fingerprint;
        /// <summary>
        /// the current state of the certificate.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The type of the certificate.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The ID of the certificate.
        /// </summary>
        public readonly string Uuid;

        [OutputConstructor]
        private GetCertificateResult(
            ImmutableArray<string> domains,

            string id,

            string name,

            string notAfter,

            string sha1Fingerprint,

            string state,

            string type,

            string uuid)
        {
            Domains = domains;
            Id = id;
            Name = name;
            NotAfter = notAfter;
            Sha1Fingerprint = sha1Fingerprint;
            State = state;
            Type = type;
            Uuid = uuid;
        }
    }
}
