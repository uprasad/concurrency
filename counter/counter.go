package counter

type Counter interface {
	Inc()
	Get() int64
	Reset()
}
