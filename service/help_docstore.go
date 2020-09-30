package service

type DocStore struct {
	DocumentId			uint64
	Directory 			string
	Filename 			string
	Format 				string
	ContentBytes 		[]byte
}
