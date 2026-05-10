package app

import (
	"robot/internal/excel"
	"robot/internal/domain"
)
func LoadFromExcel(path string) error {
	parsedRows, err := excel.ParseRows("form.xlsx")
	if err != nil {
		return err
	}
	categorias := map[string]*domain.Categoria{}
	for _, row := range parsedRows {
		equipo, err := buildTeam(row)
		if err != nil {
			return err 
		}
		if err := buildCategory(categorias, row.Categoria, equipo); err != nil {
			return err 
		}
	}
	return nil
}

func buildTeam(row excel.FormRow) (*domain.Equipo, error) {
	lider, err := domain.NewIntegrante(row.Capitan, row.Correo, true)
	if err != nil {
		return nil, err
	}

	integrantes, err := buildIntegrantes(row.Integrantes, row.CorreosIntegrantes)
	if err != nil {
		return nil, err
	}
	integrantes = append(integrantes, lider)
	
	return domain.NewEquipo(row.Equipo, row.Escuela, row.Nivel, integrantes)
}

func buildCategory(categorias map[string]*domain.Categoria, name string, equipo *domain.Equipo) error {
	categoria, exists := categorias[name]
	if !exists {
		categoria, err := domain.NewCategoria(name, equipo)
		if err != nil {
			return err
		}
		categorias[name] = categoria
		return nil
	}

	return categoria.AddEquipo(equipo)
}

func buildIntegrantes(nombres, correos []string) ([]*domain.Integrante, error){
	var integrantes []*domain.Integrante

	for index, nombre := range nombres {
		correo := ""
		if index < len(correos) {
			correo = correos[index]
		}
		persona, err := domain.NewIntegrante(nombre, correo, false)
		if err != nil {
			return integrantes, domain.ErrUnsuccessIntegrante
		}
		integrantes = append(integrantes, persona)
	}
	return integrantes, nil
}