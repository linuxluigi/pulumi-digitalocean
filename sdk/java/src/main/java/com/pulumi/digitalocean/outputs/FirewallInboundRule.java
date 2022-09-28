// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class FirewallInboundRule {
    /**
     * @return The ports on which traffic will be allowed
     * specified as a string containing a single port, a range (e.g. &#34;8000-9000&#34;),
     * or &#34;1-65535&#34; to open all ports for a protocol. Required for when protocol is
     * `tcp` or `udp`.
     * 
     */
    private @Nullable String portRange;
    /**
     * @return The type of traffic to be allowed.
     * This may be one of &#34;tcp&#34;, &#34;udp&#34;, or &#34;icmp&#34;.
     * 
     */
    private String protocol;
    /**
     * @return An array of strings containing the IPv4
     * addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs from which the
     * inbound traffic will be accepted.
     * 
     */
    private @Nullable List<String> sourceAddresses;
    /**
     * @return An array containing the IDs of
     * the Droplets from which the inbound traffic will be accepted.
     * 
     */
    private @Nullable List<Integer> sourceDropletIds;
    /**
     * @return An array containing the IDs of
     * the Kubernetes clusters from which the inbound traffic will be accepted.
     * 
     */
    private @Nullable List<String> sourceKubernetesIds;
    /**
     * @return An array containing the IDs
     * of the Load Balancers from which the inbound traffic will be accepted.
     * 
     */
    private @Nullable List<String> sourceLoadBalancerUids;
    /**
     * @return An array containing the names of Tags
     * corresponding to groups of Droplets from which the inbound traffic
     * will be accepted.
     * 
     */
    private @Nullable List<String> sourceTags;

    private FirewallInboundRule() {}
    /**
     * @return The ports on which traffic will be allowed
     * specified as a string containing a single port, a range (e.g. &#34;8000-9000&#34;),
     * or &#34;1-65535&#34; to open all ports for a protocol. Required for when protocol is
     * `tcp` or `udp`.
     * 
     */
    public Optional<String> portRange() {
        return Optional.ofNullable(this.portRange);
    }
    /**
     * @return The type of traffic to be allowed.
     * This may be one of &#34;tcp&#34;, &#34;udp&#34;, or &#34;icmp&#34;.
     * 
     */
    public String protocol() {
        return this.protocol;
    }
    /**
     * @return An array of strings containing the IPv4
     * addresses, IPv6 addresses, IPv4 CIDRs, and/or IPv6 CIDRs from which the
     * inbound traffic will be accepted.
     * 
     */
    public List<String> sourceAddresses() {
        return this.sourceAddresses == null ? List.of() : this.sourceAddresses;
    }
    /**
     * @return An array containing the IDs of
     * the Droplets from which the inbound traffic will be accepted.
     * 
     */
    public List<Integer> sourceDropletIds() {
        return this.sourceDropletIds == null ? List.of() : this.sourceDropletIds;
    }
    /**
     * @return An array containing the IDs of
     * the Kubernetes clusters from which the inbound traffic will be accepted.
     * 
     */
    public List<String> sourceKubernetesIds() {
        return this.sourceKubernetesIds == null ? List.of() : this.sourceKubernetesIds;
    }
    /**
     * @return An array containing the IDs
     * of the Load Balancers from which the inbound traffic will be accepted.
     * 
     */
    public List<String> sourceLoadBalancerUids() {
        return this.sourceLoadBalancerUids == null ? List.of() : this.sourceLoadBalancerUids;
    }
    /**
     * @return An array containing the names of Tags
     * corresponding to groups of Droplets from which the inbound traffic
     * will be accepted.
     * 
     */
    public List<String> sourceTags() {
        return this.sourceTags == null ? List.of() : this.sourceTags;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FirewallInboundRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String portRange;
        private String protocol;
        private @Nullable List<String> sourceAddresses;
        private @Nullable List<Integer> sourceDropletIds;
        private @Nullable List<String> sourceKubernetesIds;
        private @Nullable List<String> sourceLoadBalancerUids;
        private @Nullable List<String> sourceTags;
        public Builder() {}
        public Builder(FirewallInboundRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.portRange = defaults.portRange;
    	      this.protocol = defaults.protocol;
    	      this.sourceAddresses = defaults.sourceAddresses;
    	      this.sourceDropletIds = defaults.sourceDropletIds;
    	      this.sourceKubernetesIds = defaults.sourceKubernetesIds;
    	      this.sourceLoadBalancerUids = defaults.sourceLoadBalancerUids;
    	      this.sourceTags = defaults.sourceTags;
        }

        @CustomType.Setter
        public Builder portRange(@Nullable String portRange) {
            this.portRange = portRange;
            return this;
        }
        @CustomType.Setter
        public Builder protocol(String protocol) {
            this.protocol = Objects.requireNonNull(protocol);
            return this;
        }
        @CustomType.Setter
        public Builder sourceAddresses(@Nullable List<String> sourceAddresses) {
            this.sourceAddresses = sourceAddresses;
            return this;
        }
        public Builder sourceAddresses(String... sourceAddresses) {
            return sourceAddresses(List.of(sourceAddresses));
        }
        @CustomType.Setter
        public Builder sourceDropletIds(@Nullable List<Integer> sourceDropletIds) {
            this.sourceDropletIds = sourceDropletIds;
            return this;
        }
        public Builder sourceDropletIds(Integer... sourceDropletIds) {
            return sourceDropletIds(List.of(sourceDropletIds));
        }
        @CustomType.Setter
        public Builder sourceKubernetesIds(@Nullable List<String> sourceKubernetesIds) {
            this.sourceKubernetesIds = sourceKubernetesIds;
            return this;
        }
        public Builder sourceKubernetesIds(String... sourceKubernetesIds) {
            return sourceKubernetesIds(List.of(sourceKubernetesIds));
        }
        @CustomType.Setter
        public Builder sourceLoadBalancerUids(@Nullable List<String> sourceLoadBalancerUids) {
            this.sourceLoadBalancerUids = sourceLoadBalancerUids;
            return this;
        }
        public Builder sourceLoadBalancerUids(String... sourceLoadBalancerUids) {
            return sourceLoadBalancerUids(List.of(sourceLoadBalancerUids));
        }
        @CustomType.Setter
        public Builder sourceTags(@Nullable List<String> sourceTags) {
            this.sourceTags = sourceTags;
            return this;
        }
        public Builder sourceTags(String... sourceTags) {
            return sourceTags(List.of(sourceTags));
        }
        public FirewallInboundRule build() {
            final var o = new FirewallInboundRule();
            o.portRange = portRange;
            o.protocol = protocol;
            o.sourceAddresses = sourceAddresses;
            o.sourceDropletIds = sourceDropletIds;
            o.sourceKubernetesIds = sourceKubernetesIds;
            o.sourceLoadBalancerUids = sourceLoadBalancerUids;
            o.sourceTags = sourceTags;
            return o;
        }
    }
}