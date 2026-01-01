Flip - Spaced Repetition Flashcard App

Flip is a Test Driven Full-Stack web application designed to help users learn and retain information effortlessly using the SuperMemo-2 (SM-2) Spaced Repetition algorithm.

Tech Stack
    Backend
        - Language: Go 1.24
        - Framework: Gin 
        - Database: PostgreSQL
        - ORM: GORM
        - Auth: JWT & Bcrypt
    Frontend
        - Framework: Nuxt 4 (Vue.js)
        - Styling: Tailwind CSS
        - State Management: Pinia
        - Visualization: Chart.js
    Infrastructure
       - Containerization: Docker & Docker Compose
        - Gateway: Nginx 

Features
    - Secure Authentication: Login/Registration system with JWT cookies and password hashing.
    - Deck Management: Create, rename, and delete decks to organize your studies.
    - Spaced Repetition (SM-2): An algorithm that schedules reviews based on your performance.
    - CSV Import: Bulk import flashcards from other CSV files.
    - Statistics: Visual breakdown of your learning progress (New vs. Mature cards) and daily due counts.

Getting Started
    Follow these steps to get this app running on your machine

    Prequisites
        - Docker Desktop
        - Git
    
    Clone the Repository
        Run the following commands:
            - git clone https://github.com/FizzaSadath/flashcard-app.git
            - cd flashcard-app

    Configure Environments:
        - Create .env file in the root directory and paste these contents

            DB_HOST=db
            DB_PORT=5432
            DB_USER=flash_user
            DB_PASSWORD=flash_password
            DB_NAME=flashcard_db
            # Security (Change this to a random string)
            JWT_SECRET=super-secret-key-change-me
        
    Build & Run
        Run the following command:
            - docker compose up --build

    Access the App
        - Frontend: http://localhost
        - Backend: http://localhost/api/ping

How to Use
    1. Create an Account: Click Sign Up to create a new user. 
    2. Create a Deck: Go to the Dashboard and type a topic (e.g., "Web Development") to create a deck.
    3. Add Cards: 
        - Manual: Fill in the front/back and click on "Add Card"
        - Import: Use the Import CSV button. The CSV format must be:

            Front,Back
            What is Docker?,A containerization platform.
            What is a Goroutine?,A lightweight thread managed by the Go runtime.

    4. Study Mode: 
        Click Study on a deck. The app will show you cards that are "Due" today.
        When you reveal the answer, grade yourself:
            - 0-2 (Fail): You forgot. The card will reset to Day 1.
            - 3 (Hard): You remembered with effort.
            - 4 (Good): You remembered comfortably.
            - 5 (Easy): Instant recall. The interval will increase aggressively.
    
    5. View Cards:
        Click Cards on a deck. The app will show you the cards in this deck including your streak, and when to review the card next.

    6. View Stats
        Go to Stats section in the dashboard. View global card statistics as well as the deck-wise breakdown of new, learning and mature cards.
    

    Feel free to reach out to share your suggestions and ideas.

