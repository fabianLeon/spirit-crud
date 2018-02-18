package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:CombatesController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:CombatesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:CombatesController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:CombatesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:CombatesController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:CombatesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:CombatesController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:CombatesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:CombatesController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:CombatesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:EventoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:EventoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:EventoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:EventoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:EventoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:EventoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:EventoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:EventoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:EventoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:EventoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:FotoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:FotoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:FotoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:FotoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:FotoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:FotoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:FotoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:FotoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:FotoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:FotoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:LugarController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:LugarController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:LugarController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:LugarController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:LugarController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:LugarController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:LugarController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:LugarController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:LugarController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:LugarController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:MovimientoFinancieroController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:MovimientoFinancieroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:MovimientoFinancieroController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:MovimientoFinancieroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:MovimientoFinancieroController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:MovimientoFinancieroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:MovimientoFinancieroController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:MovimientoFinancieroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:MovimientoFinancieroController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:MovimientoFinancieroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaEventoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaEventoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaEventoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaEventoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaEventoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaEventoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaEventoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaEventoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaEventoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaEventoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaFotoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaFotoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaFotoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaFotoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaFotoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaFotoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaFotoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaFotoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaFotoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaFotoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaSesionClaseController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaSesionClaseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaSesionClaseController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaSesionClaseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaSesionClaseController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaSesionClaseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaSesionClaseController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaSesionClaseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaSesionClaseController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:PersonaSesionClaseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:RollController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:RollController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:RollController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:RollController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:RollController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:RollController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:RollController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:RollController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:RollController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:RollController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:RutinaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:RutinaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:RutinaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:RutinaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:RutinaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:RutinaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:RutinaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:RutinaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:RutinaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:RutinaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:SesionClaseController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:SesionClaseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:SesionClaseController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:SesionClaseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:SesionClaseController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:SesionClaseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:SesionClaseController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:SesionClaseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:SesionClaseController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:SesionClaseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:SesionClaseRutinaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:SesionClaseRutinaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:SesionClaseRutinaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:SesionClaseRutinaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:SesionClaseRutinaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:SesionClaseRutinaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:SesionClaseRutinaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:SesionClaseRutinaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:SesionClaseRutinaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:SesionClaseRutinaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoConceptoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoConceptoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoConceptoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoConceptoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoConceptoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoConceptoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoConceptoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoConceptoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoConceptoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoConceptoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoMovimientoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoMovimientoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoMovimientoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoMovimientoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoMovimientoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit-crud/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
