// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean;

import com.pulumi.core.Either;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.digitalocean.enums.Region;
import com.pulumi.digitalocean.inputs.KubernetesClusterMaintenancePolicyArgs;
import com.pulumi.digitalocean.inputs.KubernetesClusterNodePoolArgs;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class KubernetesClusterArgs extends com.pulumi.resources.ResourceArgs {

    public static final KubernetesClusterArgs Empty = new KubernetesClusterArgs();

    /**
     * A boolean value indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window.
     * 
     */
    @Import(name="autoUpgrade")
    private @Nullable Output<Boolean> autoUpgrade;

    /**
     * @return A boolean value indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window.
     * 
     */
    public Optional<Output<Boolean>> autoUpgrade() {
        return Optional.ofNullable(this.autoUpgrade);
    }

    /**
     * **Use with caution.** When set to true, all associated DigitalOcean resources created via the Kubernetes API (load balancers, volumes, and volume snapshots) will be destroyed along with the cluster when it is destroyed.
     * 
     * This resource supports customized create timeouts. The default timeout is 30 minutes.
     * 
     */
    @Import(name="destroyAllAssociatedResources")
    private @Nullable Output<Boolean> destroyAllAssociatedResources;

    /**
     * @return **Use with caution.** When set to true, all associated DigitalOcean resources created via the Kubernetes API (load balancers, volumes, and volume snapshots) will be destroyed along with the cluster when it is destroyed.
     * 
     * This resource supports customized create timeouts. The default timeout is 30 minutes.
     * 
     */
    public Optional<Output<Boolean>> destroyAllAssociatedResources() {
        return Optional.ofNullable(this.destroyAllAssociatedResources);
    }

    /**
     * Enable/disable the high availability control plane for a cluster. Once enabled for a cluster, high availability cannot be disabled. Default: false
     * 
     */
    @Import(name="ha")
    private @Nullable Output<Boolean> ha;

    /**
     * @return Enable/disable the high availability control plane for a cluster. Once enabled for a cluster, high availability cannot be disabled. Default: false
     * 
     */
    public Optional<Output<Boolean>> ha() {
        return Optional.ofNullable(this.ha);
    }

    /**
     * A block representing the cluster&#39;s maintenance window. Updates will be applied within this window. If not specified, a default maintenance window will be chosen. `auto_upgrade` must be set to `true` for this to have an effect.
     * 
     */
    @Import(name="maintenancePolicy")
    private @Nullable Output<KubernetesClusterMaintenancePolicyArgs> maintenancePolicy;

    /**
     * @return A block representing the cluster&#39;s maintenance window. Updates will be applied within this window. If not specified, a default maintenance window will be chosen. `auto_upgrade` must be set to `true` for this to have an effect.
     * 
     */
    public Optional<Output<KubernetesClusterMaintenancePolicyArgs>> maintenancePolicy() {
        return Optional.ofNullable(this.maintenancePolicy);
    }

    /**
     * A name for the Kubernetes cluster.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A name for the Kubernetes cluster.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * A block representing the cluster&#39;s default node pool. Additional node pools may be added to the cluster using the `digitalocean.KubernetesNodePool` resource. The following arguments may be specified:
     * 
     */
    @Import(name="nodePool", required=true)
    private Output<KubernetesClusterNodePoolArgs> nodePool;

    /**
     * @return A block representing the cluster&#39;s default node pool. Additional node pools may be added to the cluster using the `digitalocean.KubernetesNodePool` resource. The following arguments may be specified:
     * 
     */
    public Output<KubernetesClusterNodePoolArgs> nodePool() {
        return this.nodePool;
    }

    /**
     * The slug identifier for the region where the Kubernetes cluster will be created.
     * 
     */
    @Import(name="region", required=true)
    private Output<Either<String,Region>> region;

    /**
     * @return The slug identifier for the region where the Kubernetes cluster will be created.
     * 
     */
    public Output<Either<String,Region>> region() {
        return this.region;
    }

    /**
     * Enables or disables the DigitalOcean container registry integration for the cluster. This requires that a container registry has first been created for the account. Default: false
     * 
     */
    @Import(name="registryIntegration")
    private @Nullable Output<Boolean> registryIntegration;

    /**
     * @return Enables or disables the DigitalOcean container registry integration for the cluster. This requires that a container registry has first been created for the account. Default: false
     * 
     */
    public Optional<Output<Boolean>> registryIntegration() {
        return Optional.ofNullable(this.registryIntegration);
    }

    /**
     * Enable/disable surge upgrades for a cluster. Default: false
     * 
     */
    @Import(name="surgeUpgrade")
    private @Nullable Output<Boolean> surgeUpgrade;

    /**
     * @return Enable/disable surge upgrades for a cluster. Default: false
     * 
     */
    public Optional<Output<Boolean>> surgeUpgrade() {
        return Optional.ofNullable(this.surgeUpgrade);
    }

    /**
     * A list of tag names to be applied to the Kubernetes cluster.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return A list of tag names to be applied to the Kubernetes cluster.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The slug identifier for the version of Kubernetes used for the cluster. Use [doctl](https://github.com/digitalocean/doctl) to find the available versions `doctl kubernetes options versions`. (**Note:** A cluster may only be upgraded to newer versions in-place. If the version is decreased, a new resource will be created.)
     * 
     */
    @Import(name="version", required=true)
    private Output<String> version;

    /**
     * @return The slug identifier for the version of Kubernetes used for the cluster. Use [doctl](https://github.com/digitalocean/doctl) to find the available versions `doctl kubernetes options versions`. (**Note:** A cluster may only be upgraded to newer versions in-place. If the version is decreased, a new resource will be created.)
     * 
     */
    public Output<String> version() {
        return this.version;
    }

    /**
     * The ID of the VPC where the Kubernetes cluster will be located.
     * 
     */
    @Import(name="vpcUuid")
    private @Nullable Output<String> vpcUuid;

    /**
     * @return The ID of the VPC where the Kubernetes cluster will be located.
     * 
     */
    public Optional<Output<String>> vpcUuid() {
        return Optional.ofNullable(this.vpcUuid);
    }

    private KubernetesClusterArgs() {}

    private KubernetesClusterArgs(KubernetesClusterArgs $) {
        this.autoUpgrade = $.autoUpgrade;
        this.destroyAllAssociatedResources = $.destroyAllAssociatedResources;
        this.ha = $.ha;
        this.maintenancePolicy = $.maintenancePolicy;
        this.name = $.name;
        this.nodePool = $.nodePool;
        this.region = $.region;
        this.registryIntegration = $.registryIntegration;
        this.surgeUpgrade = $.surgeUpgrade;
        this.tags = $.tags;
        this.version = $.version;
        this.vpcUuid = $.vpcUuid;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KubernetesClusterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KubernetesClusterArgs $;

        public Builder() {
            $ = new KubernetesClusterArgs();
        }

        public Builder(KubernetesClusterArgs defaults) {
            $ = new KubernetesClusterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param autoUpgrade A boolean value indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window.
         * 
         * @return builder
         * 
         */
        public Builder autoUpgrade(@Nullable Output<Boolean> autoUpgrade) {
            $.autoUpgrade = autoUpgrade;
            return this;
        }

        /**
         * @param autoUpgrade A boolean value indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window.
         * 
         * @return builder
         * 
         */
        public Builder autoUpgrade(Boolean autoUpgrade) {
            return autoUpgrade(Output.of(autoUpgrade));
        }

        /**
         * @param destroyAllAssociatedResources **Use with caution.** When set to true, all associated DigitalOcean resources created via the Kubernetes API (load balancers, volumes, and volume snapshots) will be destroyed along with the cluster when it is destroyed.
         * 
         * This resource supports customized create timeouts. The default timeout is 30 minutes.
         * 
         * @return builder
         * 
         */
        public Builder destroyAllAssociatedResources(@Nullable Output<Boolean> destroyAllAssociatedResources) {
            $.destroyAllAssociatedResources = destroyAllAssociatedResources;
            return this;
        }

        /**
         * @param destroyAllAssociatedResources **Use with caution.** When set to true, all associated DigitalOcean resources created via the Kubernetes API (load balancers, volumes, and volume snapshots) will be destroyed along with the cluster when it is destroyed.
         * 
         * This resource supports customized create timeouts. The default timeout is 30 minutes.
         * 
         * @return builder
         * 
         */
        public Builder destroyAllAssociatedResources(Boolean destroyAllAssociatedResources) {
            return destroyAllAssociatedResources(Output.of(destroyAllAssociatedResources));
        }

        /**
         * @param ha Enable/disable the high availability control plane for a cluster. Once enabled for a cluster, high availability cannot be disabled. Default: false
         * 
         * @return builder
         * 
         */
        public Builder ha(@Nullable Output<Boolean> ha) {
            $.ha = ha;
            return this;
        }

        /**
         * @param ha Enable/disable the high availability control plane for a cluster. Once enabled for a cluster, high availability cannot be disabled. Default: false
         * 
         * @return builder
         * 
         */
        public Builder ha(Boolean ha) {
            return ha(Output.of(ha));
        }

        /**
         * @param maintenancePolicy A block representing the cluster&#39;s maintenance window. Updates will be applied within this window. If not specified, a default maintenance window will be chosen. `auto_upgrade` must be set to `true` for this to have an effect.
         * 
         * @return builder
         * 
         */
        public Builder maintenancePolicy(@Nullable Output<KubernetesClusterMaintenancePolicyArgs> maintenancePolicy) {
            $.maintenancePolicy = maintenancePolicy;
            return this;
        }

        /**
         * @param maintenancePolicy A block representing the cluster&#39;s maintenance window. Updates will be applied within this window. If not specified, a default maintenance window will be chosen. `auto_upgrade` must be set to `true` for this to have an effect.
         * 
         * @return builder
         * 
         */
        public Builder maintenancePolicy(KubernetesClusterMaintenancePolicyArgs maintenancePolicy) {
            return maintenancePolicy(Output.of(maintenancePolicy));
        }

        /**
         * @param name A name for the Kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A name for the Kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param nodePool A block representing the cluster&#39;s default node pool. Additional node pools may be added to the cluster using the `digitalocean.KubernetesNodePool` resource. The following arguments may be specified:
         * 
         * @return builder
         * 
         */
        public Builder nodePool(Output<KubernetesClusterNodePoolArgs> nodePool) {
            $.nodePool = nodePool;
            return this;
        }

        /**
         * @param nodePool A block representing the cluster&#39;s default node pool. Additional node pools may be added to the cluster using the `digitalocean.KubernetesNodePool` resource. The following arguments may be specified:
         * 
         * @return builder
         * 
         */
        public Builder nodePool(KubernetesClusterNodePoolArgs nodePool) {
            return nodePool(Output.of(nodePool));
        }

        /**
         * @param region The slug identifier for the region where the Kubernetes cluster will be created.
         * 
         * @return builder
         * 
         */
        public Builder region(Output<Either<String,Region>> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The slug identifier for the region where the Kubernetes cluster will be created.
         * 
         * @return builder
         * 
         */
        public Builder region(Either<String,Region> region) {
            return region(Output.of(region));
        }

        /**
         * @param region The slug identifier for the region where the Kubernetes cluster will be created.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Either.ofLeft(region));
        }

        /**
         * @param region The slug identifier for the region where the Kubernetes cluster will be created.
         * 
         * @return builder
         * 
         */
        public Builder region(Region region) {
            return region(Either.ofRight(region));
        }

        /**
         * @param registryIntegration Enables or disables the DigitalOcean container registry integration for the cluster. This requires that a container registry has first been created for the account. Default: false
         * 
         * @return builder
         * 
         */
        public Builder registryIntegration(@Nullable Output<Boolean> registryIntegration) {
            $.registryIntegration = registryIntegration;
            return this;
        }

        /**
         * @param registryIntegration Enables or disables the DigitalOcean container registry integration for the cluster. This requires that a container registry has first been created for the account. Default: false
         * 
         * @return builder
         * 
         */
        public Builder registryIntegration(Boolean registryIntegration) {
            return registryIntegration(Output.of(registryIntegration));
        }

        /**
         * @param surgeUpgrade Enable/disable surge upgrades for a cluster. Default: false
         * 
         * @return builder
         * 
         */
        public Builder surgeUpgrade(@Nullable Output<Boolean> surgeUpgrade) {
            $.surgeUpgrade = surgeUpgrade;
            return this;
        }

        /**
         * @param surgeUpgrade Enable/disable surge upgrades for a cluster. Default: false
         * 
         * @return builder
         * 
         */
        public Builder surgeUpgrade(Boolean surgeUpgrade) {
            return surgeUpgrade(Output.of(surgeUpgrade));
        }

        /**
         * @param tags A list of tag names to be applied to the Kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A list of tag names to be applied to the Kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags A list of tag names to be applied to the Kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param version The slug identifier for the version of Kubernetes used for the cluster. Use [doctl](https://github.com/digitalocean/doctl) to find the available versions `doctl kubernetes options versions`. (**Note:** A cluster may only be upgraded to newer versions in-place. If the version is decreased, a new resource will be created.)
         * 
         * @return builder
         * 
         */
        public Builder version(Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version The slug identifier for the version of Kubernetes used for the cluster. Use [doctl](https://github.com/digitalocean/doctl) to find the available versions `doctl kubernetes options versions`. (**Note:** A cluster may only be upgraded to newer versions in-place. If the version is decreased, a new resource will be created.)
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        /**
         * @param vpcUuid The ID of the VPC where the Kubernetes cluster will be located.
         * 
         * @return builder
         * 
         */
        public Builder vpcUuid(@Nullable Output<String> vpcUuid) {
            $.vpcUuid = vpcUuid;
            return this;
        }

        /**
         * @param vpcUuid The ID of the VPC where the Kubernetes cluster will be located.
         * 
         * @return builder
         * 
         */
        public Builder vpcUuid(String vpcUuid) {
            return vpcUuid(Output.of(vpcUuid));
        }

        public KubernetesClusterArgs build() {
            if ($.nodePool == null) {
                throw new MissingRequiredPropertyException("KubernetesClusterArgs", "nodePool");
            }
            if ($.region == null) {
                throw new MissingRequiredPropertyException("KubernetesClusterArgs", "region");
            }
            if ($.version == null) {
                throw new MissingRequiredPropertyException("KubernetesClusterArgs", "version");
            }
            return $;
        }
    }

}
