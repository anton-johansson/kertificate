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
import { Box, Button, ButtonGroup, Fab, Icon, Paper, Tab, Tabs, TextField, Typography } from '@material-ui/core';
import { ArrowDropDown } from '@material-ui/icons';
import { makeStyles } from '@material-ui/core/styles';

import { loadCommonAuthority } from '../../actions/common-authority';
import { getFingerprintSHA1, getFingerprintSHA256 } from '../../pki';
import { DateTime, SplitButton } from '../../ui';

const styles = makeStyles(theme => ({
    area: {
        margin: theme.spacing(2),
    },
    details: {
        display: 'flex',
    },
    detailsLeft: {
        width: 400,
        marginRight: theme.spacing(4),
    },
    detailsRight: {
        marginTop: theme.spacing(1),
    },
    detailsPart: {
        marginLeft: theme.spacing(2),
        display: 'flex',
        flexDirection: 'row',
    },
    detailsPartTitle: {
        width: 300,
        color: 'gray',
    },
    detailsPartValue: {
        width: '100%',
        color: 'gray',
    },
    detailsSection: {
        marginBottom: theme.spacing(4),
    },
    detailsSectionTitle: {
        marginBottom: theme.spacing(2),
    },
    paper: {
        marginTop: 0,
        display: 'flex',
        overflow: 'auto',
        flexDirection: 'column',
    },
    tabs: {
        display: 'flex',
        overflow: 'auto',
        flexDirection: 'column',
    },
    tab: {
        margin: theme.spacing(2),
    },
    tabButton: {
        marginLeft: theme.spacing(4),
    },
    pem: {
        fontFamily: 'Ubuntu Mono',
        fontSize: 16,
    },
    toolbar: {
        paddingBottom: theme.spacing(2),
        display: 'flex',
    },
    close: {
        marginLeft: 'auto',
    },
    toolbarButton: {
        marginRight: theme.spacing(1),
    },
}));

const viewOptions = ['View certificate as PEM', 'View private key as PEM'];
const exportOptions = ['Export certificate as PEM', 'Export private key as PEM', 'Export as PKCS12'];

const TabPanel = ({children, value, selectedTab}) => {
    const classes = styles();
    return <div className={classes.tab} hidden={value !== selectedTab}>{children}</div>;
}

const Edit = ({commonAuthority, match: {params: {commonAuthorityId}}, loadCommonAuthority}) => {
    loadCommonAuthority(commonAuthorityId);

    const [selectedTab, setSelectedTab] = useState("details");
    const classes = styles();
    return (
        <div className={classes.area}>
            <div className={classes.toolbar}>
                <Button className={classes.toolbarButton} variant="contained" color="primary">
                    <Icon>view_list</Icon>&nbsp;&nbsp;Save
                </Button>
                <SplitButton className={classes.toolbarButton} iconName="folder_open" options={viewOptions} width="300"/>
                <SplitButton className={classes.toolbarButton} iconName="import_export" options={exportOptions} />
                <Fab color="primary" className={classes.close}>
                    <Icon>keyboard_return</Icon>
                </Fab>
            </div>
            <Paper className={classes.paper}>
                <Paper square className={classes.tabs} elevation={1}>
                    <Tabs value={selectedTab} onChange={(_, tab) => setSelectedTab(tab)} indicatorColor="primary" textColor="primary" variant="standard">
                        <Tab className={classes.tabButton} value="details" label="Certificate details" icon={<Icon>view_list</Icon>} />
                        <Tab className={classes.tabButton} value="data" label="Certificate data" icon={<Icon>insert_drive_file</Icon>} />
                        <Tab className={classes.tabButton} value="renewal" label="Automatic renewal" icon={<Icon>autorenew</Icon>} />
                    </Tabs>
                </Paper>
                <TabPanel value="details" selectedTab={selectedTab}>
                    <div className={classes.details}>
                        <div className={classes.detailsLeft}>
                            <TextField variant="outlined" margin="dense" label="Name" fullWidth />
                            <TextField variant="outlined" margin="dense" label="Comments" fullWidth multiline />
                        </div>
                        <div className={classes.detailsRight}>
                            <div className={classes.detailsSection}>
                                <Typography className={classes.detailsSectionTitle} variant="subtitle2">Subject</Typography>
                                <div className={classes.detailsPart}>
                                    <Typography className={classes.detailsPartTitle}>Common Name (CN)</Typography>
                                    <Typography className={classes.detailsPartValue}>{commonAuthority.certificateData && commonAuthority.certificateData.subject.commonName}</Typography>
                                </div>
                                <div className={classes.detailsPart}>
                                    <Typography className={classes.detailsPartTitle}>Organization (O)</Typography>
                                    <Typography className={classes.detailsPartValue}>{commonAuthority.certificateData && commonAuthority.certificateData.subject.organizationName}</Typography>
                                </div>
                                <div className={classes.detailsPart}>
                                    <Typography className={classes.detailsPartTitle}>Organizational Unit (OU)</Typography>
                                    <Typography className={classes.detailsPartValue}>{commonAuthority.certificateData && commonAuthority.certificateData.subject.organizationalUnitName}</Typography>
                                </div>
                            </div>
                            <div className={classes.detailsSection}>
                                <Typography className={classes.detailsSectionTitle} variant="subtitle2">Validity period</Typography>
                                <div className={classes.detailsPart}>
                                    <Typography className={classes.detailsPartTitle}>Issued on</Typography>
                                    <Typography className={classes.detailsPartValue}>{commonAuthority.certificateData && <DateTime value={commonAuthority.certificateData.validity.notBefore} />}</Typography>
                                </div>
                                <div className={classes.detailsPart}>
                                    <Typography className={classes.detailsPartTitle}>Expires on</Typography>
                                    <Typography className={classes.detailsPartValue}>{commonAuthority.certificateData && <DateTime value={commonAuthority.certificateData.validity.notAfter} />}</Typography>
                                </div>
                            </div>
                            <div className={classes.detailsSection}>
                                <Typography className={classes.detailsSectionTitle} variant="subtitle2">Fingerprints</Typography>
                                <div className={classes.detailsPart}>
                                    <Typography className={classes.detailsPartTitle}>SHA-256 fingerprint</Typography>
                                    <Typography className={classes.detailsPartValue}>{getFingerprintSHA256(commonAuthority.certificateData)}</Typography>
                                </div>
                                <div className={classes.detailsPart}>
                                    <Typography className={classes.detailsPartTitle}>SHA-1 fingerprint</Typography>
                                    <Typography className={classes.detailsPartValue}>{getFingerprintSHA1(commonAuthority.certificateData)}</Typography>
                                </div>
                            </div>
                        </div>
                    </div>
                </TabPanel>
                <TabPanel value="data" selectedTab={selectedTab}>
                    <Box>
                        <Typography variant="caption" className={classes.pem}>
                            {commonAuthority.pem && commonAuthority.pem.split('\n').map((data, key) => <div key={key}>{data}</div>)}
                        </Typography>
                    </Box>
                </TabPanel>
                <TabPanel value="renewal" selectedTab={selectedTab}>
                    <Typography>Not implemented.</Typography>
                </TabPanel>
            </Paper>
        </div>
    );
}

const mapStateToProps = state => ({
    commonAuthority: state.commonAuthority.commonAuthority,
});
const mapDispatchToProps = {loadCommonAuthority};

export default connect(mapStateToProps, mapDispatchToProps)(Edit);
