package trailer

type Handler interface {
	OnStatus(*Tweet)
	OnDelete(*DeleteEvent)
	OnEvent(*Event)
}
