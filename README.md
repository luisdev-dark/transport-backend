# Transport Backend MVP

Backend MVP en Go para aplicación de transporte con rutas fijas A→B y paradas intermedias.

## 📋 Stack Tecnológico

- **Lenguaje**: Go 1.21
- **Router**: Chi
- **Base de Datos**: PostgreSQL (Neon)
- **Pool de Conexiones**: pgxpool
- **Deploy**: Vercel (Go serverless)
- **Frontend**: Expo (TypeScript)
- **API**: REST JSON

## 🚫 Restricciones del MVP

- Sin autenticación real (user_id hardcodeado)
- Sin pagos online
- Sin tracking en tiempo real
- Sin microservicios
- Enfoque en MVP primero

## 🛣️ Endpoints Disponibles

> **Nota**: Vercel agrega el prefijo `/api` automáticamente. No usar `/api` dentro del router.

- `GET /health` - Estado del servidor
- `GET /routes` - Lista de rutas activas
- `GET /routes/{id}` - Detalle de ruta con paradas
- `POST /trips` - Crear nuevo viaje
- `GET /trips/{id}` - Obtener detalles de viaje

## 📁 Estructura del Proyecto

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
├── go.mod                # Dependencias Go
└── vercel.json           # Config Vercel deployment
```

## 🔑 Reglas Importantes

1. **NO usar prefijo `/api`** dentro del router (Vercel lo agrega)
2. **Usar `r.Context()`** en todos los handlers
3. **Payment methods permitidos**: `cash`, `yape`, `plin`
4. **Siempre responder JSON**
5. **UUIDs como strings**
6. **Usuario hardcodeado**: `11111111-1111-1111-1111-111111111111`

## 🗄️ Base de Datos

### Tablas

- `users` - Pasajeros del sistema
- `routes` - Rutas fijas A→B
- `stops` - Paradas intermedias por ruta
- `trips` - Viajes solicitados por pasajeros

### Enums

- `payment_method`: cash | yape | plin
- `trip_status`: requested | confirmed | in_progress | completed | cancelled

## 🚀 Deploy en Vercel

### Configuración de Variables de Entorno

Debes configurar la siguiente variable en Vercel:

- `DATABASE_URL` - Connection string de PostgreSQL (requerido)

### Comandos de Deploy

```bash
# Instalar Vercel CLI
npm i -g vercel

# Deploy
vercel
```

## 🛠️ Desarrollo Local (Opcional)

Para testing local, puedes crear un archivo `main.go`:

```go
package main

import (
    "log"
    "net/http"
    "os"
    "transport-backend/api"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "5000"
    }
    log.Printf("Starting server on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, http.HandlerFunc(api.Handler)))
}
```

Ejecutar:

```bash
go run main.go
```

## 📝 Formato de Requests

### POST /trips

```json
{
  "route_id": "22222222-2222-2222-2222-222222222222",
  "pickup_stop_id": "44444444-4444-4444-4444-444444444444",
  "dropoff_stop_id": "55555555-5555-5555-5555-555555555555",
  "payment_method": "cash"
}
```

### Response Exitosa

```json
{
  "id": "generated-uuid",
  "passenger_id": "11111111-1111-1111-1111-111111111111",
  "route_id": "22222222-2222-2222-2222-222222222222",
  "pickup_stop_id": "44444444-4444-4444-4444-444444444444",
  "dropoff_stop_id": "55555555-5555-5555-5555-555555555555",
  "status": "requested",
  "payment_method": "cash",
  "created_at": "2026-01-14T00:00:00Z"
}
```

### Response Error

```json
{
  "error": "mensaje descriptivo del error"
}
```

## 🎯 Próximos Pasos (Post-MVP)

1. Implementar autenticación real (JWT)
2. Integrar sistema de pagos online
3. Agregar tracking en tiempo real
4. Crear endpoints para conductores
5. Implementar notificaciones push
6. Agregar geolocalización

## 📄 Licencia

Este proyecto es un MVP para propósitos educativos.
