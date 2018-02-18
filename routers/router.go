// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/fabianLeon/spirit-crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/combates",
			beego.NSInclude(
				&controllers.CombatesController{},
			),
		),

		beego.NSNamespace("/concepto",
			beego.NSInclude(
				&controllers.ConceptoController{},
			),
		),

		beego.NSNamespace("/evento",
			beego.NSInclude(
				&controllers.EventoController{},
			),
		),

		beego.NSNamespace("/foto",
			beego.NSInclude(
				&controllers.FotoController{},
			),
		),

		beego.NSNamespace("/lugar",
			beego.NSInclude(
				&controllers.LugarController{},
			),
		),

		beego.NSNamespace("/movimiento_financiero",
			beego.NSInclude(
				&controllers.MovimientoFinancieroController{},
			),
		),

		beego.NSNamespace("/persona",
			beego.NSInclude(
				&controllers.PersonaController{},
			),
		),

		beego.NSNamespace("/persona_evento",
			beego.NSInclude(
				&controllers.PersonaEventoController{},
			),
		),

		beego.NSNamespace("/persona_foto",
			beego.NSInclude(
				&controllers.PersonaFotoController{},
			),
		),

		beego.NSNamespace("/persona_sesion_clase",
			beego.NSInclude(
				&controllers.PersonaSesionClaseController{},
			),
		),

		beego.NSNamespace("/roll",
			beego.NSInclude(
				&controllers.RollController{},
			),
		),

		beego.NSNamespace("/rutina",
			beego.NSInclude(
				&controllers.RutinaController{},
			),
		),

		beego.NSNamespace("/sesion_clase",
			beego.NSInclude(
				&controllers.SesionClaseController{},
			),
		),

		beego.NSNamespace("/sesion_clase_rutina",
			beego.NSInclude(
				&controllers.SesionClaseRutinaController{},
			),
		),

		beego.NSNamespace("/tipo_concepto",
			beego.NSInclude(
				&controllers.TipoConceptoController{},
			),
		),

		beego.NSNamespace("/tipo_movimiento",
			beego.NSInclude(
				&controllers.TipoMovimientoController{},
			),
		),

		beego.NSNamespace("/tipo_sesion",
			beego.NSInclude(
				&controllers.TipoSesionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
