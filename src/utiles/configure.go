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
	"crypto/rsa"
	"io/ioutil"
	"os"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var (
	VERIFYKEY *rsa.PublicKey
	SIGNKEY   *rsa.PrivateKey
)

// =====================================================================================================================
//
// 	Funciones de configuración
//
// =====================================================================================================================

// tareas iniciales del sistema
func Configure() {
	// se leen las variables de entorno del archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// configuración del nivel de logs para el sistema
	if "DEBUG" == os.Getenv("ENVIRONMENT") {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}

	// carga de las claves de firmas de JWT, la clave privada es información sensible
	SIGNKEY = loadRSAPrivateKeyFromDisk("keys/rs512-4096-private.pem")
	VERIFYKEY = loadRSAPublicKeyFromDisk("keys/rs512-4096-public.pem")
}

// Función encargada de leer del disco la clave privada y convertirla de PEM a RSA key
func loadRSAPrivateKeyFromDisk(location string) *rsa.PrivateKey {
	keyData, err := ioutil.ReadFile(location)
	if err != nil {
		panic(err.Error())
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		panic(err.Error())
	}
	return signKey
}

// Función encargada de leer del disco la clave publica y convertirla de PEM a RSA key
func loadRSAPublicKeyFromDisk(location string) *rsa.PublicKey {
	keyData, err := ioutil.ReadFile(location)
	if err != nil {
		panic(err.Error())
	}
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		panic(err.Error())
	}
	return verifyKey
}
