const Frisbee = require('frisbee');

class Api {
    constructor(dispatch, getState) {
        this.dispatch = dispatch;
        this.state = getState()
    }

    setCustomToken(customToken) {
        this.customToken = customToken;
    }

    base() {
        const headers = {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        };

        const token = this.customToken || this.state.authentication.token;
        if (!!token) {
            headers['Authorization'] = token;
        }

        return new Frisbee({
            baseURI: process.env.API_BASE,
            mode: 'cors',
            headers,
        });
    }

    get(path) {
        return this.base().get(path);
    }

    post(path, body) {
        return this.base().post(path, { body });
    }
}

export default Api;
