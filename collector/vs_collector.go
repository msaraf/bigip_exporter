package collector

import (
	"github.com/pr8kerl/f5er/f5"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"time"
)

type vsCollector struct {
	metrics map[string]vsMetric
	bigip   *f5.Device
	partitions_list []string
}

type vsMetric struct {
	desc      *prometheus.Desc
	extract   func(f5.LBVirtualStatsInnerEntries) float64
	valueType prometheus.ValueType
}

func NewVSCollector(bigip *f5.Device, namespace string, partitions_list []string) (error, *vsCollector) {
	var (
		subsystem  = "vs"
		labelNames = []string{"partition", "vs"}
	)
	return nil, &vsCollector{
		metrics: map[string]vsMetric{
			"syncookie_accepts": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "syncookie_accepts"),
					"syncookie_accepts",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Syncookie_accepts.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"ephemeral_bitsOut": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "ephemeral_bits_out"),
					"ephemeral_bits_out",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Ephemeral_bitsOut.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"clientside_bitsOut": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "clientside_bits_out"),
					"clientside_bits_out",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Clientside_bitsOut.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"fiveMinAvgUsageRatio": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "five_min_avg_usage_ratio"),
					"five_min_avg_usage_ratio",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.FiveMinAvgUsageRatio.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"fiveSecAvgUsageRatio": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "five_sec_avg_usage_ratio"),
					"five_sec_avg_usage_ratio",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.FiveSecAvgUsageRatio.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"syncookie_syncookies": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "syncookie_syncookies"),
					"syncookie_syncookies",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Syncookie_syncookies.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"ephemeral_slowKilled": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "ephemeral_slow_killed"),
					"ephemeral_slow_killed",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Ephemeral_slowKilled.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"ephemeral_pktsOut": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "ephemeral_pkts_out"),
					"ephemeral_pkts_out",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Ephemeral_pktsOut.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"syncookie_rejects": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "syncookie_rejects"),
					"syncookie_rejects",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Syncookie_rejects.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"syncookie_syncacheCurr": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "syncookie_syncache_curr"),
					"syncookie_syncache_curr",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Syncookie_syncacheCurr.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"csMinConnDur": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "cs_min_conn_dur"),
					"cs_min_conn_dur",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.CsMinConnDur.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"csMeanConnDur": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "cs_mean_conn_dur"),
					"cs_mean_conn_dur",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.CsMeanConnDur.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"syncookie_swsyncookieInstance": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "syncookie_swsyncookie_instance"),
					"syncookie_swsyncookie_instance",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Syncookie_swsyncookieInstance.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"syncookie_syncacheOver": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "syncookie_syncache_over"),
					"syncookie_syncache_over",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Syncookie_syncacheOver.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"syncookie_hwAccepts": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "syncookie_hw_accepts"),
					"syncookie_hw_accepts",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Syncookie_hwAccepts.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"ephemeral_pktsIn": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "ephemeral_pkts_in"),
					"ephemeral_pkts_in",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Ephemeral_pktsIn.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"clientside_totConns": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "clientside_tot_conns"),
					"clientside_tot_conns",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Clientside_totConns.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"ephemeral_curConns": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "ephemeral_cur_conns"),
					"ephemeral_cur_conns",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Ephemeral_curConns.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"clientside_evictedConns": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "clientside_evicted_conns"),
					"clientside_evicted_conns",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Clientside_evictedConns.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"oneMinAvgUsageRatio": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "one_min_avg_usage_ratio"),
					"one_min_avg_usage_ratio",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.OneMinAvgUsageRatio.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"ephemeral_evictedConns": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "ephemeral_evicted_conns"),
					"ephemeral_evicted_conns",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Ephemeral_evictedConns.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"clientside_slowKilled": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "clientside_slow_killed"),
					"clientside_slow_killed",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Clientside_slowKilled.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"clientside_bitsIn": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "clientside_bits_in"),
					"clientside_bits_in",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Clientside_bitsIn.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"ephemeral_maxConns": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "ephemeral_max_conns"),
					"ephemeral_max_conns",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Ephemeral_maxConns.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"syncookie_hwsyncookieInstance": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "syncookie_hwsyncookie_instance"),
					"syncookie_hwsyncookie_instance",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Syncookie_hwsyncookieInstance.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"clientside_pktsOut": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "clientside_pkts_out"),
					"clientside_pkts_out",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Clientside_pktsOut.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"clientside_curConns": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "clientside_cur_conns"),
					"clientside_cur_conns",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Clientside_curConns.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"ephemeral_bitsIn": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "ephemeral_bits_in"),
					"ephemeral_bits_in",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Ephemeral_bitsIn.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"clientside_pktsIn": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "clientside_pkts_in"),
					"clientside_pkts_in",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Clientside_pktsIn.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"totRequests": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "tot_requests"),
					"tot_requests",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.TotRequests.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"csMaxConnDur": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "cs_max_conn_dur"),
					"cs_max_conn_dur",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.CsMaxConnDur.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"syncookie_hwSyncookies": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "syncookie_hw_syncookies"),
					"syncookie_hw_syncookies",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Syncookie_hwSyncookies.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"clientside_maxConns": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "clientside_max_conns"),
					"clientside_max_conns",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Clientside_maxConns.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"ephemeral_totConns": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "ephemeral_tot_conns"),
					"ephemeral_tot_conns",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					return float64(entries.Ephemeral_totConns.Value)
				},
				valueType: prometheus.CounterValue,
			},
			"status_availabilityState": {
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, subsystem, "status_availability_state"),
					"status_availability_state",
					labelNames,
					nil,
				),
				extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
					if entries.Status_availabilityState.Description == "available" {
						return 1
					}
					return 0
				},
				valueType: prometheus.CounterValue,
			},
		},
		bigip: bigip,
		partitions_list: partitions_list,
	}
}

func (c *vsCollector) Collect(ch chan<- prometheus.Metric) {
	start := time.Now()
	err, virtualServers := c.bigip.ShowVirtuals()
	if err != nil {
		log.Fatal(err)
	}
	for _, virtualServer := range virtualServers.Items {
		if c.partitions_list != nil && !stringInSlice(virtualServer.Partition, c.partitions_list) {
			continue
		}
		err, virtualStats := c.bigip.ShowVirtualStats("/" + virtualServer.Partition + "/" + virtualServer.Name)
		if err != nil {
			log.Fatal(err)
		}
		lables := []string{virtualServer.Partition, virtualServer.Name}
		urlKey := "https://localhost/mgmt/tm/ltm/virtual/~" + virtualServer.Partition + "~" + virtualServer.Name + "/~" + virtualServer.Partition + "~" + virtualServer.Name + "/stats"
		for _, metric := range c.metrics {
			ch <- prometheus.MustNewConstMetric(metric.desc, metric.valueType, metric.extract(virtualStats.Entries[urlKey].NestedStats.Entries), lables...)
		}
	}
	elapsed := time.Since(start)
	log.Printf("Getting stats took %s", elapsed)
}

func (c *vsCollector) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range c.metrics {
		ch <- metric.desc
	}
}