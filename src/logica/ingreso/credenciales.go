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
	"github.com/dgrijalva/jwt-go/v4"
	"kerigma/src/utiles"
	"time"
)

// Credenciales estructura del dispositivo a nivel de base de datos pero también a nivel de JWT
type Credenciales struct {
	ID         string `json:"id,omitempty"`         // uuid del token, relacionaría la sesión
	Clave      string `json:"-"`                    // Se debe omitir siempre en los json
	Expiracion int64  `json:"expiracion,omitempty"` // Solo se una a nivel de JWT
	Alcance    int8   `json:"alcance,omitempty" `   // alcance de privilegios del usuario
	Estado     int8   `json:"estado,omitempty"`     // bandera para saber si el usuario esta activo en la plataforma
}

// =====================================================================================================================
//
// 	Funciones encargadas de los JWT
//
// =====================================================================================================================

// ElaborarToken función generadora de JWT para un dispositivo especifico
func (credencial *Credenciales) ElaborarToken() *string {
	// hora actual para el conteo del token en UTC
	now := time.Now().UTC()
	claims := jwt.MapClaims{
		"id":  credencial.ID,                  // uuid del token
		"alc": credencial.Alcance,             // alcance de privilegios del dispositivo
		"nad": now.Unix(),                     // no valido antes de
		"exp": now.Add(24 * time.Hour).Unix(), // fecha de expiración de 24 horas
	}

	// generación del token
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, &claims)
	// firma del token
	tokenSigned, err := token.SignedString(utiles.SIGNKEY)
	if err != nil {
		panic(err.Error())
	}

	return &tokenSigned
}

// ValidarToken validamos el token de autenticación
func (credencial *Credenciales) ValidarToken(tokenString string, scope int8) error {
	// convertimos el token y validamos con la clave pública
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return utiles.VERIFYKEY, nil
	})
	if err != nil || !token.Valid {
		return err
	}
	// extraemos los datos del token
	claims := token.Claims.(jwt.MapClaims)

	// convertimos las variables de los claims a la estructura, si hay algún problema las credenciales son nil
	var ok bool
	var tmp float64
	if tmp, ok = claims["alc"].(float64); !ok {
		return errors.New("alcance inválido")
	}
	credencial.Alcance = int8(tmp)
	if credencial.Alcance < scope {
		return errors.New("alcance inválido")
	}

	if credencial.ID, ok = claims["id"].(string); !ok {
		return errors.New("token inválido")
	}
	if tmp, ok = claims["exp"].(float64); !ok {
		return errors.New("token expirado")
	}
	credencial.Expiracion = int64(tmp)

	// si todo esta ok se regresa error nulo
	return nil
}
