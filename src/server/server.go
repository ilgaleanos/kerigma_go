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

package main

import (
	"runtime"

	"github.com/buaazp/fasthttprouter"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"kerigma/src/controladores"
	"kerigma/src/utiles"
)

// =====================================================================================================================
//
// 	Sistema básico de CORS
//
// =====================================================================================================================

type intercesor map[string]fasthttp.RequestHandler

func (maria intercesor) intercesora(xto *fasthttp.RequestCtx) {
	handler := maria["Handler"]
	handler(xto)
	xto.Response.Header.Set("Access-Control-Allow-Origin", "*")
	xto.Response.Header.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, withCredentials")
	xto.Response.Header.Set("Access-Control-Allow-Credentials", "true")
	xto.Response.Header.Set("Server", "ofs")
}

// =====================================================================================================================
//
// 	Función Principal
//
// =====================================================================================================================

// Por aquí se inicia todo el proceso
func main() {
	// se generan las variables globales a partir de variables del entorno
	utiles.Configure()

	maxConn := runtime.NumCPU() * 4
	if err := utiles.InitDB(maxConn); err != nil {
		panic(err.Error())
	}
	defer utiles.CloseDB()

	router := fasthttprouter.New()
	router.GET("/", controladores.Default)
	router.GET("/credenciales/", controladores.IngresarGet)
	router.POST("/credenciales/", controladores.IngresarPost)

	maria := make(intercesor)
	maria["Handler"] = router.Handler

	log.Fatal(fasthttp.ListenAndServe(":8080", maria.intercesora))
}
