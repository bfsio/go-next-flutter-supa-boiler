# 📚 Pet Rock – Architectural & Design Rationale

## 🎯 Overview

The Pet Rock application is a modern, multi-tenant SaaS platform targeting **Android, iOS, and Web**, using a **monorepo**, **multi-language polyglot stack**, and **schema-isolated PostgreSQL** backends. This architecture supports strong domain encapsulation, scalable team ownership, and clean separation of concerns while maintaining developer ergonomics and user security.

---

## ⚙️ Architectural Decisions

### 1. **Monorepo using `pnpm`**

#### ✅ Pros:
- Centralized versioning and dependencies
- Cross-platform coordination (web, mobile, backend)
- Easier refactoring across boundaries
- Supported by most modern tooling (e.g., Turborepo, Vite, Next.js)

#### ❌ Cons:
- Slightly more complex tooling (e.g., workspace hoisting)
- Potentially slower install on CI (mitigated by caching)

#### 🔍 Why:
Monorepos optimize for **team scalability** and **reuse**, aligning with **Open/Closed Principle** (OCP) — shared packages can be extended without modifying individual apps.

---

### 2. **PostgreSQL + Schema-based Multi-tenancy + RLS**

#### ✅ Pros:
- True tenant isolation via `search_path`
- Cleaner backup and migration paths per tenant
- RLS ensures least-privilege, row-level access
- Works well with Supabase’s policy engine

#### ❌ Cons:
- Schema management overhead
- More complex onboarding for new devs

#### 🔍 Why:
This model adheres to the **Single Responsibility Principle** (SRP) at the data layer. Each schema encapsulates tenant data, simplifying reasoning and isolation. RLS handles **authorization declaratively**, shifting security left into the data layer, which is especially important in SaaS.

---

### 3. **Go (Gin) Backend**

#### ✅ Pros:
- Small binary size, blazing fast performance
- First-class concurrency support (goroutines)
- Rich ecosystem (GIN, sqlc, pgx, migrate, etc.)
- Ideal for low-latency, stateless APIs

#### ❌ Cons:
- More verbose than scripting languages
- No generics until recently (Go 1.18+)

#### 🔍 Why:
Go is the right tool for **backend concurrency, schema management, and performance-critical services**. Its static typing and simple interfaces support **Interface Segregation** and **Dependency Inversion** principles. Gin provides minimalist routing with middleware support for JWT and tenant resolution.

---

### 4. **Supabase (Auth only)**

#### ✅ Pros:
- Managed JWT-based auth with session support
- Easy integration with web and mobile clients
- Public key available for server-side validation

#### ❌ Cons:
- Less customizable than building a full auth stack
- Ties some logic to their JWT structure

#### 🔍 Why:
Supabase is used **only for authentication** to decouple identity from business logic. This aligns with **Separation of Concerns** and allows for potential migration or fallbacks. Supabase’s built-in RLS support complements PostgreSQL policies.

---

### 5. **Next.js (w/ Vite, Tailwind, ShadCN UI)**

#### ✅ Pros:
- Next.js: Server-side rendering, routing, and API endpoints out of the box
- Vite: Instant dev refresh, blazing-fast build
- TailwindCSS: Utility-first, responsive-by-default
- ShadCN UI: Headless + accessible UI primitives built on Radix UI

#### ❌ Cons:
- Steeper initial learning curve (Tailwind config, headless UI)
- ShadCN requires manual updates (though customizable)

#### 🔍 Why:
This stack is built for **developer velocity and end-user experience**:
- **Tailwind** enforces consistent design system at scale.
- **ShadCN** offers composition over configuration — you build real components, not override defaults.
- **Vite** + **Next.js** provide the best DX-performance tradeoff, especially for SSR and SEO-critical routes (e.g., landing pages, blog).

Supports **Liskov Substitution Principle**: UI components can be extended, swapped, or themed without rewriting logic.

---

### 6. **Flutter for Mobile (iOS, Android, Web)**

#### ✅ Pros:
- One codebase → three platforms
- Modern UI engine (Skia) for pixel-perfect rendering
- Great tooling and community
- Supabase Flutter SDK available

#### ❌ Cons:
- Larger binary size vs native
- Slower cold start (especially on web)
- Limited third-party native modules (some edge cases)

#### 🔍 Why:
Flutter aligns with your **cross-platform delivery goals** without fragmenting the team. For a V1 MVP, fast UI iteration and brand consistency outweigh native performance. Ideal for CRUD-heavy SaaS apps that don’t need native hardware access.

---

## 🧠 Domain Modeling Choices

### 1. **Tenant-per-schema**

- Clean physical and logical separation of data
- Eliminates cross-tenant leakage risks
- Improves backup/restore, test fixture generation
- Supports strong auditability (per-schema logs)

### 2. **Supabase UID-based user mapping**

- Centralizes identity across all layers
- Auth roles and claims passed via JWT
- Schema logic depends only on tenant ID and UID

### 3. **Rocks & Activity Logs**

- Rocks = primary resource (scoped to user and tenant)
- Activity Logs = eventual audit layer, useful for analytics, undo, moderation

---

## 🔐 Security Posture

| Layer     | Strategy                                                    |
|-----------|-------------------------------------------------------------|
| Client    | Supabase session tokens (JWT), HTTPS-only communication     |
| Backend   | JWT middleware + subdomain → schema resolution              |
| Database  | `search_path` + RLS policies + tenant isolation by schema   |
| CI/CD     | Secrets managed via `.env` or GitHub Actions Secrets        |

---

## 💡 SOLID Principles in Action

| Principle                   | Implementation                                                   |
|----------------------------|------------------------------------------------------------------|
| Single Responsibility       | Backend handles APIs, frontend handles UI, DB handles security  |
| Open/Closed                 | UI and services designed to be extended without rewriting        |
| Liskov Substitution         | Swappable UI components (ShadCN), replaceable auth               |
| Interface Segregation       | Clear domain interfaces in Go and TypeScript                    |
| Dependency Inversion        | Clients depend on contracts, not implementation specifics        |

---

## 🧱 Conclusion

Each decision made here prioritizes:
- 🔒 **Security by design**
- 💨 **Developer velocity**
- 📦 **Modularity and reusability**
- 📈 **Scalability with clear tenant boundaries**
