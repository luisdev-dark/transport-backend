package routes

import (
	"net/http"
	"regexp"

	"transport-backend/internal/db"
	"transport-backend/internal/httpx"
	"transport-backend/internal/models"

	"github.com/go-chi/chi/v5"
)

var uuidRegex = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

func isValidUUID(s string) bool {
	return uuidRegex.MatchString(s)
}

func ListRoutes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pool, err := db.GetPool(ctx)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, "database connection error")
		return
	}

	query := `
		SELECT id, name, origin_name, destination_name, base_price_cents, currency, is_active, created_at, updated_at
		FROM routes
		WHERE is_active = true
		ORDER BY name
	`

	rows, err := pool.Query(ctx, query)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, "failed to fetch routes")
		return
	}
	defer rows.Close()

	var routes []models.Route
	for rows.Next() {
		var route models.Route
		err := rows.Scan(
			&route.ID,
			&route.Name,
			&route.OriginName,
			&route.DestinationName,
			&route.BasePriceCents,
			&route.Currency,
			&route.IsActive,
			&route.CreatedAt,
			&route.UpdatedAt,
		)
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "failed to scan route")
			return
		}
		routes = append(routes, route)
	}

	if routes == nil {
		routes = []models.Route{}
	}

	httpx.JSON(w, http.StatusOK, routes)
}

func GetRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	routeID := chi.URLParam(r, "id")

	if !isValidUUID(routeID) {
		httpx.Error(w, http.StatusBadRequest, "invalid route id format")
		return
	}

	pool, err := db.GetPool(ctx)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, "database connection error")
		return
	}

	routeQuery := `
		SELECT id, name, origin_name, destination_name, base_price_cents, currency, is_active, created_at, updated_at
		FROM routes
		WHERE id = $1
	`

	var route models.Route
	err = pool.QueryRow(ctx, routeQuery, routeID).Scan(
		&route.ID,
		&route.Name,
		&route.OriginName,
		&route.DestinationName,
		&route.BasePriceCents,
		&route.Currency,
		&route.IsActive,
		&route.CreatedAt,
		&route.UpdatedAt,
	)
	if err != nil {
		httpx.Error(w, http.StatusNotFound, "route not found")
		return
	}

	stopsQuery := `
		SELECT id, route_id, name, stop_order, latitude, longitude, created_at, updated_at
		FROM stops
		WHERE route_id = $1
		ORDER BY stop_order
	`

	rows, err := pool.Query(ctx, stopsQuery, routeID)
	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, "failed to fetch stops")
		return
	}
	defer rows.Close()

	var stops []models.Stop
	for rows.Next() {
		var stop models.Stop
		err := rows.Scan(
			&stop.ID,
			&stop.RouteID,
			&stop.Name,
			&stop.StopOrder,
			&stop.Latitude,
			&stop.Longitude,
			&stop.CreatedAt,
			&stop.UpdatedAt,
		)
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "failed to scan stop")
			return
		}
		stops = append(stops, stop)
	}

	if stops == nil {
		stops = []models.Stop{}
	}

	response := models.RouteWithStops{
		Route: route,
		Stops: stops,
	}

	httpx.JSON(w, http.StatusOK, response)
}
