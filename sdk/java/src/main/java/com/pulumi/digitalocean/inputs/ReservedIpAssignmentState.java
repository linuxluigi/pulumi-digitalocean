// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ReservedIpAssignmentState extends com.pulumi.resources.ResourceArgs {

    public static final ReservedIpAssignmentState Empty = new ReservedIpAssignmentState();

    /**
     * The ID of Droplet that the reserved IP will be assigned to.
     * 
     */
    @Import(name="dropletId")
    private @Nullable Output<Integer> dropletId;

    /**
     * @return The ID of Droplet that the reserved IP will be assigned to.
     * 
     */
    public Optional<Output<Integer>> dropletId() {
        return Optional.ofNullable(this.dropletId);
    }

    /**
     * The reserved IP to assign to the Droplet.
     * 
     */
    @Import(name="ipAddress")
    private @Nullable Output<String> ipAddress;

    /**
     * @return The reserved IP to assign to the Droplet.
     * 
     */
    public Optional<Output<String>> ipAddress() {
        return Optional.ofNullable(this.ipAddress);
    }

    private ReservedIpAssignmentState() {}

    private ReservedIpAssignmentState(ReservedIpAssignmentState $) {
        this.dropletId = $.dropletId;
        this.ipAddress = $.ipAddress;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ReservedIpAssignmentState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ReservedIpAssignmentState $;

        public Builder() {
            $ = new ReservedIpAssignmentState();
        }

        public Builder(ReservedIpAssignmentState defaults) {
            $ = new ReservedIpAssignmentState(Objects.requireNonNull(defaults));
        }

        /**
         * @param dropletId The ID of Droplet that the reserved IP will be assigned to.
         * 
         * @return builder
         * 
         */
        public Builder dropletId(@Nullable Output<Integer> dropletId) {
            $.dropletId = dropletId;
            return this;
        }

        /**
         * @param dropletId The ID of Droplet that the reserved IP will be assigned to.
         * 
         * @return builder
         * 
         */
        public Builder dropletId(Integer dropletId) {
            return dropletId(Output.of(dropletId));
        }

        /**
         * @param ipAddress The reserved IP to assign to the Droplet.
         * 
         * @return builder
         * 
         */
        public Builder ipAddress(@Nullable Output<String> ipAddress) {
            $.ipAddress = ipAddress;
            return this;
        }

        /**
         * @param ipAddress The reserved IP to assign to the Droplet.
         * 
         * @return builder
         * 
         */
        public Builder ipAddress(String ipAddress) {
            return ipAddress(Output.of(ipAddress));
        }

        public ReservedIpAssignmentState build() {
            return $;
        }
    }

}