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

package mac_test

import (
	"context"
	"fmt"
	"math"
	"testing"

	"github.com/smartystreets/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/encoding/lorawan"
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	. "go.thethings.network/lorawan-stack/v3/pkg/networkserver/internal"
	. "go.thethings.network/lorawan-stack/v3/pkg/networkserver/mac"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestNeedsNewChannelReq(t *testing.T) {
	for _, tc := range []struct {
		Name        string
		InputDevice *ttnpb.EndDevice
		Needs       bool
	}{
		{
			Name:        "no MAC state",
			InputDevice: &ttnpb.EndDevice{},
		},
		{
			Name: "current(channels:[(123,1-5),nil,(128,2-4)]),desired(channels:[(123,1-5),nil,(128,2-4)])",
			InputDevice: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{
					CurrentParameters: ttnpb.MACParameters{
						Channels: []*ttnpb.MACParameters_Channel{
							{
								UplinkFrequency:  123,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_5,
							},
							nil,
							{
								UplinkFrequency:  128,
								MinDataRateIndex: ttnpb.DATA_RATE_2,
								MaxDataRateIndex: ttnpb.DATA_RATE_4,
							},
						},
					},
					DesiredParameters: ttnpb.MACParameters{
						Channels: []*ttnpb.MACParameters_Channel{
							{
								UplinkFrequency:  123,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_5,
							},
							nil,
							{
								UplinkFrequency:  128,
								MinDataRateIndex: ttnpb.DATA_RATE_2,
								MaxDataRateIndex: ttnpb.DATA_RATE_4,
							},
						},
					},
				},
			},
		},
		{
			Name: "current(channels:[(123,1-5),(124,1-3),(128,2-4)]),desired(channels:[(123,1-5),(124,1-3),(128,2-4)])",
			InputDevice: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{
					CurrentParameters: ttnpb.MACParameters{
						Channels: []*ttnpb.MACParameters_Channel{
							{
								UplinkFrequency:  123,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_5,
							},
							{
								UplinkFrequency:  124,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_3,
							},
							{
								UplinkFrequency:  128,
								MinDataRateIndex: ttnpb.DATA_RATE_2,
								MaxDataRateIndex: ttnpb.DATA_RATE_4,
							},
						},
					},
					DesiredParameters: ttnpb.MACParameters{
						Channels: []*ttnpb.MACParameters_Channel{
							{
								UplinkFrequency:  123,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_5,
							},
							{
								UplinkFrequency:  124,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_3,
							},
							{
								UplinkFrequency:  128,
								MinDataRateIndex: ttnpb.DATA_RATE_2,
								MaxDataRateIndex: ttnpb.DATA_RATE_4,
							},
						},
					},
				},
			},
		},
		{
			Name: "current(channels:[(123,1-5),(124,1-3),(128,2-4)]),desired(channels:[(123,1-5),nil,(128,2-4)])",
			InputDevice: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{
					CurrentParameters: ttnpb.MACParameters{
						Channels: []*ttnpb.MACParameters_Channel{
							{
								UplinkFrequency:  123,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_5,
							},
							{
								UplinkFrequency:  124,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_3,
							},
							{
								UplinkFrequency:  128,
								MinDataRateIndex: ttnpb.DATA_RATE_2,
								MaxDataRateIndex: ttnpb.DATA_RATE_4,
							},
						},
					},
					DesiredParameters: ttnpb.MACParameters{
						Channels: []*ttnpb.MACParameters_Channel{
							{
								UplinkFrequency:  123,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_5,
							},
							nil,
							{
								UplinkFrequency:  128,
								MinDataRateIndex: ttnpb.DATA_RATE_2,
								MaxDataRateIndex: ttnpb.DATA_RATE_4,
							},
						},
					},
				},
			},
			Needs: true,
		},
		{
			Name: "current(channels:[(123,1-5),(124,1-3),(128,2-4)]),desired(channels:[(123,1-5),(124,1-3)])",
			InputDevice: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{
					CurrentParameters: ttnpb.MACParameters{
						Channels: []*ttnpb.MACParameters_Channel{
							{
								UplinkFrequency:  123,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_5,
							},
							{
								UplinkFrequency:  124,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_3,
							},
							{
								UplinkFrequency:  128,
								MinDataRateIndex: ttnpb.DATA_RATE_2,
								MaxDataRateIndex: ttnpb.DATA_RATE_4,
							},
						},
					},
					DesiredParameters: ttnpb.MACParameters{
						Channels: []*ttnpb.MACParameters_Channel{
							{
								UplinkFrequency:  123,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_5,
							},
							{
								UplinkFrequency:  124,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_3,
							},
						},
					},
				},
			},
			Needs: true,
		},
		{
			Name: "current(channels:[(123,1-5),(124,1-3),(128,2-4)]),desired(channels:[(123,1-5),(124,1-3),(128,2-4),(129,2-3)])",
			InputDevice: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{
					CurrentParameters: ttnpb.MACParameters{
						Channels: []*ttnpb.MACParameters_Channel{
							{
								UplinkFrequency:  123,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_5,
							},
							{
								UplinkFrequency:  124,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_3,
							},
							{
								UplinkFrequency:  128,
								MinDataRateIndex: ttnpb.DATA_RATE_2,
								MaxDataRateIndex: ttnpb.DATA_RATE_4,
							},
						},
					},
					DesiredParameters: ttnpb.MACParameters{
						Channels: []*ttnpb.MACParameters_Channel{
							{
								UplinkFrequency:  123,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_5,
							},
							{
								UplinkFrequency:  124,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_3,
							},
							{
								UplinkFrequency:  128,
								MinDataRateIndex: ttnpb.DATA_RATE_2,
								MaxDataRateIndex: ttnpb.DATA_RATE_4,
							},
							{
								UplinkFrequency:  129,
								MinDataRateIndex: ttnpb.DATA_RATE_2,
								MaxDataRateIndex: ttnpb.DATA_RATE_3,
							},
						},
					},
				},
			},
			Needs: true,
		},
		{
			Name: "current(channels:[(123,1-5),(124,1-3),(128,2-4)]),desired(channels:[(123,1-5),(124,1-3),(128,2-3)])",
			InputDevice: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{
					CurrentParameters: ttnpb.MACParameters{
						Channels: []*ttnpb.MACParameters_Channel{
							{
								UplinkFrequency:  123,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_5,
							},
							{
								UplinkFrequency:  124,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_3,
							},
							{
								UplinkFrequency:  128,
								MinDataRateIndex: ttnpb.DATA_RATE_2,
								MaxDataRateIndex: ttnpb.DATA_RATE_4,
							},
						},
					},
					DesiredParameters: ttnpb.MACParameters{
						Channels: []*ttnpb.MACParameters_Channel{
							{
								UplinkFrequency:  123,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_5,
							},
							{
								UplinkFrequency:  124,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_3,
							},
							{
								UplinkFrequency:  128,
								MinDataRateIndex: ttnpb.DATA_RATE_2,
								MaxDataRateIndex: ttnpb.DATA_RATE_3,
							},
						},
					},
				},
			},
			Needs: true,
		},
		{
			Name: "current(channels:[(123,1-5),(124,1-3),(128,2-4)]),desired(channels:[(123,1-5),(124,1-3),(127,2-4)])",
			InputDevice: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{
					CurrentParameters: ttnpb.MACParameters{
						Channels: []*ttnpb.MACParameters_Channel{
							{
								UplinkFrequency:  123,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_5,
							},
							{
								UplinkFrequency:  124,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_3,
							},
							{
								UplinkFrequency:  128,
								MinDataRateIndex: ttnpb.DATA_RATE_2,
								MaxDataRateIndex: ttnpb.DATA_RATE_4,
							},
						},
					},
					DesiredParameters: ttnpb.MACParameters{
						Channels: []*ttnpb.MACParameters_Channel{
							{
								UplinkFrequency:  123,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_5,
							},
							{
								UplinkFrequency:  124,
								MinDataRateIndex: ttnpb.DATA_RATE_1,
								MaxDataRateIndex: ttnpb.DATA_RATE_3,
							},
							{
								UplinkFrequency:  127,
								MinDataRateIndex: ttnpb.DATA_RATE_2,
								MaxDataRateIndex: ttnpb.DATA_RATE_4,
							},
						},
					},
				},
			},
			Needs: true,
		},
	} {
		tc := tc
		test.RunSubtest(t, test.SubtestConfig{
			Name:     tc.Name,
			Parallel: true,
			Func: func(ctx context.Context, t *testing.T, a *assertions.Assertion) {
				dev := CopyEndDevice(tc.InputDevice)
				res := DeviceNeedsNewChannelReq(dev)
				if tc.Needs {
					a.So(res, should.BeTrue)
				} else {
					a.So(res, should.BeFalse)
				}
				a.So(dev, should.Resemble, tc.InputDevice)
			},
		})
	}
}

