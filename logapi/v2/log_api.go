package logapi

// LogCreatingV2 generates log messages
type LogCreatingV2 interface {
	Info(msg LogAwareStringing, parms ...interface{})
	Error(msg LogAwareStringing, parms ...interface{})
	Warn(msg LogAwareStringing, parms ...interface{})
	Debug(msg LogAwareStringing, parms ...interface{})
}

// ToString represents a late binding on a string value
type ToString interface {
	String() string
}

// //Stereotype establishes the mode using the formating.
// type Stereotype int

// const (
// 	//DefaultStereo denotes the standard formating
// 	DefaultStereo Stereotype = iota
// 	//TimeStereo denotes is the logging option above info
// 	TimeStereo
// 	//OKStereo is the logging option above info
// 	OKStereo
// 	//WorkingStereo is the logging option above warn
// 	WorkingStereo
// 	//WaitingForUserStereo is the logging option above error
// 	WaitingForUserStereo
// 	//ExitStereo is the logging option above error
// 	ExitStereo
// 	//NotImplementedStereo is the logging option above error
// 	NotImplementedStereo
// )

// //IntP returns the uint for level
// func (s Stereotype) IntP() *uint {
// 	myInt := uint(s)
// 	return &myInt
// }

// func (s Stereotype) String() string {
// 	switch s {
// 	case TimeStereo:
// 		return "TIME"
// 	case OKStereo:
// 		return "OK"
// 	case WorkingStereo:
// 		return "WORKING"
// 	case WaitingForUserStereo:
// 		return "USERINPUT"
// 	case ExitStereo:
// 		return "EXIT"
// 	case NotImplementedStereo:
// 		return "UNFINISHED"
// 	default:
// 		return ""
// 	}

// }
