package dto

type CategoriaDto struct {
	Nombre string `json:"nombre"`
}
type ProductoDto struct {
	Nombre      string `json:"nombre"`
	Precio      int    `json:"precio"`
	Stock       int    `json:"stock"`
	Descripcion string `json:"descripcion"`
	CategoriaID uint   `json:"categoria_id"`
}
type PerfilDto struct {
	Nombre string `json:"nombre"`
}
type UsuarioDto struct {
	Nombre   string `json:"nombre"`
	Correo   string `json:"correo"`
	Telefono string `json:"telefono"`
	Password string `json:"password"`
	PerfilID uint   `json:"perfil_id"`
}
type LoginDto struct {
	Correo   string `json:"correo"`
	Password string `json:"password"`
}

type LoginRespuestaDto struct {
	Nombre string `json:"nombre"`
	Token  string `json:"token"`
}
