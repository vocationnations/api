package model

import "fmt"

type CultureEntry struct {
	tableName struct{} `pg:"cultureentries"`
	Id        int      `json:"id"`
	Clan      float64  `json:"clan"`
	Adhocracy float64  `json:"adhocracy"`
	Market    float64  `json:"market"`
	Hierarchy float64  `json:"hierarchy"`
	RawEntry  string   `json:"raw_entry"`
	UserId    int      `json:"user_id"`
}

func (cc CultureEntry) String() string {
	return fmt.Sprintf("ID: <%d> Clan: <%f> Adhocracy: <%f> Market <%f> Hierarchy <%f>", cc.Id, cc.Clan, cc.Adhocracy, cc.Market, cc.Hierarchy)
}
