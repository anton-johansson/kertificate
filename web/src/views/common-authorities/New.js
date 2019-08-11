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
import { makeStyles } from '@material-ui/core/styles';
import { Button, Dialog, DialogTitle, DialogContent, DialogActions, FormControlLabel, LinearProgress, MenuItem, Radio, RadioGroup, Stepper, Step, StepLabel, TextField, Typography } from '@material-ui/core';

const STEP_TEMPLATE = 1;
const STEP_SUBJECT = 2;
const STEP_SETTINGS = 3;
const STEP_GENERATE = 4;
const STEP_GENERATED = 5;

const styles = makeStyles(theme => ({
    dialog: {
        minHeight: 600,
        maxHeight: 600,
    },
    generate: {
        paddingTop: theme.spacing(4),
        textAlign: 'center',
    },
}));

const StepTemplate = () => {
    const [useTemplate, setUseTemplate] = useState(false);
    const [template, setTemplate] = useState('');
    return (
        <div>
            <RadioGroup onChange={event => setUseTemplate('true' === event.target.value)} value={useTemplate}>
                <FormControlLabel label="Do not use template" value={false} control={<Radio color="primary" />} />
                <FormControlLabel label="Use template" value={true} control={<Radio color="primary" />} />
            </RadioGroup>
            <TextField label="Template" variant="outlined" margin="dense" fullWidth select value={template} onChange={event => setTemplate(event.target.value)} disabled={!useTemplate}>
                <MenuItem value="">
                    <em>None</em>
                </MenuItem>
                <MenuItem value={1}>Anton home</MenuItem>
                <MenuItem value={2}>Anton office</MenuItem>
            </TextField>
        </div>
    );
}

const StepSubject = () => {
    return (
        <div>
            <TextField label="Country" variant="outlined" margin="dense" fullWidth />
            <TextField label="Province" variant="outlined" margin="dense" fullWidth />
            <TextField label="Locality" variant="outlined" margin="dense" fullWidth />
            <TextField label="Street address" variant="outlined" margin="dense" fullWidth />
            <TextField label="Postal code" variant="outlined" margin="dense" fullWidth />
            <TextField label="Organization" variant="outlined" margin="dense" fullWidth />
            <TextField label="Organizational unit" variant="outlined" margin="dense" fullWidth />
        </div>
    );
}

const StepSettings = () => {
    const [keyType, setKeyType] = useState('');
    const [keyLength, setKeyLength] = useState('');
    return (
        <div>
            <TextField label="Common name" variant="outlined" margin="dense" fullWidth />
            <TextField label="Key type" variant="outlined" margin="dense" fullWidth select value={keyType} onChange={event => setKeyType(event.target.value)}>
                <MenuItem value="">
                    <em>None</em>
                </MenuItem>
                <MenuItem value="rsa">RSA</MenuItem>
                <MenuItem value="dsa">DSA</MenuItem>
            </TextField>
            <TextField label="Key length" variant="outlined" margin="dense" fullWidth select value={keyLength} onChange={event => setKeyLength(event.target.value)}>
                <MenuItem value="">
                    <em>None</em>
                </MenuItem>
                <MenuItem value={1024}>1024</MenuItem>
                <MenuItem value={2048}>2048</MenuItem>
                <MenuItem value={4096}>4096</MenuItem>
            </TextField>
            <TextField label="Optional passphrase" variant="outlined" margin="dense" type="password" fullWidth />
            <TextField label="Repeat optional passphrase" variant="outlined" margin="dense" type="password" fullWidth />
        </div>
    );
}

const StepGenerate = ({generated}) => {
    const classes = styles();
    const message = generated ? 'Certificate generated successfully!' : 'Generating certificate...';
    return (
        <div className={classes.generate}>
            <Typography variant="h6" paragraph>{message}</Typography>
            <LinearProgress />
        </div>
    );
}

const getStepContent = stepNumber => {
    switch (stepNumber) {
        case STEP_TEMPLATE: return <StepTemplate />
        case STEP_SUBJECT: return <StepSubject />
        case STEP_SETTINGS: return <StepSettings />
        case STEP_GENERATE: return <StepGenerate generated={false} />
        case STEP_GENERATED: return <StepGenerate generated={true} />
    }
}

const New = ({open, setNewDialogOpened}) => {
    const classes = styles();
    const [activeStep, setActiveStep] = useState(STEP_TEMPLATE);
    const isInitialStep = activeStep === STEP_TEMPLATE;
    const previousDisabled = activeStep > STEP_SETTINGS;
    const previousLabel = isInitialStep ? 'Cancel' : 'Back';
    const nextDisabled = activeStep === STEP_GENERATE;
    const nextLabel = activeStep === STEP_SETTINGS ? 'Generate' : activeStep >= STEP_GENERATE ? 'Open' : 'Next';
    const previous = () => isInitialStep ? setNewDialogOpened(false) : setActiveStep(activeStep - 1);
    const next = () => setActiveStep(activeStep + 1);
    return (
        <Dialog open={open} maxWidth="sm" fullWidth={true} classes={{ paper: classes.dialog }}>
            <DialogTitle>Create new common authority</DialogTitle>
            <DialogContent dividers>
                <Stepper activeStep={activeStep - 1}>
                    <Step>
                        <StepLabel>Template</StepLabel>
                    </Step>
                    <Step>
                        <StepLabel>Subject</StepLabel>
                    </Step>
                    <Step>
                        <StepLabel>Settings</StepLabel>
                    </Step>
                    <Step>
                        <StepLabel>Generate</StepLabel>
                    </Step>
                </Stepper>
                {getStepContent(activeStep)}
            </DialogContent>
            <DialogActions>
                <Button color="primary" disabled={previousDisabled} onClick={previous}>{previousLabel}</Button>
                <Button color="primary" disabled={nextDisabled} onClick={next} variant="contained">{nextLabel}</Button>
            </DialogActions>
        </Dialog>
    );
}

export default New;
