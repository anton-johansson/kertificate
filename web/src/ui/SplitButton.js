import React, { useState, useRef } from 'react';
import { Button, ButtonGroup, ClickAwayListener, Grow, Icon, MenuItem, MenuList, Paper, Popper } from '@material-ui/core';
import { ArrowDropDown } from '@material-ui/icons';

const SplitButton = ({className, iconName, options}) => {
    const [open, setOpen] = useState(false);
    const [selectedIndex, setSelectedIndex] = useState(0);
    const anchorRef = useRef(null);

    function handleToggle() {
        setOpen(open => !open);
    }

    function handleClose(event) {
        if (anchorRef.current && anchorRef.current.contains(event.target)) {
            return;
        }

        setOpen(false);
    }

    function handleMenuItemClick(index) {
        setSelectedIndex(index);
        setOpen(false);
    }

    return (
        <div>
            <ButtonGroup variant="contained" color="primary" ref={anchorRef} className={className}>
                <Button>
                    <Icon>{iconName}</Icon>
                    &nbsp;&nbsp;
                    {options[selectedIndex]}
                </Button>
                <Button color="primary" size="small" onClick={handleToggle}>
                    <ArrowDropDown />
                </Button>
            </ButtonGroup>
            <Popper open={open} style={{ zIndex: 1 }} anchorEl={anchorRef.current} transition disablePortal>
                {({ TransitionProps, placement }) => (
                    <Grow {...TransitionProps} style={{ transformOrigin: placement === 'bottom' ? 'center top' : 'center bottom' }}>
                        <Paper>
                            <ClickAwayListener onClickAway={handleClose}>
                                <MenuList>
                                    {options.map((option, index) => (
                                        <MenuItem key={option} selected={index === selectedIndex} onClick={() => handleMenuItemClick(index)}>
                                            {option}
                                        </MenuItem>
                                    ))}
                                </MenuList>
                            </ClickAwayListener>
                        </Paper>
                    </Grow>
                )}
            </Popper>
        </div>
    )
}

export default SplitButton;
