package api

var socketChannels =  make(map[string]chan interface{})

func Log(data interface{}) {
	for  k := range socketChannels {
		socketChannels[k] <- data
	}
}
