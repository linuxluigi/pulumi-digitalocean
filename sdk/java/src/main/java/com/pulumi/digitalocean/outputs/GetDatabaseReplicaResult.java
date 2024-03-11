// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetDatabaseReplicaResult {
    private String clusterId;
    /**
     * @return Name of the replica&#39;s default database.
     * 
     */
    private String database;
    /**
     * @return Database replica&#39;s hostname.
     * 
     */
    private String host;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String name;
    /**
     * @return Password for the replica&#39;s default user.
     * 
     */
    private String password;
    /**
     * @return Network port that the database replica is listening on.
     * 
     */
    private Integer port;
    /**
     * @return Same as `host`, but only accessible from resources within the account and in the same region.
     * 
     */
    private String privateHost;
    private String privateNetworkUuid;
    /**
     * @return Same as `uri`, but only accessible from resources within the account and in the same region.
     * 
     */
    private String privateUri;
    private String region;
    private String storageSizeMib;
    /**
     * @return A list of tag names to be applied to the database replica.
     * 
     */
    private @Nullable List<String> tags;
    /**
     * @return The full URI for connecting to the database replica.
     * 
     */
    private String uri;
    /**
     * @return Username for the replica&#39;s default user.
     * 
     */
    private String user;
    /**
     * @return The UUID of the database replica.
     * 
     */
    private String uuid;

    private GetDatabaseReplicaResult() {}
    public String clusterId() {
        return this.clusterId;
    }
    /**
     * @return Name of the replica&#39;s default database.
     * 
     */
    public String database() {
        return this.database;
    }
    /**
     * @return Database replica&#39;s hostname.
     * 
     */
    public String host() {
        return this.host;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String name() {
        return this.name;
    }
    /**
     * @return Password for the replica&#39;s default user.
     * 
     */
    public String password() {
        return this.password;
    }
    /**
     * @return Network port that the database replica is listening on.
     * 
     */
    public Integer port() {
        return this.port;
    }
    /**
     * @return Same as `host`, but only accessible from resources within the account and in the same region.
     * 
     */
    public String privateHost() {
        return this.privateHost;
    }
    public String privateNetworkUuid() {
        return this.privateNetworkUuid;
    }
    /**
     * @return Same as `uri`, but only accessible from resources within the account and in the same region.
     * 
     */
    public String privateUri() {
        return this.privateUri;
    }
    public String region() {
        return this.region;
    }
    public String storageSizeMib() {
        return this.storageSizeMib;
    }
    /**
     * @return A list of tag names to be applied to the database replica.
     * 
     */
    public List<String> tags() {
        return this.tags == null ? List.of() : this.tags;
    }
    /**
     * @return The full URI for connecting to the database replica.
     * 
     */
    public String uri() {
        return this.uri;
    }
    /**
     * @return Username for the replica&#39;s default user.
     * 
     */
    public String user() {
        return this.user;
    }
    /**
     * @return The UUID of the database replica.
     * 
     */
    public String uuid() {
        return this.uuid;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDatabaseReplicaResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clusterId;
        private String database;
        private String host;
        private String id;
        private String name;
        private String password;
        private Integer port;
        private String privateHost;
        private String privateNetworkUuid;
        private String privateUri;
        private String region;
        private String storageSizeMib;
        private @Nullable List<String> tags;
        private String uri;
        private String user;
        private String uuid;
        public Builder() {}
        public Builder(GetDatabaseReplicaResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterId = defaults.clusterId;
    	      this.database = defaults.database;
    	      this.host = defaults.host;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.password = defaults.password;
    	      this.port = defaults.port;
    	      this.privateHost = defaults.privateHost;
    	      this.privateNetworkUuid = defaults.privateNetworkUuid;
    	      this.privateUri = defaults.privateUri;
    	      this.region = defaults.region;
    	      this.storageSizeMib = defaults.storageSizeMib;
    	      this.tags = defaults.tags;
    	      this.uri = defaults.uri;
    	      this.user = defaults.user;
    	      this.uuid = defaults.uuid;
        }

        @CustomType.Setter
        public Builder clusterId(String clusterId) {
            if (clusterId == null) {
              throw new MissingRequiredPropertyException("GetDatabaseReplicaResult", "clusterId");
            }
            this.clusterId = clusterId;
            return this;
        }
        @CustomType.Setter
        public Builder database(String database) {
            if (database == null) {
              throw new MissingRequiredPropertyException("GetDatabaseReplicaResult", "database");
            }
            this.database = database;
            return this;
        }
        @CustomType.Setter
        public Builder host(String host) {
            if (host == null) {
              throw new MissingRequiredPropertyException("GetDatabaseReplicaResult", "host");
            }
            this.host = host;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetDatabaseReplicaResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetDatabaseReplicaResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder password(String password) {
            if (password == null) {
              throw new MissingRequiredPropertyException("GetDatabaseReplicaResult", "password");
            }
            this.password = password;
            return this;
        }
        @CustomType.Setter
        public Builder port(Integer port) {
            if (port == null) {
              throw new MissingRequiredPropertyException("GetDatabaseReplicaResult", "port");
            }
            this.port = port;
            return this;
        }
        @CustomType.Setter
        public Builder privateHost(String privateHost) {
            if (privateHost == null) {
              throw new MissingRequiredPropertyException("GetDatabaseReplicaResult", "privateHost");
            }
            this.privateHost = privateHost;
            return this;
        }
        @CustomType.Setter
        public Builder privateNetworkUuid(String privateNetworkUuid) {
            if (privateNetworkUuid == null) {
              throw new MissingRequiredPropertyException("GetDatabaseReplicaResult", "privateNetworkUuid");
            }
            this.privateNetworkUuid = privateNetworkUuid;
            return this;
        }
        @CustomType.Setter
        public Builder privateUri(String privateUri) {
            if (privateUri == null) {
              throw new MissingRequiredPropertyException("GetDatabaseReplicaResult", "privateUri");
            }
            this.privateUri = privateUri;
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            if (region == null) {
              throw new MissingRequiredPropertyException("GetDatabaseReplicaResult", "region");
            }
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder storageSizeMib(String storageSizeMib) {
            if (storageSizeMib == null) {
              throw new MissingRequiredPropertyException("GetDatabaseReplicaResult", "storageSizeMib");
            }
            this.storageSizeMib = storageSizeMib;
            return this;
        }
        @CustomType.Setter
        public Builder tags(@Nullable List<String> tags) {

            this.tags = tags;
            return this;
        }
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }
        @CustomType.Setter
        public Builder uri(String uri) {
            if (uri == null) {
              throw new MissingRequiredPropertyException("GetDatabaseReplicaResult", "uri");
            }
            this.uri = uri;
            return this;
        }
        @CustomType.Setter
        public Builder user(String user) {
            if (user == null) {
              throw new MissingRequiredPropertyException("GetDatabaseReplicaResult", "user");
            }
            this.user = user;
            return this;
        }
        @CustomType.Setter
        public Builder uuid(String uuid) {
            if (uuid == null) {
              throw new MissingRequiredPropertyException("GetDatabaseReplicaResult", "uuid");
            }
            this.uuid = uuid;
            return this;
        }
        public GetDatabaseReplicaResult build() {
            final var _resultValue = new GetDatabaseReplicaResult();
            _resultValue.clusterId = clusterId;
            _resultValue.database = database;
            _resultValue.host = host;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.password = password;
            _resultValue.port = port;
            _resultValue.privateHost = privateHost;
            _resultValue.privateNetworkUuid = privateNetworkUuid;
            _resultValue.privateUri = privateUri;
            _resultValue.region = region;
            _resultValue.storageSizeMib = storageSizeMib;
            _resultValue.tags = tags;
            _resultValue.uri = uri;
            _resultValue.user = user;
            _resultValue.uuid = uuid;
            return _resultValue;
        }
    }
}
