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

package responses

import (
	"github.com/Burmuley/go-vkapi/objects"
)

/////////////////////////////////////////////////////////////
// `polls` group of responses
/////////////////////////////////////////////////////////////

// PollsAddvote type represents `polls_addVote_response` API response object
type PollsAddvote objects.BaseBoolInt // Result

// PollsCreate type represents `polls_create_response` API response object
type PollsCreate objects.PollsPoll

// PollsDeletevote type represents `polls_deleteVote_response` API response object
type PollsDeletevote objects.BaseBoolInt // Result

// PollsGetbyid type represents `polls_getById_response` API response object
type PollsGetbyid objects.PollsPoll

// PollsGetvoters type represents `polls_getVoters_response` API response object
type PollsGetvoters objects.PollsVoters
