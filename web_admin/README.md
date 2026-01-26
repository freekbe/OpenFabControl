# Web Admin

Admin panel for OpenFabControl.

## Prerequisites

- Node.js 18+
- npm

## Development

```bash
# Install dependencies
npm install

# Start dev server
npm run dev
```

Access at `http://localhost:5173`

## Build

```bash
# Build for production
npm run build
```

## Linting

```bash
npm run lint
```

## Production Run

The built files are in `dist/` and served by nginx via docker-compose.

```bash
# From web_admin/
npm run build        # Build the app
cd ..
make up              # Start the stack
```

Access at `https://localhost:4080`

## API Endpoints

API requests are proxied by nginx:

- `/web-admin-api/*` → Backend admin routes
- `/web-user-api/*` → Backend user routes
- `/machine-api/*` → Backend machine routes
