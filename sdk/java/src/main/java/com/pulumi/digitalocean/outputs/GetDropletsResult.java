// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.digitalocean.outputs.GetDropletsDroplet;
import com.pulumi.digitalocean.outputs.GetDropletsFilter;
import com.pulumi.digitalocean.outputs.GetDropletsSort;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetDropletsResult {
    /**
     * @return A list of Droplets satisfying any `filter` and `sort` criteria. Each Droplet has the following attributes:
     * 
     */
    private List<GetDropletsDroplet> droplets;
    private @Nullable List<GetDropletsFilter> filters;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable List<GetDropletsSort> sorts;

    private GetDropletsResult() {}
    /**
     * @return A list of Droplets satisfying any `filter` and `sort` criteria. Each Droplet has the following attributes:
     * 
     */
    public List<GetDropletsDroplet> droplets() {
        return this.droplets;
    }
    public List<GetDropletsFilter> filters() {
        return this.filters == null ? List.of() : this.filters;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public List<GetDropletsSort> sorts() {
        return this.sorts == null ? List.of() : this.sorts;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDropletsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetDropletsDroplet> droplets;
        private @Nullable List<GetDropletsFilter> filters;
        private String id;
        private @Nullable List<GetDropletsSort> sorts;
        public Builder() {}
        public Builder(GetDropletsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.droplets = defaults.droplets;
    	      this.filters = defaults.filters;
    	      this.id = defaults.id;
    	      this.sorts = defaults.sorts;
        }

        @CustomType.Setter
        public Builder droplets(List<GetDropletsDroplet> droplets) {
            this.droplets = Objects.requireNonNull(droplets);
            return this;
        }
        public Builder droplets(GetDropletsDroplet... droplets) {
            return droplets(List.of(droplets));
        }
        @CustomType.Setter
        public Builder filters(@Nullable List<GetDropletsFilter> filters) {
            this.filters = filters;
            return this;
        }
        public Builder filters(GetDropletsFilter... filters) {
            return filters(List.of(filters));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder sorts(@Nullable List<GetDropletsSort> sorts) {
            this.sorts = sorts;
            return this;
        }
        public Builder sorts(GetDropletsSort... sorts) {
            return sorts(List.of(sorts));
        }
        public GetDropletsResult build() {
            final var o = new GetDropletsResult();
            o.droplets = droplets;
            o.filters = filters;
            o.id = id;
            o.sorts = sorts;
            return o;
        }
    }
}