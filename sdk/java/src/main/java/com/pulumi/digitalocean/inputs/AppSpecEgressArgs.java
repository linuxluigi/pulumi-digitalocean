// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AppSpecEgressArgs extends com.pulumi.resources.ResourceArgs {

    public static final AppSpecEgressArgs Empty = new AppSpecEgressArgs();

    /**
     * The app egress type: `AUTOASSIGN`, `DEDICATED_IP`
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The app egress type: `AUTOASSIGN`, `DEDICATED_IP`
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private AppSpecEgressArgs() {}

    private AppSpecEgressArgs(AppSpecEgressArgs $) {
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppSpecEgressArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppSpecEgressArgs $;

        public Builder() {
            $ = new AppSpecEgressArgs();
        }

        public Builder(AppSpecEgressArgs defaults) {
            $ = new AppSpecEgressArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param type The app egress type: `AUTOASSIGN`, `DEDICATED_IP`
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The app egress type: `AUTOASSIGN`, `DEDICATED_IP`
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public AppSpecEgressArgs build() {
            return $;
        }
    }

}