func TestEnqueueNewChannelReq(t *testing.T) {
	for _, tc := range []struct {
		Name                                 string
		CurrentParameters, DesiredParameters ttnpb.MACParameters
		ExpectedRequests                     []*ttnpb.MACCommand_NewChannelReq
	}{
		{
			Name: "no NewChannelReq necessary",
			CurrentParameters: ttnpb.MACParameters{
				Channels: []*ttnpb.MACParameters_Channel{
					{
						UplinkFrequency:  123,
						MinDataRateIndex: ttnpb.DATA_RATE_1,
						MaxDataRateIndex: ttnpb.DATA_RATE_5,
					},
					{
						UplinkFrequency:  124,
						MinDataRateIndex: ttnpb.DATA_RATE_1,
						MaxDataRateIndex: ttnpb.DATA_RATE_3,
					},
					{
						UplinkFrequency:  128,
						MinDataRateIndex: ttnpb.DATA_RATE_2,
						MaxDataRateIndex: ttnpb.DATA_RATE_4,
					},
				},
			},
			DesiredParameters: ttnpb.MACParameters{
				Channels: []*ttnpb.MACParameters_Channel{
					{
						UplinkFrequency:  123,
						MinDataRateIndex: ttnpb.DATA_RATE_1,
						MaxDataRateIndex: ttnpb.DATA_RATE_5,
					},
					{
						UplinkFrequency:  124,
						MinDataRateIndex: ttnpb.DATA_RATE_1,
						MaxDataRateIndex: ttnpb.DATA_RATE_3,
					},
					{
						UplinkFrequency:  128,
						MinDataRateIndex: ttnpb.DATA_RATE_2,
						MaxDataRateIndex: ttnpb.DATA_RATE_4,
					},
				},
			},
		},
		{
			Name: "4 NewChannelReq necessary",
			CurrentParameters: ttnpb.MACParameters{
				Channels: []*ttnpb.MACParameters_Channel{
					{
						UplinkFrequency:  124,
						MinDataRateIndex: ttnpb.DATA_RATE_1,
						MaxDataRateIndex: ttnpb.DATA_RATE_3,
					},
					nil,
					{
						UplinkFrequency:  123,
						MinDataRateIndex: ttnpb.DATA_RATE_1,
						MaxDataRateIndex: ttnpb.DATA_RATE_5,
					},
					{
						UplinkFrequency:  129,
						MinDataRateIndex: ttnpb.DATA_RATE_2,
						MaxDataRateIndex: ttnpb.DATA_RATE_4,
					},
					{
						UplinkFrequency:  130,
						MinDataRateIndex: ttnpb.DATA_RATE_2,
						MaxDataRateIndex: ttnpb.DATA_RATE_5,
					},
				},
			},
			DesiredParameters: ttnpb.MACParameters{
				Channels: []*ttnpb.MACParameters_Channel{
					nil,
					{
						UplinkFrequency:  128,
						MinDataRateIndex: ttnpb.DATA_RATE_2,
						MaxDataRateIndex: ttnpb.DATA_RATE_4,
					},
					{
						UplinkFrequency:  123,
						MinDataRateIndex: ttnpb.DATA_RATE_1,
						MaxDataRateIndex: ttnpb.DATA_RATE_5,
					},
					{
						UplinkFrequency:  130,
						MinDataRateIndex: ttnpb.DATA_RATE_2,
						MaxDataRateIndex: ttnpb.DATA_RATE_5,
					},
				},
			},
			ExpectedRequests: []*ttnpb.MACCommand_NewChannelReq{
				{},
				{
					ChannelIndex:     1,
					Frequency:        128,
					MinDataRateIndex: ttnpb.DATA_RATE_2,
					MaxDataRateIndex: ttnpb.DATA_RATE_4,
				},
				{
					ChannelIndex:     3,
					Frequency:        130,
					MinDataRateIndex: ttnpb.DATA_RATE_2,
					MaxDataRateIndex: ttnpb.DATA_RATE_5,
				},
				{
					ChannelIndex: 4,
				},
			},
		},
	} {
		tc := tc
		test.RunSubtest(t, test.SubtestConfig{
			Name: tc.Name,
			Func: func(ctx context.Context, t *testing.T, a *assertions.Assertion) {
				downlinkLength := 1 + lorawan.DefaultMACCommands[ttnpb.CID_NEW_CHANNEL].DownlinkLength
				uplinkLength := 1 + lorawan.DefaultMACCommands[ttnpb.CID_NEW_CHANNEL].UplinkLength

				type TestConf struct {
					MaxDownlinkLength, MaxUplinkLength uint16
					ExpectedCount                      int
				}
				confs := []TestConf{
					{},
					{
						MaxUplinkLength: math.MaxUint16,
					},
					{
						MaxDownlinkLength: math.MaxUint16,
					},
					{
						MaxDownlinkLength: math.MaxUint16,
						MaxUplinkLength:   math.MaxUint16,
						ExpectedCount:     len(tc.ExpectedRequests),
					},
				}
				for i := range tc.ExpectedRequests {
					for j := 0; j <= i; j++ {
						confs = append(confs, TestConf{
							MaxDownlinkLength: uint16(i+1) * downlinkLength,
							MaxUplinkLength:   uint16(j+1) * uplinkLength,
							ExpectedCount:     j + 1,
						})
					}
				}

				for _, conf := range confs {
					for _, pendingReqs := range [][]*ttnpb.MACCommand{
						nil,
						{
							{},
						},
					} {
						conf := conf
						pendingReqs := pendingReqs
						test.RunSubtest(t, test.SubtestConfig{
							Name:     fmt.Sprintf("max_downlink_len:%d,max_uplink_len:%d,pending_requests:%d", conf.MaxDownlinkLength, conf.MaxUplinkLength, len(pendingReqs)),
							Parallel: true,
							Func: func(ctx context.Context, t *testing.T, a *assertions.Assertion) {
								dev := &ttnpb.EndDevice{
									MACState: &ttnpb.MACState{
										CurrentParameters: tc.CurrentParameters,
										DesiredParameters: tc.DesiredParameters,
										PendingRequests:   pendingReqs,
									},
								}
								reqs := tc.ExpectedRequests[:conf.ExpectedCount]
								expectedDev := CopyEndDevice(dev)
								var expectedEvs events.Builders
								for _, req := range reqs {
									expectedDev.MACState.PendingRequests = append(expectedDev.MACState.PendingRequests, req.MACCommand())
									expectedEvs = append(expectedEvs, EvtEnqueueNewChannelRequest.With(events.WithData(req)))
								}

								st := EnqueueNewChannelReq(ctx, dev, conf.MaxDownlinkLength, conf.MaxUplinkLength)
								a.So(dev, should.Resemble, expectedDev)
								a.So(st.QueuedEvents, should.ResembleEventBuilders, expectedEvs)
								a.So(st, should.Resemble, EnqueueState{
									MaxDownLen:   conf.MaxDownlinkLength - uint16(conf.ExpectedCount)*downlinkLength,
									MaxUpLen:     conf.MaxUplinkLength - uint16(conf.ExpectedCount)*uplinkLength,
									Ok:           len(tc.ExpectedRequests) == conf.ExpectedCount,
									QueuedEvents: st.QueuedEvents,
								})
							},
						})
					}
				}
			},
		})
	}
}

