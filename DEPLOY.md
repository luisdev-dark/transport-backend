# 🚀 Guía de Deploy en Vercel

## ✅ Checklist Pre-Deploy

- [x] Código subido a GitHub: `luisdev-dark/transport-backend`
- [x] `vercel.json` configurado
- [x] `api/index.go` como entry point
- [ ] Variable de entorno `DATABASE_URL` (configurar en Vercel)

## 📝 Paso a Paso

### 1️⃣ Ir a Vercel

1. Abre: **https://vercel.com**
2. Inicia sesión con GitHub

### 2️⃣ Importar Proyecto

1. Click en **"Add New..."** → **"Project"**
2. Busca y selecciona: `transport-backend`
3. Click en **"Import"**

### 3️⃣ Configurar Variables de Entorno (IMPORTANTE)

**En la sección "Environment Variables":**

```
Name:  DATABASE_URL
Value: postgresql://neondb_owner:npg_V3JXSoBD8CmW@ep-proud-cell-ahy9f4zl-pooler.c-3.us-east-1.aws.neon.tech/neondb?sslmode=require
```

**📌 IMPORTANTE:** Marca las 3 opciones:
- ✅ Production
- ✅ Preview
- ✅ Development

### 4️⃣ Deploy

1. Click en **"Deploy"**
2. Espera 1-2 minutos
3. ¡Listo! 🎉

## 🔗 URLs Finales

Después del deploy, Vercel te dará:

### URL de Producción
```
https://transport-backend.vercel.app
```

### Tus Endpoints Funcionando
```
GET  https://transport-backend.vercel.app/api/health
GET  https://transport-backend.vercel.app/api/routes
GET  https://transport-backend.vercel.app/api/routes/{id}
POST https://transport-backend.vercel.app/api/trips
GET  https://transport-backend.vercel.app/api/trips/{id}
```

## 🧪 Probar el Deploy

### Test 1: Health Check
```bash
curl https://transport-backend.vercel.app/api/health
```

Respuesta esperada:
```json
{"status":"ok"}
```

### Test 2: Obtener Rutas
```bash
curl https://transport-backend.vercel.app/api/routes
```

Respuesta esperada:
```json
[
  {
    "id": "22222222-2222-2222-2222-222222222222",
    "name": "Ruta Norte-Sur",
    "origin_name": "Terminal Norte",
    "destination_name": "Terminal Sur",
    ...
  }
]
```

## ⚠️ Si Algo Sale Mal

### Error: No DATABASE_URL
- Ve a **Settings** → **Environment Variables**
- Verifica que `DATABASE_URL` esté configurada
- Click **"Redeploy"**

### Error 500: Database connection failed
- Verifica que tu DATABASE_URL sea correcta
- Asegúrate que incluya `?sslmode=require`
- Revisa los logs en Vercel Dashboard

### Ver Logs en Tiempo Real
1. En Vercel Dashboard → Tu proyecto
2. Click en el deployment más reciente
3. Tab **"Functions"** → Ver logs de errores

## 🎯 Resumen Ultra-Rápido

1. **Vercel.com** → Import `transport-backend`
2. **Environment Variables** → `DATABASE_URL` = tu connection string de Neon
3. **Deploy** → Espera → ✅ Listo!

## 📱 Próximo Paso

Una vez deployed, actualiza tu frontend Expo con la URL:

```typescript
// En tu app Expo
const API_URL = "https://transport-backend.vercel.app/api";

// Ejemplo fetch
const routes = await fetch(`${API_URL}/routes`);
```

---

**¿Dudas?** Todos los endpoints están documentados en el README.md principal.
