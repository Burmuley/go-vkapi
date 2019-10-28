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
// Source schema can be found at https://github.com/VKCOM/vk-api-schema/blob/master/methods.json           //
// Code generator location: https://gitlab.com/Burmuley/go-vkapi-gen                                       //
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

package go_vkapi

import (
	"gitlab.com/Burmuley/go-vkapi/responses"
)

type Pages struct {
	*VKApi
}

/////////////////////////////////////////////////////////////
// `Pages` methods
/////////////////////////////////////////////////////////////

// Clearcache - Allows to clear the cache of particular 'external' pages which may be attached to VK posts.
// Parameters:
//   * url - Address of the page where you need to refesh the cached version
func (p Pages) Clearcache(url string) (resp responses.Ok, err error) {
	params := map[string]interface{}{}

	params["url"] = url

	err = p.SendObjRequest("pages.clearCache", params, &resp)

	return
}

// Get - Returns information about a wiki page.
// Parameters:
//   * ownerId - Page owner ID.
//   * pageId - Wiki page ID.
//   * global - '1' — to return information about a global wiki page
//   * sitePreview - '1' — resulting wiki page is a preview for the attached link
//   * title - Wiki page title.
//   * needSource - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * needHtml - '1' — to return the page as HTML,
func (p Pages) Get(ownerId int, pageId int, global bool, sitePreview bool, title string, needSource bool, needHtml bool) (resp responses.PagesGet, err error) {
	params := map[string]interface{}{}

	if ownerId > 0 {
		params["owner_id"] = ownerId
	}

	if pageId > 0 {
		params["page_id"] = pageId
	}

	params["global"] = global

	params["site_preview"] = sitePreview

	if title != "" {
		params["title"] = title
	}

	params["need_source"] = needSource

	params["need_html"] = needHtml

	err = p.SendObjRequest("pages.get", params, &resp)

	return
}

// Gethistory - Returns a list of all previous versions of a wiki page.
// Parameters:
//   * pageId - Wiki page ID.
//   * groupId - ID of the community that owns the wiki page.
//   * userId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
func (p Pages) Gethistory(pageId int, groupId int, userId int) (resp responses.PagesGethistory, err error) {
	params := map[string]interface{}{}

	params["page_id"] = pageId

	if groupId > 0 {
		params["group_id"] = groupId
	}

	if userId > 0 {
		params["user_id"] = userId
	}

	err = p.SendObjRequest("pages.getHistory", params, &resp)

	return
}

// Gettitles - Returns a list of wiki pages in a group.
// Parameters:
//   * groupId - ID of the community that owns the wiki page.
func (p Pages) Gettitles(groupId int) (resp responses.PagesGettitles, err error) {
	params := map[string]interface{}{}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = p.SendObjRequest("pages.getTitles", params, &resp)

	return
}

// Getversion - Returns the text of one of the previous versions of a wiki page.
// Parameters:
//   * versionId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * groupId - ID of the community that owns the wiki page.
//   * userId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * needHtml - '1' — to return the page as HTML
func (p Pages) Getversion(versionId int, groupId int, userId int, needHtml bool) (resp responses.PagesGetversion, err error) {
	params := map[string]interface{}{}

	params["version_id"] = versionId

	if groupId > 0 {
		params["group_id"] = groupId
	}

	if userId > 0 {
		params["user_id"] = userId
	}

	params["need_html"] = needHtml

	err = p.SendObjRequest("pages.getVersion", params, &resp)

	return
}

// Parsewiki - Returns HTML representation of the wiki markup.
// Parameters:
//   * text - Text of the wiki page.
//   * groupId - ID of the group in the context of which this markup is interpreted.
func (p Pages) Parsewiki(text string, groupId int) (resp responses.PagesParsewiki, err error) {
	params := map[string]interface{}{}

	params["text"] = text

	if groupId > 0 {
		params["group_id"] = groupId
	}

	err = p.SendObjRequest("pages.parseWiki", params, &resp)

	return
}

// Save - Saves the text of a wiki page.
// Parameters:
//   * text - Text of the wiki page in wiki-format.
//   * pageId - Wiki page ID. The 'title' parameter can be passed instead of 'pid'.
//   * groupId - ID of the community that owns the wiki page.
//   * userId - User ID
//   * title - Wiki page title.
func (p Pages) Save(text string, pageId int, groupId int, userId int, title string) (resp responses.PagesSave, err error) {
	params := map[string]interface{}{}

	if text != "" {
		params["text"] = text
	}

	if pageId > 0 {
		params["page_id"] = pageId
	}

	if groupId > 0 {
		params["group_id"] = groupId
	}

	if userId > 0 {
		params["user_id"] = userId
	}

	if title != "" {
		params["title"] = title
	}

	err = p.SendObjRequest("pages.save", params, &resp)

	return
}

// Saveaccess - Saves modified read and edit access settings for a wiki page.
// Parameters:
//   * pageId - Wiki page ID.
//   * groupId - ID of the community that owns the wiki page.
//   * userId - !!! NO DESCRIPTION IN JSON SCHEMA !!!
//   * view - Who can view the wiki page: '1' — only community members, '2' — all users can view the page, '0' — only community managers
//   * edit - Who can edit the wiki page: '1' — only community members, '2' — all users can edit the page, '0' — only community managers
func (p Pages) Saveaccess(pageId int, groupId int, userId int, view int, edit int) (resp responses.PagesSaveaccess, err error) {
	params := map[string]interface{}{}

	params["page_id"] = pageId

	if groupId > 0 {
		params["group_id"] = groupId
	}

	if userId > 0 {
		params["user_id"] = userId
	}

	if view > 0 {
		params["view"] = view
	}

	if edit > 0 {
		params["edit"] = edit
	}

	err = p.SendObjRequest("pages.saveAccess", params, &resp)

	return
}
