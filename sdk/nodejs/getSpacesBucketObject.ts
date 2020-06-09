// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The Spaces object data source allows access to the metadata and
 * _optionally_ (see below) content of an object stored inside a Spaces bucket.
 *
 * > **Note:** The content of an object (`body` field) is available only for objects which have a human-readable
 * `Content-Type` (`text/*` and `application/json`). This is to prevent printing unsafe characters and potentially
 * downloading large amount of data which would be thrown away in favor of metadata.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const bootstrapScript = digitalocean.getSpacesBucketObject({
 *     bucket: "ourcorp-deploy-config",
 *     region: "nyc3",
 *     key: "droplet-bootstrap-script.sh",
 * });
 * const web = new digitalocean.Droplet("web", {
 *     image: "ubuntu-18-04-x64",
 *     region: "nyc2",
 *     size: "s-1vcpu-1gb",
 *     userData: bootstrapScript.then(bootstrapScript => bootstrapScript.body),
 * });
 * ```
 */
export function getSpacesBucketObject(args: GetSpacesBucketObjectArgs, opts?: pulumi.InvokeOptions): Promise<GetSpacesBucketObjectResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("digitalocean:index/getSpacesBucketObject:getSpacesBucketObject", {
        "bucket": args.bucket,
        "key": args.key,
        "range": args.range,
        "region": args.region,
        "versionId": args.versionId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSpacesBucketObject.
 */
export interface GetSpacesBucketObjectArgs {
    /**
     * The name of the bucket to read the object from.
     */
    readonly bucket: string;
    /**
     * The full path to the object inside the bucket
     */
    readonly key: string;
    readonly range?: string;
    /**
     * The slug of the region where the bucket is stored.
     */
    readonly region: string;
    /**
     * Specific version ID of the object returned (defaults to latest version)
     */
    readonly versionId?: string;
}

/**
 * A collection of values returned by getSpacesBucketObject.
 */
export interface GetSpacesBucketObjectResult {
    /**
     * Object data (see **limitations above** to understand cases in which this field is actually available)
     */
    readonly body: string;
    readonly bucket: string;
    /**
     * Specifies caching behavior along the request/reply chain.
     */
    readonly cacheControl: string;
    /**
     * Specifies presentational information for the object.
     */
    readonly contentDisposition: string;
    /**
     * Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.
     */
    readonly contentEncoding: string;
    /**
     * The language the content is in.
     */
    readonly contentLanguage: string;
    /**
     * Size of the body in bytes.
     */
    readonly contentLength: number;
    /**
     * A standard MIME type describing the format of the object data.
     */
    readonly contentType: string;
    /**
     * [ETag](https://en.wikipedia.org/wiki/HTTP_ETag) generated for the object (an MD5 sum of the object content in case it's not encrypted)
     */
    readonly etag: string;
    /**
     * If the object expiration is configured (see [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html)), the field includes this header. It includes the expiry-date and rule-id key value pairs providing object expiration information. The value of the rule-id is URL encoded.
     */
    readonly expiration: string;
    /**
     * The date and time at which the object is no longer cacheable.
     */
    readonly expires: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly key: string;
    /**
     * Last modified date of the object in RFC1123 format (e.g. `Mon, 02 Jan 2006 15:04:05 MST`)
     */
    readonly lastModified: string;
    /**
     * A map of metadata stored with the object in Spaces
     */
    readonly metadata: {[key: string]: any};
    readonly range?: string;
    readonly region: string;
    /**
     * The latest version ID of the object returned.
     */
    readonly versionId: string;
    /**
     * If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. Spaces stores the value of this header in the object metadata.
     */
    readonly websiteRedirectLocation: string;
}
