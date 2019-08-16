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

import React, { useState } from 'react';
import { connect } from 'react-redux';
import { makeStyles } from '@material-ui/core/styles';
import { Fab, Icon, InputBase, Paper, Table, TableBody, TableCell, TableHead, TablePagination, TableRow, Tooltip } from '@material-ui/core';
import { Add } from '@material-ui/icons';

import DateTime from '../../ui/DateTime';
import { loadCommonAuthorities, setSearchString } from '../../actions/common-authority';

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
        //display: 'block',
        marginBottom: theme.spacing(2),
        padding: theme.spacing(1),
        display: 'flex',
        //flexDirection: 'row',
        width: 400,
        //verticalAlign: 'center',
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
    },
    topBar: {
        display: 'flex',
        alignItems: 'center',
    },
}));

const CommonAuthority = ({name, notBefore, notAfter, onClick}) => {
    return (
        <TableRow hover onClick={onClick}>
            <TableCell>
                {name}
            </TableCell>
            <TableCell>
                <DateTime value={notBefore} /></TableCell>
            <TableCell>
                <DateTime value={notAfter} />
            </TableCell>
        </TableRow>
    );
}

const Search = ({commonAuthorities, searchString, onSelectCA, onNewCA, loadCommonAuthorities, setSearchString}) => {
    loadCommonAuthorities();

    const [rowsPerPage, setRowsPerPage] = useState(10);
    const [page, setPage] = useState(0);

    const onChangePage = (_, newPage) => setPage(newPage);
    const onChangeRowsPerPage = event => {
        setPage(0);
        setRowsPerPage(event.target.value);
    };
    const onChangeSearchString = event => {
        setPage(0);
        setSearchString(event.target.value);
    }

    const filteredList = !!searchString ? commonAuthorities.filter(commonAuthority => commonAuthority.name.toLowerCase().includes(searchString.toLowerCase())) : commonAuthorities;

    const classes = styles();
    return (
        <div>
            <div className={classes.topBar}>
                <Paper className={classes.search}>
                    <Icon className={classes.searchIcon}>search</Icon>
                    <InputBase className={classes.searchInput} placeholder="Search common authority..." onChange={onChangeSearchString} />
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
                        {
                            filteredList.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage).map(commonAuthority => (
                                <CommonAuthority
                                    key={`common-authority-${commonAuthority.commonAuthorityId}`}
                                    name={commonAuthority.name}
                                    notBefore={commonAuthority.notBefore}
                                    notAfter={commonAuthority.notAfter}
                                    onClick={() => onSelectCA(commonAuthority.commonAuthorityId)}
                                    />
                            ))
                        }
                    </TableBody>
                </Table>
                <TablePagination
                    component="div"
                    count={filteredList.length}
                    rowsPerPageOptions={[5, 10, 25]}
                    rowsPerPage={rowsPerPage}
                    onChangeRowsPerPage={onChangeRowsPerPage}
                    page={page}
                    onChangePage={onChangePage}
                    />
            </Paper>
        </div>
    );
}

const mapStateToProps = state => ({
    searchString: state.commonAuthority.searchString,
    commonAuthorities: state.commonAuthority.commonAuthorities,
});
const mapDispatchToProps = {loadCommonAuthorities, setSearchString};

export default connect(mapStateToProps, mapDispatchToProps)(Search);
