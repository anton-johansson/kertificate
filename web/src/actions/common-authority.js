/**
 * Copyright 2019 Anton Johansson
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import Api from '../api';
import { pki } from 'node-forge';

export const LOAD_COMMON_AUTHORITIES = 'LOAD_COMMON_AUTHORITIES';
export const LOAD_COMMON_AUTHORITIES_SUCCESS = 'LOAD_COMMON_AUTHORITIES_SUCCESS';
export const LOAD_COMMON_AUTHORITIES_FAILED = 'LOAD_COMMON_AUTHORITIES_FAILED';
export const loadCommonAuthorities = force => {
    return async (dispatch, getState) => {
        const doLoad = force || getState().commonAuthority.loadState === 'not-loaded';
        if (!doLoad) {
            return;
        }

        dispatch({type: LOAD_COMMON_AUTHORITIES});
        const api = new Api(dispatch, getState);
        const response = await api.get('/v1/common-authorities');

        if (response.status === 200) {
            dispatch({
                type: LOAD_COMMON_AUTHORITIES_SUCCESS,
                payload: {
                    commonAuthorities: response.body,
                },
            });
            return;
        } else {
            console.log('Unknown status when loading common authorities:', response.status);
            dispatch({type: LOAD_COMMON_AUTHORITIES_FAILED});
        }
    };
}

export const SET_SEARCH_STRING = 'SET_SEARCH_STRING';
export const setSearchString = searchString => {
    return async dispatch => {
        dispatch({
            type: SET_SEARCH_STRING,
            payload: {
                searchString,
            },
        });
    };
}

export const LOAD_COMMON_AUTHORITY_SUCCESS = 'LOAD_COMMON_AUTHORITY_SUCCESS';
export const loadCommonAuthority = commonAuthorityId => {
    return async (dispatch, getState) => {
        const commonAuthority = getState().commonAuthority.commonAuthority;
        if (commonAuthority.commonAuthorityId === commonAuthorityId) {
            return;
        }

        const responseData = {
            commonAuthorityId,
            name: 'Anton CA',
            comments: 'This is an amazing common authority.',
            certificateData: '-----BEGIN CERTIFICATE-----\n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\nWAiMFhit/vE=\n-----END CERTIFICATE-----\n',
        };

        const certificateData = pki.certificateFromPem(responseData.certificateData);
        const subject = certificateData.subject.attributes.reduce((current, item) => {
            if (!item.name) {
                return current;
            }

            return ({...current, [item.name]: item.value});
        }, {});
        certificateData.subject = subject;

        dispatch({
            type: LOAD_COMMON_AUTHORITY_SUCCESS,
            payload: {
                ...responseData,
                certificateData,
                pem: responseData.certificateData,
            },
        });
    };
}
