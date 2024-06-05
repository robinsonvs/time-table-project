package entity

type Professor struct {
	ID              int64  `json:"id"`
	UUID            string `json:"uuid"`
	Name            string `json:"name"`
	HoursToAllocate int    `json:"hours_to_allocate"`
}
