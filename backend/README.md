# Questionnaire System Backend

This is the backend service for the Questionnaire System, built with Golang and the Hertz framework.

## Project Structure

```
backend/
├── app/
│   ├── api/        # API handlers
│   ├── model/      # Database models
│   └── service/    # Business logic
├── config/         # Configuration files
├── db/             # Database initialization scripts
├── router/         # Route definitions
├── main.go         # Application entry point
└── go.mod          # Go module file
```

## Getting Started

### Prerequisites

- Go 1.21 or later
- PostgreSQL database

### Running the Application

1. Make sure your PostgreSQL database is running
2. Update the configuration in `config/config.yaml` if needed
3. Run the application:

```bash
go run main.go
```

The server will start on the configured port (default: 8080).

## Data Model

The application uses the following data model:

- **Form**: Represents a questionnaire or survey form
- **Section**: Represents a group of related questions in a form
- **Question**: Represents a question in a section (types: radio, checkbox, text)
- **QuestionOption**: Represents an option for a question with associated score
- **Response**: Represents a user's submission of a form
- **Answer**: Represents a user's answer to a question

## API Endpoints

- `GET /api/v1/health` - Health check endpoint
- `GET /api/v1/forms` - List all forms
- `POST /api/v1/forms` - Create a new form
- `GET /api/v1/forms/:id` - Get a form by ID
- `PUT /api/v1/forms/:id` - Update a form
- `DELETE /api/v1/forms/:id` - Delete a form
- `GET /api/v1/forms/:formId/sections` - List sections of a form
- `POST /api/v1/forms/:formId/sections` - Create a new section in a form
- `GET /api/v1/forms/:formId/sections/:id` - Get a section by ID
- `PUT /api/v1/forms/:formId/sections/:id` - Update a section
- `DELETE /api/v1/forms/:formId/sections/:id` - Delete a section

## Development

To add new features or endpoints, follow these steps:

1. Create models in the `app/model` directory
2. Implement business logic in the `app/service` directory
3. Add API handlers in the `app/api` directory
4. Register routes in the `router` package
