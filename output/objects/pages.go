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
// `pages` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// PagesPrivacySettings type represents `pages_privacy_settings` API object
type PagesPrivacySettings int

// PagesWikipage type represents `pages_wikipage` API object
type PagesWikipage struct {
	CreatorId   int                  `json:"creator_id"`   // Page creator ID
	CreatorName int                  `json:"creator_name"` // Page creator name
	EditorId    int                  `json:"editor_id"`    // Last editor ID
	EditorName  string               `json:"editor_name"`  // Last editor name
	GroupId     int                  `json:"group_id"`     // Community ID
	Id          int                  `json:"id"`           // Page ID
	Title       string               `json:"title"`        // Page title
	Views       int                  `json:"views"`        // Views number
	WhoCanEdit  PagesPrivacySettings `json:"who_can_edit"` // Edit settings of the page
	WhoCanView  PagesPrivacySettings `json:"who_can_view"` // View settings of the page
}

// PagesWikipageFull type represents `pages_wikipage_full` API object
type PagesWikipageFull struct {
	Created                  int                  `json:"created"`                      // Date when the page has been created in Unixtime
	CreatorId                int                  `json:"creator_id"`                   // Page creator ID
	CurrentUserCanEdit       BaseBoolInt          `json:"current_user_can_edit"`        // Information whether current user can edit the page
	CurrentUserCanEditAccess BaseBoolInt          `json:"current_user_can_edit_access"` // Information whether current user can edit the page access settings
	Edited                   int                  `json:"edited"`                       // Date when the page has been edited in Unixtime
	EditorId                 int                  `json:"editor_id"`                    // Last editor ID
	GroupId                  int                  `json:"group_id"`                     // Community ID
	Html                     string               `json:"html"`                         // Page content, HTML
	Id                       int                  `json:"id"`                           // Page ID
	Source                   string               `json:"source"`                       // Page content, wiki
	Title                    string               `json:"title"`                        // Page title
	ViewUrl                  string               `json:"view_url"`                     // URL of the page preview
	Views                    int                  `json:"views"`                        // Views number
	WhoCanEdit               PagesPrivacySettings `json:"who_can_edit"`                 // Edit settings of the page
	WhoCanView               PagesPrivacySettings `json:"who_can_view"`                 // View settings of the page
}

// PagesWikipageHistory type represents `pages_wikipage_history` API object
type PagesWikipageHistory struct {
	Date       int    `json:"date"`        // Date when the page has been edited in Unixtime
	EditorId   int    `json:"editor_id"`   // Last editor ID
	EditorName string `json:"editor_name"` // Last editor name
	Id         int    `json:"id"`          // Version ID
	Length     int    `json:"length"`      // Page size in bytes
}
