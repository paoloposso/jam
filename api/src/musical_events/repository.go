package musicalevents

type Repository interface {
	Create(MusicalEvent) (string, error)
}
