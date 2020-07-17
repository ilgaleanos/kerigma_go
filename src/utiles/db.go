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
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

// =====================================================================================================================
//
// 	Funciones de conexión a base de datos
//
// =====================================================================================================================

// puntero al pool de conexiones a la base de datos
var DB *pgxpool.Pool

// Función exportada para abrir la conexión
func InitDB(maxConn int) error {
	pgx, err := newPGX(maxConn)
	if err != nil {
		return err
	}
	DB = pgx
	return nil
}

// CloseDBes la encargada de cerrar todo el pool de conexiones
func CloseDB() {
	DB.Close()
}

// newPGX encargada de abrir un nuevo pool de conexiones
func newPGX(maxConn int) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s pool_max_conns=%d",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		maxConn,
	)

	return pgxpool.Connect(context.Background(), dsn)
}
