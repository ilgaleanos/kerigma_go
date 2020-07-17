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
	"errors"

	guuid "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Login estructura del para realizar login y generación de tokens
type Login struct {
	Usuario string  `json:"usuario,omitempty"` // _id de mongo para login
	Clave   string  `json:"clave,omitempty"`   // Se debe omitir siempre en los json
	Token   *string `json:"token,omitempty"`   // JWT generado
}

// =====================================================================================================================
//
// 	Flujos
//
// =====================================================================================================================

// FlujoDeIngreso permite validar si usuario es o no valido en base de datos con las credenciales provistas
func (login *Login) FlujoDeIngreso() error {
	if login.Clave == "" {
		return errors.New("parámetro requerido: clave")
	}

	credenciales := Credenciales{ID: guuid.New().String()}
	credenciales.consultarUsuario(&login.Usuario)

	// validamos que la contraseña concuerde, en caso de que no retornamos un error estándar por seguridad
	if ok := checkPasswordHash(login.Clave, credenciales.Clave); !ok {
		return errors.New("parámetro requerido: clave")
	}
	if credenciales.Estado != 1 {
		return errors.New("Usuario no disponible")
	}

	login.Token = credenciales.ElaborarToken()
	return nil
}

// =====================================================================================================================
//
// 	Funciones HASH
//
// =====================================================================================================================

// función para generar un hash a partir de un password plano
func hashPassword(password string) (string, error) {
	// se emplea u factor de trabajo de 10 para balance seguridad rendimiento
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}

// función que compara un password plano con su hash para verificar su coincidencia
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
