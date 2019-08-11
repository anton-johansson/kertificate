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
import { makeStyles } from '@material-ui/core/styles';
import { Fab, Icon, InputBase, Paper, Table, TableBody, TableCell, TableHead, TablePagination, TableRow, Tooltip } from '@material-ui/core';
import { Add } from '@material-ui/icons';

const styles = makeStyles(theme => ({
    addNew: {
        marginLeft: 'auto',
        margin: theme.spacing(2),
    },
    paper: {
        margin: theme.spacing(2),
        marginTop: 0,
        display: 'flex',
        overflow: 'auto',
        flexDirection: 'column',
    },
    search: {
        margin: theme.spacing(2),
        display: 'block',
        marginBottom: theme.spacing(2),
        padding: theme.spacing(1),
        display: 'flex',
        flexDirection: 'auto',
        width: 400,
        verticalAlign: 'center',
    },
    searchIcon: {
      //height: '100%',
      position: 'absolute',
      pointerEvents: 'none',
      //display: 'flex',
      alignItems: 'center',
      justifyContent: 'center',
    },
    searchInput: {
      padding: theme.spacing(0, 0, 0, 4),
      transition: theme.transitions.create('width'),
      width: '100%',
      [theme.breakpoints.up('md')]: {
        width: 200,
      },
    },
    topBar: {
        display: 'flex',
        alignItems: 'center',
    },
}));

const CommonAuthority = ({name, validFrom, validTo, onClick}) => {
    return (
        <TableRow hover onClick={onClick}>
            <TableCell>{name}</TableCell>
            <TableCell>{validFrom}</TableCell>
            <TableCell>{validTo}</TableCell>
        </TableRow>
    );
}

const Search = ({onSelectCA, onNewCA}) => {
    const classes = styles();
    return (
        <div>
            <div className={classes.topBar}>
                <Paper className={classes.search}>
                    <Icon className={classes.searchIcon}>search</Icon>
                    <InputBase className={classes.searchInput} placeholder="Search authority..." />
                </Paper>
                <Tooltip title="Create new common authority" placement="left">
                    <Fab className={classes.addNew} color="primary" onClick={onNewCA}>
                        <Add />
                    </Fab>
                </Tooltip>
            </div>
            <Paper className={classes.paper}>
                <Table>
                    <TableHead>
                        <TableRow>
                            <TableCell>Name</TableCell>
                            <TableCell>Valid from</TableCell>
                            <TableCell>Valid to</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        <CommonAuthority name="Antons CA" validFrom="2019-08-08" validTo="2020-08-08" onClick={() => onSelectCA(1)} />
                        <CommonAuthority name="Some random CA" validFrom="2019-06-01" validTo="2019-09-01" onClick={() => onSelectCA(2)} />
                    </TableBody>
                </Table>
                <TablePagination
                    component="div"
                    rowsPerPageOptions={[10, 25, 50]}
                    rowsPerPage={25}
                    onChangePage={() => console.log('changed page')}
                    count={32}
                    page={0} />
            </Paper>
        </div>
    );
}

export default Search;
