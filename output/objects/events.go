/*
Copyright 2019 Konstantin Vasilev (burmuley@gmail.com)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
// WARNING! AUTOMATICALLY GENERATED CONTENT! DON'T CHANGE IT MANUALLY!                                     //
// Source schema can be found at https://github.com/VKCOM/vk-api-schema/blob/master/responses.json         //
// Code generator location: https://github.com/Burmuley/go-vkapi-gen                                       //
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

package objects

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
// `events` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// EventsEventAttach type represents `events_event_attach` API object
type EventsEventAttach struct {
	Address      string                      `json:"address"`       // address of event
	ButtonText   string                      `json:"button_text"`   // text of attach
	Friends      []int                       `json:"friends"`       // array of friends ids
	Id           int                         `json:"id"`            // event ID
	IsFavorite   bool                        `json:"is_favorite"`   // is favorite
	MemberStatus GroupsGroupFullMemberStatus `json:"member_status"` // Current user's member status
	Text         string                      `json:"text"`          // text of attach
	Time         int                         `json:"time"`          // event start time
}
