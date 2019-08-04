import { CHECK_REMEMBERED_TOKEN_FAILURE, LOGIN_SUCCESS, LOGIN_FAILED, LOGOUT, CHECK_REMEMBERED_TOKEN_SUCCESS } from '../actions/authentication';

const initial = {
    checkingRememberedAuthentication: true,
    token: '',
    information: {
        username: '',
        firstName: '',
        lastName: '',
    },
};

export default function authenticationReducer(state = initial, action) {
    switch (action.type) {
        case CHECK_REMEMBERED_TOKEN_SUCCESS:
            return {
                ...state,
                checkingRememberedAuthentication: false,
                information: {
                    username: action.username,
                    firstName: action.firstName,
                    lastName: action.lastName,
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
                },
            };
        case LOGIN_FAILED:
            console.log('login failed!');
            return state;
        case LOGOUT:
            return {
                ...state,
                token: '',
                information: {
                    username: '',
                    firstName: '',
                    lastName: '',
                },
            }
        default:
            return state;
    }
}
