// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.digitalocean.outputs.GetRecordsFilter;
import com.pulumi.digitalocean.outputs.GetRecordsRecord;
import com.pulumi.digitalocean.outputs.GetRecordsSort;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetRecordsResult {
    /**
     * @return Domain of the DNS record.
     * 
     */
    private String domain;
    private @Nullable List<GetRecordsFilter> filters;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<GetRecordsRecord> records;
    private @Nullable List<GetRecordsSort> sorts;

    private GetRecordsResult() {}
    /**
     * @return Domain of the DNS record.
     * 
     */
    public String domain() {
        return this.domain;
    }
    public List<GetRecordsFilter> filters() {
        return this.filters == null ? List.of() : this.filters;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public List<GetRecordsRecord> records() {
        return this.records;
    }
    public List<GetRecordsSort> sorts() {
        return this.sorts == null ? List.of() : this.sorts;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRecordsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String domain;
        private @Nullable List<GetRecordsFilter> filters;
        private String id;
        private List<GetRecordsRecord> records;
        private @Nullable List<GetRecordsSort> sorts;
        public Builder() {}
        public Builder(GetRecordsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.domain = defaults.domain;
    	      this.filters = defaults.filters;
    	      this.id = defaults.id;
    	      this.records = defaults.records;
    	      this.sorts = defaults.sorts;
        }

        @CustomType.Setter
        public Builder domain(String domain) {
            this.domain = Objects.requireNonNull(domain);
            return this;
        }
        @CustomType.Setter
        public Builder filters(@Nullable List<GetRecordsFilter> filters) {
            this.filters = filters;
            return this;
        }
        public Builder filters(GetRecordsFilter... filters) {
            return filters(List.of(filters));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder records(List<GetRecordsRecord> records) {
            this.records = Objects.requireNonNull(records);
            return this;
        }
        public Builder records(GetRecordsRecord... records) {
            return records(List.of(records));
        }
        @CustomType.Setter
        public Builder sorts(@Nullable List<GetRecordsSort> sorts) {
            this.sorts = sorts;
            return this;
        }
        public Builder sorts(GetRecordsSort... sorts) {
            return sorts(List.of(sorts));
        }
        public GetRecordsResult build() {
            final var o = new GetRecordsResult();
            o.domain = domain;
            o.filters = filters;
            o.id = id;
            o.records = records;
            o.sorts = sorts;
            return o;
        }
    }
}
