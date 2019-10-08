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
// `pages` group of responses
/////////////////////////////////////////////////////////////

// PagesSaveAccess type represents `pages_saveAccess_response` API response object
type PagesSaveAccess int // Page ID

// PagesGetTitles type represents `pages_getTitles_response` API response object
type PagesGetTitles objects.PagesWikipage

// PagesGetVersion type represents `pages_getVersion_response` API response object
type PagesGetVersion objects.PagesWikipageFull

// PagesSave type represents `pages_save_response` API response object
type PagesSave int // Page ID

// PagesParseWiki type represents `pages_parseWiki_response` API response object
type PagesParseWiki string // HTML source

// PagesGet type represents `pages_get_response` API response object
type PagesGet objects.PagesWikipageFull

// PagesGetHistory type represents `pages_getHistory_response` API response object
type PagesGetHistory objects.PagesWikipageHistory
