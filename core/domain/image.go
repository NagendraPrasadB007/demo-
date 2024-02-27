package domain

import "time"

type Image struct {
	Imageid         int       `json:"imageid"`
	Filename        string    `json:"filename"`
	Size            int       `json:"size"`
	Mimetype        string    `json:"mimetype"`
	Createddatetime time.Time `json:"createddatetime"`
	Filepath        string    `json:"filepath"`
	Pickuprequestid int       `json:"pickuprequestid"`
}
