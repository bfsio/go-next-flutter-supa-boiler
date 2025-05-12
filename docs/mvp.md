
# 🪨 Pet Rock – Architectural & MVP Goals

## 🎯 MVP Architectural Goals

### 🏘️ 1. Multi-Tenant via Wildcard Subdomains

- **Behavior**: Each tenant is accessed through a unique subdomain (e.g., `acme.rockthepet.com`)
- **Mapping**: Subdomains are mapped 1:1 to PostgreSQL schemas (`acme`, `beta`, etc.)
- **Backend Middleware**: Extracts subdomain from request → sets `search_path` for schema isolation
- **Benefit**: Strong isolation, native PostgreSQL multi-tenancy, backup/restore flexibility

---

### 🔐 2. Centralized Authentication (`auth.rockthepet.com`)

- **Supabase Hosted Auth**: All auth flows (login, signup, password reset) occur via shared domain
- **JWT Tokens**: Passed to the main application via secure headers
- **Reasoning**:
  - Avoids auth repetition per tenant
  - Cleanly separates identity concerns
  - Reduces surface area for security issues
- **UX Consideration**: Smooth redirect and return to proper tenant after login

---

### 🚧 3. No Cross-Tenant Data Access

- Tenants are **hard-separated** at:
  - Routing layer (subdomains)
  - DB schema (`search_path`)
  - Application logic (each request scoped to tenant)
- Ensures compliance with **zero-trust** and **SOC2/GDPR**-friendly data segmentation

---

### 🧩 4. Per-Tenant Business Logic (Pluggable Pattern)

- **Mechanism**: Tenants may override default behavior by placing a file in:
  ```
  /apps/backend/tenants/{tenant}.go
  ```
- **Convention**:
  - File name matches subdomain
  - Function signature must match a known interface

- **Example**:
  ```go
  func AcmeOverride(ctx context.Context, input RockInput) (RockOutput, error) { ... }
  ```

- **⚠️ Open Discussion**:
  - This resembles a *plugin system* but lacks isolation.
  - Might violate **Open/Closed Principle** and increase runtime complexity.
  - Consider moving to a `plugin` interface or gRPC microservice per tenant if needed later.

---

### 🏭 5. Tenant Creation via Factory Pattern

- **Purpose**: Encapsulate all steps of tenant provisioning in a single call
- **Process**:
  1. Validate tenant subdomain
  2. Create new PostgreSQL schema
  3. Run migration scripts (`up`) into new schema
  4. Seed default tenant-specific data

- **Interface**:
  ```go
  func CreateTenant(tenantName string) error
  ```

- **Why Factory?**:
  - Keeps creation logic isolated and testable
  - Easy to plug into CLI, web admin, or automation

---

### 🗃️ 6. Disaster Recovery & Migrations

- **Migration Tooling**:
  - `migrate up/down` support for:
    - Core (shared) schema
    - All tenants (looped per schema)
    - Selected list of tenants (CLI-flagged)
  - Scripts stored under:
    ```
    /apps/backend/migrations/{core|tenants}
    ```

- **Seed Scripts**:
  - Support for fresh and incremental seeds
  - Can bootstrap default users, roles, rocks, etc.

- **Recovery Scenarios**:
  - ✅ Full app with no tenant data
  - ✅ Full app with tenant data
  - ✅ Only tenant(s)
  - ✅ Selective recovery by tenant list

---

### 🎨 7. Unified Design Language Across Web and Mobile

- **Goals**:
  - Consistent UI/UX across Next.js (Web) and Flutter (Mobile)
  - Inclusive design for accessibility and branding
  - Unified design tokens: spacing, color, typography

- **Strategy**:
  - Use **Figma** as a shared design source of truth
  - Extract tokens via plugins like **Figma Tokens**
  - Use Tailwind + Flutter equivalents:
    - Tailwind: web atomic design
    - Flutter: `ThemeData` mapped to Figma tokens

- **Design Tools**:
  - [Figma Tokens Plugin](https://www.figma.com/community/plugin/843461159747178978/Figma-Tokens)
  - [Style Dictionary](https://amzn.github.io/style-dictionary/#/)
  - Optional: ThemeSync, Flutter’s [figma-to-flutter](https://pub.dev/packages/figma_to_flutter)

- **Open Consideration**:
  - Would a shared design-tokens repository (`packages/tokens`) be valuable for auto-export to both environments?

---

## 📌 Summary of Key Design Advantages

| Design Element             | Rationale                                                                 |
|----------------------------|---------------------------------------------------------------------------|
| **Wildcard Subdomains**    | Enables clean tenant routing and isolation                                |
| **Schema-per-Tenant**      | Strong security, backup/recovery, operational clarity                     |
| **Supabase Auth Only**     | Delegates auth, simplifies implementation, separates concerns             |
| **Per-Tenant Logic**       | Early-stage flexibility for customization (to be reviewed later)          |
| **Factory Pattern**        | Clean tenant provisioning, reusability across CLI and backend             |
| **Unified Design Language**| Supports responsive and accessible UI across platforms                    |

---

## 🧠 Open Discussions

1. 🔄 Is per-tenant business logic better served via external services or plugins?
2. 🎨 Should we extract and version design tokens in a shared library?
3. 🧪 Should migrations and seeds be dry-runnable for CI tests?
4. 🔌 Should we abstract DB access with `sqlc`, `gorm`, or go full PGX?

---w