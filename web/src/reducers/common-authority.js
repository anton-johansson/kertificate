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

import { LOAD_COMMON_AUTHORITIES, LOAD_COMMON_AUTHORITIES_SUCCESS, LOAD_COMMON_AUTHORITIES_FAILED, SET_SEARCH_STRING, LOAD_COMMON_AUTHORITY_SUCCESS } from '../actions/common-authority';

const INITIAL_STATE = {
    loadState: 'not-loaded',
    searchString: '',
    commonAuthorities: [],
    commonAuthority: {
        commonAuthorityId: -1,
        certificateData: undefined,
    },
};

export default function commonAuthoritySearchReducer(state = INITIAL_STATE, action) {
    switch (action.type) {
        case LOAD_COMMON_AUTHORITIES:
            return {
                ...state,
                loadState: 'loading',
            };
        case LOAD_COMMON_AUTHORITIES_SUCCESS:
            return {
                ...state,
                loadState: 'loaded',
                commonAuthorities: action.payload.commonAuthorities,
            };
        case LOAD_COMMON_AUTHORITIES_FAILED:
            return {
                ...state,
                loadState: 'loaded',
                commonAuthorities: [],
            };
        case SET_SEARCH_STRING:
            return {
                ...state,
                searchString: action.payload.searchString,
            };
        case LOAD_COMMON_AUTHORITY_SUCCESS:
            return {
                ...state,
                commonAuthority: {
                    commonAuthorityId: action.payload.commonAuthorityId,
                    name: action.payload.name,
                    comments: action.payload.comments,
                    certificateData: action.payload.certificateData,
                    pem: action.payload.pem,
                },
            };
        default:
            return state;
    }
}
