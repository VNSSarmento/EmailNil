package campaign

// O repositorio sempre vai ser quem vai ter as assinaturas dos metodos de requisição ao banco de dados
type Repository interface {
	Save(campaign *Campaign) error
	Get() ([]Campaign, error)
}
