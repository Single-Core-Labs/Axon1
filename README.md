# Axon 🚀

**Axon** is a high-performance, open-source AI Gateway and Observability platform built for scaling AI applications with confidence. It acts as a resilient middleman between your application and various AI providers (OpenAI, Anthropic, Gemini, Mistral, etc.), offering unified APIs, automatic failover, and deep performance insights.

---

## ✨ Key Features

- **🛡️ Unified AI Gateway**: One API to rule them all. Access multiple LLM providers through a single endpoint.
- **🔄 Smart Fallback & Retries**: Automatically switch to a backup provider if your primary goes down.
- **📊 Real-time Observability**: Track costs, latency, and success rates in a beautiful dashboard.
- **🔍 Trace Explorer**: Deep-dive into prompt/response logs for debugging and optimization.
- **⚡ High Performance**: Low-latency core built in Go with asynchronous ClickHouse tracing.
- **📦 Multi-SDK**: Native support for Python and TypeScript/JavaScript.

---

## 🏗️ Architecture & Stack

Axon is built as a robust monorepo:

- **Gateway (`apps/gateway`)**: Go + Fiber + ClickHouse (Observability).
- **Platform API (`apps/api`)**: Go + Fiber + Postgres (Metadata).
- **Dashboard (`apps/dashboard`)**: Next.js 15 + Tailwind CSS + Recharts.
- **SDKs (`packages/`)**:
  - `axon-ai` (Python Pydantic v2)
  - `@1corelabs/axon` (TypeScript/ESM/CJS)

---

## 🚀 Getting Started

### Prerequisites

- Docker & Docker Compose
- Go 1.22+ (for local dev)
- Node.js 20+ (for dashboard dev)

### Quick Start with Docker

```bash
git clone https://github.com/Single-Core-Labs/Axon1.git
cd Axon
docker-compose up -d
```

Your Axon instance will be available at:
- **Dashboard**: `http://localhost:3000`
- **Gateway**: `http://localhost:8080`
- **API**: `http://localhost:8081`

---

## 🛠️ Development

### Monorepo Management
This project uses **Turbo** for monorepo orchestration.

```bash
# Install dependencies
npm install

# Build all apps and packages
npx turbo build

# Run tests across the monorepo
npx turbo test
```

### Environment Setup
Copy the `.env.example` file to `.env` and fill in your provider API keys.

```bash
cp .env.example .env
```

---

## 📜 License

Distributed under the **MIT License**. See `LICENSE` for more information.

---

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

Build with ❤️ by **1Core Labs**.
