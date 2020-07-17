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

package ingreso

import (
	"context"

	log "github.com/sirupsen/logrus"
	"kerigma/src/utiles"
)

// =====================================================================================================================
//
// 	Funciones SQL
//
// =====================================================================================================================

// Función que consulta al usuario a base de datos
func (credenciales *Credenciales) consultarUsuario(usuario *string) {
	err := utiles.DB.QueryRow(
		context.Background(),
		"SELECT clave, estado, alcance FROM usuarios WHERE usuario = $1",
		usuario,
	).Scan(&credenciales.Clave, &credenciales.Estado, &credenciales.Alcance)

	if err != nil {
		log.Error(err.Error())
	}
}
