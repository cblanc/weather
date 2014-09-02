# Weather

Utility to retrieve 5 day weather forecast for a location via [http://openweathermap.org/](http://openweathermap.org/)

## Usage

Add binary to $PATH

```
$ weather london

or

$ weather london,uk

or 

$ weather new york
```

```
5 day forecast for New York, US (40.71427, -74.00597)

+------------------+------------------------+----------------+--------------------+------------------+--------------+
|       DAY        |        FORECAST        |  TEMP (RANGE)  | CLOUD COVERAGE (%) | WIND SPEED (M/S) | HUMIDITY (%) |
+------------------+------------------------+----------------+--------------------+------------------+--------------+
| Monday, 09/01    | Clouds (broken clouds) | 23C (23C-23C)  | 68%                |              1.0 | 89           |
| Tuesday, 09/02   | Clouds (few clouds)    | 27C (22C-29C)  | 12%                |              1.5 | 83           |
| Wednesday, 09/03 | Clear (sky is clear)   | 25C (19C-27C)  | 0%                 |              1.8 | 65           |
| Thursday, 09/04  | Rain (light rain)      | 27C (20C-27C)  | 16%                |              2.4 | 0            |
| Friday, 09/05    | Rain (light rain)      | 29C (23C-29C)  | 0%                 |              3.1 | 0            |
+------------------+------------------------+----------------+--------------------+------------------+--------------+
```