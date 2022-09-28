// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetAppArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAppArgs Empty = new GetAppArgs();

    /**
     * The ID of the app to retrieve information about.
     * 
     */
    @Import(name="appId", required=true)
    private Output<String> appId;

    /**
     * @return The ID of the app to retrieve information about.
     * 
     */
    public Output<String> appId() {
        return this.appId;
    }

    private GetAppArgs() {}

    private GetAppArgs(GetAppArgs $) {
        this.appId = $.appId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAppArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAppArgs $;

        public Builder() {
            $ = new GetAppArgs();
        }

        public Builder(GetAppArgs defaults) {
            $ = new GetAppArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param appId The ID of the app to retrieve information about.
         * 
         * @return builder
         * 
         */
        public Builder appId(Output<String> appId) {
            $.appId = appId;
            return this;
        }

        /**
         * @param appId The ID of the app to retrieve information about.
         * 
         * @return builder
         * 
         */
        public Builder appId(String appId) {
            return appId(Output.of(appId));
        }

        public GetAppArgs build() {
            $.appId = Objects.requireNonNull($.appId, "expected parameter 'appId' to be non-null");
            return $;
        }
    }

}