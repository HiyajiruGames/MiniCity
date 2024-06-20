package store

type SystemData struct {
}

type SystemDataStore struct {
}

func (s *SystemDataStore) Write(systemData *SystemData) {

}

func (s *SystemDataStore) Load() SystemData {
	return SystemData{}
}
