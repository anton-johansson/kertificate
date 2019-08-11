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
import Button from '@material-ui/core/Button';
import Icon from '@material-ui/core/Icon';

const styles = makeStyles(theme => ({
    button: {
        padding: 10,
        justifyContent: 'left',
        textTransform: 'none',
        color: theme.palette.grey[600],
    },
    link: {
        textDecoration: 'none',
    },
    icon: {
        paddingRight: 10,
    },
    selected: {
        color: theme.palette.primary.main,
    },
}));

const MenuItem = ({iconName, exact, path, title}) => {
    const classes = styles({match: false});
    return (
        <div>
            <Route exact={exact} path={path} children={({match}) => {
                const buttonClassName = (match ? [classes.selected, classes.button] : [classes.button]).join(" ");
                return (
                    <Link to={path} tabIndex={-1} className={classes.link}>
                        <Button fullWidth={true} className={buttonClassName}>
                            <Icon className={classes.icon}>{iconName}</Icon>
                            {title}
                        </Button>
                    </Link>
                );
            }} />
        </div>
    );
}

export default MenuItem;
