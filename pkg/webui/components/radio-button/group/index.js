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
import bind from 'autobind-decorator'
import classnames from 'classnames'

import PropTypes from '@ttn-lw/lib/prop-types'

import style from './group.styl'

export const RadioGroupContext = React.createContext()

const findCheckedRadio = children => {
  let value
  let matched = false

  React.Children.forEach(children, radio => {
    if (radio && radio.props && !matched && radio.props.checked) {
      value = radio.props.value
      matched = true
    }
  })

  return value
}

class RadioGroup extends React.Component {
  constructor(props) {
    super(props)

    let value
    if ('value' in props) {
      value = props.value
    } else if ('initialValue' in props) {
      value = props.initialValue
    } else {
      value = findCheckedRadio(props.children)
    }

    this.state = { value }
  }

  static getDerivedStateFromProps(props, state) {
    if ('value' in props) {
      return { value: props.value }
    }

    const value = findCheckedRadio(props.children)
    if (value && value !== state.value) {
      return { value }
    }

    return null
  }

  @bind
  handleRadioChange(event) {
    const { onChange } = this.props
    const { target } = event

    // Retain boolean type if the value was initially provided as boolean.
    const value = typeof this.props.value === 'boolean' ? target.value === 'true' : target.value

    if (!('value' in this.props)) {
      this.setState({ value })
    }

    onChange(value)
  }

  render() {
    const { className, name, disabled, horizontal, children } = this.props
    const { value } = this.state

    const ctx = {
      className: style.groupRadio,
      onChange: this.handleRadioChange,
      disabled,
      value,
      name,
    }

    const cls = classnames(className, style.group, {
      [style.horizontal]: horizontal,
    })

    return (
      <div className={cls}>
        <RadioGroupContext.Provider value={ctx}>{children}</RadioGroupContext.Provider>
      </div>
    )
  }
}

RadioGroup.propTypes = {
  children: PropTypes.oneOfType([PropTypes.arrayOf(PropTypes.node), PropTypes.node]).isRequired,
  className: PropTypes.string,
  disabled: PropTypes.bool,
  horizontal: PropTypes.bool,
  initialValue: PropTypes.string,
  name: PropTypes.string.isRequired,
  onChange: PropTypes.func,
  value: PropTypes.oneOfType([PropTypes.string, PropTypes.bool]),
}

RadioGroup.defaultProps = {
  className: undefined,
  disabled: false,
  initialValue: undefined,
  value: undefined,
  horizontal: false,
  onChange: () => null,
}

export default RadioGroup
