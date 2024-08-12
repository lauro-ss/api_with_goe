package service

type UserRepository interface {
	Create()
	Update()
	Get()
	List()
	Delete()
}
