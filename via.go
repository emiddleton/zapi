package zapi

import ()

type Via struct {
	Channel string                 // 	This tells you how the ticket or event was created
	Source  map[string]interface{} // 	For some channels a source object gives more information about how or why the ticket or event was created
}
