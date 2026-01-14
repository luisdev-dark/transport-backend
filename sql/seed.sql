-- Seed data for Transport App MVP
-- Fixed UUIDs for consistent testing

-- Passenger user (MVP hardcoded user)
-- UUID: 11111111-1111-1111-1111-111111111111
INSERT INTO users (id, name, phone, email)
VALUES (
    '11111111-1111-1111-1111-111111111111',
    'Passenger MVP',
    '+51999999999',
    'passenger@mvp.test'
) ON CONFLICT (id) DO NOTHING;

-- Route 1: Lima Centro → Miraflores
-- UUID: 22222222-2222-2222-2222-222222222222
INSERT INTO routes (id, name, origin_name, destination_name, base_price_cents, currency)
VALUES (
    '22222222-2222-2222-2222-222222222222',
    'Ruta Lima Centro - Miraflores',
    'Lima Centro',
    'Miraflores',
    500,
    'PEN'
) ON CONFLICT (id) DO NOTHING;

-- Route 2: San Isidro → Barranco
-- UUID: 33333333-3333-3333-3333-333333333333
INSERT INTO routes (id, name, origin_name, destination_name, base_price_cents, currency)
VALUES (
    '33333333-3333-3333-3333-333333333333',
    'Ruta San Isidro - Barranco',
    'San Isidro',
    'Barranco',
    400,
    'PEN'
) ON CONFLICT (id) DO NOTHING;

-- Stops for Route 1 (Lima Centro → Miraflores)
-- Stop 1: Plaza San Martin (origin)
-- UUID: aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa
INSERT INTO stops (id, route_id, name, stop_order, latitude, longitude)
VALUES (
    'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa',
    '22222222-2222-2222-2222-222222222222',
    'Plaza San Martin',
    1,
    -12.0519,
    -77.0349
) ON CONFLICT (id) DO NOTHING;

-- Stop 2: Av. Arequipa (intermediate)
-- UUID: bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb
INSERT INTO stops (id, route_id, name, stop_order, latitude, longitude)
VALUES (
    'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb',
    '22222222-2222-2222-2222-222222222222',
    'Av. Arequipa',
    2,
    -12.0889,
    -77.0359
) ON CONFLICT (id) DO NOTHING;

-- Stop 3: Ovalo Miraflores (destination)
-- UUID: cccccccc-cccc-cccc-cccc-cccccccccccc
INSERT INTO stops (id, route_id, name, stop_order, latitude, longitude)
VALUES (
    'cccccccc-cccc-cccc-cccc-cccccccccccc',
    '22222222-2222-2222-2222-222222222222',
    'Ovalo Miraflores',
    3,
    -12.1199,
    -77.0289
) ON CONFLICT (id) DO NOTHING;

-- Stops for Route 2 (San Isidro → Barranco)
-- Stop 1: Parque El Olivar (origin)
-- UUID: dddddddd-dddd-dddd-dddd-dddddddddddd
INSERT INTO stops (id, route_id, name, stop_order, latitude, longitude)
VALUES (
    'dddddddd-dddd-dddd-dddd-dddddddddddd',
    '33333333-3333-3333-3333-333333333333',
    'Parque El Olivar',
    1,
    -12.0989,
    -77.0369
) ON CONFLICT (id) DO NOTHING;

-- Stop 2: Larcomar (intermediate)
-- UUID: eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee
INSERT INTO stops (id, route_id, name, stop_order, latitude, longitude)
VALUES (
    'eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee',
    '33333333-3333-3333-3333-333333333333',
    'Larcomar',
    2,
    -12.1319,
    -77.0279
) ON CONFLICT (id) DO NOTHING;

-- Stop 3: Plaza Barranco (destination)
-- UUID: ffffffff-ffff-ffff-ffff-ffffffffffff
INSERT INTO stops (id, route_id, name, stop_order, latitude, longitude)
VALUES (
    'ffffffff-ffff-ffff-ffff-ffffffffffff',
    '33333333-3333-3333-3333-333333333333',
    'Plaza Barranco',
    3,
    -12.1489,
    -77.0229
) ON CONFLICT (id) DO NOTHING;
