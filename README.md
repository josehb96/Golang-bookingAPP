# Golang Booking Application

This is a simple booking application built using the Go programming language. The application allows users to book tickets for a conference, with features such as:

- Greeting users and displaying available tickets
- Validating user input for first name, last name, email, and ticket quantity
- Booking tickets and updating the remaining tickets
- Displaying the first names of all bookings

## Features

- **User Input Validation**: The application validates user input to ensure that the first and last names are at least 2 characters long, the email address contains an "@" symbol, and the requested number of tickets is valid and available.
- **Ticket Booking**: Users can book tickets for the conference, and the remaining ticket count is updated accordingly.
- **Booking History**: The application keeps track of all bookings and displays the first names of all attendees.

## Usage

1. Clone the repository:

   ```
   git clone https://github.com/JoseHurtadoBaeza/Golang-bookingAPP-TechWorld-with-Nana.git
   ```

2. Navigate to the project directory:

   ```
   cd Golang-bookingAPP-TechWorld-with-Nana
   ```

3. Run the application:

   ```
   go run main.go
   ```

4. Follow the on-screen instructions to book your tickets.

## Dependencies

This project does not have any external dependencies. It uses the standard Go library.
