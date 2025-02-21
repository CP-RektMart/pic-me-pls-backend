package object

import "strings"

type Folder string

const (
	FolderGallery       Folder = "GALLERY"
	FolderVerifyCitizen Folder = "VERIFY_CITIZENCARD"
	FolderProfileImage  Folder = "PROFILE_IMAGE"
)

func (f Folder) GetFullPath(fileName string) string {
	switch f {
	case FolderGallery, FolderVerifyCitizen, FolderProfileImage:
		return strings.ToLower(string(f)) + "/" + fileName
	}
	return "others/" + fileName
}
