// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.digitalocean.outputs.GetAppSpec;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetAppResult {
    /**
     * @return The ID the app&#39;s currently active deployment.
     * 
     */
    private String activeDeploymentId;
    private String appId;
    /**
     * @return The date and time of when the app was created.
     * 
     */
    private String createdAt;
    /**
     * @return The default URL to access the app.
     * 
     */
    private String defaultIngress;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The live URL of the app.
     * 
     */
    private String liveUrl;
    /**
     * @return A DigitalOcean App spec describing the app.
     * 
     */
    private List<GetAppSpec> specs;
    /**
     * @return The date and time of when the app was last updated.
     * 
     */
    private String updatedAt;
    /**
     * @return The uniform resource identifier for the app.
     * 
     */
    private String urn;

    private GetAppResult() {}
    /**
     * @return The ID the app&#39;s currently active deployment.
     * 
     */
    public String activeDeploymentId() {
        return this.activeDeploymentId;
    }
    public String appId() {
        return this.appId;
    }
    /**
     * @return The date and time of when the app was created.
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return The default URL to access the app.
     * 
     */
    public String defaultIngress() {
        return this.defaultIngress;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The live URL of the app.
     * 
     */
    public String liveUrl() {
        return this.liveUrl;
    }
    /**
     * @return A DigitalOcean App spec describing the app.
     * 
     */
    public List<GetAppSpec> specs() {
        return this.specs;
    }
    /**
     * @return The date and time of when the app was last updated.
     * 
     */
    public String updatedAt() {
        return this.updatedAt;
    }
    /**
     * @return The uniform resource identifier for the app.
     * 
     */
    public String urn() {
        return this.urn;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAppResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String activeDeploymentId;
        private String appId;
        private String createdAt;
        private String defaultIngress;
        private String id;
        private String liveUrl;
        private List<GetAppSpec> specs;
        private String updatedAt;
        private String urn;
        public Builder() {}
        public Builder(GetAppResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.activeDeploymentId = defaults.activeDeploymentId;
    	      this.appId = defaults.appId;
    	      this.createdAt = defaults.createdAt;
    	      this.defaultIngress = defaults.defaultIngress;
    	      this.id = defaults.id;
    	      this.liveUrl = defaults.liveUrl;
    	      this.specs = defaults.specs;
    	      this.updatedAt = defaults.updatedAt;
    	      this.urn = defaults.urn;
        }

        @CustomType.Setter
        public Builder activeDeploymentId(String activeDeploymentId) {
            this.activeDeploymentId = Objects.requireNonNull(activeDeploymentId);
            return this;
        }
        @CustomType.Setter
        public Builder appId(String appId) {
            this.appId = Objects.requireNonNull(appId);
            return this;
        }
        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            this.createdAt = Objects.requireNonNull(createdAt);
            return this;
        }
        @CustomType.Setter
        public Builder defaultIngress(String defaultIngress) {
            this.defaultIngress = Objects.requireNonNull(defaultIngress);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder liveUrl(String liveUrl) {
            this.liveUrl = Objects.requireNonNull(liveUrl);
            return this;
        }
        @CustomType.Setter
        public Builder specs(List<GetAppSpec> specs) {
            this.specs = Objects.requireNonNull(specs);
            return this;
        }
        public Builder specs(GetAppSpec... specs) {
            return specs(List.of(specs));
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            this.updatedAt = Objects.requireNonNull(updatedAt);
            return this;
        }
        @CustomType.Setter
        public Builder urn(String urn) {
            this.urn = Objects.requireNonNull(urn);
            return this;
        }
        public GetAppResult build() {
            final var o = new GetAppResult();
            o.activeDeploymentId = activeDeploymentId;
            o.appId = appId;
            o.createdAt = createdAt;
            o.defaultIngress = defaultIngress;
            o.id = id;
            o.liveUrl = liveUrl;
            o.specs = specs;
            o.updatedAt = updatedAt;
            o.urn = urn;
            return o;
        }
    }
}