// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react'
import { Switch, Route } from 'react-router-dom'

import { withBreadcrumb } from '@ttn-lw/components/breadcrumbs/context'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'

import withFeatureRequirement from '@console/lib/components/with-feature-requirement'

import Gateway from '@console/views/gateway'
import GatewayAdd from '@console/views/gateway-add'
import GatewaysList from '@console/views/gateways-list'

import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'
import { pathId as pathIdRegexp } from '@ttn-lw/lib/regexp'

import { mayViewGateways } from '@console/lib/feature-checks'

@withFeatureRequirement(mayViewGateways, { redirect: '/' })
@withBreadcrumb('gateways', () => <Breadcrumb path="/gateways" content={sharedMessages.gateways} />)
export default class Gateways extends React.Component {
  static propTypes = {
    match: PropTypes.match.isRequired,
  }

  render() {
    const { path } = this.props.match

    return (
      <Switch>
        <Route exact path={`${path}`} component={GatewaysList} />
        <Route exact path={`${path}/add`} component={GatewayAdd} />
        <Route path={`${path}/:gtwId${pathIdRegexp}`} component={Gateway} sensitive />
      </Switch>
    )
  }
}
