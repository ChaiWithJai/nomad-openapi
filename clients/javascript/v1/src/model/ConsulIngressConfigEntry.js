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
import ConsulGatewayTLSConfig from './ConsulGatewayTLSConfig';
import ConsulIngressListener from './ConsulIngressListener';

/**
 * The ConsulIngressConfigEntry model module.
 * @module model/ConsulIngressConfigEntry
 * @version 1.1.4
 */
class ConsulIngressConfigEntry {
    /**
     * Constructs a new <code>ConsulIngressConfigEntry</code>.
     * @alias module:model/ConsulIngressConfigEntry
     */
    constructor() { 
        
        ConsulIngressConfigEntry.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ConsulIngressConfigEntry</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ConsulIngressConfigEntry} obj Optional instance to populate.
     * @return {module:model/ConsulIngressConfigEntry} The populated <code>ConsulIngressConfigEntry</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ConsulIngressConfigEntry();

            if (data.hasOwnProperty('Listeners')) {
                obj['Listeners'] = ApiClient.convertToType(data['Listeners'], [ConsulIngressListener]);
            }
            if (data.hasOwnProperty('TLS')) {
                obj['TLS'] = ConsulGatewayTLSConfig.constructFromObject(data['TLS']);
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<module:model/ConsulIngressListener>} Listeners
 */
ConsulIngressConfigEntry.prototype['Listeners'] = undefined;

/**
 * @member {module:model/ConsulGatewayTLSConfig} TLS
 */
ConsulIngressConfigEntry.prototype['TLS'] = undefined;






export default ConsulIngressConfigEntry;
