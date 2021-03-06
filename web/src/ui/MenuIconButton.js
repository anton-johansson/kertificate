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
import { Route, Link } from 'react-router-dom';
import { makeStyles } from '@material-ui/core/styles';
import Icon from '@material-ui/core/Icon';
import IconButton from '@material-ui/core/IconButton';
import Tooltip from '@material-ui/core/Tooltip';

const styles = makeStyles(theme => ({
    appBarIcon: {
        color: theme.palette.common.white,
    },
    appBarIconSelected: {
        color: theme.palette.secondary.light,
    },
}));

const MenuItem = ({iconName, tooltip, exact, path}) => {
    const classes = styles();
    return (
        <Route exact={exact} path={path} children={({history, match}) => (
            <Tooltip title={tooltip}>
                <IconButton className={match ? classes.appBarIconSelected : classes.appBarIcon} onClick={() => history.push(path)}>
                    <Icon>{iconName}</Icon>
                </IconButton>
            </Tooltip>
        )} />
    );
}

export default MenuItem;
