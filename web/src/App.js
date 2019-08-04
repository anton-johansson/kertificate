import React from 'react';
import { connect } from 'react-redux';

import Main from './Main';
import SignIn from './SignIn';
import Loader from './ui/Loader';

const App = props => {
    if (props.checkingRememberedAuthentication) {
        return <Loader />
    }

    if (props.username) {
        return <Main />
    }

    return <SignIn />
}

const mapStateToProps = state => ({
    checkingRememberedAuthentication: state.authentication.checkingRememberedAuthentication,
    username: state.authentication.information.username,
});

export default connect(mapStateToProps)(App);
