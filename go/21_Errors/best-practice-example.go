package main

import (
    "errors"
    "fmt"
)


// ##### Error handling Best Practice
// - Check errors immediately
// - Wrap context around errors : Use fmt.Errorf("context: %w", err) to provide context when propagating errors. It helps in debugging.
// - Use sentinel errors


// ##### Common Pitfalls
// - Silent error handling: Make sure every error is either handled, returened, or logged
// - Misusing defer and recover:  only use recover in exceptional situations where the program can continue meaningfully
// Inconsistent error messages: Error messages should be very clear and concise


var ErrNotFound = errors.New("resource not found")

type ValidationError struct {
    Field string
    Msg   string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Msg)
}

func processData(id int) error {
    if id == 0 {
        return fmt.Errorf("invalid ID %d: %w", id, ErrNotFound)
    } else if id < 0 {
        return &ValidationError{Field: "id", Msg: "must be positive"}
    }
    return nil
}

func main() {
    if err := processData(-1); err != nil {
        var vErr *ValidationError
        if errors.As(err, &vErr) {
            fmt.Println("Validation error:", vErr)
        } else if errors.Is(err, ErrNotFound) {
            fmt.Println("Not found error:", err)
        } else {
            fmt.Println("General error:", err)
        }
    }
}


// - sentinel errors (ErrNotFound) should be the last error in the error chain
// - Custom error types (ValidatorError)
// - Error unwrapping and checking (errors.Is and errors.As)
