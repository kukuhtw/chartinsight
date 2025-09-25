# frontend/Dockerfile
# --- build stage ---
FROM node:20-alpine AS build
WORKDIR /app

# salin manifest & lockfiles (kalau ada)
COPY package.json pnpm-lock.yaml* yarn.lock* package-lock.json* ./
RUN corepack enable

# install deps (pnpm > yarn > npm)
RUN if [ -f pnpm-lock.yaml ]; then pnpm install --frozen-lockfile; \
    elif [ -f yarn.lock ]; then yarn install --frozen-lockfile; \
    elif [ -f package-lock.json ]; then npm ci; \
    else npm install --no-audit --no-fund; fi

# salin source lalu build
COPY . .
RUN if [ -f pnpm-lock.yaml ]; then pnpm build; \
    elif [ -f yarn.lock ]; then yarn build; \
    else npm run build; fi

# --- runtime stage (nginx) ---
FROM nginx:1.27-alpine
WORKDIR /usr/share/nginx/html
COPY --from=build /app/dist ./
# gunakan proxy & SPA fallback (butuh file nginx.conf di project)
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
