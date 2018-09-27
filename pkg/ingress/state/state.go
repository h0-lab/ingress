//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2018] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package state

import (
	"github.com/lastbackend/lastbackend/pkg/distribution/types"
)

const logLevel = 3

type State struct {
	ingress   *IngressState
	networks  *NetworkState
	routes    *RouteState
	endpoints *EndpointState
}

type IngressState struct {
	Info   types.IngressInfo
	Status types.IngressStatus
}

func (s *State) Ingress() *IngressState {
	return s.ingress
}

func (s *State) Networks() *NetworkState {
	return s.networks
}

func (s *State) Routes() *RouteState {
	return s.routes
}

func (s *State) Endpoints() *EndpointState {
	return s.endpoints
}

func New() *State {

	state := State{
		ingress: new(IngressState),
		networks: &NetworkState{
			subnets: make(map[string]types.NetworkState, 0),
		},
		routes: &RouteState{
			routes: make(map[string]*types.RouteManifest, 0),
		},
		endpoints: &EndpointState{
			endpoints: make(map[string]*types.EndpointState, 0),
		},
	}

	return &state
}
