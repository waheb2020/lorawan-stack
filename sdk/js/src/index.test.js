// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
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

import { Applications, Application } from './entity/applications'
import { Devices, Device } from './entity/devices'
import TtnLw from '.'

const mockApplicationData = {
  ids: {
    application_id: 'test',
  },
  created_at: '2018-08-29T14:00:20.793Z',
  updated_at: '2018-08-29T14:00:20.793Z',
  name: 'string',
  description: 'string',
  attributes: {
    additionalProp1: 'string',
    additionalProp2: 'string',
    additionalProp3: 'string',
  },
  contact_info: [
    {
      contact_type: 'CONTACT_TYPE_OTHER',
      contact_method: 'CONTACT_METHOD_OTHER',
      value: 'string',
      'public': true,
      validated_at: '2018-08-29T14:00:20.793Z',
    },
  ],
  default_formatters: {
    up_formatter: 'FORMATTER_JAVASCRIPT',
    up_formatter_parameter: 'function Up(Bytes, Port) {}',
    down_formatter: 'FORMATTER_JAVASCRIPT',
    down_formatter_parameter: 'function Down(Bytes, Port) {}',
  },
}

const mockDeviceData = {
  ids: {
    device_id: 'test-device',
    application_ids: {
      application_id: 'test',
    },
    dev_eui: 'string',
    join_eui: 'string',
    dev_addr: 'string',
  },
}

jest.mock('./api', function () {
  return jest.fn().mockImplementation(function () {
    return {
      GetApplication: jest.fn().mockResolvedValue(mockApplicationData),
      ListApplications: jest.fn().mockResolvedValue([ mockApplicationData ]),
      GetDevice: jest.fn().mockResolvedValue(mockDeviceData),
    }
  })
})

describe('SDK class', function () {
  const token = 'faketoken'
  const ttn = new TtnLw( token, {
    connectionType: 'http',
    baseURL: 'http://localhost:1885/api/v3',
  })

  test('instance instanciates successfully', async function () {
    expect(ttn).toBeDefined()
    expect(ttn).toBeInstanceOf(TtnLw)
    expect(ttn.Applications).toBeInstanceOf(Applications)
  })

  test('retrieves application instance correctly', async function () {
    const app = await ttn.Applications.getById('test')
    expect(app).toBeDefined()
    expect(app).toBeInstanceOf(Application)
  })

  test('retrieves device via app instance correctly', async function () {
    const app = await ttn.Applications.getById('test')
    const device = await app.Devices.getById('test-device')

    expect(app.Devices).toBeInstanceOf(Devices)
    expect(device).toBeDefined()
    expect(device).toBeInstanceOf(Device)
  })

  test('retrieves device via shorthand correctly', async function () {
    const device = await ttn.Applications.withId('test').getDevice('test-device')

    expect(device).toBeDefined()
    expect(device).toBeInstanceOf(Device)
  })

})
