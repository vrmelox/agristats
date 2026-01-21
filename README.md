# AgriStats - Agricultural Harvest Analysis Tool

A command-line tool written in Go that analyzes harvest data from CSV files to help farmers optimize their planting decisions through yield analysis, seasonal patterns, and environmental correlations.

## 📋 Description

AgriStats processes agricultural harvest data to provide insights on:
- Crop yield performance per hectare
- Optimal planting periods
- Weather pattern correlations (rainfall and temperature)
- Simple predictive analytics for better farming decisions

## 🚀 Features

- 📊 **Yield Analysis**: Calculate average yield per hectare for each crop type
- 📅 **Seasonal Insights**: Identify the best planting months based on historical performance
- 🌧️ **Weather Correlation**: Analyze relationships between rainfall, temperature, and crop yield
- 📈 **Trend Visualization**: ASCII charts for visual data representation
- 📄 **Multiple Output Formats**: Generate reports in text or JSON format
- 🔍 **Flexible Filtering**: Filter data by crop type and year

## 📥 Input Format

### CSV File Structure

The tool expects a CSV file (`harvests.csv`) with the following columns:

```csv
date,crop,hectares,yield_tons,rainfall_mm,temperature_avg
2024-01-15,maize,5.2,18.4,45,28
2024-01-20,cassava,3.1,12.7,52,27
2024-02-10,yam,4.5,13.2,38,29
```

**Column Descriptions:**
- `date`: Harvest date (YYYY-MM-DD format)
- `crop`: Crop type name
- `hectares`: Area harvested in hectares
- `yield_tons`: Total yield in tons
- `rainfall_mm`: Average rainfall in millimeters
- `temperature_avg`: Average temperature in degrees Celsius

## 🖥️ Command-Line Arguments

```bash
agristats -file <path> -crop <name> -year <YYYY> -report <json|text>
```

### Arguments

| Argument | Description | Required | Example |
|----------|-------------|----------|---------|
| `-file` | Path to the CSV file | Yes | `-file harvests.csv` |
| `-crop` | Filter by specific crop name | No | `-crop maize` |
| `-year` | Filter by year | No | `-year 2024` |
| `-report` | Output format (`text` or `json`) | No | `-report json` |

### Usage Examples

```bash
# Analyze all data with text report
./agristats -file harvests.csv

# Analyze only maize crops
./agristats -file harvests.csv -crop maize

# Generate JSON report for 2024
./agristats -file harvests.csv -year 2024 -report json

# Analyze cassava in 2024 with JSON output
./agristats -file harvests.csv -crop cassava -year 2024 -report json
```

## 📤 Output Format

### Text Report

```
═══════════════════════════════════════════
AGRISTATS ANALYSIS REPORT
═══════════════════════════════════════════
Period: 2024-01-01 to 2024-12-31
Total Harvests: 47

TOP 3 CROPS BY YIELD:
1. Maize: 3.54 tons/hectare (avg)
2. Cassava: 4.10 tons/hectare (avg)
3. Yam: 2.89 tons/hectare (avg)

BEST PLANTING MONTH: March (avg yield +23%)

RAINFALL CORRELATION: 0.67 (strong positive)
```

### JSON Report

When using `-report json`, a `report.json` file is generated with structured data:

```json
{
  "period": {
    "start": "2024-01-01",
    "end": "2024-12-31"
  },
  "total_harvests": 47,
  "top_crops": [
    {
      "crop": "maize",
      "avg_yield_per_hectare": 3.54
    },
    {
      "crop": "cassava",
      "avg_yield_per_hectare": 4.10
    },
    {
      "crop": "yam",
      "avg_yield_per_hectare": 2.89
    }
  ],
  "best_month": {
    "month": "March",
    "yield_improvement": 23
  },
  "correlations": {
    "rainfall": 0.67,
    "temperature": -0.12
  }
}
```

### ASCII Charts

Visual representation of trends and patterns in the data.

## 🏗️ Project Structure

```
agristats/
├── main.go           # Entry point and CLI argument handling
├── harvest.go        # Harvest struct and methods
├── parser.go         # CSV file reading and parsing
├── analyzer.go       # Statistical calculations
├── reporter.go       # Report generation (text/JSON)
├── go.mod            # Go module file
└── testdata/
    └── sample.csv    # Sample data for testing
```

## 🎓 Go Concepts Used

This project demonstrates beginner-level Go concepts:

- ✅ Basic types (`string`, `int`, `float64`, `bool`)
- ✅ Structs for data modeling
- ✅ Slices for data collections
- ✅ Maps for data aggregation (crop → statistics)
- ✅ Loops (`for`, `range`)
- ✅ Conditionals (`if/else`, `switch`)
- ✅ Functions and parameters
- ✅ File operations (`os.Open`, `bufio.Scanner`)
- ✅ CSV parsing (`encoding/csv`)
- ✅ Error handling (`if err != nil`)
- ✅ CLI arguments (`flag` package)
- ✅ String formatting (`fmt.Printf`)
- ✅ Type conversions (`strconv`)

## 🛠️ Installation

### Prerequisites

- Go 1.20 or higher

### Build from Source

```bash
# Clone the repository
git clone https://github.com/yourusername/agristats.git
cd agristats

# Build the binary
go build -o agristats

# Run the tool
./agristats -file harvests.csv
```

## ✅ Success Criteria

1. ✅ **Parse Performance**: Correctly processes 1000+ CSV lines
2. ✅ **Accurate Statistics**: Precise calculations for averages, max, and min values
3. ✅ **Data Validation**: Handles missing or invalid data gracefully
4. ✅ **Readable Reports**: Well-formatted, easy-to-understand output
5. ✅ **Clean Code**: Organized into reusable functions and modules

## 📊 Example Analysis

Given a dataset with various crops planted throughout 2024, AgriStats can identify:

- **Best performing crop**: Cassava with 4.10 tons/hectare average yield
- **Optimal planting time**: March shows 23% higher yields compared to annual average
- **Weather insights**: Strong positive correlation (0.67) between rainfall and yield

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 👨‍🌾 Use Cases

- Small-scale farmers optimizing crop selection
- Agricultural cooperatives analyzing regional performance
- Farm management planning seasonal activities
- Agricultural students learning data analysis

## 🔮 Future Enhancements

- Machine learning predictions for future yields
- Multi-year trend analysis
- Soil quality factor integration
- Interactive web dashboard
- Export to Excel format
- Support for multiple CSV files

---

**Built with ❤️ for farmers and agricultural communities**