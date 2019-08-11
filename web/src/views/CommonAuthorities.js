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
import { Route } from 'react-router-dom';
import { Edit, New, Search } from './common-authorities';

const CommonAuthorities = ({match, history}) => {
    const [isNewDialogOpen, setNewDialogOpened] = useState(false);
    const onSelectCA = commonAuthorityId => history.push(`${match.url}/${commonAuthorityId}`);
    const onNewCA = () => setNewDialogOpened(true);
    return (
        <div>
            <Route exact path={`${match.url}`} component={() => <Search onSelectCA={onSelectCA} onNewCA={onNewCA} />} />
            <Route exact path={`${match.url}/:commonAuthorityId`} component={Edit} />
            <New open={isNewDialogOpen} setNewDialogOpened={setNewDialogOpened} />
        </div>
    );
}

export default CommonAuthorities;
