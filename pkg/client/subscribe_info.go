// Copyright 2022, OpenSergo Authors
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

package client

import (
	"reflect"

	"github.com/opensergo/opensergo-go/pkg/transport/subscribe"
)

type SubscribeInfo struct {
	subscribeKey subscribe.SubscribeKey
	subscribers  []subscribe.Subscriber
}

func NewSubscribeInfo(subscribeKey *subscribe.SubscribeKey) *SubscribeInfo {
	return &SubscribeInfo{
		subscribeKey: *subscribeKey,
		subscribers:  make([]subscribe.Subscriber, 0),
	}
}

func (subscribeInfo *SubscribeInfo) AppendSubscriber(subscriber subscribe.Subscriber) {
	if len(subscribeInfo.subscribers) == 0 {
		subscribeInfo.subscribers = append(subscribeInfo.subscribers, subscriber)
		return
	}

	for _, item := range subscribeInfo.subscribers {
		if reflect.ValueOf(item).Interface() != reflect.ValueOf(subscriber).Interface() {
			subscribeInfo.subscribers = append(subscribeInfo.subscribers, subscriber)
		}
	}
}
