/*
 * Copyright (c) 2019  InterDigital Communications, Inc
 *
 * Licensed under the Apache License, Version 2.0 (the \"License\");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an \"AS IS\" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * AdvantEDGE Platform Controller REST API
 * This API is the main platform API and mainly used by the AdvantEDGE frontend to interact with scenarios <p>**Micro-service**<br>[meep-ctrl-engine](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-ctrl-engine) <p>**Type & Usage**<br>Platform main interface used by controller software that want to interact with the AdvantEDGE platform <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address:30000/api_ <p>**Default Port**<br>`30000` 
 *
 * OpenAPI spec version: 1.0.0
 * Contact: AdvantEDGE@InterDigital.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.9
 *
 * Do not edit the class manually.
 *
 */

(function(factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/ActivationInfo', 'model/Deployment', 'model/Domain', 'model/EgressService', 'model/Event', 'model/EventMobility', 'model/EventNetworkCharacteristicsUpdate', 'model/EventOther', 'model/EventPoasInRange', 'model/ExternalConfig', 'model/GpuConfig', 'model/IngressService', 'model/NetworkCharacteristics', 'model/NetworkLocation', 'model/NodeServiceMaps', 'model/PhysicalLocation', 'model/PodStatus', 'model/PodsStatus', 'model/Process', 'model/Replay', 'model/ReplayEvent', 'model/ReplayFileList', 'model/ReplayInfo', 'model/ReplayStatus', 'model/Scenario', 'model/ScenarioConfig', 'model/ScenarioList', 'model/ServiceConfig', 'model/ServicePort', 'model/Zone', 'api/EventReplayApi', 'api/PodStatesApi', 'api/ScenarioConfigurationApi', 'api/ScenarioExecutionApi'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('./ApiClient'), require('./model/ActivationInfo'), require('./model/Deployment'), require('./model/Domain'), require('./model/EgressService'), require('./model/Event'), require('./model/EventMobility'), require('./model/EventNetworkCharacteristicsUpdate'), require('./model/EventOther'), require('./model/EventPoasInRange'), require('./model/ExternalConfig'), require('./model/GpuConfig'), require('./model/IngressService'), require('./model/NetworkCharacteristics'), require('./model/NetworkLocation'), require('./model/NodeServiceMaps'), require('./model/PhysicalLocation'), require('./model/PodStatus'), require('./model/PodsStatus'), require('./model/Process'), require('./model/Replay'), require('./model/ReplayEvent'), require('./model/ReplayFileList'), require('./model/ReplayInfo'), require('./model/ReplayStatus'), require('./model/Scenario'), require('./model/ScenarioConfig'), require('./model/ScenarioList'), require('./model/ServiceConfig'), require('./model/ServicePort'), require('./model/Zone'), require('./api/EventReplayApi'), require('./api/PodStatesApi'), require('./api/ScenarioConfigurationApi'), require('./api/ScenarioExecutionApi'));
  }
}(function(ApiClient, ActivationInfo, Deployment, Domain, EgressService, Event, EventMobility, EventNetworkCharacteristicsUpdate, EventOther, EventPoasInRange, ExternalConfig, GpuConfig, IngressService, NetworkCharacteristics, NetworkLocation, NodeServiceMaps, PhysicalLocation, PodStatus, PodsStatus, Process, Replay, ReplayEvent, ReplayFileList, ReplayInfo, ReplayStatus, Scenario, ScenarioConfig, ScenarioList, ServiceConfig, ServicePort, Zone, EventReplayApi, PodStatesApi, ScenarioConfigurationApi, ScenarioExecutionApi) {
  'use strict';

  /**
   * This_API_is_the_main_platform_API_and_mainly_used_by_the_AdvantEDGE_frontend_to_interact_with_scenarios_pMicro_servicebr_meep_ctrl_engine_httpsgithub_comInterDigitalIncAdvantEDGEtreemastergo_appsmeep_ctrl_engine_pType__UsagebrPlatform_main_interface_used_by_controller_software_that_want_to_interact_with_the_AdvantEDGE_platform_pDetailsbrAPI_details_available_at__your_AdvantEDGE_ip_address30000api__pDefault_Portbr30000_.<br>
   * The <code>index</code> module provides access to constructors for all the classes which comprise the public API.
   * <p>
   * An AMD (recommended!) or CommonJS application will generally do something equivalent to the following:
   * <pre>
   * var AdvantEdgePlatformControllerRestApi = require('index'); // See note below*.
   * var xxxSvc = new AdvantEdgePlatformControllerRestApi.XxxApi(); // Allocate the API class we're going to use.
   * var yyyModel = new AdvantEdgePlatformControllerRestApi.Yyy(); // Construct a model instance.
   * yyyModel.someProperty = 'someValue';
   * ...
   * var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
   * ...
   * </pre>
   * <em>*NOTE: For a top-level AMD script, use require(['index'], function(){...})
   * and put the application logic within the callback function.</em>
   * </p>
   * <p>
   * A non-AMD browser application (discouraged) might do something like this:
   * <pre>
   * var xxxSvc = new AdvantEdgePlatformControllerRestApi.XxxApi(); // Allocate the API class we're going to use.
   * var yyy = new AdvantEdgePlatformControllerRestApi.Yyy(); // Construct a model instance.
   * yyyModel.someProperty = 'someValue';
   * ...
   * var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
   * ...
   * </pre>
   * </p>
   * @module index
   * @version 1.0.0
   */
  var exports = {
    /**
     * The ApiClient constructor.
     * @property {module:ApiClient}
     */
    ApiClient: ApiClient,
    /**
     * The ActivationInfo model constructor.
     * @property {module:model/ActivationInfo}
     */
    ActivationInfo: ActivationInfo,
    /**
     * The Deployment model constructor.
     * @property {module:model/Deployment}
     */
    Deployment: Deployment,
    /**
     * The Domain model constructor.
     * @property {module:model/Domain}
     */
    Domain: Domain,
    /**
     * The EgressService model constructor.
     * @property {module:model/EgressService}
     */
    EgressService: EgressService,
    /**
     * The Event model constructor.
     * @property {module:model/Event}
     */
    Event: Event,
    /**
     * The EventMobility model constructor.
     * @property {module:model/EventMobility}
     */
    EventMobility: EventMobility,
    /**
     * The EventNetworkCharacteristicsUpdate model constructor.
     * @property {module:model/EventNetworkCharacteristicsUpdate}
     */
    EventNetworkCharacteristicsUpdate: EventNetworkCharacteristicsUpdate,
    /**
     * The EventOther model constructor.
     * @property {module:model/EventOther}
     */
    EventOther: EventOther,
    /**
     * The EventPoasInRange model constructor.
     * @property {module:model/EventPoasInRange}
     */
    EventPoasInRange: EventPoasInRange,
    /**
     * The ExternalConfig model constructor.
     * @property {module:model/ExternalConfig}
     */
    ExternalConfig: ExternalConfig,
    /**
     * The GpuConfig model constructor.
     * @property {module:model/GpuConfig}
     */
    GpuConfig: GpuConfig,
    /**
     * The IngressService model constructor.
     * @property {module:model/IngressService}
     */
    IngressService: IngressService,
    /**
     * The NetworkCharacteristics model constructor.
     * @property {module:model/NetworkCharacteristics}
     */
    NetworkCharacteristics: NetworkCharacteristics,
    /**
     * The NetworkLocation model constructor.
     * @property {module:model/NetworkLocation}
     */
    NetworkLocation: NetworkLocation,
    /**
     * The NodeServiceMaps model constructor.
     * @property {module:model/NodeServiceMaps}
     */
    NodeServiceMaps: NodeServiceMaps,
    /**
     * The PhysicalLocation model constructor.
     * @property {module:model/PhysicalLocation}
     */
    PhysicalLocation: PhysicalLocation,
    /**
     * The PodStatus model constructor.
     * @property {module:model/PodStatus}
     */
    PodStatus: PodStatus,
    /**
     * The PodsStatus model constructor.
     * @property {module:model/PodsStatus}
     */
    PodsStatus: PodsStatus,
    /**
     * The Process model constructor.
     * @property {module:model/Process}
     */
    Process: Process,
    /**
     * The Replay model constructor.
     * @property {module:model/Replay}
     */
    Replay: Replay,
    /**
     * The ReplayEvent model constructor.
     * @property {module:model/ReplayEvent}
     */
    ReplayEvent: ReplayEvent,
    /**
     * The ReplayFileList model constructor.
     * @property {module:model/ReplayFileList}
     */
    ReplayFileList: ReplayFileList,
    /**
     * The ReplayInfo model constructor.
     * @property {module:model/ReplayInfo}
     */
    ReplayInfo: ReplayInfo,
    /**
     * The ReplayStatus model constructor.
     * @property {module:model/ReplayStatus}
     */
    ReplayStatus: ReplayStatus,
    /**
     * The Scenario model constructor.
     * @property {module:model/Scenario}
     */
    Scenario: Scenario,
    /**
     * The ScenarioConfig model constructor.
     * @property {module:model/ScenarioConfig}
     */
    ScenarioConfig: ScenarioConfig,
    /**
     * The ScenarioList model constructor.
     * @property {module:model/ScenarioList}
     */
    ScenarioList: ScenarioList,
    /**
     * The ServiceConfig model constructor.
     * @property {module:model/ServiceConfig}
     */
    ServiceConfig: ServiceConfig,
    /**
     * The ServicePort model constructor.
     * @property {module:model/ServicePort}
     */
    ServicePort: ServicePort,
    /**
     * The Zone model constructor.
     * @property {module:model/Zone}
     */
    Zone: Zone,
    /**
     * The EventReplayApi service constructor.
     * @property {module:api/EventReplayApi}
     */
    EventReplayApi: EventReplayApi,
    /**
     * The PodStatesApi service constructor.
     * @property {module:api/PodStatesApi}
     */
    PodStatesApi: PodStatesApi,
    /**
     * The ScenarioConfigurationApi service constructor.
     * @property {module:api/ScenarioConfigurationApi}
     */
    ScenarioConfigurationApi: ScenarioConfigurationApi,
    /**
     * The ScenarioExecutionApi service constructor.
     * @property {module:api/ScenarioExecutionApi}
     */
    ScenarioExecutionApi: ScenarioExecutionApi
  };

  return exports;
}));
