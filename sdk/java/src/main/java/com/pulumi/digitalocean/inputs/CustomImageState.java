// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.digitalocean.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CustomImageState extends com.pulumi.resources.ResourceArgs {

    public static final CustomImageState Empty = new CustomImageState();

    /**
     * A time value given in ISO8601 combined date and time format that represents when the image was created.
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return A time value given in ISO8601 combined date and time format that represents when the image was created.
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    /**
     * An optional description for the image.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return An optional description for the image.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * An optional distribution name for the image. Valid values are documented [here](https://docs.digitalocean.com/reference/api/api-reference/#operation/create_custom_image)
     * 
     */
    @Import(name="distribution")
    private @Nullable Output<String> distribution;

    /**
     * @return An optional distribution name for the image. Valid values are documented [here](https://docs.digitalocean.com/reference/api/api-reference/#operation/create_custom_image)
     * 
     */
    public Optional<Output<String>> distribution() {
        return Optional.ofNullable(this.distribution);
    }

    /**
     * A unique number that can be used to identify and reference a specific image.
     * 
     */
    @Import(name="imageId")
    private @Nullable Output<Integer> imageId;

    /**
     * @return A unique number that can be used to identify and reference a specific image.
     * 
     */
    public Optional<Output<Integer>> imageId() {
        return Optional.ofNullable(this.imageId);
    }

    /**
     * The minimum disk size in GB required for a Droplet to use this image.
     * 
     */
    @Import(name="minDiskSize")
    private @Nullable Output<Integer> minDiskSize;

    /**
     * @return The minimum disk size in GB required for a Droplet to use this image.
     * 
     */
    public Optional<Output<Integer>> minDiskSize() {
        return Optional.ofNullable(this.minDiskSize);
    }

    /**
     * A name for the Custom Image.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A name for the Custom Image.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Indicates whether the image in question is public or not.
     * 
     */
    @Import(name="public")
    private @Nullable Output<Boolean> public_;

    /**
     * @return Indicates whether the image in question is public or not.
     * 
     */
    public Optional<Output<Boolean>> public_() {
        return Optional.ofNullable(this.public_);
    }

    /**
     * A list of regions. (Currently only one is supported).
     * 
     */
    @Import(name="regions")
    private @Nullable Output<List<String>> regions;

    /**
     * @return A list of regions. (Currently only one is supported).
     * 
     */
    public Optional<Output<List<String>>> regions() {
        return Optional.ofNullable(this.regions);
    }

    /**
     * The size of the image in gigabytes.
     * 
     */
    @Import(name="sizeGigabytes")
    private @Nullable Output<Double> sizeGigabytes;

    /**
     * @return The size of the image in gigabytes.
     * 
     */
    public Optional<Output<Double>> sizeGigabytes() {
        return Optional.ofNullable(this.sizeGigabytes);
    }

    /**
     * A uniquely identifying string for each image.
     * 
     */
    @Import(name="slug")
    private @Nullable Output<String> slug;

    /**
     * @return A uniquely identifying string for each image.
     * 
     */
    public Optional<Output<String>> slug() {
        return Optional.ofNullable(this.slug);
    }

    /**
     * A status string indicating the state of a custom image.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return A status string indicating the state of a custom image.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * A list of optional tags for the image.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return A list of optional tags for the image.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Describes the kind of image.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Describes the kind of image.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * A URL from which the custom Linux virtual machine image may be retrieved.
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return A URL from which the custom Linux virtual machine image may be retrieved.
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    private CustomImageState() {}

    private CustomImageState(CustomImageState $) {
        this.createdAt = $.createdAt;
        this.description = $.description;
        this.distribution = $.distribution;
        this.imageId = $.imageId;
        this.minDiskSize = $.minDiskSize;
        this.name = $.name;
        this.public_ = $.public_;
        this.regions = $.regions;
        this.sizeGigabytes = $.sizeGigabytes;
        this.slug = $.slug;
        this.status = $.status;
        this.tags = $.tags;
        this.type = $.type;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CustomImageState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CustomImageState $;

        public Builder() {
            $ = new CustomImageState();
        }

        public Builder(CustomImageState defaults) {
            $ = new CustomImageState(Objects.requireNonNull(defaults));
        }

        /**
         * @param createdAt A time value given in ISO8601 combined date and time format that represents when the image was created.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt A time value given in ISO8601 combined date and time format that represents when the image was created.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        /**
         * @param description An optional description for the image.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description An optional description for the image.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param distribution An optional distribution name for the image. Valid values are documented [here](https://docs.digitalocean.com/reference/api/api-reference/#operation/create_custom_image)
         * 
         * @return builder
         * 
         */
        public Builder distribution(@Nullable Output<String> distribution) {
            $.distribution = distribution;
            return this;
        }

        /**
         * @param distribution An optional distribution name for the image. Valid values are documented [here](https://docs.digitalocean.com/reference/api/api-reference/#operation/create_custom_image)
         * 
         * @return builder
         * 
         */
        public Builder distribution(String distribution) {
            return distribution(Output.of(distribution));
        }

        /**
         * @param imageId A unique number that can be used to identify and reference a specific image.
         * 
         * @return builder
         * 
         */
        public Builder imageId(@Nullable Output<Integer> imageId) {
            $.imageId = imageId;
            return this;
        }

        /**
         * @param imageId A unique number that can be used to identify and reference a specific image.
         * 
         * @return builder
         * 
         */
        public Builder imageId(Integer imageId) {
            return imageId(Output.of(imageId));
        }

        /**
         * @param minDiskSize The minimum disk size in GB required for a Droplet to use this image.
         * 
         * @return builder
         * 
         */
        public Builder minDiskSize(@Nullable Output<Integer> minDiskSize) {
            $.minDiskSize = minDiskSize;
            return this;
        }

        /**
         * @param minDiskSize The minimum disk size in GB required for a Droplet to use this image.
         * 
         * @return builder
         * 
         */
        public Builder minDiskSize(Integer minDiskSize) {
            return minDiskSize(Output.of(minDiskSize));
        }

        /**
         * @param name A name for the Custom Image.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A name for the Custom Image.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param public_ Indicates whether the image in question is public or not.
         * 
         * @return builder
         * 
         */
        public Builder public_(@Nullable Output<Boolean> public_) {
            $.public_ = public_;
            return this;
        }

        /**
         * @param public_ Indicates whether the image in question is public or not.
         * 
         * @return builder
         * 
         */
        public Builder public_(Boolean public_) {
            return public_(Output.of(public_));
        }

        /**
         * @param regions A list of regions. (Currently only one is supported).
         * 
         * @return builder
         * 
         */
        public Builder regions(@Nullable Output<List<String>> regions) {
            $.regions = regions;
            return this;
        }

        /**
         * @param regions A list of regions. (Currently only one is supported).
         * 
         * @return builder
         * 
         */
        public Builder regions(List<String> regions) {
            return regions(Output.of(regions));
        }

        /**
         * @param regions A list of regions. (Currently only one is supported).
         * 
         * @return builder
         * 
         */
        public Builder regions(String... regions) {
            return regions(List.of(regions));
        }

        /**
         * @param sizeGigabytes The size of the image in gigabytes.
         * 
         * @return builder
         * 
         */
        public Builder sizeGigabytes(@Nullable Output<Double> sizeGigabytes) {
            $.sizeGigabytes = sizeGigabytes;
            return this;
        }

        /**
         * @param sizeGigabytes The size of the image in gigabytes.
         * 
         * @return builder
         * 
         */
        public Builder sizeGigabytes(Double sizeGigabytes) {
            return sizeGigabytes(Output.of(sizeGigabytes));
        }

        /**
         * @param slug A uniquely identifying string for each image.
         * 
         * @return builder
         * 
         */
        public Builder slug(@Nullable Output<String> slug) {
            $.slug = slug;
            return this;
        }

        /**
         * @param slug A uniquely identifying string for each image.
         * 
         * @return builder
         * 
         */
        public Builder slug(String slug) {
            return slug(Output.of(slug));
        }

        /**
         * @param status A status string indicating the state of a custom image.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status A status string indicating the state of a custom image.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags A list of optional tags for the image.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A list of optional tags for the image.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags A list of optional tags for the image.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param type Describes the kind of image.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Describes the kind of image.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param url A URL from which the custom Linux virtual machine image may be retrieved.
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url A URL from which the custom Linux virtual machine image may be retrieved.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public CustomImageState build() {
            return $;
        }
    }

}
