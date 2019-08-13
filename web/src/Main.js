/**
 * Copyright 2019 Anton Johansson
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import React from 'react';
import { connect } from 'react-redux';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar'
import Avatar from '@material-ui/core/Avatar';
import Drawer from '@material-ui/core/Drawer';
import Badge from '@material-ui/core/Badge';
import Icon from '@material-ui/core/Icon';
import IconButton from '@material-ui/core/IconButton';
import Toolbar from '@material-ui/core/Toolbar';
import Tooltip from '@material-ui/core/Tooltip';
import Typography from '@material-ui/core/Typography';

import MenuDivider from './ui/MenuDivider';
import MenuIconButton from './ui/MenuIconButton';
import MenuItem from './ui/MenuItem';

import { Account, CommonAuthorities, Dashboard, Users } from './views';

import { logout } from './actions/authentication';

const drawerWidth = 280;

const styles = makeStyles(theme => ({
    '@global': {
        body: {
            backgroundColor: '#f5f5f5',
        },
    },
    appBar: {
        zIndex: theme.zIndex.drawer + 1,
    },
    appBarIcon: {
        color: theme.palette.common.white,
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
    grow: {
        flexGrow: 1,
    },
    toolbar: theme.mixins.toolbar,
}));

const Main = ({firstName, lastName, logout}) => {
    const classes = styles();
    return (
        <div>
            <Router>
                <AppBar className={classes.appBar}>
                    <Toolbar>
                        <Typography variant="h5">
                            Kertificate
                        </Typography>
                        <div className={classes.grow} />
                        <Tooltip title="Notifications">
                            <IconButton className={classes.appBarIcon}>
                                <Badge badgeContent={2} color="secondary">
                                    <Icon>notifications</Icon>
                                </Badge>
                            </IconButton>
                        </Tooltip>
                        <MenuIconButton path="/account" tooltip="Account" iconName="account_circle" exact={true} />
                        <Tooltip title="Logout">
                            <IconButton className={classes.appBarIcon} onClick={logout}>
                                <Icon>exit_to_app</Icon>
                            </IconButton>
                        </Tooltip>
                    </Toolbar>
                </AppBar>
                <Drawer classes={{paper: classes.drawer}} open={true} variant="persistent">
                    <div className={classes.toolbar} />
                    <div className={classes.avatarFrame}>
                        <Avatar className={classes.avatar} alt={`${firstName} ${lastName}`} src="https://avatars1.githubusercontent.com/u/6347803" />
                        <Typography>{`${firstName} ${lastName}`}</Typography>
                    </div>
                    <MenuDivider />
                    <MenuItem title="Dashboard" path="/" exact={true} iconName="dashboard" />
                    <MenuItem title="Account" path="/account" exact={true} iconName="account_circle" />
                    <MenuItem title="Users" path="/users" iconName="supervisor_account" />
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
                    <Route exact path="/account" component={Account} />
                    <Route exact path="/" component={Dashboard} />
                    <Route path="/users" component={Users} />
                    <Route path="/common-authorities" component={CommonAuthorities} />
                </main>
            </Router>
        </div>
    )
};

const mapStateToProps = state => ({
    firstName: state.authentication.information.firstName,
    lastName: state.authentication.information.lastName,
});
const mapDispatchToProps = {logout};

export default connect(mapStateToProps, mapDispatchToProps)(Main);
