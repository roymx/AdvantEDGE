/*
 * MEEP Controller REST API
 * Copyright (c) 2019  InterDigital Communications, Inc Licensed under the Apache License, Version 2.0 (the \"License\"); you may not use this file except in compliance with the License. You may obtain a copy of the License at      http://www.apache.org/licenses/LICENSE-2.0  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an \"AS IS\" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License. 
 *
 * OpenAPI spec version: 1.0.0
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.9
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/ExternalConfig', 'model/GpuConfig', 'model/ServiceConfig'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./ExternalConfig'), require('./GpuConfig'), require('./ServiceConfig'));
  } else {
    // Browser globals (root is window)
    if (!root.MeepControllerRestApi) {
      root.MeepControllerRestApi = {};
    }
    root.MeepControllerRestApi.Process = factory(root.MeepControllerRestApi.ApiClient, root.MeepControllerRestApi.ExternalConfig, root.MeepControllerRestApi.GpuConfig, root.MeepControllerRestApi.ServiceConfig);
  }
}(this, function(ApiClient, ExternalConfig, GpuConfig, ServiceConfig) {
  'use strict';

  /**
   * The Process model module.
   * @module model/Process
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>Process</code>.
   * Application or service object
   * @alias module:model/Process
   * @class
   */
  var exports = function() {
  };

  /**
   * Constructs a <code>Process</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/Process} obj Optional instance to populate.
   * @return {module:model/Process} The populated <code>Process</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('id'))
        obj.id = ApiClient.convertToType(data['id'], 'String');
      if (data.hasOwnProperty('name'))
        obj.name = ApiClient.convertToType(data['name'], 'String');
      if (data.hasOwnProperty('type'))
        obj.type = ApiClient.convertToType(data['type'], 'String');
      if (data.hasOwnProperty('isExternal'))
        obj.isExternal = ApiClient.convertToType(data['isExternal'], 'Boolean');
      if (data.hasOwnProperty('image'))
        obj.image = ApiClient.convertToType(data['image'], 'String');
      if (data.hasOwnProperty('environment'))
        obj.environment = ApiClient.convertToType(data['environment'], 'String');
      if (data.hasOwnProperty('commandArguments'))
        obj.commandArguments = ApiClient.convertToType(data['commandArguments'], 'String');
      if (data.hasOwnProperty('commandExe'))
        obj.commandExe = ApiClient.convertToType(data['commandExe'], 'String');
      if (data.hasOwnProperty('serviceConfig'))
        obj.serviceConfig = ServiceConfig.constructFromObject(data['serviceConfig']);
      if (data.hasOwnProperty('gpuConfig'))
        obj.gpuConfig = GpuConfig.constructFromObject(data['gpuConfig']);
      if (data.hasOwnProperty('externalConfig'))
        obj.externalConfig = ExternalConfig.constructFromObject(data['externalConfig']);
      if (data.hasOwnProperty('status'))
        obj.status = ApiClient.convertToType(data['status'], 'String');
      if (data.hasOwnProperty('userChartLocation'))
        obj.userChartLocation = ApiClient.convertToType(data['userChartLocation'], 'String');
      if (data.hasOwnProperty('userChartAlternateValues'))
        obj.userChartAlternateValues = ApiClient.convertToType(data['userChartAlternateValues'], 'String');
      if (data.hasOwnProperty('userChartGroup'))
        obj.userChartGroup = ApiClient.convertToType(data['userChartGroup'], 'String');
      if (data.hasOwnProperty('meta'))
        obj.meta = ApiClient.convertToType(data['meta'], {'String': 'String'});
      if (data.hasOwnProperty('userMeta'))
        obj.userMeta = ApiClient.convertToType(data['userMeta'], {'String': 'String'});
      if (data.hasOwnProperty('appLatency'))
        obj.appLatency = ApiClient.convertToType(data['appLatency'], 'Number');
      if (data.hasOwnProperty('appLatencyVariation'))
        obj.appLatencyVariation = ApiClient.convertToType(data['appLatencyVariation'], 'Number');
      if (data.hasOwnProperty('appThroughput'))
        obj.appThroughput = ApiClient.convertToType(data['appThroughput'], 'Number');
      if (data.hasOwnProperty('appPacketLoss'))
        obj.appPacketLoss = ApiClient.convertToType(data['appPacketLoss'], 'Number');
      if (data.hasOwnProperty('placementId'))
        obj.placementId = ApiClient.convertToType(data['placementId'], 'String');
    }
    return obj;
  }

  /**
   * Unique process ID
   * @member {String} id
   */
  exports.prototype.id = undefined;

  /**
   * Process name
   * @member {String} name
   */
  exports.prototype.name = undefined;

  /**
   * Process type
   * @member {module:model/Process.TypeEnum} type
   */
  exports.prototype.type = undefined;

  /**
   * true: process is external to MEEP false: process is internal to MEEP
   * @member {Boolean} isExternal
   */
  exports.prototype.isExternal = undefined;

  /**
   * Docker image to deploy inside MEEP
   * @member {String} image
   */
  exports.prototype.image = undefined;

  /**
   * Environment variables using the format NAME=\"value\",NAME=\"value\",NAME=\"value\"
   * @member {String} environment
   */
  exports.prototype.environment = undefined;

  /**
   * Arguments to command executable
   * @member {String} commandArguments
   */
  exports.prototype.commandArguments = undefined;

  /**
   * Executable to invoke at container start up
   * @member {String} commandExe
   */
  exports.prototype.commandExe = undefined;

  /**
   * @member {module:model/ServiceConfig} serviceConfig
   */
  exports.prototype.serviceConfig = undefined;

  /**
   * @member {module:model/GpuConfig} gpuConfig
   */
  exports.prototype.gpuConfig = undefined;

  /**
   * @member {module:model/ExternalConfig} externalConfig
   */
  exports.prototype.externalConfig = undefined;

  /**
   * Process status
   * @member {String} status
   */
  exports.prototype.status = undefined;

  /**
   * Chart location for the deployment of the chart provided by the user
   * @member {String} userChartLocation
   */
  exports.prototype.userChartLocation = undefined;

  /**
   * Chart values.yaml file location for the deployment of the chart provided by the user
   * @member {String} userChartAlternateValues
   */
  exports.prototype.userChartAlternateValues = undefined;

  /**
   * Chart supplemental information related to the group (service)
   * @member {String} userChartGroup
   */
  exports.prototype.userChartGroup = undefined;

  /**
   * Key/Value Pair Map (string, string)
   * @member {Object.<String, String>} meta
   */
  exports.prototype.meta = undefined;

  /**
   * Key/Value Pair Map (string, string)
   * @member {Object.<String, String>} userMeta
   */
  exports.prototype.userMeta = undefined;

  /**
   * Latency in ms caused by the application
   * @member {Number} appLatency
   */
  exports.prototype.appLatency = undefined;

  /**
   * Latency variation in ms caused by the application
   * @member {Number} appLatencyVariation
   */
  exports.prototype.appLatencyVariation = undefined;

  /**
   * The limit of the traffic supported by the application
   * @member {Number} appThroughput
   */
  exports.prototype.appThroughput = undefined;

  /**
   * Packet lost (in terms of percentage) caused by the application
   * @member {Number} appPacketLoss
   */
  exports.prototype.appPacketLoss = undefined;

  /**
   * Identifier used for process placement in AdvantEDGE cluster
   * @member {String} placementId
   */
  exports.prototype.placementId = undefined;


  /**
   * Allowed values for the <code>type</code> property.
   * @enum {String}
   * @readonly
   */
  exports.TypeEnum = {
    /**
     * value: "UE-APP"
     * @const
     */
    UE_APP: "UE-APP",

    /**
     * value: "EDGE-APP"
     * @const
     */
    EDGE_APP: "EDGE-APP",

    /**
     * value: "MEC-SVC"
     * @const
     */
    MEC_SVC: "MEC-SVC",

    /**
     * value: "CLOUD-APP"
     * @const
     */
    CLOUD_APP: "CLOUD-APP"
  };

  return exports;

}));
