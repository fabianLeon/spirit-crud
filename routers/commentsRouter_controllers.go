package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:CombatesController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:CombatesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:CombatesController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:CombatesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:CombatesController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:CombatesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:CombatesController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:CombatesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:CombatesController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:CombatesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:EventoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:EventoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:EventoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:EventoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:EventoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:EventoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:EventoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:EventoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:EventoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:EventoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:FotoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:FotoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:FotoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:FotoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:FotoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:FotoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:FotoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:FotoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:FotoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:FotoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:LugarController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:LugarController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:LugarController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:LugarController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:LugarController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:LugarController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:LugarController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:LugarController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:LugarController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:LugarController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:MovimientoFinancieroController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:MovimientoFinancieroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:MovimientoFinancieroController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:MovimientoFinancieroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:MovimientoFinancieroController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:MovimientoFinancieroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:MovimientoFinancieroController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:MovimientoFinancieroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:MovimientoFinancieroController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:MovimientoFinancieroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaEventoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaEventoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaEventoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaEventoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaEventoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaEventoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaEventoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaEventoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaEventoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaEventoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaFotoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaFotoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaFotoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaFotoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaFotoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaFotoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaFotoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaFotoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaFotoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaFotoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaSesionClaseController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaSesionClaseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaSesionClaseController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaSesionClaseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaSesionClaseController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaSesionClaseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaSesionClaseController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaSesionClaseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaSesionClaseController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:PersonaSesionClaseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:RollController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:RollController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:RollController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:RollController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:RollController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:RollController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:RollController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:RollController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:RollController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:RollController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:RutinaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:RutinaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:RutinaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:RutinaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:RutinaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:RutinaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:RutinaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:RutinaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:RutinaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:RutinaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:SesionClaseController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:SesionClaseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:SesionClaseController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:SesionClaseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:SesionClaseController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:SesionClaseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:SesionClaseController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:SesionClaseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:SesionClaseController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:SesionClaseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:SesionClaseRutinaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:SesionClaseRutinaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:SesionClaseRutinaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:SesionClaseRutinaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:SesionClaseRutinaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:SesionClaseRutinaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:SesionClaseRutinaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:SesionClaseRutinaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:SesionClaseRutinaController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:SesionClaseRutinaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoConceptoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoConceptoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoConceptoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoConceptoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoConceptoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoConceptoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoConceptoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoConceptoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoConceptoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoConceptoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoMovimientoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoMovimientoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoMovimientoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoMovimientoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoMovimientoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["github.com/fabianLeon/spirit_api/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
