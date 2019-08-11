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
