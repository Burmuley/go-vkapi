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
// `groups` group of responses
/////////////////////////////////////////////////////////////

// GroupsAddcallbackserver type represents `groups_addCallbackServer_response` API response object
type GroupsAddcallbackserver struct {
	ServerId int `json:"server_id"`
}

// GroupsAddlink type represents `groups_addLink_response` API response object
type GroupsAddlink objects.GroupsGroupLink

// GroupsAddAddress type represents `groups_add_address_response` API response object
type GroupsAddAddress objects.GroupsAddress

// GroupsCreate type represents `groups_create_response` API response object
type GroupsCreate objects.GroupsGroup

// GroupsEditaddress type represents `groups_editAddress_response` API response object
type GroupsEditaddress objects.GroupsAddress // Result

// GroupsGetaddresses type represents `groups_getAddresses_response` API response object
type GroupsGetaddresses struct {
	Count int                     `json:"count"` // Total count of addresses
	Items []objects.GroupsAddress `json:"items"`
}

// GroupsGetbanned type represents `groups_getBanned_response` API response object
type GroupsGetbanned struct {
	Count int                        `json:"count"` // Total users number
	Items []objects.GroupsBannedItem `json:"items"`
}

// GroupsGetbyid type represents `groups_getById_response` API response object
type GroupsGetbyid objects.GroupsGroupFull

// GroupsGetcallbackconfirmationcode type represents `groups_getCallbackConfirmationCode_response` API response object
type GroupsGetcallbackconfirmationcode struct {
	Code string `json:"code"` // Confirmation code
}

// GroupsGetcallbackservers type represents `groups_getCallbackServers_response` API response object
type GroupsGetcallbackservers struct {
	Count int                            `json:"count"`
	Items []objects.GroupsCallbackServer `json:"items"`
}

// GroupsGetcallbacksettings type represents `groups_getCallbackSettings_response` API response object
type GroupsGetcallbacksettings objects.GroupsCallbackSettings

// GroupsGetcataloginfoExtended type represents `groups_getCatalogInfo_extended_response` API response object
type GroupsGetcataloginfoExtended struct {
	Categories []objects.GroupsGroupCategoryFull `json:"categories"`
	Enabled    int                               `json:"enabled"` // Information whether catalog is enabled for current user
}

// GroupsGetcataloginfo type represents `groups_getCatalogInfo_response` API response object
type GroupsGetcataloginfo struct {
	Categories []objects.GroupsGroupCategory `json:"categories"`
	Enabled    int                           `json:"enabled"` // Information whether catalog is enabled for current user
}

// GroupsGetcatalog type represents `groups_getCatalog_response` API response object
type GroupsGetcatalog struct {
	Count int                   `json:"count"` // Total communities number
	Items []objects.GroupsGroup `json:"items"`
}

// GroupsGetinvitedusers type represents `groups_getInvitedUsers_response` API response object
type GroupsGetinvitedusers struct {
	Count int                     `json:"count"` // Total communities number
	Items []objects.UsersUserFull `json:"items"`
}

// GroupsGetinvitesExtended type represents `groups_getInvites_extended_response` API response object
type GroupsGetinvitesExtended struct {
	Count    int                               `json:"count"` // Total communities number
	Groups   []objects.GroupsGroupFull         `json:"groups"`
	Items    []objects.GroupsGroupXtrInvitedBy `json:"items"`
	Profiles []objects.UsersUserMin            `json:"profiles"`
}

// GroupsGetinvites type represents `groups_getInvites_response` API response object
type GroupsGetinvites struct {
	Count int                               `json:"count"` // Total communities number
	Items []objects.GroupsGroupXtrInvitedBy `json:"items"`
}

// GroupsGetlongpollserver type represents `groups_getLongPollServer_response` API response object
type GroupsGetlongpollserver objects.GroupsLongPollServer

// GroupsGetlongpollsettings type represents `groups_getLongPollSettings_response` API response object
type GroupsGetlongpollsettings objects.GroupsLongPollSettings

// GroupsGetmembersFields type represents `groups_getMembers_fields_response` API response object
type GroupsGetmembersFields struct {
	Count int                         `json:"count"` // Total members number
	Items []objects.GroupsUserXtrRole `json:"items"`
}

// GroupsGetmembersFilter type represents `groups_getMembers_filter_response` API response object
type GroupsGetmembersFilter struct {
	Count int                        `json:"count"` // Total members number
	Items []objects.GroupsMemberRole `json:"items"`
}

// GroupsGetmembers type represents `groups_getMembers_response` API response object
type GroupsGetmembers struct {
	Count int   `json:"count"` // Total members number
	Items []int `json:"items"`
}

// GroupsGetrequestsFields type represents `groups_getRequests_fields_response` API response object
type GroupsGetrequestsFields struct {
	Count int                     `json:"count"` // Total communities number
	Items []objects.UsersUserFull `json:"items"`
}

// GroupsGetrequests type represents `groups_getRequests_response` API response object
type GroupsGetrequests struct {
	Count int   `json:"count"` // Total communities number
	Items []int `json:"items"`
}

// GroupsGetsettings type represents `groups_getSettings_response` API response object
type GroupsGetsettings objects.GroupsGroupSettings

// GroupsGettokenpermissions type represents `groups_getTokenPermissions_response` API response object
type GroupsGettokenpermissions struct {
	Mask        int                                    `json:"mask"`
	Permissions []objects.GroupsTokenPermissionSetting `json:"permissions"`
}

// GroupsGetExtended type represents `groups_get_extended_response` API response object
type GroupsGetExtended struct {
	Count int                       `json:"count"` // Total communities number
	Items []objects.GroupsGroupFull `json:"items"`
}

// GroupsGet type represents `groups_get_response` API response object
type GroupsGet struct {
	Count int   `json:"count"` // Total communities number
	Items []int `json:"items"`
}

// GroupsIsmemberExtended type represents `groups_isMember_extended_response` API response object
type GroupsIsmemberExtended struct {
	CanInvite  objects.BaseBoolInt `json:"can_invite"` // Information whether user can be invited
	CanRecall  objects.BaseBoolInt `json:"can_recall"` // Information whether user's invite to the group can be recalled
	Invitation objects.BaseBoolInt `json:"invitation"` // Information whether user has been invited to the group
	Member     objects.BaseBoolInt `json:"member"`     // Information whether user is a member of the group
	Request    objects.BaseBoolInt `json:"request"`    // Information whether user has sent request to the group
}

// GroupsIsmember type represents `groups_isMember_response` API response object
type GroupsIsmember objects.BaseBoolInt // Information whether user is a member of the group

// GroupsIsmemberUserIdsExtended type represents `groups_isMember_user_ids_extended_response` API response object
type GroupsIsmemberUserIdsExtended objects.GroupsMemberStatusFull

// GroupsIsmemberUserIds type represents `groups_isMember_user_ids_response` API response object
type GroupsIsmemberUserIds objects.GroupsMemberStatus

// GroupsSearch type represents `groups_search_response` API response object
type GroupsSearch struct {
	Count int                   `json:"count"` // Total communities number
	Items []objects.GroupsGroup `json:"items"`
}
