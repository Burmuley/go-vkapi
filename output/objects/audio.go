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
// `audio` group of objects
/////////////////////////////////////////////////////////////////////////////////////////////////////////////

// AudioAudio type represents `audio_audio` API object
type AudioAudio struct {
	AlbumId   int    `json:"album_id"`  // Album ID
	Artist    string `json:"artist"`    // Artist name
	Date      int    `json:"date"`      // Date when uploaded
	Duration  int    `json:"duration"`  // Duration in seconds
	GenreId   int    `json:"genre_id"`  // Genre ID
	Id        int    `json:"id"`        // Audio ID
	Performer string `json:"performer"` // Performer name
	Title     string `json:"title"`     // Title
	Url       string `json:"url"`       // URL of mp3 file
}
