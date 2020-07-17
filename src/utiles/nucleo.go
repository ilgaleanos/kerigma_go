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

package utiles

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
	"gopkg.in/mgo.v2/bson"
)

var (
	metodoNoPermitido = []byte(`{"err":"metodo no permitido"}`)
	peticionErrada    = []byte(`{"err":"petición errada"}`)
	respuestaErrada   = []byte(`{"err":"respuesta errada"}`)
	noAutorizado      = []byte(`{"err":"no autorizado"}`)
)

// =====================================================================================================================
//
// 	Funciones base de entrega de contenido
//
// =====================================================================================================================

// MethodNotAllowed respuesta JSON en caso de no tener implementado el método solicitado
func MethodNotAllowed(c *fasthttp.RequestCtx) {
	c.SetStatusCode(405)
	c.SetContentType("application/json")
	c.Write(metodoNoPermitido)
}

// BadRequest respuesta JSON en caso de una petición mal formada
func BadRequest(c *fasthttp.RequestCtx) {
	c.SetStatusCode(400)
	c.SetContentType("application/json")
	c.Write(peticionErrada)
}

// Forbidden respuesta JSON en caso de no contar con las credenciales adecuadas
func Forbidden(c *fasthttp.RequestCtx) {
	c.SetStatusCode(403)
	c.SetContentType("application/json")
	c.Write(noAutorizado)
}

// SendJSON respuesta JSON a partir de una estructura flexible
func SendJSON(c *fasthttp.RequestCtx, j *bson.M) {
	jb, err := json.Marshal(j)
	if err != nil {
		c.SetStatusCode(500)
		c.SetContentType("application/json")
		c.Write(respuestaErrada)
		return
	}
	c.SetContentType("application/json")
	c.Write(jb)
}

// SendTEXT respuesta tipo texto plano con alguna cadena en particular
func SendTEXT(c *fasthttp.RequestCtx, s string) {
	c.SetContentType("text/plain")
	c.WriteString(s)
}
