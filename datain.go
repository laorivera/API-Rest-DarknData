package main

type Selection struct {
	Class    string   `json:"class"`
	Race     string   `json:"race"`
	ItemSlot ItemSlot `json:"itemSlot"`
}

type ItemSlot struct {
	Head      ItemSelect `json:"head"`
	Chest     ItemSelect `json:"chest"`
	Foot      ItemSelect `json:"foot"`
	Hands     ItemSelect `json:"hands"`
	Pants     ItemSelect `json:"pants"`
	Back      ItemSelect `json:"back"`
	Necklace  ItemSelect `json:"necklace"`
	RingOne   ItemSelect `json:"ringOne"`
	RingTwo   ItemSelect `json:"ringTwo"`
	WeaponOne ItemSelect `json:"weaponOne"`
	WeaponTwo ItemSelect `json:"weaponTwo"`
}

type ItemSelect struct {
	Id      string        `json:"id"`
	Name    string        `json:"name"`
	Rarity  string        `json:"rarity"`
	Enchant EnchantSelect `json:"enchant"`
	Rating  string        `json:"rating"`
}

type EnchantSelect struct {
	TypeU  string `json:"typeu"`
	ValueU string `json:"valueu"`
	TypeR  string `json:"typer"`
	ValueR string `json:"valuer"`
	TypeE  string `json:"typee"`
	ValueE string `json:"valuee"`
	TypeL  string `json:"typel"`
	ValueL string `json:"valuel"`
	TypeQ  string `json:"typeq"`
	ValueQ string `json:"valueq"`
	TypeA  string `json:"typea"`
	ValueA string `json:"valuea"`
}
