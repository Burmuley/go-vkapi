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
// Code generator location: https://gitlab.com/Burmuley/go-vkapi-gen                                       //
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

package responses

import (
	"gitlab.com/Burmuley/go-vkapi/objects"
)

/////////////////////////////////////////////////////////////
// `polls` group of responses
/////////////////////////////////////////////////////////////

// PollsAddvot type represents `polls_addVote_response` API response object
type PollsAddvot *objects.BaseBoolInt // Result

// PollsCreat type represents `polls_create_response` API response object
type PollsCreat *objects.PollsPoll

// PollsDeletevot type represents `polls_deleteVote_response` API response object
type PollsDeletevot *objects.BaseBoolInt // Result

// PollsGetbyid type represents `polls_getById_response` API response object
type PollsGetbyid *objects.PollsPoll

// PollsGetvoter type represents `polls_getVoters_response` API response object
type PollsGetvoter *objects.PollsVoters
