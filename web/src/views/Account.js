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
import { Avatar, Button, Divider, Grid, Paper, TextField, Typography } from '@material-ui/core';

const styles = makeStyles(theme => ({
    profileFrame: {
        display: 'flex',
    },
    profileName: {
        margin: theme.spacing(2),
    },
    profileAvatar: {
        margin: theme.spacing(2),
        width: 100,
        height: 100,
        marginLeft: 'auto',
    },
    paper: {
      margin: theme.spacing(2),
      display: 'flex',
      overflow: 'auto',
      flexDirection: 'column',
    },
    textField: {
        width: '100%',
    },
    container: {
        paddingLeft: theme.spacing(2),
        paddingTop: theme.spacing(2),
    },
    grid: {
        paddingRight: theme.spacing(2),
        paddingBottom: theme.spacing(2),
    },
    button: {
        margin: theme.spacing(1),
    },
}));

export default function Dashboard() {
    const classes = styles();
    return (
        <div>
            <Grid container>
                <Grid item xs={4}>
                    <Paper className={classes.paper}>
                        <div className={classes.profileFrame}>
                            <Typography className={classes.profileName} variant="h5">Anton Johansson</Typography>
                            <Avatar className={classes.profileAvatar} alt={`Anton Johansson`} src="https://avatars1.githubusercontent.com/u/6347803" />
                        </div>
                        <Divider />
                        <div>
                            <Button className={classes.button} fullWidth={false} color="primary" disabled={false}>Upload avatar</Button>
                            <Button className={classes.button} fullWidth={false} color="default" disabled={false}>Remove avatar</Button>
                        </div>
                    </Paper>
                </Grid>
                <Grid item xs={8}>
                    <Paper className={classes.paper}>
                        <div className={classes.container}>
                            <Typography variant="h6">
                                Profile
                            </Typography>
                            <Typography paragraph variant="caption">
                                Some fields may not be editable if accounts are created using an external provider.
                            </Typography>
                        </div>
                        <Divider />
                        <Grid container className={classes.container}>
                            <Grid item xs={6} className={classes.grid}>
                                <TextField className={classes.textField} label="Username" variant="outlined" margin="dense" defaultValue="anton3" disabled={true} />
                            </Grid>
                            <Grid item xs={6} className={classes.grid}>
                                <TextField className={classes.textField} label="Email address" variant="outlined" margin="dense" defaultValue="antoon.johansson@gmail.com" />
                            </Grid>
                            <Grid item xs={6} className={classes.grid}>
                                <TextField className={classes.textField} label="First name" variant="outlined" margin="dense" defaultValue="Anton" />
                            </Grid>
                            <Grid item xs={6} className={classes.grid}>
                                <TextField className={classes.textField} label="Last name" variant="outlined" margin="dense" defaultValue="Johansson" />
                            </Grid>
                        </Grid>
                        <Divider />
                        <div>
                            <Button className={classes.button} fullWidth={false} variant="contained" color="primary" disabled={false}>Save details</Button>
                        </div>
                    </Paper>
                </Grid>
            </Grid>
        </div>
    );
}
