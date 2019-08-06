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

import { CHECK_REMEMBERED_TOKEN, CHECK_REMEMBERED_TOKEN_FAILURE, LOGIN_SUCCESS, LOGIN_FAILED, LOGOUT, CHECK_REMEMBERED_TOKEN_SUCCESS } from '../actions/authentication';
import { REFRESH_TOKEN } from '../api';

const INITIAL_STATE = {
    checkingRememberedAuthentication: false,
    token: '',
    information: {
        username: '',
        firstName: '',
        lastName: '',
        emailAddress: '',
    },
};

export default function authenticationReducer(state = INITIAL_STATE, action) {
    switch (action.type) {
        case CHECK_REMEMBERED_TOKEN:
            return {
                ...state,
                checkingRememberedAuthentication: true,
            };
        case CHECK_REMEMBERED_TOKEN_SUCCESS:
            return {
                ...state,
                checkingRememberedAuthentication: false,
                token: action.token,
                information: {
                    username: action.username,
                    firstName: action.firstName,
                    lastName: action.lastName,
                    emailAddress: action.emailAddress,
                },
            };
        case CHECK_REMEMBERED_TOKEN_FAILURE:
            return {
                ...state,
                checkingRememberedAuthentication: false,
            };
        case LOGIN_SUCCESS:
            return {
                ...state,
                token: action.token,
                information: {
                    username: action.username,
                    firstName: action.firstName,
                    lastName: action.lastName,
                    emailAddress: action.emailAddress,
                },
            };
        case LOGIN_FAILED:
            console.log('login failed!');
            return state;
        case LOGOUT:
            return INITIAL_STATE;
        case REFRESH_TOKEN:
            return {
                ...state,
                token: action.newToken,
            };
        default:
            return state;
    }
}
