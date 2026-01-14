# Transport Backend MVP

Backend MVP en Go para app de transporte con rutas fijas A→B y paradas intermedias.

## Overview

- **Purpose**: API REST para gestionar rutas de transporte y viajes de pasajeros
- **Stack**: Go 1.21, Chi router, PostgreSQL (Neon), pgxpool
- **Target**: Deploy en Vercel (Go serverless) con frontend Expo (TypeScript)

## Current State

MVP funcional con todos los endpoints implementados:
- GET /health - Estado del servidor
- GET /routes - Lista de rutas activas
- GET /routes/{id} - Detalle de ruta con paradas
- POST /trips - Crear viaje
- GET /trips/{id} - Obtener viaje

## Project Architecture

```
├── api/
│   └── index.go          # Entry point Vercel (Handler function)
├── internal/
│   ├── db/
│   │   └── db.go         # Pool singleton pgxpool serverless-friendly
│   ├── httpx/
│   │   └── respond.go    # Respuestas JSON estándar
│   ├── models/
│   │   └── models.go     # Structs de datos
│   ├── routes/
│   │   └── handlers.go   # Handlers GET /routes, GET /routes/{id}
│   └── trips/
│       └── handlers.go   # Handlers POST /trips, GET /trips/{id}
├── sql/
│   ├── schema.sql        # DDL + enums + triggers
│   └── seed.sql          # Data inicial con UUIDs fijos
├── main.go               # Entry point local (puerto 5000)
├── go.mod                # Dependencias Go
└── vercel.json           # Config Vercel deployment
```

## Key Design Decisions

1. **Usuario hardcodeado**: `11111111-1111-1111-1111-111111111111` (MVP sin auth)
2. **Payment methods**: Solo `cash`, `yape`, `plin`
3. **Rutas fijas**: 2 rutas con 3 paradas cada una
4. **Sin prefijo /api**: Vercel lo agrega automáticamente

## Database

PostgreSQL con las siguientes tablas:
- `users` - Pasajeros
- `routes` - Rutas fijas A→B
- `stops` - Paradas intermedias por ruta
- `trips` - Viajes solicitados

Enums: `payment_method`, `trip_status`

## Recent Changes

- 2026-01-14: Implementación inicial completa del MVP
  - Creado schema SQL con enums y triggers
  - Implementados todos los handlers
  - Configurado pool serverless-friendly
  - Validaciones: payment_method, UUID, pickup != dropoff

## Running Locally

```bash
go run main.go
```

El servidor escucha en puerto 5000.

## Environment Variables

- `DATABASE_URL` - Connection string PostgreSQL (requerido)
- `PORT` - Puerto del servidor (default: 5000)

## Next Steps (Post-MVP)

1. Autenticación real (JWT)
2. Sistema de pagos online
3. Tracking en tiempo real
4. Endpoints para conductores
5. Notificaciones push
