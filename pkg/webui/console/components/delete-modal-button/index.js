// Copyright © 2021 The Things Network Foundation, The Things Industries B.V.
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
import { defineMessages } from 'react-intl'

import Input from '@ttn-lw/components/input'
import ModalButton from '@ttn-lw/components/button/modal-button'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './delete-modal-button.styl'

const m = defineMessages({
  modalDefaultWarning:
    'Are you sure you want to delete "{entityName}"? This action cannot be undone and it will not be possible to reuse the entity ID.',
  modalPurgeWarning:
    'Are you sure you want to delete "{entityName}"? This action cannot be undone, but it will be possible to reuse the entity ID.',
  confirmMessage:
    'This will permanently delete the "{entityName}" and all associated entities, as well as all collaborator associations. Please type <strong>{entityId}</strong> to confirm.',
})

const DeleteModalButton = props => {
  const { entityId, entityName, onApprove, onCancel, shouldConfirm, shouldPurge, message } = props
  const name = entityName ? entityName : entityId

  const [confirmId, setConfirmId] = React.useState('')

  return (
    <ModalButton
      type="button"
      icon="delete"
      danger
      naked
      onApprove={onApprove}
      onCancel={onCancel}
      message={message}
      modalData={{
        approveButtonProps: { disabled: shouldConfirm && confirmId !== entityId },
        children: (
          <div>
            <Message
              className={style.message}
              content={shouldPurge ? m.modalPurgeWarning : m.modalDefaultWarning}
              values={{ entityName: name }}
              component="p"
            />
            {shouldConfirm && (
              <>
                <Message
                  content={m.confirmMessage}
                  values={{ entityId, entityName: name, strong: id => <strong>{id}</strong> }}
                  component="p"
                />
                <Input value={confirmId} onChange={setConfirmId} />
              </>
            )}
          </div>
        ),
      }}
    />
  )
}

DeleteModalButton.propTypes = {
  entityId: PropTypes.string.isRequired,
  entityName: PropTypes.string,
  message: PropTypes.message.isRequired,
  onApprove: PropTypes.func,
  onCancel: PropTypes.func,
  shouldConfirm: PropTypes.bool,
  shouldPurge: PropTypes.bool,
}

DeleteModalButton.defaultProps = {
  entityName: undefined,
  onApprove: undefined,
  onCancel: undefined,
  shouldConfirm: false,
  shouldPurge: false,
}

export default DeleteModalButton
