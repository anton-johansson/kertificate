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
        const token = getToken();
        if (!token) {
            console.log('did not find token');
            dispatch({type: CHECK_REMEMBERED_TOKEN_FAILURE});
            return;
        }

        console.log('found existing token, checking it against api:', token);
        let api = new Frisbee({
            baseURI: process.env.API_BASE,
            mode: 'cors',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
                'Authorization': token,
            },
        });

        const response = await api.get('/v1/authentication/me');

        if (response.status === 200) {
            const {username, firstName, lastName, emailAddress} = response.body;
            dispatch({
                type: CHECK_REMEMBERED_TOKEN_SUCCESS,
                username,
                firstName,
                lastName,
                emailAddress,
            });
            return;
        } else if (response.status !== 401) {
            console.log('Unknown status when logging in:', response.status);
        }
        dispatch({type: CHECK_REMEMBERED_TOKEN_FAILURE});
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
                'Content-Type': 'application/json'
            },
        })

        /*
        const { token } = getState().authentication;
        if (!!token) {
            api.auth(token);
        }
        */

        const response = await api.post('/v1/authentication/authenticate', {
            body: {
                username,
                password,
            },
        });

        if (response.status === 200) {
            const token = response.headers.get('X-Set-Authorization');
            const {username, firstName, lastName, emailAddress} = response.body;
            setToken(token);
            dispatch({
                type: LOGIN_SUCCESS,
                token,
                username,
                firstName,
                lastName,
                emailAddress,
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
