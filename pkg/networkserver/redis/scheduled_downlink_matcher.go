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

package redis

import (
	"context"
	"regexp"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/networkserver/internal/time"
	ttnredis "go.thethings.network/lorawan-stack/v3/pkg/redis"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// Duration before metadata for scheduled downlinks expire. TxAcknowledgements received
// after this TTL will be considered invalid.
const ttl = time.Minute

// ScheduledDownlinkMatcher is an implementation of networkserver.ScheduledDownlinkMatcher.
type ScheduledDownlinkMatcher struct {
	Redis *ttnredis.Client
}

func (m *ScheduledDownlinkMatcher) cidToKey(cid string) string {
	return m.Redis.Key("cid", cid)
}

var parseDownlinkIDRegex = regexp.MustCompile(`^ns:downlink:([0-9a-zA-Z]+)$`)

// parseDownlinkCorrelationID matches the ns:downlink:XXXXXXX correlation ID.
func parseDownlinkCorrelationID(cids []string) (string, bool) {
	for _, cid := range cids {
		matches := parseDownlinkIDRegex.FindStringSubmatch(cid)
		if len(matches) != 2 || matches[1] == "" {
			continue
		}
		return matches[1], true
	}
	return "", false
}

var errMissingDownlinkCorrelationID = errors.DefineInternal("missing_downlink_correlation_id", "missing identifier correlation ID on downlink message")

func (m *ScheduledDownlinkMatcher) Add(ctx context.Context, down *ttnpb.DownlinkMessage) error {
	cid, ok := parseDownlinkCorrelationID(down.GetCorrelationIDs())
	if !ok {
		return errMissingDownlinkCorrelationID.New()
	}
	_, err := ttnredis.SetProto(ctx, m.Redis, m.cidToKey(cid), down, ttl)
	return err
}

var errNoDownlinkMatch = errors.DefineNotFound("no_downlink_match", "no downlink matched")

func (m *ScheduledDownlinkMatcher) Match(ctx context.Context, ack *ttnpb.TxAcknowledgment) (*ttnpb.DownlinkMessage, error) {
	cid, ok := parseDownlinkCorrelationID(ack.GetDownlinkMessage().GetCorrelationIDs())
	if !ok {
		return nil, errMissingDownlinkCorrelationID.New()
	}
	pb := &ttnpb.DownlinkMessage{}
	if err := ttnredis.GetProto(ctx, m.Redis, m.cidToKey(cid)).ScanProto(pb); err != nil {
		return nil, errNoDownlinkMatch.WithCause(err)
	}
	return pb, nil
}
