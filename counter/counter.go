package counter

type Counter interface {
	Inc()
	Dec()
	Get() int64
	Reset()
}
