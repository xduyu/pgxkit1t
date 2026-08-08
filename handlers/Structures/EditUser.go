package structures

import gstructures "v1/test/fiber/db/helpers/GStructures"

type HReqEditUser struct {
	Editable_UserId int                   `json:"id"`
	Data            gstructures.GEditUser `json:"data"`
}
