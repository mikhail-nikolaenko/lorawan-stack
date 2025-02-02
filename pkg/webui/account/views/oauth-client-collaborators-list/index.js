// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
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
import { Container, Row, Col } from 'react-grid-system'
import { connect } from 'react-redux'

import PAGE_SIZES from '@ttn-lw/constants/page-sizes'

import { useBreadcrumbs } from '@ttn-lw/components/breadcrumbs/context'
import Breadcrumb from '@ttn-lw/components/breadcrumbs/breadcrumb'

import IntlHelmet from '@ttn-lw/lib/components/intl-helmet'

import CollaboratorsTable from '@account/containers/collaborators-table'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'

import { selectSelectedClientId } from '@account/store/selectors/clients'

const OAuthClientCollaboratorsList = props => {
  const { clientId } = props

  useBreadcrumbs(
    'clients.single.collaborators',
    <Breadcrumb
      path={`/oauth-clients/${clientId}/collaborators`}
      content={sharedMessages.collaborators}
    />,
  )

  return (
    <Container>
      <Row>
        <IntlHelmet title={sharedMessages.collaborators} />
        <Col>
          <CollaboratorsTable pageSize={PAGE_SIZES.REGULAR} />
        </Col>
      </Row>
    </Container>
  )
}

OAuthClientCollaboratorsList.propTypes = {
  clientId: PropTypes.string.isRequired,
}

export default connect(state => ({
  clientId: selectSelectedClientId(state),
}))(OAuthClientCollaboratorsList)
