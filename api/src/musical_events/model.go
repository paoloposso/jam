package musicalevents

type MusicalEvent struct {
	Location    Location
	Name        string
	Description string
	ID          string
	CreatorID   string
}

type Location struct {
	Latitude  float64
	Longitude float64
}
