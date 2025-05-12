# 🪨 Pet Rock – Architecture & Domain Model

## ✅ Architectural Decisions

### 1. Monorepo
- **Tooling**: `pnpm` workspaces
- **Structure**:
  - `apps/` – application code (backend, web, mobile)
  - `packages/` – shared libraries, UI components
  - `scripts/` – devops and CLI scripts
- **Benefits**:
  - Centralized dependency management
  - Easier refactoring and deployment
  - Unified tooling

---

### 2. Authentication
- **Provider**: Supabase Auth
- **Auth method**: Email/password or OAuth with JWT
- **JWT validation**: Handled in Go backend using Supabase public key
- **Authorization**:
  - Enforced at the DB level via RLS
  - Context passed from frontend → backend → DB

---

### 3. Backend API
- **Language**: Go
- **Framework**: Gin
- **Database**: PostgreSQL
- **Migration Management**: Go-based scripts (e.g., `golang-migrate` or custom)
- **Multi-tenancy**:
  - Tenants mapped to PostgreSQL **schemas**
  - Requests routed via **wildcard subdomains** (e.g., `tenant1.petrock.app`)
  - Gin middleware extracts subdomain and sets PostgreSQL `search_path`
- **RLS**:
  - Enforced via PostgreSQL policies
  - Each query context-aware to authenticated user

---

### 4. Web Frontend
- **Framework**: Next.js (via Vite)
- **Languages**: TypeScript
- **Styling**: TailwindCSS
- **UI Library**: ShadCN UI
- **Auth Integration**: Supabase JS client
- **Routing**:
  - Wildcard subdomain detection for tenant context
  - Supabase session passed via headers to backend

---

### 5. Mobile Frontend
- **Framework**: Flutter
- **Platforms**: iOS, Android, Web
- **Auth**: Supabase Flutter SDK
- **Backend Integration**: REST or gRPC (TBD)

---

### 6. DevOps & Tooling
- **Local Dev**: Docker Compose (PostgreSQL, Supabase emulator optional)
- **Env Management**: `.env` files
- **DX Tooling**:
  - `Makefile` or shell scripts for common commands (`make dev`, `make migrate-up`)
  - Potential GitHub Actions workflows for CI/CD

---

## 🧠 Domain Model (MVP)

### 🧑‍🤝‍🧑 Users

Each user is authenticated by Supabase, but user metadata is mirrored inside the PostgreSQL schema for access control.

| Field        | Type      | Notes                          |
|--------------|-----------|--------------------------------|
| `id`         | UUID      | Supabase Auth UID              |
| `email`      | TEXT      | Auth field                     |
| `tenant_id`  | TEXT      | From `subdomain`               |
| `role`       | TEXT      | `admin`, `user`, etc.          |

---

### 🪨 Rocks

The core item in the app. Each rock is owned by a user and isolated to their tenant.

| Field        | Type      | Notes                          |
|--------------|-----------|--------------------------------|
| `id`         | UUID      | Primary Key                    |
| `name`       | TEXT      | Rock name                      |
| `type`       | TEXT      | e.g., lava, quartz, granite    |
| `owner_id`   | UUID      | Foreign key → `users.id`       |
| `created_at` | TIMESTAMP | Timestamp                      |

---

### 📝 Activity Log (Optional)

For auditing or logs per tenant.

| Field        | Type      | Notes                          |
|--------------|-----------|--------------------------------|
| `id`         | UUID      | Primary Key                    |
| `user_id`    | UUID      | Who performed the action       |
| `action`     | TEXT      | e.g., `added_rock`, `deleted`  |
| `data`       | JSONB     | Optional metadata              |
| `timestamp`  | TIMESTAMP |                                |

---

### 🔐 Row-Level Security (RLS)

- Each PostgreSQL **schema** includes:
  - `users`, `rocks`, `activity_log`, etc.
- Access control logic:
  - Backend extracts tenant from subdomain (e.g., `tenant1` from `tenant1.petrock.app`)
  - Middleware sets:  
    ```sql
    SET search_path TO tenant1;
    ```
  - Policies enforce user-specific access, e.g.:
    ```sql
    CREATE POLICY "Only view own rocks"
    ON rocks
    FOR SELECT
    USING (owner_id = auth.uid());
    ```

---

## 📁 Project Structure

```
pet-rock/
│
├── apps/
│ ├── backend/ # Go (Gin) API service — includes migrations
│ ├── web/ # Next.js frontend (Vite, Tailwind, ShadCN, TS)
│ └── mobile/ # Flutter app
│
├── packages/ # Shared code and utilities
│ ├── ts-config/ # Shared TypeScript config
│ └── ui/ # Shared ShadCN components
│
├── scripts/ # Dev/ops scripts (e.g., DB setup, tenant init)
│
├── .env # Root env config
├── docker-compose.yml # Local dev services
├── Makefile # Task runner
└── pnpm-workspace.yaml # Monorepo configuration
```