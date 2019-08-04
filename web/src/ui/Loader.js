import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import CircularProgress from '@material-ui/core/CircularProgress';

const styles = makeStyles(theme => ({
    container: {
        marginTop: theme.spacing(16),
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
    },
}));

const Loader = () => {
    const classes = styles();
    return (
        <div className={classes.container}>
            <CircularProgress variant="indeterminate" />
        </div>
    );
}

export default Loader;
