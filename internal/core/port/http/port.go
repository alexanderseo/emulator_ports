package port

type ServiceHandler interface {
	Read()
	Write()
}
