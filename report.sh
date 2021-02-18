rm report.txt
for TOTAL in {20..400..20}
    do
        export TOTAL_RECORDS=$TOTAL
        go test -v -bench=. >> report.txt
        echo "" >> report.txt
    done

# remove 2 last empty lines
head -n -2 report.txt > .tmp.txt; mv .tmp.txt report.txt
python3 ./process.py
