import { CHECK_REMEMBERED_TOKEN, CHECK_REMEMBERED_TOKEN_FAILURE, LOGIN_SUCCESS, LOGIN_FAILED, LOGOUT, CHECK_REMEMBERED_TOKEN_SUCCESS } from '../actions/authentication';

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
        default:
            return state;
    }
}
