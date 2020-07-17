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
	"kerigma/src/utiles"

	"github.com/valyala/fasthttp"
)

const kerigma = "  _  __         _                       \n" +
	" | |/ /        (_)                      \n" +
	" | ' / ___ _ __ _  __ _ _ __ ___   __ _ \n" +
	" |  < / _ \\ '__| |/ _` | '_ ` _ \\ / _` |\n" +
	" | . \\  __/ |  | | (_| | | | | | | (_| |\n" +
	" |_|\\_\\___|_|  |_|\\__, |_| |_| |_|\\__,_|\n" +
	"                   __/ |                \n" +
	"                  |___/                 "

// =====================================================================================================================
//
// 	Controladores
//
// =====================================================================================================================

// Default todo lo que no coincida con una ruta llega aquí
func Default(c *fasthttp.RequestCtx) {
	utiles.SendTEXT(c, kerigma)
}
