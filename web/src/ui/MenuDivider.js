import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Divider from '@material-ui/core/Divider';

const styles = makeStyles(theme => ({
    divider: {
        marginTop: theme.spacing(1),
        marginBottom: theme.spacing(1),
    },
}));

const MenuDivider = () => {
    const classes = styles();
    return <Divider className={classes.divider} />;
}

export default MenuDivider;
