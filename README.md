## Personal Expense Analyzer

A simple command-line tool written in Go to track and analyze personal expenses.

### Features

- Add expenses with amount, category, and description
- List all recorded expenses
- Generate summary reports
- Generate category breakdown reports

### Installation

1. Ensure you have Go installed (version 1.21 or later)
2. Clone or download the files to a directory
3. Build the tool:

```bash
go build -o expense main.go expense.go
```

### Usage

The tool provides the following commands:

#### Add an Expense
```bash
./expense add <amount> <category> <description>
```
Example:
```bash
./expense add 25.50 food "Lunch at restaurant"
./expense add 45.00 transportation "Bus pass"
```

#### List All Expenses
```bash
./expense list
```

#### Generate Summary Report
```bash
./expense report
```

#### Generate Category Breakdown Report
```bash
./expense report category
```

### Example Session

```bash
# Add some expenses
./expense add 15.99 food "Coffee and sandwich"
./expense add 45.00 transportation "Monthly bus pass"
./expense add 8.50 food "Snacks"
./expense add 29.99 entertainment "Movie ticket"

# List all expenses
./expense list

# Generate reports
./expense report
./expense report category
```

### Notes

- Expenses are stored in memory and will be lost when the program exits
- For persistent storage, you would need to add file-based or database storage
- Amounts should be provided as numbers (e.g., "25.50" not "$25.50")
- Categories are case-sensitive

### Future Enhancements

- Add persistent storage (JSON file, database)
- Add date filtering for reports
- Add export to CSV functionality
- Add budget tracking and alerts
- Add monthly/yearly reporting

This tool provides a solid foundation for personal expense tracking with clean, modular Go code that's easy to extend.