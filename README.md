## Personal Expense Analyzer Tool

A simple command-line tool written in Go to track and analyze personal expenses.

### Installation

1. Make sure you have Go installed (version 1.16 or higher)
2. Clone or download these files into a directory
3. Build the tool:
   ```bash
   go build -o expense-analyzer
   ```

### Usage

#### Add an Expense
```bash
./expense-analyzer add <amount> <category> <description>
```
Example:
```bash
./expense-analyzer add 25.50 food "Lunch at restaurant"
./expense-analyzer add 45.00 transportation "Bus pass"
./expense-analyzer add 12.99 entertainment "Movie ticket"
```

#### Generate Reports

**Summary Report** (shows spending by category):
```bash
./expense-analyzer report
```

**Detailed Report** (shows all transactions):
```bash
./expense-analyzer report detailed
```

**Quick Summary** (shows overview statistics):
```bash
./expense-analyzer summary
```

### Features

- ✅ Add expenses with amount, category, and description
- ✅ Automatic timestamp recording
- ✅ Summary report with category breakdown and percentages
- ✅ Detailed transaction history
- ✅ Quick statistics overview
- ✅ Error handling for invalid inputs

### Example Session

```bash
$ ./expense-analyzer add 15.00 food "Coffee and snack"
✓ Added expense: $15.00 for food (Coffee and snack)

$ ./expense-analyzer add 50.00 shopping "New shoes"
✓ Added expense: $50.00 for shopping (New shoes)

$ ./expense-analyzer summary

=== QUICK SUMMARY ===
Total spent: $65.00
Transactions: 2
Categories: 2
Highest spending category: shopping ($50.00)

$ ./expense-analyzer report

=== EXPENSE SUMMARY REPORT ===
Total expenses: $65.00
Number of transactions: 2

Spending by category:
  food: $15.00 (23.1%)
  shopping: $50.00 (76.9%)
```

### Notes

- Data is stored in memory and will be lost when the program exits
- For persistent storage, you could extend the tool to save to a file or database
- Categories are case-sensitive
- Amounts should be in decimal format (e.g., 25.50, 100, 15.75)