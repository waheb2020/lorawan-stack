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

import React, { useEffect } from 'react'
import { defineMessages } from 'react-intl'

import Field from '@ttn-lw/components/form/field'

import PropTypes from '@ttn-lw/lib/prop-types'

import connect from './connect'
import JoinEUIPrefixesInput from './join-eui-prefixes-input'

const m = defineMessages({
  prefixesFetchingFailure: 'Prefixes unavailable',
})

const JoinEUIPrefixesField = function({ error, getPrefixes, ...rest }) {
  useEffect(() => {
    getPrefixes()
  }, [getPrefixes])

  return (
    <Field
      {...rest}
      component={JoinEUIPrefixesInput}
      warning={Boolean(error) ? m.prefixesFetchingFailure : undefined}
    />
  )
}

const { component, ...fieldPropTypes } = Field.propTypes
const { id, ...inputPropTypes } = JoinEUIPrefixesInput.propTypes

JoinEUIPrefixesField.propTypes = {
  ...inputPropTypes,
  ...fieldPropTypes,
  fetching: PropTypes.bool.isRequired,
  getPrefixes: PropTypes.func.isRequired,
  prefixes: PropTypes.arrayOf(
    PropTypes.shape({
      prefix: PropTypes.string,
      length: PropTypes.number,
    }),
  ),
  showPrefixes: PropTypes.bool,
}

JoinEUIPrefixesField.defaultProps = {
  prefixes: [],
  showPrefixes: true,
}

export default connect(JoinEUIPrefixesField)
