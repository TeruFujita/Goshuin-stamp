# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial project setup with Next.js frontend and Go backend
- Docker environment configuration
- MySQL database with temples and goshuin_collections tables
- Basic API endpoints for temples and goshuin collections
- Frontend landing page with Japanese cultural theme
- Responsive design with Tailwind CSS
- Framer Motion animations

### Changed
- Switched from Gin framework to Go standard library (net/http)
- Implemented manual Ent client instead of auto-generated code

### Fixed
- Go 1.24 compatibility issues with Ent schema generation
- Database connection and CRUD operations

## [0.1.0] - 2024-08-11

### Added
- Project initialization
- Basic directory structure
- Docker Compose configuration
- Environment variable setup
- Go backend with standard library HTTP server
- Next.js frontend with TypeScript
- MySQL database integration
- Sample temple data (浅草寺, 明治神宮, 金閣寺)

### Technical Decisions
- Use Go standard library instead of Gin framework
- Manual Ent client implementation for database operations
- Docker-first development environment
- TypeScript + Tailwind CSS for frontend

---

## Version History

- **0.1.0** (2024-08-11): Initial MVP with basic functionality
  - Backend API with 10 endpoints
  - Frontend with 4 main pages
  - Database with 2 tables
  - Docker environment ready
