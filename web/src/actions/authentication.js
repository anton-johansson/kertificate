import Api, { getToken, setToken } from '../api';

export const CHECK_REMEMBERED_TOKEN = 'CHECK_REMEMBERED_TOKEN';
export const CHECK_REMEMBERED_TOKEN_SUCCESS = 'CHECK_REMEMBERED_TOKEN_SUCCESS';
export const CHECK_REMEMBERED_TOKEN_FAILURE = 'CHECK_REMEMBERED_TOKEN_FAILURE';
export const checkRememberedToken = () => {
    return async (dispatch, getState) => {
        dispatch({type: CHECK_REMEMBERED_TOKEN});
        const token = getToken();
        if (!token) {
            console.log('did not find token');
            dispatch({type: CHECK_REMEMBERED_TOKEN_FAILURE});
            return;
        }

        console.log('found existing token, checking it against api:', token);
        const api = new Api(dispatch, getState);
        api.setToken(token);

        const response = await api.get('/v1/authentication/me');
        if (response.status === 200) {
            const {username, firstName, lastName, emailAddress} = response.body;
            dispatch({
                type: CHECK_REMEMBERED_TOKEN_SUCCESS,
                token: api.getToken(),
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
        const api = new Api(dispatch, getState);
        const response = await api.post('/v1/authentication/authenticate', { username, password });

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
