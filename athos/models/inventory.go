/*
 * Copyright (C) 2017 Nethesis S.r.l.
 * http://www.nethesis.it - info@nethesis.it
 *
 * This file is part of Icaro project.
 *
 * Icaro is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License,
 * or any later version.
 *
 * Icaro is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Icaro.  If not, see COPYING.
 *
 * author: Edoardo Spadoni <edoardo.spadoni@nethesis.it>
 */

package models

import (
	"time"
)

type Inventory struct {
	ID        int       `db:"id" json:"-"`
	Data      JSONB     `db:"data" json:"data"`
	Timestamp time.Time `db:"timestamp" json:"timestamp"`

	System   System `json:"system"`
	SystemID int    `db:"system_id" json:"-"`
}

type InventoryHistory struct {
	ID        int       `db:"id" json:"-"`
	Data      JSONB     `db:"data" json:"data"`
	Timestamp time.Time `db:"timestamp" json:"timestamp"`

	System   System `json:"system"`
	SystemID int    `db:"system_id" json:"-"`
}

type InventoryJSON struct {
	Data struct {
		SystemID string `json:"lk"`
		Data     JSONB  `json:"data"`
	} `json:"data"`
}
