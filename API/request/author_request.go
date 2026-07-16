package request

type AuthorRequestCreate struct {
	Name       string `json:"name" binding:"required"`
	Gender     int    `json:"gender" binding:"oneof=0 1 2"`
	Note       string `json:"note"`
	FacultyIDs []*int `json:"faculty_ids"`
}

type AuthorRequestUpdate struct {
	Name       *string `json:"name" binding:"required"`
	Gender     *int    `json:"gender" binding:"oneof=0 1 2"`
	Note       *string `json:"note"`
	FacultyIDs []*int  `json:"faculty_ids"`
}
