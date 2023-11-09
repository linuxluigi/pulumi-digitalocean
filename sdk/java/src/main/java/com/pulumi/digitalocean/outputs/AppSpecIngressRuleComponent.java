// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class AppSpecIngressRuleComponent {
    /**
     * @return The name of the component.
     * 
     */
    private @Nullable String name;
    /**
     * @return An optional flag to preserve the path that is forwarded to the backend service.
     * 
     */
    private @Nullable Boolean preservePathPrefix;
    /**
     * @return An optional field that will rewrite the path of the component to be what is specified here. This is mutually exclusive with `preserve_path_prefix`.
     * 
     */
    private @Nullable String rewrite;

    private AppSpecIngressRuleComponent() {}
    /**
     * @return The name of the component.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return An optional flag to preserve the path that is forwarded to the backend service.
     * 
     */
    public Optional<Boolean> preservePathPrefix() {
        return Optional.ofNullable(this.preservePathPrefix);
    }
    /**
     * @return An optional field that will rewrite the path of the component to be what is specified here. This is mutually exclusive with `preserve_path_prefix`.
     * 
     */
    public Optional<String> rewrite() {
        return Optional.ofNullable(this.rewrite);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AppSpecIngressRuleComponent defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String name;
        private @Nullable Boolean preservePathPrefix;
        private @Nullable String rewrite;
        public Builder() {}
        public Builder(AppSpecIngressRuleComponent defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.preservePathPrefix = defaults.preservePathPrefix;
    	      this.rewrite = defaults.rewrite;
        }

        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder preservePathPrefix(@Nullable Boolean preservePathPrefix) {
            this.preservePathPrefix = preservePathPrefix;
            return this;
        }
        @CustomType.Setter
        public Builder rewrite(@Nullable String rewrite) {
            this.rewrite = rewrite;
            return this;
        }
        public AppSpecIngressRuleComponent build() {
            final var _resultValue = new AppSpecIngressRuleComponent();
            _resultValue.name = name;
            _resultValue.preservePathPrefix = preservePathPrefix;
            _resultValue.rewrite = rewrite;
            return _resultValue;
        }
    }
}