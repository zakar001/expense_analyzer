## Personal Expense Analyzer

A simple command-line tool written in Go to track and analyze personal expenses.

### Features

- Add expenses with category, amount, and description
- Generate detailed expense reports (all expenses or filtered by category)
- View spending summary with category breakdown and percentages
- Automatic date tracking

### How to Use

1. **Build the tool:**
   ```bash
   go build -o expense-analyzer
   ```

2. **Add an expense:**
   ```bash
   ./expense-analyzer add <category> <amount> <description>
   ```
   Example:
   ```bash
   ./expense-analyzer add food 25.50 "Lunch at restaurant"
   ./expense-analyzer add transport 15.00 "Bus fare"
   ./expense-analyzer add entertainment 45.00 "Movie tickets"
   ```

3. **Generate a report:**
   ```bash
   # All expenses
   ./expense-analyzer report
   
   # Expenses for a specific category
   ./expense-analyzer report food
   ```

4. **View spending summary:**
   ```bash
   ./expense-analyzer summary
   ```

### Example Usage Session

```bash
# Add some expenses
./expense-analyzer add food 25.50 "Lunch"
./expense-analyzer add transport 15.00 "Bus"
./expense-analyzer add food 8.75 "Coffee"
./expense-analyzer add entertainment 45.00 "Movie"

# Generate reports
./expense-analyzer report
./expense-analyzer report food
./expense-analyzer summary
```

### Categories

You can use any category names you prefer. Common examples:
- `food`, `transport`, `entertainment`, `utilities`, `shopping`, `health`

### Notes

- Expenses are stored in memory and will be lost when the program exits
- For persistent storage, you would need to add file/database integration
- All amounts should be positive numbers
- Dates are automatically set to the current date/time when adding expenses

This tool provides a solid foundation that can be extended with features like monthly reports, expense editing, data persistence, or graphical reports.