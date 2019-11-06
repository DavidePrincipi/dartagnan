<?php

/*
 * Copyright (C) 2019 Nethesis S.r.l.
 * http://www.nethesis.it - nethserver@nethesis.it
 *
 * This script is part of Dartagnan.
 *
 * Dartagnan is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License,
 * or any later version.
 *
 * Dartagnan is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Dartagnan.  If not, see COPYING.
 */

//
// Test the automatic tier assignment procedure.
//
// 1. Read a list of system_id keys from standard input
// 2. Print the system_id with resulting tier_id
//

while($system_id = trim(fgets(STDIN))) {

    $hash = 0;

    foreach(str_split($system_id) as $c) {
        $hash += ord($c);
    }
    $hash = $hash % 10;
    if($hash < 1) { // 10%
        $tier_id = 0;
    } elseif($hash < 3) { // +20% = 30%
        $tier_id = 1;
    } else { // +70% = 100%
        $tier_id = 2;
    }

    echo $system_id . ' ' . $tier_id . "\n";
}
