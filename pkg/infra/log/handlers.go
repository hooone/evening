package log

//DisposableHandler : interface for logger that can be Close
type DisposableHandler interface {
	Close()
}

//ReloadableHandler : interface for logger that can be reload
type ReloadableHandler interface {
	Reload()
}
