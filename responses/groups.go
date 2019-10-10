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

import "gitlab.com/Burmuley/go-vkapi/objects"

/////////////////////////////////////////////////////////////
// `groups` group of responses
/////////////////////////////////////////////////////////////

// GroupsAddCallbackServer type represents `groups_addCallbackServer_response` API response object
type GroupsAddCallbackServer struct {
	ServerId int `json:"server_id"`
}

// GroupsAddLink type represents `groups_addLink_response` API response object
type GroupsAddLink objects.GroupsGroupLink

// GroupsAddAddress type represents `groups_add_address_response` API response object
type GroupsAddAddress objects.GroupsAddress

// GroupsCreate type represents `groups_create_response` API response object
type GroupsCreate objects.GroupsGroup

// GroupsEditAddress type represents `groups_editAddress_response` API response object
type GroupsEditAddress objects.GroupsAddress // Result

// GroupsGetAddresses type represents `groups_getAddresses_response` API response object
type GroupsGetAddresses struct {
	Count int                     `json:"count"` // Total count of addresses
	Items []objects.GroupsAddress `json:"items"`
}

// GroupsGetBanned type represents `groups_getBanned_response` API response object
type GroupsGetBanned struct {
	Count int                        `json:"count"` // Total users number
	Items []objects.GroupsBannedItem `json:"items"`
}

// GroupsGetById type represents `groups_getById_response` API response object
type GroupsGetById objects.GroupsGroupFull

// GroupsGetCallbackConfirmationCode type represents `groups_getCallbackConfirmationCode_response` API response object
type GroupsGetCallbackConfirmationCode struct {
	Code string `json:"code"` // Confirmation code
}

// GroupsGetCallbackServers type represents `groups_getCallbackServers_response` API response object
type GroupsGetCallbackServers struct {
	Count int                            `json:"count"`
	Items []objects.GroupsCallbackServer `json:"items"`
}

// GroupsGetCallbackSettings type represents `groups_getCallbackSettings_response` API response object
type GroupsGetCallbackSettings objects.GroupsCallbackSettings

// GroupsGetCatalogInfoExtended type represents `groups_getCatalogInfo_extended_response` API response object
type GroupsGetCatalogInfoExtended struct {
	Categories []objects.GroupsGroupCategoryFull `json:"categories"`
	Enabled    int                               `json:"enabled"` // Information whether catalog is enabled for current user
}

// GroupsGetCatalogInfo type represents `groups_getCatalogInfo_response` API response object
type GroupsGetCatalogInfo struct {
	Categories []objects.GroupsGroupCategory `json:"categories"`
	Enabled    int                           `json:"enabled"` // Information whether catalog is enabled for current user
}

// GroupsGetCatalog type represents `groups_getCatalog_response` API response object
type GroupsGetCatalog struct {
	Count int                   `json:"count"` // Total communities number
	Items []objects.GroupsGroup `json:"items"`
}

// GroupsGetInvitedUsers type represents `groups_getInvitedUsers_response` API response object
type GroupsGetInvitedUsers struct {
	Count int                     `json:"count"` // Total communities number
	Items []objects.UsersUserFull `json:"items"`
}

// GroupsGetInvitesExtended type represents `groups_getInvites_extended_response` API response object
type GroupsGetInvitesExtended struct {
	Count    int                               `json:"count"` // Total communities number
	Groups   []objects.GroupsGroupFull         `json:"groups"`
	Items    []objects.GroupsGroupXtrInvitedBy `json:"items"`
	Profiles []objects.UsersUserMin            `json:"profiles"`
}

// GroupsGetInvites type represents `groups_getInvites_response` API response object
type GroupsGetInvites struct {
	Count int                               `json:"count"` // Total communities number
	Items []objects.GroupsGroupXtrInvitedBy `json:"items"`
}

// GroupsGetLongPollServer type represents `groups_getLongPollServer_response` API response object
type GroupsGetLongPollServer objects.GroupsLongPollServer

// GroupsGetLongPollSettings type represents `groups_getLongPollSettings_response` API response object
type GroupsGetLongPollSettings objects.GroupsLongPollSettings

// GroupsGetMembersFields type represents `groups_getMembers_fields_response` API response object
type GroupsGetMembersFields struct {
	Count int                         `json:"count"` // Total members number
	Items []objects.GroupsUserXtrRole `json:"items"`
}

// GroupsGetMembersFilter type represents `groups_getMembers_filter_response` API response object
type GroupsGetMembersFilter struct {
	Count int                        `json:"count"` // Total members number
	Items []objects.GroupsMemberRole `json:"items"`
}

// GroupsGetMembers type represents `groups_getMembers_response` API response object
type GroupsGetMembers struct {
	Count int   `json:"count"` // Total members number
	Items []int `json:"items"`
}

// GroupsGetRequestsFields type represents `groups_getRequests_fields_response` API response object
type GroupsGetRequestsFields struct {
	Count int                     `json:"count"` // Total communities number
	Items []objects.UsersUserFull `json:"items"`
}

// GroupsGetRequests type represents `groups_getRequests_response` API response object
type GroupsGetRequests struct {
	Count int   `json:"count"` // Total communities number
	Items []int `json:"items"`
}

// GroupsGetSettings type represents `groups_getSettings_response` API response object
type GroupsGetSettings objects.GroupsGroupSettings

// GroupsGetTokenPermissions type represents `groups_getTokenPermissions_response` API response object
type GroupsGetTokenPermissions struct {
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

// GroupsIsMemberExtended type represents `groups_isMember_extended_response` API response object
type GroupsIsMemberExtended struct {
	CanInvite  objects.BaseBoolInt `json:"can_invite"` // Information whether user can be invited
	CanRecall  objects.BaseBoolInt `json:"can_recall"` // Information whether user's invite to the group can be recalled
	Invitation objects.BaseBoolInt `json:"invitation"` // Information whether user has been invited to the group
	Member     objects.BaseBoolInt `json:"member"`     // Information whether user is a member of the group
	Request    objects.BaseBoolInt `json:"request"`    // Information whether user has sent request to the group
}

// GroupsIsMember type represents `groups_isMember_response` API response object
type GroupsIsMember objects.BaseBoolInt // Information whether user is a member of the group

// GroupsIsMemberUserIdsExtended type represents `groups_isMember_user_ids_extended_response` API response object
type GroupsIsMemberUserIdsExtended objects.GroupsMemberStatusFull

// GroupsIsMemberUserIds type represents `groups_isMember_user_ids_response` API response object
type GroupsIsMemberUserIds objects.GroupsMemberStatus

// GroupsSearch type represents `groups_search_response` API response object
type GroupsSearch struct {
	Count int                   `json:"count"` // Total communities number
	Items []objects.GroupsGroup `json:"items"`
}
