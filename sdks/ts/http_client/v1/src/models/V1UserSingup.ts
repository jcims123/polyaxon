// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.1.8-rc0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface V1UserSingup
 */
export interface V1UserSingup {
    /**
     * 
     * @type {string}
     * @memberof V1UserSingup
     */
    username?: string;
    /**
     * 
     * @type {string}
     * @memberof V1UserSingup
     */
    email?: string;
    /**
     * 
     * @type {string}
     * @memberof V1UserSingup
     */
    organization?: string;
    /**
     * 
     * @type {string}
     * @memberof V1UserSingup
     */
    password?: string;
}

export function V1UserSingupFromJSON(json: any): V1UserSingup {
    return V1UserSingupFromJSONTyped(json, false);
}

export function V1UserSingupFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1UserSingup {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'username': !exists(json, 'username') ? undefined : json['username'],
        'email': !exists(json, 'email') ? undefined : json['email'],
        'organization': !exists(json, 'organization') ? undefined : json['organization'],
        'password': !exists(json, 'password') ? undefined : json['password'],
    };
}

export function V1UserSingupToJSON(value?: V1UserSingup | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'username': value.username,
        'email': value.email,
        'organization': value.organization,
        'password': value.password,
    };
}


