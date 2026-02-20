# Design Document – Student Grade Management System API

This document explains the design decisions, architecture, and implementation approach used for the Student Grade Management System API.

## Design Goals

The primary goals of this system are:

- Implement strict role-based access control
- Follow all given technical constraints
- Use only the Go standard library
- Keep the system simple, transparent, and easy to explain
- Avoid databases, authentication, and external dependencies
- Ensure the project is suitable for academic evaluation and live demonstration

## High-Level Architecture

The application follows a layered architecture:

HTTP Request  
→ Role Enforcement Middleware  
→ Role-Specific Handlers (Admin / Teacher / Student)  
→ In-Memory Storage  

Each layer has a single responsibility, which improves clarity and maintainability.

## Role Enforcement Strategy

Authentication mechanisms were explicitly disallowed.  
Therefore, role enforcement is implemented using HTTP request headers.

### Enforced Headers

- X-Role
- X-User-ID

The middleware validates:
- Presence of required headers
- Valid role values (Admin, Teacher, Student)
- Access permissions for each endpoint

Requests failing validation are rejected before reaching handlers.

## Role Enforcement Strategy

Authentication mechanisms were explicitly disallowed.  
Therefore, role enforcement is implemented using HTTP request headers.

### Enforced Headers

- X-Role
- X-User-ID

The middleware validates:
- Presence of required headers
- Valid role values (Admin, Teacher, Student)
- Access permissions for each endpoint

Requests failing validation are rejected before reaching handlers.

## GPA Calculation

GPA calculation logic is isolated in a utility layer.

Steps:
1. Convert numeric scores to grade points
2. Compute the arithmetic mean
3. Return GPA as a float value

This separation keeps business rules clean and testable.

## Error Handling Strategy

- No panics are used
- All invalid requests return proper HTTP status codes
- Errors are handled as early as possible
- Responses are consistent across endpoints

## Conclusion

This design prioritizes correctness, clarity, and strict adherence to constraints.

The system is intentionally simple, deterministic, and easy to explain, making it well suited for academic evaluation and technical discussion.