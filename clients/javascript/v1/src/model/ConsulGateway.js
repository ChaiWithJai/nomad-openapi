/**
 * Nomad
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.4
 * Contact: support@hashicorp.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import ConsulGatewayProxy from './ConsulGatewayProxy';
import ConsulIngressConfigEntry from './ConsulIngressConfigEntry';
import ConsulTerminatingConfigEntry from './ConsulTerminatingConfigEntry';

/**
 * The ConsulGateway model module.
 * @module model/ConsulGateway
 * @version 1.1.4
 */
class ConsulGateway {
    /**
     * Constructs a new <code>ConsulGateway</code>.
     * @alias module:model/ConsulGateway
     */
    constructor() { 
        
        ConsulGateway.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ConsulGateway</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ConsulGateway} obj Optional instance to populate.
     * @return {module:model/ConsulGateway} The populated <code>ConsulGateway</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ConsulGateway();

            if (data.hasOwnProperty('Ingress')) {
                obj['Ingress'] = ConsulIngressConfigEntry.constructFromObject(data['Ingress']);
            }
            if (data.hasOwnProperty('Mesh')) {
                obj['Mesh'] = ApiClient.convertToType(data['Mesh'], Object);
            }
            if (data.hasOwnProperty('Proxy')) {
                obj['Proxy'] = ConsulGatewayProxy.constructFromObject(data['Proxy']);
            }
            if (data.hasOwnProperty('Terminating')) {
                obj['Terminating'] = ConsulTerminatingConfigEntry.constructFromObject(data['Terminating']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/ConsulIngressConfigEntry} Ingress
 */
ConsulGateway.prototype['Ingress'] = undefined;

/**
 * @member {Object} Mesh
 */
ConsulGateway.prototype['Mesh'] = undefined;

/**
 * @member {module:model/ConsulGatewayProxy} Proxy
 */
ConsulGateway.prototype['Proxy'] = undefined;

/**
 * @member {module:model/ConsulTerminatingConfigEntry} Terminating
 */
ConsulGateway.prototype['Terminating'] = undefined;






export default ConsulGateway;
