package maps

type rwLockMap struct {
	ma map[string]string
	rw *RWLock
}

func (m *rwLockMap) lookup(key string) string {
	m.rw.startRead()
	defer m.rw.doneRead()
	return m.ma[key]
}

func (m *rwLockMap) insert(key, value string) {
	m.rw.startWrite()
	defer m.rw.doneWrite()
	m.ma[key] = value
}
