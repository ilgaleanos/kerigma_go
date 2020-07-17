/*
 * Copyright 2020 Leito. All Rights Reserved.
 * <p>
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * <p>
 * http://www.apache.org/licenses/LICENSE-2.0
 * <p>
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package controladores

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
	"gopkg.in/mgo.v2/bson"
	"kerigma/src/logica/ingreso"
	"kerigma/src/utiles"
)

// =====================================================================================================================
//
// 	Controladores
//
// =====================================================================================================================

// LoginHandlerGet controlador del método de ingreso a la plataforma
func IngresarGet(xto *fasthttp.RequestCtx) {
	var err error
	credenciales := &ingreso.Credenciales{}
	if err = credenciales.ValidarToken(string(xto.Request.Header.Peek("authorization")), 0); credenciales == nil || err != nil {
		utiles.Forbidden(xto)
		return
	}

	utiles.SendJSON(xto, &bson.M{"aceptado": true})
}

// LoginHandlerPost de login recibe username y password y genera un JWT que es enviado como json
func IngresarPost(xto *fasthttp.RequestCtx) {
	// estructura para convertir la entrada, se espera un json válido
	login := ingreso.Login{}
	if err := json.Unmarshal(xto.PostBody(), &login); err != nil {
		utiles.BadRequest(xto)
		return
	}

	// cotejamos con base de datos los datos, el token lo empaquetamos en la variable login de regreso
	if err := login.FlujoDeIngreso(); err != nil {
		utiles.SendJSON(xto, &bson.M{"err": err.Error()})
		return
	}

	// enviamos el token de acceso al usuario
	utiles.SendJSON(xto, &bson.M{"token": login.Token})
}
