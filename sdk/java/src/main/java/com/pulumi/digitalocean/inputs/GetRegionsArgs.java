// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.digitalocean.inputs.GetRegionsFilterArgs;
import com.pulumi.digitalocean.inputs.GetRegionsSortArgs;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRegionsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRegionsArgs Empty = new GetRegionsArgs();

    /**
     * Filter the results.
     * The `filter` block is documented below.
     * 
     */
    @Import(name="filters")
    private @Nullable Output<List<GetRegionsFilterArgs>> filters;

    /**
     * @return Filter the results.
     * The `filter` block is documented below.
     * 
     */
    public Optional<Output<List<GetRegionsFilterArgs>>> filters() {
        return Optional.ofNullable(this.filters);
    }

    /**
     * Sort the results.
     * The `sort` block is documented below.
     * 
     */
    @Import(name="sorts")
    private @Nullable Output<List<GetRegionsSortArgs>> sorts;

    /**
     * @return Sort the results.
     * The `sort` block is documented below.
     * 
     */
    public Optional<Output<List<GetRegionsSortArgs>>> sorts() {
        return Optional.ofNullable(this.sorts);
    }

    private GetRegionsArgs() {}

    private GetRegionsArgs(GetRegionsArgs $) {
        this.filters = $.filters;
        this.sorts = $.sorts;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRegionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRegionsArgs $;

        public Builder() {
            $ = new GetRegionsArgs();
        }

        public Builder(GetRegionsArgs defaults) {
            $ = new GetRegionsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filters Filter the results.
         * The `filter` block is documented below.
         * 
         * @return builder
         * 
         */
        public Builder filters(@Nullable Output<List<GetRegionsFilterArgs>> filters) {
            $.filters = filters;
            return this;
        }

        /**
         * @param filters Filter the results.
         * The `filter` block is documented below.
         * 
         * @return builder
         * 
         */
        public Builder filters(List<GetRegionsFilterArgs> filters) {
            return filters(Output.of(filters));
        }

        /**
         * @param filters Filter the results.
         * The `filter` block is documented below.
         * 
         * @return builder
         * 
         */
        public Builder filters(GetRegionsFilterArgs... filters) {
            return filters(List.of(filters));
        }

        /**
         * @param sorts Sort the results.
         * The `sort` block is documented below.
         * 
         * @return builder
         * 
         */
        public Builder sorts(@Nullable Output<List<GetRegionsSortArgs>> sorts) {
            $.sorts = sorts;
            return this;
        }

        /**
         * @param sorts Sort the results.
         * The `sort` block is documented below.
         * 
         * @return builder
         * 
         */
        public Builder sorts(List<GetRegionsSortArgs> sorts) {
            return sorts(Output.of(sorts));
        }

        /**
         * @param sorts Sort the results.
         * The `sort` block is documented below.
         * 
         * @return builder
         * 
         */
        public Builder sorts(GetRegionsSortArgs... sorts) {
            return sorts(List.of(sorts));
        }

        public GetRegionsArgs build() {
            return $;
        }
    }

}