package objects

import "strings"

type Folder string

const (
	FolderPackage       Folder = "PACKAGE"
	FolderVerifyCitizen Folder = "VERIFY_CITIZENCARD"
	FolderProfileImage  Folder = "PROFILE_IMAGE"
)

func (f Folder) GetFullPath(fileName string) string {
	switch f {
	case FolderPackage, FolderVerifyCitizen, FolderProfileImage:
		return strings.ToLower(string(f)) + "/" + fileName
	}
	return "others/" + fileName
}
