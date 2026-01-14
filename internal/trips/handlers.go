package trips

import (
        "encoding/json"
        "net/http"
        "regexp"

        "transport-backend/internal/db"
        "transport-backend/internal/httpx"
        "transport-backend/internal/models"

        "github.com/go-chi/chi/v5"
)

const HardcodedPassengerID = "11111111-1111-1111-1111-111111111111"

var uuidRegex = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

func isValidUUID(s string) bool {
        return uuidRegex.MatchString(s)
}

func isValidPaymentMethod(method string) bool {
        return method == "cash" || method == "yape" || method == "plin"
}

func CreateTrip(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()

        var req models.CreateTripRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
                httpx.Error(w, http.StatusBadRequest, "invalid request body")
                return
        }

        if !isValidUUID(req.RouteID) {
                httpx.Error(w, http.StatusBadRequest, "invalid route_id format")
                return
        }

        if !isValidPaymentMethod(req.PaymentMethod) {
                httpx.Error(w, http.StatusBadRequest, "payment_method must be cash, yape, or plin")
                return
        }

        if req.PickupStopID != nil && !isValidUUID(*req.PickupStopID) {
                httpx.Error(w, http.StatusBadRequest, "invalid pickup_stop_id format")
                return
        }

        if req.DropoffStopID != nil && !isValidUUID(*req.DropoffStopID) {
                httpx.Error(w, http.StatusBadRequest, "invalid dropoff_stop_id format")
                return
        }

        if req.PickupStopID != nil && req.DropoffStopID != nil && *req.PickupStopID == *req.DropoffStopID {
                httpx.Error(w, http.StatusBadRequest, "pickup_stop_id and dropoff_stop_id cannot be the same")
                return
        }

        pool, err := db.GetPool(ctx)
        if err != nil {
                httpx.Error(w, http.StatusInternalServerError, "database connection error")
                return
        }

        var routeExists bool
        err = pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM routes WHERE id = $1)", req.RouteID).Scan(&routeExists)
        if err != nil {
                httpx.Error(w, http.StatusInternalServerError, "database error")
                return
        }
        if !routeExists {
                httpx.Error(w, http.StatusBadRequest, "route not found")
                return
        }

        if req.PickupStopID != nil {
                var stopBelongsToRoute bool
                err = pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM stops WHERE id = $1 AND route_id = $2)", *req.PickupStopID, req.RouteID).Scan(&stopBelongsToRoute)
                if err != nil || !stopBelongsToRoute {
                        httpx.Error(w, http.StatusBadRequest, "pickup_stop_id does not belong to this route")
                        return
                }
        }

        if req.DropoffStopID != nil {
                var stopBelongsToRoute bool
                err = pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM stops WHERE id = $1 AND route_id = $2)", *req.DropoffStopID, req.RouteID).Scan(&stopBelongsToRoute)
                if err != nil || !stopBelongsToRoute {
                        httpx.Error(w, http.StatusBadRequest, "dropoff_stop_id does not belong to this route")
                        return
                }
        }

        query := `
                INSERT INTO trips (passenger_id, route_id, pickup_stop_id, dropoff_stop_id, status, payment_method, notes)
                VALUES ($1, $2, $3, $4, 'requested', $5, $6)
                RETURNING id, passenger_id, route_id, pickup_stop_id, dropoff_stop_id, status, payment_method, fare_cents, notes, created_at, updated_at
        `

        var trip models.Trip
        err = pool.QueryRow(ctx, query,
                HardcodedPassengerID,
                req.RouteID,
                req.PickupStopID,
                req.DropoffStopID,
                req.PaymentMethod,
                req.Notes,
        ).Scan(
                &trip.ID,
                &trip.PassengerID,
                &trip.RouteID,
                &trip.PickupStopID,
                &trip.DropoffStopID,
                &trip.Status,
                &trip.PaymentMethod,
                &trip.FareCents,
                &trip.Notes,
                &trip.CreatedAt,
                &trip.UpdatedAt,
        )
        if err != nil {
                httpx.Error(w, http.StatusInternalServerError, "failed to create trip")
                return
        }

        httpx.JSON(w, http.StatusCreated, trip)
}

func GetTrip(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()
        tripID := chi.URLParam(r, "id")

        if !isValidUUID(tripID) {
                httpx.Error(w, http.StatusBadRequest, "invalid trip id format")
                return
        }

        pool, err := db.GetPool(ctx)
        if err != nil {
                httpx.Error(w, http.StatusInternalServerError, "database connection error")
                return
        }

        query := `
                SELECT id, passenger_id, route_id, pickup_stop_id, dropoff_stop_id, status, payment_method, fare_cents, notes, created_at, updated_at
                FROM trips
                WHERE id = $1
        `

        var trip models.Trip
        err = pool.QueryRow(ctx, query, tripID).Scan(
                &trip.ID,
                &trip.PassengerID,
                &trip.RouteID,
                &trip.PickupStopID,
                &trip.DropoffStopID,
                &trip.Status,
                &trip.PaymentMethod,
                &trip.FareCents,
                &trip.Notes,
                &trip.CreatedAt,
                &trip.UpdatedAt,
        )
        if err != nil {
                httpx.Error(w, http.StatusNotFound, "trip not found")
                return
        }

        httpx.JSON(w, http.StatusOK, trip)
}
