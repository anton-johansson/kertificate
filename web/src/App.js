import React from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar'
import Avatar from '@material-ui/core/Avatar';
import Drawer from '@material-ui/core/Drawer';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';

import MenuItem from './ui/MenuItem';
import MenuDivider from './ui/MenuDivider';
import Dashboard from './views/Dashboard';
import Users from './views/Users';

const drawerWidth = 280;

const styles = makeStyles(theme => ({
    '@global': {
        body: {
            backgroundColor: '#f4f6f8',
        },
    },
    appBar: {
        zIndex: theme.zIndex.drawer + 1,
    },
    avatarFrame: {
        marginTop: theme.spacing(1),
        marginBottom: theme.spacing(1),
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
    },
    avatar: {
        marginTop: theme.spacing(1),
        marginBottom: theme.spacing(1),
        width: 60,
        height: 60,
    },
    content: {
      padding: theme.spacing(1),
      marginLeft: drawerWidth + theme.spacing(4),
    },
    drawer: {
        width: drawerWidth,
        paddingTop: theme.spacing(1),
        paddingLeft: theme.spacing(2),
        paddingRight: theme.spacing(2),
    },
    toolbar: theme.mixins.toolbar,
}));

const App = () => {
    const classes = styles();
    return (
        <div>
            <Router>
                <AppBar className={classes.appBar}>
                    <Toolbar>
                        <Typography variant="h5">
                            Kertificate
                        </Typography>
                    </Toolbar>
                </AppBar>
                <Drawer classes={{paper: classes.drawer}} open={true} variant="persistent">
                    <div className={classes.toolbar} />
                    <div className={classes.avatarFrame}>
                        <Avatar className={classes.avatar} alt="Anton Johansson" src="https://avatars1.githubusercontent.com/u/6347803" />
                        <Typography>
                            Anton Johansson
                        </Typography>
                    </div>
                    <MenuDivider />
                    <MenuItem title="Dashboard" path="/" exact={true} iconName="dashboard" selected={true} />
                    <MenuItem title="Users"  path="/users" iconName="supervisor_account" />
                    <MenuItem title="Configuration" path="/configuration" iconName="settings" />
                    <MenuDivider />
                    <MenuItem title="Common authorities" path="/common-authorities" iconName="assignment_turned_in" />
                    <MenuItem title="Certificates" path="/certificates" iconName="assignment" />
                    <MenuDivider />
                    <MenuItem title="Consumer types" path="/consumer-types" iconName="data_usage" />
                    <MenuItem title="Certificate templates" path="/certificate-templates" iconName="list" />
                </Drawer>
                <main className={classes.content}>
                    <div className={classes.toolbar} />
                    <Route exact path="/" component={Dashboard} />
                    <Route path="/users" component={Users} />
                </main>
            </Router>
        </div>
    )
};

export default App;
