package zapi

import ()

type Attachment struct {
	Id          int64        // yes 	Automatically assigned when created
	FileName    string       //	yes 	The name of the image file
	ContentUrl  string       //	yes 	A full URL where the attachment image file can be downloaded
	ContentType string       //	yes 	The content type of the image. Example value: image/png
	Size        integer      //	yes 	The size of the image file in bytes
	Thumbnails  []Attachment // yes 	An array of Photo objects. Note that thumbnails do not have thumbnails.
	Inline      bool         // yes 	If true, the attachment is excluded from the attachment list and the attachment's URL can be referenced within the comment of a ticket. Default is false
}

func Show(attachmentId int64) (Attachment, error) {
	return Attachment{}, nil
}
