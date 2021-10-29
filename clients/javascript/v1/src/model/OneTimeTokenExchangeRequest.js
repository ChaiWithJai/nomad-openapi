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

/**
 * The OneTimeTokenExchangeRequest model module.
 * @module model/OneTimeTokenExchangeRequest
 * @version 1.1.4
 */
class OneTimeTokenExchangeRequest {
    /**
     * Constructs a new <code>OneTimeTokenExchangeRequest</code>.
     * @alias module:model/OneTimeTokenExchangeRequest
     */
    constructor() { 
        
        OneTimeTokenExchangeRequest.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>OneTimeTokenExchangeRequest</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/OneTimeTokenExchangeRequest} obj Optional instance to populate.
     * @return {module:model/OneTimeTokenExchangeRequest} The populated <code>OneTimeTokenExchangeRequest</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new OneTimeTokenExchangeRequest();

            if (data.hasOwnProperty('OneTimeSecretID')) {
                obj['OneTimeSecretID'] = ApiClient.convertToType(data['OneTimeSecretID'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} OneTimeSecretID
 */
OneTimeTokenExchangeRequest.prototype['OneTimeSecretID'] = undefined;






export default OneTimeTokenExchangeRequest;
