[![Go](https://img.shields.io/github/go-mod/go-version/criteo/varnishncsa-exporter)](https://github.com/criteo/varnishncsa-exporter)
[![status](https://img.shields.io/badge/status-template-blue)](https://github.com/criteo/varnishncsa-exporter)
[![CI](https://github.com/criteo/varnishncsa-exporter/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/criteo/varnishncsa-exporter/actions/workflows/ci.yml)
[![GitHub](https://img.shields.io/github/license/criteo/varnishncsa-exporter)](https://github.com/criteo/varnishncsa-exporter/blob/main/LICENSE)

# Varnishncsa Exporter

Processes varnishncsa structured logs output and exposes prometheus metrics.
The tool starts a varnishncsa with options in daemon mode and pipes it stdout
in order to process it: logs it and exposes prometheus metrics by starting a
web server.

## How to start the process

Process help
```
$ sudo ./varnishncsa_exporter -h
NAME:
   varnishncsa_exporter - Exposes prometheus metrics for Varnish by parsing structured logs output from varnishncsa.

USAGE:
   varnishncsa_exporter [global options] command [command options]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help

   Miscellaneous:

   --debug, -D  show debug output (default: false)
   --version    Print version (default: false)

   Piped daemon:

   --binary value, -b value     Binary to run and pipe output (default: /usr/bin/varnishncsa)
   --directory value, -n value  Varnishd working directory (default: /run/varnish/)
   --format value, -F value     Set the output log format string (default: {"Timestamp": "%t", "Handling": "%{Varnish:handling}x", "Bytes": "%b", "X-Real-Host": "%{x-real-host}i", "X-Frontend-Id": "%{x-frontend-id}i"})
   --labels value, -L value     Prometheus labels mapping key, value represented  in json (default: {"X-Real-Host": "host", "X-Frontend-Id": "frontend"})

   Prometheus server:

   --httpd_address value, -a value  Prometheus HTTP server address (default: 127.0.0.1)
   --httpd_port value, -p value     Prometheus HTTP server port (default: 8080)
```

Start with default options: the process requires priviledged permissions in order to start 
```
$ sudo ./varnishncsa_exporter
INFO[0000] Starting prometheus server on 127.0.0.1:8080
INFO[0000] Executing command:
        /usr/bin/varnishncsa -n /run/varnish/ -F {"Timestamp": "%t", "Handling": "%{Varnish:handling}x", "Bytes": "%b", "X-Real-Host": "%{x-real-host}i", "X-Frontend-Id": "%{x-frontend-id}i"}
```

Using the default options, the prometheus output looks like
```
$ curl -s http://127.0.0.1:8080/metrics
# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 1.4667e-05
go_gc_duration_seconds{quantile="0.25"} 1.4667e-05
go_gc_duration_seconds{quantile="0.5"} 1.4878e-05
go_gc_duration_seconds{quantile="0.75"} 1.4878e-05
go_gc_duration_seconds{quantile="1"} 1.4878e-05
go_gc_duration_seconds_sum 2.9545e-05
go_gc_duration_seconds_count 2
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 7
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.21.10"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 2.603048e+06
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 8.921264e+06
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 4462
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 149266
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 3.667344e+06
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 2.603048e+06
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 4.333568e+06
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 3.694592e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 47774
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 3.407872e+06
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 8.02816e+06
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 1.7169035150380206e+09
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 0
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 197040
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 1200
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 15600
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 66696
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 97776
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.194304e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 416306
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 327680
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 327680
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 1.2557328e+07
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 6
# HELP hit_count_total Hit count.
# TYPE hit_count_total counter
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net"} 2140
# HELP miss_count_total Miss count.
# TYPE miss_count_total counter
miss_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net"} 51
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 0.07
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 524288
# HELP process_open_fds Number of open file descriptors.
# TYPE process_open_fds gauge
process_open_fds 10
# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 1.5335424e+07
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.71690350777e+09
# HELP process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE process_virtual_memory_bytes gauge
process_virtual_memory_bytes 1.575911424e+09
# HELP process_virtual_memory_max_bytes Maximum amount of virtual memory available in bytes.
# TYPE process_virtual_memory_max_bytes gauge
process_virtual_memory_max_bytes 1.8446744073709552e+19
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 1
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
# HELP synth_count_total Synth count.
# TYPE synth_count_total counter
synth_count_total{frontend="-",host="-"} 1
```

One can change the format and labels options in order to retrieve and expose his intended varnish field.
Below is an example where the *User-Agent* HTTP header field is extracted and mapped into prometheus label *userAgent*
**Disclaimer**: using the user agent as prometheus metric is not a good idea knowing the cardinaity of the the user agent field.
```
$ sudo ./varnishncsa_exporter -F '{"Handling": "%{Varnish:handling}x", "X-Real-Host": "%{x-real-host}i", "X-Frontend-Id": "%{x-frontend-id}i", "User-Agent": "%{User-agent}i"}' -L '{"X-Real-Host": "host", "X-Frontend-Id": "frontend", "User-Agent": "userAgent"}
INFO[0000] Starting prometheus server on 127.0.0.1:8080
INFO[0000] Executing command:
        /usr/bin/varnishncsa -n /run/varnish/ -F {"Timestamp": "%t", "Handling": "%{Varnish:handling}x", "Bytes": "%b", "X-Real-Host": "%{x-real-host}i", "X-Frontend-Id": "%{x-frontend-id}i", "User-Agent": "%{User-agent}i"}
```

And the prometheus output looks like
```
$ curl -s http://127.0.0.1:8080/metrics
# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 1.2664e-05
go_gc_duration_seconds{quantile="0.25"} 1.2664e-05
go_gc_duration_seconds{quantile="0.5"} 2.1671e-05
go_gc_duration_seconds{quantile="0.75"} 2.1671e-05
go_gc_duration_seconds{quantile="1"} 2.1671e-05
go_gc_duration_seconds_sum 3.4335e-05
go_gc_duration_seconds_count 2
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 7
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.21.10"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 3.158392e+06
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 8.263792e+06
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 4462
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 98137
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 3.86484e+06
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 3.158392e+06
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 3.710976e+06
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 4.349952e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 53627
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 3.530752e+06
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 8.060928e+06
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 1.716903784494539e+09
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 0
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 151764
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 1200
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 15600
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 80808
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 97776
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.194304e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 448186
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 327680
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 327680
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 1.2819472e+07
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 6
# HELP hit_count_total Hit count.
# TYPE hit_count_total counter
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="%E5%81%A5%E5%BA%B72.0/349 CFNetwork/1494.0.7 Darwin/23.4.0"} 2
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="%E7%88%86%E6%96%99%E5%85%AC%E7%A4%BE/20240425001 CFNetwork/1496.0.7 Darwin/23.5.0"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="CapCut/11.1.1.20 CFNetwork/1408.0.4 Darwin/22.5.0"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 10; M2006C3LII MIUI/V12.0.25.0.QCDINXM)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 10; M2006C3MG MIUI/V12.0.23.0.QCRMIXM)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 10; Redmi Note 9 Pro Max MIUI/V12.0.3.0.QJXINXM)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 11; CPH1933 Build/RKQ1.200903.002)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 11; CPH1937 Build/RKQ1.200903.002)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 11; CPH1989 Build/RP1A.200720.011)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 11; CPH2059 Build/RKQ1.200903.002)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 11; Infinix X6812B Build/RP1A.200720.011)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 11; Infinix X688B Build/RP1A.200720.011)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 11; M2006C3LG Build/RP1A.200720.011)"} 2
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 11; Mi 9T Build/RKQ1.200826.002)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 11; RMX1851 Build/RKQ1.201217.002)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 11; RMX2040 Build/RP1A.200720.011)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 11; RMX2193 Build/RP1A.200720.011)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 11; RMX3231 Build/RP1A.201005.001)"} 2
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 11; RMX3235 Build/RP1A.201005.001)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 11; RMX3261 Build/RP1A.201005.001)"} 2
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 11; RMX3263 Build/RP1A.201005.001)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 11; RMX3581 Build/RP1A.201005.001)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 11; SM-A125F Build/RP1A.200720.012)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 11; SM-M215F Build/RP1A.200720.012)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 11; vivo 1819 Build/RP1A.200720.012)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 11; vivo 1904 Build/RP1A.200720.012)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 11; vivo 1906 Build/RP1A.200720.012)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 11; vivo 2007 Build/RP1A.200720.012)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 12; CPH2477 Build/SP1A.210812.016)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 12; Infinix X6515 Build/SP1A.210812.016)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 12; M2004J19C Build/SP1A.210812.016)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 12; M2007J20CG Build/SKQ1.211019.001)"} 2
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 12; M2010J19CG Build/SKQ1.211202.001)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 12; V2134 Build/SP1A.210812.003)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 12; V2217 Build/SP1A.210812.003)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 12; vivo 1920 Build/SP1A.210812.003)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 12; vivo 2019 Build/SP1A.210812.003)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 13; 21121119SG Build/TP1A.220624.014)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 13; 220333QAG Build/TKQ1.221114.001)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 13; 2302EPCC4I Build/TP1A.220624.014)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 13; 23106RN0DA Build/TP1A.220624.014)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 13; Infinix X6525 Build/TP1A.220624.014)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 13; Infinix X6731 Build/TP1A.220624.014)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 13; Infinix X678B Build/TP1A.220624.014)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 13; M2103K19G Build/TP1A.220624.014)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 13; RMX3085 Build/SP1A.210812.016)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 13; RMX3630 Build/TP1A.220905.001)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 13; RMX3760 Build/TP1A.220624.014)"} 4
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 13; RMX3762 Build/TP1A.220624.014)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 13; RMX3830 Build/TP1A.220624.014)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 13; RMX3834 Build/TP1A.220624.014)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 13; SM-G781B Build/TP1A.220624.014)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 13; V2310 Build/TP1A.220624.014)"} 2
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 13; V2312 Build/TP1A.220624.014_MOD1)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 13; vivo 2018 Build/TP1A.220624.014)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 14; CPH2363 Build/TP1A.220905.001)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 14; CPH2381 Build/UKQ1.230924.001)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 14; CPH2387 Build/TP1A.220905.001)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 14; CPH2483 Build/UP1A.230620.001)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 14; CPH2495 Build/UP1A.230620.001)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 14; CPH2505 Build/RKQ1.211119.001)"} 2
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 14; SM-A155F Build/UP1A.231005.007)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 14; SM-A235F Build/UP1A.231005.007)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 14; SM-A525F Build/UP1A.231005.007)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 14; SM-A546E Build/UP1A.231005.007)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 14; SM-G990E Build/UP1A.231005.007)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 14; SM-S911B Build/UP1A.231005.007)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 14; SM-X210 Build/UP1A.231005.007)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 14; V2201 Build/UP1A.231005.007)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 14; V2303 Build/UP1A.231005.007)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 8.0.0; RNE-L21 Build/HUAWEIRNE-L21)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 8.1.0; vivo 1811 Build/OPM1.171019.026)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 8.1.0; vivo 1820 Build/O11019)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 9; CPH2015 Build/PPR1.180610.011)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dcard/3406 CFNetwork/1490.0.4 Darwin/23.2.0"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="GetHornet-AppStore/4229 CFNetwork/1494.0.7 Darwin/23.4.0"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; Active 3 Build/QP1A.190711.020; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/124.0.6367.179 Mobile Safari/537.36"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; CLT-L29 Build/HUAWEICLT-L29; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/125.0.6422.53 Mobile Safari/537.36 (Mobile; afma-sdk-a-v241199999.231004000.1)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; CPH1823 Build/QP1A.190711.020; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/124.0.6367.179 Mobile Safari/537.36 (Mobile; afma-sdk-a-v241199999.214106000.1)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; CPH1911 Build/QP1A.190711.020; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/124.0.6367.179 Mobile Safari/537.36 (Mobile; afma-sdk-a-v241199999.232400000.1)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; CPH2127 Build/QKQ1.200614.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/124.0.6367.180 Mobile Safari/537.36 (Mobile; afma-sdk-a-v241199999.231700000.1)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; CPH2179 Build/QP1A.190711.020; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/124.0.6367.179 Mobile Safari/537.36 (Mobile; afma-sdk-a-v241199999.232400000.1)"} 3
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; CPH2185 Build/QP1A.190711.020; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/124.0.6367.179 Mobile Safari/537.36"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; CPH2185 Build/QP1A.190711.020; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/124.0.6367.179 Mobile Safari/537.36 (Mobile; afma-sdk-a-v234310999.223104000.1)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; CPH2185 Build/QP1A.190711.020; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/124.0.6367.180 Mobile Safari/537.36"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; CPH2185 Build/QP1A.190711.020; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/124.0.6367.180 Mobile Safari/537.36 (Mobile; afma-sdk-a-v241199999.240304000.1)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; CPH2239 Build/QP1A.190711.020; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/124.0.6367.179 Mobile Safari/537.36 (Mobile; afma-sdk-a-v241199999.234310000.1)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; Infinix X682C Build/QP1A.190711.020; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/124.0.6367.180 Mobile Safari/537.36 (Mobile; afma-sdk-a-v241199999.240304000.1)"} 2
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; Infinix X688C Build/QP1A.190711.020; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/124.0.6367.179 Mobile Safari/537.36 (Mobile; afma-sdk-a-v241199999.232400000.1)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; Infinix X692 Build/QP1A.190711.020; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/124.0.6367.179 Mobile Safari/537.36 (Mobile; afma-sdk-a-v241199999.232400000.1)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Mobile Safari/537.36"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Mobile Safari/537.36"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Mobile Safari/537.36"} 2
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Mobile Safari/537.36"} 3
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Mobile Safari/537.36"} 3
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Mobile Safari/537.36"} 4
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Mobile Safari/537.36"} 3
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Mobile Safari/537.36"} 21
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Mobile Safari/537.36"} 84
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36"} 3
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPad; CPU OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/125.0.6422.80 Mobile/15E148 Safari/604.1"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPad; CPU OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/123.0.6312.52 Mobile/15E148 Safari/604.1"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPad; CPU OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148"} 5
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPad; CPU OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 Safari Line/13.3.0"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPad; CPU OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/21E236 [FBAN/FBIOS;FBAV/465.0.1.41.103;FBBV/602060281;FBDV/iPad13,16;FBMD/iPad;FBSN/iPadOS;FBSV/17.4.1;FBSS/2;FBID/tablet;FBLC/zh_TW;FBOP/5;FBRV/6"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPad; CPU OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 Mobile/15E148 Safari/604.1"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPad; CPU OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 12_5_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/36.0  Mobile/15E148 Safari/605.1.15"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 15_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/125.0.6422.80 Mobile/15E148 Safari/604.1"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/125.0.6422.80 Mobile/15E148 Safari/604.1"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 15_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/104.0.5112.71 Mobile/15E148 Safari/604.1"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 15_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 15_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.6.1 Mobile/15E148 Safari/604.1"} 3
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 15_7_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 15_8 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/124.0.6367.88 Mobile/15E148 Safari/604.1"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 15_8_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 15_8_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/19H384 [FBAN/FBIOS;FBAV/465.0.1.41.103;FBBV/602060281;FBDV/iPhone9,3;FBMD/iPhone;FBSN/iOS;FBSV/15.8.2;FBSS/2;FBID/phone;FBLC/en_GB;FBOP/5"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 16_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.0 Mobile/15E148 Safari/604.1"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 16_0_2 like Mac OS X) AppleWebKit/614.1.25.0.31 (KHTML, like Gecko) Mobile/20A380 BiliApp/2890100 mobi_app/bstar_i channel/AppStore Buvid/Y947E99E39AC6531420882512C76732E8946 c_locale/en_PH s_locale/en_PH"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 16_0_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/20A392 [FBAN/FBIOS;FBAV/465.0.1.41.103;FBBV/602060281;FBDV/iPhone15,3;FBMD/iPhone;FBSN/iOS;FBSV/16.0.3;FBSS/3;FBID/phone;FBLC/en_Qaau_GB;"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 16_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/125.0.6422.80 Mobile/15E148 Safari/604.1"} 2
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 16_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) GSA/317.0.634488990 Mobile/15E148 Safari/604.1"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 16_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/20B82 [FBAN/FBIOS;FBAV/461.0.3.32.103;FBBV/591474467;FBDV/iPhone14,2;FBMD/iPhone;FBSN/iOS;FBSV/16.1;FBSS/3;FBID/phone;FBLC/zh_TW;FBOP/5;FBR"} 2
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 16_1_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/20B101 [FBAN/FBIOS;FBAV/465.0.1.41.103;FBBV/602060281;FBDV/iPhone12,1;FBMD/iPhone;FBSN/iOS;FBSV/16.1.1;FBSS/2;FBID/phone;FBLC/zh_TW;FBOP/"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 16_1_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Mobile/15E148 Safari/604.1"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 16_1_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/20B110 [FBAN/FBIOS;FBAV/456.0.0.43.107;FBBV/578727508;FBDV/iPhone14,5;FBMD/iPhone;FBSN/iOS;FBSV/16.1.2;FBSS/3;FBID/phone;FBLC/zh_TW;FBOP/"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 16_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 16_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/20C65 [FBAN/FBIOS;FBDV/iPhone12,5;FBMD/iPhone;FBSN/iOS;FBSV/16.2;FBSS/3;FBID/phone;FBLC/zh_TW;FBOP/5]"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 16_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.2 Mobile/15E148 Safari/604.1"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 16_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.2 Mobile/15E148 Safari/605.1 NAVER(inapp; search; 2000; 12.5.3; 11)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) GSA/317.0.634488990 Mobile/15E148 Safari/604.1"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/20D47 [FBAN/FBIOS;FBAV/464.0.0.39.106;FBBV/599886571;FBDV/iPhone13,2;FBMD/iPhone;FBSN/iOS;FBSV/16.3;FBSS/3;FBID/phone;FBLC/zh_TW;FBOP/5;FBR"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 Mobile/15E148 Safari/604.1"} 2
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 Mobile/15E148 Safari/604.1"} 4
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 16_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.5 Mobile/15E148 Safari/604.1"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 16_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/20F75 [FBAN/FBIOS;FBAV/465.0.1.41.103;FBBV/602060281;FBDV/iPhone15,2;FBMD/iPhone;FBSN/iOS;FBSV/16.5.1;FBSS/3;FBID/phone;FBLC/zh_TW;FBOP/5"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/21E236 [FBAN/FBIOS;FBAV/465.0.1.41.103;FBBV/602060281;FBDV/iPhone16,2;FBMD/iPhone;FBSN/iOS;FBSV/17.4.1;FBSS/3;FBID/phone;FBLC/en_Qaau_GB;"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/21E236 [FBAN/FBIOS;FBAV/465.0.1.41.103;FBBV/602060281;FBDV/iPhone16,2;FBMD/iPhone;FBSN/iOS;FBSV/17.4.1;FBSS/3;FBID/phone;FBLC/en_US;FBOP/"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/21E236 [FBAN/FBIOS;FBDV/iPhone11,2;FBMD/iPhone;FBSN/iOS;FBSV/17.4.1;FBSS/3;FBID/phone;FBLC/zh_TW;FBOP/5]"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/21E236 [FBAN/FBIOS;FBDV/iPhone11,6;FBMD/iPhone;FBSN/iOS;FBSV/17.4.1;FBSS/3;FBID/phone;FBLC/zh_TW;FBOP/5]"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 Mobile/15E148 Safari/604.1"} 41
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/125.0.6422.80 Mobile/15E148 Safari/604.1"} 3
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/21F79 [FBAN/FBIOS;FBAV/461.0.3.32.103;FBBV/591474467;FBDV/iPhone14,5;FBMD/iPhone;FBSN/iOS;FBSV/17.5;FBSS/3;FBID/phone;FBLC/zh_TW;FBOP/5;FBR"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148"} 3
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 Safari Line/14.7.0"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/21F90 [FBAN/FBIOS;FBAV/465.0.1.41.103;FBBV/602060281;FBDV/iPhone14,3;FBMD/iPhone;FBSN/iOS;FBSV/17.5.1;FBSS/3;FBID/phone;FBLC/en_US;FBOP/5"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 Mobile/15E148 Safari/604.1"} 5
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Pureit/1020302 CFNetwork/1333.0.4 Darwin/21.5.0"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Rave/2340 CFNetwork/1390 Darwin/22.0.0"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="UnityPlayer/2020.3.43f1 (UnityWebRequest/1.0, libcurl/7.84.0-DEV)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="UnityPlayer/2020.3.48f1 (UnityWebRequest/1.0, libcurl/7.84.0-DEV)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="UnityPlayer/2021.3.20f1 (UnityWebRequest/1.0, libcurl/7.84.0-DEV)"} 3
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="UnityPlayer/2021.3.24f1 (UnityWebRequest/1.0, libcurl/7.84.0-DEV)"} 2
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="UnityPlayer/2021.3.36f1 (UnityWebRequest/1.0, libcurl/8.5.0-DEV)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="UnityPlayer/2021.3.37f1 (UnityWebRequest/1.0, libcurl/8.5.0-DEV)"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="UnityPlayer/2021.3.38f1 (UnityWebRequest/1.0, libcurl/8.5.0-DEV)"} 2
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="UnityPlayer/2022.3.20f1 (UnityWebRequest/1.0, libcurl/8.5.0-DEV)"} 2
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="aegon-android/3.12.1-48-g0db27b40-nodiag-nolto"} 1
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="okhttp/4.12.0"} 185
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="okhttp/4.8.0"} 8
hit_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="python-requests/2.27.1"} 10
# HELP miss_count_total Miss count.
# TYPE miss_count_total counter
miss_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Dalvik/2.1.0 (Linux; U; Android 10; RMX2032 Build/QKQ1.200209.002)"} 1
miss_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Mobile Safari/537.36"} 2
miss_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Mobile Safari/537.36"} 7
miss_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 13; CPH2159 Build/TP1A.220905.001; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/124.0.6367.179 Mobile Safari/537.36 (Mobile; afma-sdk-a-v241199999.223104000.1)"} 1
miss_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 13; CPH2365 Build/TP1A.220905.001; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/124.0.6367.180 Mobile Safari/537.36"} 1
miss_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 13; SM-A226B Build/TP1A.220624.014; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/124.0.6367.179 Mobile Safari/537.36"} 1
miss_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 13; moto g62 5G Build/T1SSIS33.1-75-7-5; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/124.0.6367.179 Mobile Safari/537.36 (Mobile; afma-sdk-a-v241199999.232400000.1)"} 1
miss_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Linux; Android 8.1.0; vivo 1816 Build/O11019; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/87.0.4280.141 Mobile Safari/537.36 (Mobile; afma-sdk-a-v241199999.232400000.1)"} 1
miss_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36"} 2
miss_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.6.1 Safari/605.1.15"} 1
miss_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"} 1
miss_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36"} 6
miss_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.0.0"} 2
miss_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148"} 1
miss_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148"} 2
miss_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 Mobile/15E148 Safari/604.1"} 2
miss_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="aegon-android/3.12.1-3-g81450974-nodiag-nolto"} 1
miss_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="okhttp/4.12.0"} 2
miss_count_total{frontend="myapp.criteo.net",host="myapp.eu.criteo.net",userAgent="python-requests/2.27.1"} 2
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 0.06
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 524288
# HELP process_open_fds Number of open file descriptors.
# TYPE process_open_fds gauge
process_open_fds 10
# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 1.54624e+07
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.71690377928e+09
# HELP process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE process_virtual_memory_bytes gauge
process_virtual_memory_bytes 1.576173568e+09
# HELP process_virtual_memory_max_bytes Maximum amount of virtual memory available in bytes.
# TYPE process_virtual_memory_max_bytes gauge
process_virtual_memory_max_bytes 1.8446744073709552e+19
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 1
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
```
