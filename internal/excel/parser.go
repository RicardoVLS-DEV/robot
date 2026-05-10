package excel

import (
	"strings"

	"github.com/xuri/excelize/v2"
)

const (
	HeaderMarcaTemporal     = "Marca temporal"
	HeaderCorreo            = "Correo electrónico"
	HeaderCategoria         = "Categoría a participar"
	HeaderEquipo            = "Nombre del equipo"
	HeaderIntegrantes       = "Nombre de los integrantes del equipo. Los equipos deberán estar conformados de mínimo dos y hasta cuatro integrantes."
	HeaderCorreoIntegrantes = "Correos de los Integrantes del equipo"
	HeaderEscuela      		= "Si eres alumno de la UTNC captura el nombre de tu carrera y tu grupo, si eres de otra institución compártenos donde estudias"
	HeaderNivel       		= "¿Nivel de escolaridad que cursa?"
	HeaderCapitan           = "Nombre del capitán del equipo"
	HeaderAsesor            = "Nombre del asesor"
)

type FormRow struct {
	Correo         		string
	Categoria          	string
	Equipo       		string
	Integrantes        	[]string
	CorreosIntegrantes 	[]string
	Escuela      		string
	Nivel 				string
	Capitan            	string
	Asesor             	string
}

func ParseRows(path string) ([]FormRow, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sheet := f.GetSheetName(1)
	return getRows(f, sheet)
}

func getRows(f *excelize.File, sheet string) ([]FormRow, error) {
	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return []FormRow{}, nil
	}

	headers := headerIndexes(rows[0])
	formRows := make([]FormRow, 0, len(rows)-1)

	for _, row := range rows[1:] {
		formRows = append(formRows, FormRow{
			Correo:         	cell(row, headers, HeaderCorreo),
			Categoria:          cell(row, headers, HeaderCategoria),
			Equipo:       		cell(row, headers, HeaderEquipo),
			Integrantes:        splitCell(cell(row, headers, HeaderIntegrantes)),
			CorreosIntegrantes: splitCell(cell(row, headers, HeaderCorreoIntegrantes)),
			Escuela:        	cell(row, headers, HeaderEscuela),
			Nivel:        		cell(row, headers, HeaderNivel),
			Capitan:            cell(row, headers, HeaderCapitan),
			Asesor:             cell(row, headers, HeaderAsesor),
		})
	}

	return formRows, nil
}

func headerIndexes(headers []string) map[string]int {
	indexes := make(map[string]int, len(headers))
	for index, header := range headers {
		header = strings.TrimSpace(header)
		indexes[header] = index
	}

	return indexes
}

func cell(row []string, headers map[string]int, header string) string {
	index, ok := headers[header]
	if !ok || index >= len(row) {
		return ""
	}

	return strings.TrimSpace(row[index])
}

func splitCell(value string) []string {
	parts := strings.Split(value, ",")
	values := make([]string, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		values = append(values, part)
	}

	return values
}
