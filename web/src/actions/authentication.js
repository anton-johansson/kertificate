const Frisbee = require('frisbee');

const tokenKey = 'X-Authentication-Token';
const getToken = () => localStorage.getItem(tokenKey);
const setToken = token => {
    if (!!token) {
        localStorage.setItem(tokenKey, token);
    } else {
        localStorage.removeItem(tokenKey);
    }
}

export const CHECK_REMEMBERED_TOKEN = 'CHECK_REMEMBERED_TOKEN';
export const CHECK_REMEMBERED_TOKEN_SUCCESS = 'CHECK_REMEMBERED_TOKEN_SUCCESS';
export const CHECK_REMEMBERED_TOKEN_FAILURE = 'CHECK_REMEMBERED_TOKEN_FAILURE';
export const checkRememberedToken = () => {
    return async dispatch => {
        dispatch({type: CHECK_REMEMBERED_TOKEN});
        setTimeout(() => {
            const token = getToken();
            if (!!token) {
                console.log('found existing token:', token);
                dispatch({
                    type: CHECK_REMEMBERED_TOKEN_SUCCESS,
                    username: 'anton',
                    firstName: 'Anton',
                    lastName: 'Johansson',
                });
            } else {
                console.log('did not find token');
                dispatch({type: CHECK_REMEMBERED_TOKEN_FAILURE});
            }
        }, 2000);
    };
}

export const LOGIN = 'LOGIN';
export const LOGIN_SUCCESS = 'LOGIN_SUCCESS';
export const LOGIN_FAILED = 'LOGIN_FAILED';
export const login = (username, password) => {
    return async (dispatch, getState) => {
        let api = new Frisbee({
            baseURI: process.env.API_BASE,
            mode: 'cors',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
            },
        })

        const { token } = getState().authentication;
        if (!!token) {
            api.auth(token);
        }

        const response = await api.post('/v1/authentication/authenticate', {
            body: {
                username,
                password,
            },
        });

        if (response.status === 200) {
            const token = response.headers.get('X-Set-Authorization');
            setToken(token);
            dispatch({
                type: LOGIN_SUCCESS,
                username: 'anton3',
                firstName: 'Anton',
                lastName: 'Johansson',
                token,
            });
            return;
        } else if (response.status !== 401) {
            console.log('Unknown status when logging in:', response.status);
        }
        dispatch({type: LOGIN_FAILED});
    };
}

export const LOGOUT = 'LOGOUT';
export const logout = () => {
    return async dispatch => {
        setToken();
        dispatch({type: LOGOUT});
    }
}