func TestHandleNewChannelAns(t *testing.T) {
	for _, tc := range []struct {
		Name             string
		Device, Expected *ttnpb.EndDevice
		Payload          *ttnpb.MACCommand_NewChannelAns
		Events           events.Builders
		Error            error
	}{
		{
			Name: "nil payload",
			Device: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{},
			},
			Expected: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{},
			},
			Payload: nil,
			Error:   ErrNoPayload,
		},
		{
			Name: "no request",
			Device: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{},
			},
			Expected: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{},
			},
			Payload: &ttnpb.MACCommand_NewChannelAns{
				FrequencyAck: true,
				DataRateAck:  true,
			},
			Events: events.Builders{
				EvtReceiveNewChannelAccept.With(events.WithData(&ttnpb.MACCommand_NewChannelAns{
					FrequencyAck: true,
					DataRateAck:  true,
				})),
			},
			Error: ErrRequestNotFound,
		},
		{
			Name: "frequency nack/data rate ack/no rejections",
			Device: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{
					PendingRequests: []*ttnpb.MACCommand{
						(&ttnpb.MACCommand_NewChannelReq{
							ChannelIndex:     4,
							Frequency:        42,
							MinDataRateIndex: 2,
							MaxDataRateIndex: 3,
						}).MACCommand(),
					},
				},
			},
			Expected: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{
					PendingRequests:     []*ttnpb.MACCommand{},
					RejectedFrequencies: []uint64{42},
				},
			},
			Payload: &ttnpb.MACCommand_NewChannelAns{
				DataRateAck: true,
			},
			Events: events.Builders{
				EvtReceiveNewChannelReject.With(events.WithData(&ttnpb.MACCommand_NewChannelAns{
					DataRateAck: true,
				})),
			},
		},
		{
			Name: "frequency nack/data rate nack/rejected frequencies:(1,2,100)",
			Device: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{
					PendingRequests: []*ttnpb.MACCommand{
						(&ttnpb.MACCommand_NewChannelReq{
							ChannelIndex:     4,
							Frequency:        42,
							MinDataRateIndex: 2,
							MaxDataRateIndex: 3,
						}).MACCommand(),
					},
					RejectedFrequencies: []uint64{1, 2, 100},
				},
			},
			Expected: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{
					PendingRequests:     []*ttnpb.MACCommand{},
					RejectedFrequencies: []uint64{1, 2, 42, 100},
				},
			},
			Payload: &ttnpb.MACCommand_NewChannelAns{},
			Events: events.Builders{
				EvtReceiveNewChannelReject.With(events.WithData(&ttnpb.MACCommand_NewChannelAns{})),
			},
		},
		{
			Name: "both ack",
			Device: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{
					PendingRequests: []*ttnpb.MACCommand{
						(&ttnpb.MACCommand_NewChannelReq{
							ChannelIndex:     4,
							Frequency:        42,
							MinDataRateIndex: 2,
							MaxDataRateIndex: 3,
						}).MACCommand(),
					},
				},
			},
			Expected: &ttnpb.EndDevice{
				MACState: &ttnpb.MACState{
					CurrentParameters: ttnpb.MACParameters{
						Channels: []*ttnpb.MACParameters_Channel{
							nil,
							nil,
							nil,
							nil,
							{
								DownlinkFrequency: 42,
								UplinkFrequency:   42,
								MinDataRateIndex:  2,
								MaxDataRateIndex:  3,
								EnableUplink:      true,
							},
						},
					},
					PendingRequests: []*ttnpb.MACCommand{},
				},
			},
			Payload: &ttnpb.MACCommand_NewChannelAns{
				FrequencyAck: true,
				DataRateAck:  true,
			},
			Events: events.Builders{
				EvtReceiveNewChannelAccept.With(events.WithData(&ttnpb.MACCommand_NewChannelAns{
					FrequencyAck: true,
					DataRateAck:  true,
				})),
			},
		},
	} {
		tc := tc
		test.RunSubtest(t, test.SubtestConfig{
			Name:     tc.Name,
			Parallel: true,
			Func: func(ctx context.Context, t *testing.T, a *assertions.Assertion) {
				dev := CopyEndDevice(tc.Device)

				evs, err := HandleNewChannelAns(ctx, dev, tc.Payload)
				if tc.Error != nil && !a.So(err, should.EqualErrorOrDefinition, tc.Error) ||
					tc.Error == nil && !a.So(err, should.BeNil) {
					t.FailNow()
				}
				a.So(dev, should.Resemble, tc.Expected)
				a.So(evs, should.ResembleEventBuilders, tc.Events)
			},
		})
	}
}
