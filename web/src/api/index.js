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

import Frisbee from 'frisbee';

const tokenKey = 'X-Authentication-Token';
export const getToken = () => localStorage.getItem(tokenKey);
export const setToken = token => {
    if (!!token) {
        localStorage.setItem(tokenKey, token);
    } else {
        localStorage.removeItem(tokenKey);
    }
}

export const REFRESH_TOKEN = 'REFRESH_TOKEN';
const refreshToken = newToken => {
    return async dispatch => {
        dispatch({
            type: REFRESH_TOKEN,
            newToken,
        });
    }
}

class Api {
    constructor(dispatch, getState) {
        this.dispatch = dispatch;
        this.token = getState().authentication.token;
        this.checkResponse = this.checkResponse.bind(this);
    }

    setToken(token) {
        this.token = token;
    }

    getToken() {
        return this.token;
    }

    base() {
        const headers = {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        };

        if (!!this.token) {
            headers['Authorization'] = this.token;
        }

        return new Frisbee({
            baseURI: process.env.API_BASE,
            mode: 'cors',
            headers,
        });
    }

    checkResponse(response) {
        const newToken = response.headers.get('X-Set-Authorization');
        if (!!newToken) {
            console.log('got new token:', newToken);
            this.setToken(newToken);
            setToken(newToken);
            this.dispatch(refreshToken(newToken));
        }
        return response;
    }

    get(path) {
        return this.base()
                .get(path)
                .then(this.checkResponse);
    }

    post(path, body) {
        return this.base()
                .post(path, { body })
                .then(this.checkResponse);
    }
}

export default Api;
