# CLI-Automated-Application-Tracker

The CLI-Automated-Application-Tracker is a command-line tool that automatically scrapes your Gmail inbox for job and internship applications, stores them in a PostgreSQL database, and lets you visualize the results for better tracking of your applications.
This project helps you keep an organized, queryable record of where you’ve applied, so you don’t lose track of your progress.

# Features

🔍 Scrape Gmail: Automatically detect job application confirmation emails\n
🗄 Database Storage: Save results in a PostgreSQL-hosted database\n
📊 Visualization: Query and visualize your application history\n

# Example Visualizations

⚡ Fast & Automated: No more manual spreadsheets or sticky notes
Example visualization #1:
<img width="1814" height="1080" alt="image" src="https://github.com/user-attachments/assets/8f90cc7d-b1e1-42a3-bc5e-9b90a2927fe0" />

Example Visualization #2:
<img width="1513" height="1033" alt="image" src="https://github.com/user-attachments/assets/870ab181-9ead-4dbe-9087-733100a17d0a" />

# Requirements
Go (backend scraper logic)
PostgreSQL (local or hosted)
Gmail API credentials (OAuth2)

# Setup

1. Clone the repository

git clone https://github.com/<your-username>/CLI-Automated-Application-Tracker.git
cd CLI-Automated-Application-Tracker


2. Install dependencies

go mod tidy


3. Set up PostgreSQL
   
Create a database and update your .env:
DATABASE_URL=postgres://username:password@localhost:5432/applications


4. Enable Gmail API

Go to Google Cloud Console
Create OAuth2 credentials
Download credentials.json and place it in your project root

5. Run the scraper
   
go run client.go


6. Visualize results
   
You can query via psql or plug into Grafana / pgAdmin.

# Contributing

Contributions are welcome! 🎉
Here’s how you can help:

Fork the repository
Create a feature branch (git checkout -b feature/my-feature)
Commit your changes (git commit -m 'Add some feature')
Push to your branch (git push origin feature/my-feature)
Open a Pull Request
If you’re not sure where to start, check the Issues tab for open tasks or feature requests.

# License

This project is licensed under the MIT License. See the LICENSE
 file for details.

# Code of Conduct

Please be respectful when contributing. By participating in this project, you agree to abide by the Contributor Covenant Code of Conduct.
