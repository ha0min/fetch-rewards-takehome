# Receipt Processor

A REST API service that processes receipts and calculates points based on specific rules.

## How to Run

### Prerequisites
- Go 1.23.2 or higher
- Git

### Running the Project

1. Clone the repository:
```bash
git clone https://github.com/ha0min/fetch-rewards-takehome.git
```

2. Navigate to the project directory:
```bash
cd fetch-rewards-takehome
```

3. Run the server:
```bash
go run main.go
```

The server will start on `localhost:8080`.

### Running Tests

To run the unit tests:
```bash
go test ./...
```

## Design Approach

### Chain of Responsibility Pattern

The project implements a variation of the Chain of Responsibility pattern for calculating receipt points. This design choice offers several benefits:

1. **Modularity**: Each rule is encapsulated in its own struct that implements the `Rule` interface:
   ```go
   type Rule interface {
       Calculate(receipt *models.Receipt) int
   }
   ```

2. **Easy to Add New Rules**: New rules can be added by creating a new struct that implements the `Rule` interface. For example:
   ```go
   type RuleTwo struct{}
   func (r *RuleTwo) Calculate(receipt *models.Receipt) int {
       // RuleTwo implementation
   }
   ```

3. **Independent Rule Processing**: Each rule processes the receipt independently and returns its points, making the system more maintainable and testable.

### Project Structure

```plaintext
.
├── api.yml           # OpenAPI specification
├── handlers/         # HTTP request handlers
├── models/          # Data models
└── services/        # Business logic and rules
```

- **handlers**: Contains the HTTP handlers for processing receipts and retrieving points
- **models**: Defines the data structures for receipts and responses
- **services**: Implements the business logic and point calculation rules

### Testing Strategy

The project includes comprehensive unit tests that cover:

1. Individual rule calculations
2. HTTP endpoint functionality
3. End-to-end receipt processing workflow

Key test files:
`handlers/receipts_test.go`: Tests for HTTP handlers and request processing

The tests ensure that:
- Receipt processing returns valid IDs
- Points calculation works correctly
- Error cases are handled appropriately

## API Endpoints

1. Process Receipt
```
POST /receipts/process
```

2. Get Points
```
GET /receipts/{id}/points
```

For detailed API specifications, refer to the `api.yml` file.
```