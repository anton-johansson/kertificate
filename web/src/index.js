import React from 'react';
import { render } from 'react-dom';
import { Provider } from 'react-redux';
import { applyMiddleware, combineReducers, createStore } from 'redux';
import thunk from 'redux-thunk'
import logger from 'redux-logger'

import App from './App';
import authenticationReducer from './reducers/authentication';
import dummyReducer from './reducers/dummy';
import { checkRememberedToken } from './actions/authentication';

const reducer = combineReducers({
    authentication: authenticationReducer,
    dummy: dummyReducer
});

const store = createStore(reducer, applyMiddleware(thunk, logger));
store.dispatch(checkRememberedToken());

render(
    <Provider store={store}>
        <App />
    </Provider>,
    document.getElementById('root')
);

module.hot.accept();
