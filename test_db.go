package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"transport-backend/pkg/db"
)

func main() {
	// Cargar DATABASE_URL del ambiente
	// En Windows PowerShell: $env:DATABASE_URL="tu_connection_string"
	// En bash/zsh: export DATABASE_URL="tu_connection_string"

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("❌ DATABASE_URL no está configurada. Usa: $env:DATABASE_URL=\"your_connection_string\"")
	}

	fmt.Println("🔌 Probando conexión a Neon PostgreSQL...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	pool, err := db.GetPool(ctx)
	if err != nil {
		log.Fatalf("❌ Error conectando a la base de datos: %v", err)
	}

	// Verificar conexión
	var version string
	err = pool.QueryRow(ctx, "SELECT version()").Scan(&version)
	if err != nil {
		log.Fatalf("❌ Error ejecutando query: %v", err)
	}

	fmt.Println("✅ Conexión exitosa!")
	fmt.Printf("📊 PostgreSQL version: %s\n", version[:50]+"...")

	// Verificar si existen las tablas del schema
	var tableCount int
	err = pool.QueryRow(ctx, `
		SELECT COUNT(*) 
		FROM information_schema.tables 
		WHERE table_schema = 'public' 
		AND table_name IN ('users', 'routes', 'stops', 'trips')
	`).Scan(&tableCount)

	if err != nil {
		log.Fatalf("❌ Error verificando tablas: %v", err)
	}

	if tableCount == 4 {
		fmt.Println("✅ Todas las tablas del schema existen (users, routes, stops, trips)")
	} else {
		fmt.Printf("⚠️  Solo %d/4 tablas encontradas. Ejecuta schema.sql en Neon.\n", tableCount)
	}

	// Verificar rutas de ejemplo
	var routeCount int
	err = pool.QueryRow(ctx, "SELECT COUNT(*) FROM routes").Scan(&routeCount)
	if err == nil {
		fmt.Printf("✅ Rutas en BD: %d\n", routeCount)
	} else {
		fmt.Println("⚠️  Tabla 'routes' existe pero está vacía. Ejecuta seed.sql")
	}

	fmt.Println("\n🎉 Test de conexión completado!")
}
