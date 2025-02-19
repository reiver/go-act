package asns

import (
	"github.com/reiver/go-opt"
	"github.com/reiver/go-jsonld"
)

// Object represents the ActivityStreams object name-space used in a JSON-LD document.
//
// Reference:
// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-object
//
// Example usage:
//
//	import (
//		"github.com/reiver/go-asns"
//		"github.com/reiver/go-jsonld"
//	)
//	
//	// ...
//	
//	var object asns.Object
//	
//	// ...
//	
//	bytes, err := jsonld.Marshal(object)
//
// More likely you would mix this with other JSON-LD name-spaces, with code similar to the following:
//
//	import (
//		"github.com/reiver/go-asns"
//		"github.com/reiver/go-tootns"
//		"github.com/reiver/go-jsonld"
//	)
//	
//	// ...
//	
//	var actor asns.Actor
//	var object asns.Object
//	var toot tootns.Toot
//	
//	// ...
//	
//	bytes, err := jsonld.Marshal(actor, object, toot)
type Object struct {
	NameSpace jsonld.NameSpace `jsonld:"https://www.w3.org/ns/activitystreams"`
	Prefix    jsonld.Prefix    `jsonld:"as"`

	AlsoKnownAs []string             `json:"alsoKnownAs,omitempty"`
	Attachment  []Attachment         `json:"attachment,omitempty"`
	Content     opt.Optional[string] `json:"content,omitempty"`
	ID          opt.Optional[string] `json:"id,omitempty"`
	Image	    Image                `json:"image,omitempty"`
	Icon	    Icon                 `json:"icon,omitempty"`
	MediaType   opt.Optional[string] `json:"mediaType,omitempty"`
	MovedTo     opt.Optional[string] `json:"movedTo,omitempty"`
	Name        opt.Optional[string] `json:"name,omitempty"`
	Summary     opt.Optional[string] `json:"summary,omitempty"`
	Tag         []HashTag            `json:"tag,omitempty"`
	Type        opt.Optional[string] `json:"type,omitempty"`
	URL         opt.Optional[string] `json:"url,omitempty"`
}
