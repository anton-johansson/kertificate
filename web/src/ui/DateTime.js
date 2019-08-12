import React from 'react';
import dateFormat from 'dateformat';

const DateTime = ({value}) => <span>{dateFormat(new Date(value), "dddd, mmmm d, yyyy, HH:MM:ss")}</span>

export default DateTime;
