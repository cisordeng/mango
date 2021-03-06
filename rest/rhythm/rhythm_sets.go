package rhythm

import (
	"github.com/cisordeng/beego/xenon"
	bRhythm "nature/business/rhythm"
)

type RhythmSets struct {
	xenon.RestResource
}

func init () {
	xenon.RegisterResource(new(RhythmSets))
}

func (this *RhythmSets) Resource() string {
	return "rhythm.rhythm_sets"
}

func (this *RhythmSets) Params() map[string][]string {
	return map[string][]string{
		"GET":  []string{},
	}
}

func (this *RhythmSets) Get() {
	page := this.GetPage()
	rhythmSets, pageInfo := bRhythm.GetPagedRhythmSets(page, xenon.Map{}, "-index")
	data := bRhythm.EncodeManyRhythmSet(rhythmSets)
	this.ReturnJSON(xenon.Map{
		"rhythm_sets": data,
		"page_info": pageInfo.ToMap(),
	})
}
