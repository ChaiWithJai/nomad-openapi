/*
 * Nomad
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.4
 * Contact: support@hashicorp.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package io.nomadproject.client.models;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * RestartPolicy
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class RestartPolicy {
  public static final String SERIALIZED_NAME_ATTEMPTS = "Attempts";
  @SerializedName(SERIALIZED_NAME_ATTEMPTS)
  private Integer attempts;

  public static final String SERIALIZED_NAME_DELAY = "Delay";
  @SerializedName(SERIALIZED_NAME_DELAY)
  private Long delay;

  public static final String SERIALIZED_NAME_INTERVAL = "Interval";
  @SerializedName(SERIALIZED_NAME_INTERVAL)
  private Long interval;

  public static final String SERIALIZED_NAME_MODE = "Mode";
  @SerializedName(SERIALIZED_NAME_MODE)
  private String mode;


  public RestartPolicy attempts(Integer attempts) {
    
    this.attempts = attempts;
    return this;
  }

   /**
   * Get attempts
   * @return attempts
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getAttempts() {
    return attempts;
  }


  public void setAttempts(Integer attempts) {
    this.attempts = attempts;
  }


  public RestartPolicy delay(Long delay) {
    
    this.delay = delay;
    return this;
  }

   /**
   * Get delay
   * @return delay
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getDelay() {
    return delay;
  }


  public void setDelay(Long delay) {
    this.delay = delay;
  }


  public RestartPolicy interval(Long interval) {
    
    this.interval = interval;
    return this;
  }

   /**
   * Get interval
   * @return interval
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getInterval() {
    return interval;
  }


  public void setInterval(Long interval) {
    this.interval = interval;
  }


  public RestartPolicy mode(String mode) {
    
    this.mode = mode;
    return this;
  }

   /**
   * Get mode
   * @return mode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMode() {
    return mode;
  }


  public void setMode(String mode) {
    this.mode = mode;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    RestartPolicy restartPolicy = (RestartPolicy) o;
    return Objects.equals(this.attempts, restartPolicy.attempts) &&
        Objects.equals(this.delay, restartPolicy.delay) &&
        Objects.equals(this.interval, restartPolicy.interval) &&
        Objects.equals(this.mode, restartPolicy.mode);
  }

  @Override
  public int hashCode() {
    return Objects.hash(attempts, delay, interval, mode);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class RestartPolicy {\n");
    sb.append("    attempts: ").append(toIndentedString(attempts)).append("\n");
    sb.append("    delay: ").append(toIndentedString(delay)).append("\n");
    sb.append("    interval: ").append(toIndentedString(interval)).append("\n");
    sb.append("    mode: ").append(toIndentedString(mode)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}
