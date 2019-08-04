const tokenKey = 'X-Authentication-Token';
const getToken = () => localStorage.getItem(tokenKey);
const setToken = token => {
    if (!!token) {
        localStorage.setItem(tokenKey, token);
    } else {
        localStorage.removeItem(tokenKey);
    }
}

export const CHECK_REMEMBERED_TOKEN_SUCCESS = 'CHECK_REMEMBERED_TOKEN_SUCCESS';
export const CHECK_REMEMBERED_TOKEN_FAILURE = 'CHECK_REMEMBERED_TOKEN_FAILURE';
export const checkRememberedToken = () => {
    return async dispatch => {
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
    return async dispatch => {
        setTimeout(() => {
            if (username === 'anton' && password === 's3cr3t') {
                setToken('abc123');
                dispatch({
                    type: LOGIN_SUCCESS,
                    username: 'anton',
                    firstName: 'Anton',
                    lastName: 'Johansson',
                    token: 'abc123',
                });
            } else {
                dispatch({type: LOGIN_FAILED});
            }
        }, 2000);
    };
}

export const LOGOUT = 'LOGOUT';
export const logout = () => {
    return async dispatch => {
        setToken();
        dispatch({type: LOGOUT});
    }
}
