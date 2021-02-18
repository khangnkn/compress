import csv
import re

test_cases = ['TestProtobuf', 'TestJSON', 'TestLzw',
              'TestFlate', 'TestGzip', 'TestZlib']
csv_headers = test_cases.copy()
csv_headers.insert(0, '')

size_file = open('report/size.csv', 'w')
time_file = open('report/time.csv', 'w')
mem_file = open('report/mem.csv', 'w')
alloc_file = open('report/alloc.csv', 'w')

size_writer = csv.writer(size_file)
size_writer.writerow(csv_headers)
time_writer = csv.writer(time_file)
time_writer.writerow(csv_headers)
mem_writer = csv.writer(mem_file)
mem_writer.writerow(csv_headers)
alloc_writer = csv.writer(alloc_file)
alloc_writer.writerow(csv_headers)


def process_results(test_results=[]):
    for result in test_results:
        process_result(result)


def process_result(test_result=''):
    len_tc = len(test_cases)
    sep = '\n'
    lines = test_result.splitlines()
    test_input = lines[0]
    size_result = sep.join(lines[1:len_tc * 3])
    bench_result = sep.join(
        lines[len_tc * 3 + 1: len_tc * 3 + 1 + 3 + len_tc * 2])
    num_of_input = process_testcase(test_input)
    size_result = process_size(size_result)
    size_result.insert(0, num_of_input)
    time_result = process_time(bench_result)
    time_result.insert(0, num_of_input)
    mem_result = process_mem(bench_result)
    mem_result.insert(0, num_of_input)
    alloc_result = process_alloc(bench_result)
    alloc_result.insert(0, num_of_input)
    size_writer.writerow(size_result)
    time_writer.writerow(time_result)
    mem_writer.writerow(mem_result)
    alloc_writer.writerow(alloc_result)


def process_testcase(record=''):
    result = re.findall(
        "(?<=generating test data of )(([0-9])\w+)(?= records)", record)
    return result[0][0]


def process_size(raw=''):
    size_data = re.findall("(?<= )(([0-9])\w+)(?= )", raw)
    res = []
    for size in size_data:
        res.append(size[0])
    return res


def process_time(raw=''):
    time_data = re.findall("(?<= )(([0-9])\w+)(?= ns\/op)", raw)
    res = []
    for time in time_data:
        res.append(time[0])
    return res


def process_mem(raw=''):
    mem_data = re.findall("(?<= )(([0-9])+)(?= B\/op)", raw)
    res = []
    for mem in mem_data:
        res.append(mem[0])
    return res


def process_alloc(raw=''):
    alloc_data = re.findall("(?<= )(([0-9])+)(?= allocs\/op)", raw)
    res = []
    for alloc in alloc_data:
        res.append(alloc[0])
    return res


f = open('/home/nguyen/go/src/github.com/nkhang/compress/report.txt', 'r')
content = f.read()
splitted = content.split('\n\n')
process_results(splitted)
