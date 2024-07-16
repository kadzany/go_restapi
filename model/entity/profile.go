package entity

import (
	"github.com/dyaksa/encryption-pii/crypt/types"
	"github.com/google/uuid"
)

type Profile struct {
	ID        uuid.UUID
	Nik       byte[] `enc:"sha256cbc" bidx_col:"nik_bidx" txt_heap_table:"nik_text_heap"`
	Name      byte[] `enc:"sha256cbc" bidx_col:"name_bidx" txt_heap_table:"name_text_heap"`
	Phone     byte[] `enc:"sha256cbc" bidx_col:"phone_bidx" txt_heap_table:"phone_text_heap"`
	Email     byte[] `enc:"sha256cbc" bidx_col:"email_bidx" txt_heap_table:"email_text_heap"`
	DOB       types.AEADString
}

type FetchProfileRow struct {
	ID    uuid.UUID
	Nik   types.AEADString
	Name  types.AEADString
	Phone types.AEADString
	Email types.AEADString
	DOB   types.AEADString
}

type FetchProfileParams struct {
	ID uuid.UUID
}

// type FindProfileByIDParams struct {
// 	ID []uuid.UUID
// }

// type FindProfileByNameParams struct {
// 	NameBidx types.BIDXString
// }

// type FindProfileByEmailParams struct {
// 	EmailBidx types.BIDXString
// }

type FindProfilesByNameRow struct {
	ID    uuid.UUID
	Nik   types.AEADString
	Name  types.AEADString
	Phone types.AEADString
	Email types.AEADString
	Dob   types.AEADString
}

type FindProfileByParams struct {
	Type    string
	Content string
}

type FindProfileByBIDXParams struct {
	ColumnHeap string
	Hash       []string
}
