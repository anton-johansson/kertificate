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
import { Paper, Typography } from '@material-ui/core';

const styles = makeStyles(theme => ({
    paper: {
      padding: theme.spacing(2),
      display: 'flex',
      overflow: 'auto',
      flexDirection: 'column',
    },
}));

export default function Dashboard() {
    const classes = styles();
    return (
        <div>
            <Paper className={classes.paper}>
                <Typography>Dashboard</Typography>
            </Paper>
            <Paper style={{margin: '10px 0px'}} className={classes.paper}>
                <Typography>Second card</Typography>
            </Paper>
            <Typography paragraph>
                Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt
                ut labore et dolore magna aliqua. Rhoncus dolor purus non enim praesent elementum
                facilisis leo vel. Risus at ultrices mi tempus imperdiet. Semper risus in hendrerit
                gravida rutrum quisque non tellus. Convallis convallis tellus id interdum velit laoreet id
                donec ultrices. Odio morbi quis commodo odio aenean sed adipiscing. Amet nisl suscipit
                adipiscing bibendum est ultricies integer quis. Cursus euismod quis viverra nibh cras.
                Metus vulputate eu scelerisque felis imperdiet proin fermentum leo. Mauris commodo quis
                imperdiet massa tincidunt. Cras tincidunt lobortis feugiat vivamus at augue. At augue eget
                arcu dictum varius duis at consectetur lorem. Velit sed ullamcorper morbi tincidunt. Lorem
                donec massa sapien faucibus et molestie ac.
            </Typography>
            <Typography paragraph>
                Consequat mauris nunc congue nisi vitae suscipit. Fringilla est ullamcorper eget nulla
                facilisi etiam dignissim diam. Pulvinar elementum integer enim neque volutpat ac
                tincidunt. Ornare suspendisse sed nisi lacus sed viverra tellus. Purus sit amet volutpat
                consequat mauris. Elementum eu facilisis sed odio morbi. Euismod lacinia at quis risus sed
                vulputate odio. Morbi tincidunt ornare massa eget egestas purus viverra accumsan in. In
                hendrerit gravida rutrum quisque non tellus orci ac. Pellentesque nec nam aliquam sem et
                tortor. Habitant morbi tristique senectus et. Adipiscing elit duis tristique sollicitudin
                nibh sit. Ornare aenean euismod elementum nisi quis eleifend. Commodo viverra maecenas
                accumsan lacus vel facilisis. Nulla posuere sollicitudin aliquam ultrices sagittis orci a.
            </Typography>
        </div>
    );
}
