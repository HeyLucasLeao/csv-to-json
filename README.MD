# CSV-to-JSON in Golang

## Overview
This little project is a personal journey in learning how to read and iterate multiple csv files and convert them into many arrays of JSON data.

### Usage
Setting .env

To configure your environment, you’ll need to specify three variables:

- MAX_RECORDS: This variable determines the maximum number of JSON records to keep in a file before partitioning it into another one.

- CSV_COMMA: This variable specifies the comma separator for every CSV file you want to convert. Note that each CSV file must use the same separator.

- CSV_FILENAME: This variable specifies the name of the CSV file inside the data folder that you want to select. It accepts pattern syntax. For example, if you want to convert every file that starts with “dataset” concurrently, you can use “dataset*.csv”.

### Contact
Lucas Leão - heylucasleao@gmail.com