/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task_stats.proto

package firmament

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TaskStats struct {
	TaskId    uint64 `protobuf:"varint,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	Hostname  string `protobuf:"bytes,2,opt,name=hostname" json:"hostname,omitempty"`
	Timestamp uint64 `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
	// CPU stats in millicores.
	CpuLimit   int64 `protobuf:"varint,4,opt,name=cpu_limit,json=cpuLimit" json:"cpu_limit,omitempty"`
	CpuRequest int64 `protobuf:"varint,5,opt,name=cpu_request,json=cpuRequest" json:"cpu_request,omitempty"`
	CpuUsage   int64 `protobuf:"varint,6,opt,name=cpu_usage,json=cpuUsage" json:"cpu_usage,omitempty"`
	// Memory stats in Kb.
	MemLimit            int64   `protobuf:"varint,7,opt,name=mem_limit,json=memLimit" json:"mem_limit,omitempty"`
	MemRequest          int64   `protobuf:"varint,8,opt,name=mem_request,json=memRequest" json:"mem_request,omitempty"`
	MemUsage            int64   `protobuf:"varint,9,opt,name=mem_usage,json=memUsage" json:"mem_usage,omitempty"`
	MemRss              int64   `protobuf:"varint,10,opt,name=mem_rss,json=memRss" json:"mem_rss,omitempty"`
	MemCache            int64   `protobuf:"varint,11,opt,name=mem_cache,json=memCache" json:"mem_cache,omitempty"`
	MemWorkingSet       int64   `protobuf:"varint,12,opt,name=mem_working_set,json=memWorkingSet" json:"mem_working_set,omitempty"`
	MemPageFaults       int64   `protobuf:"varint,13,opt,name=mem_page_faults,json=memPageFaults" json:"mem_page_faults,omitempty"`
	MemPageFaultsRate   float64 `protobuf:"fixed64,14,opt,name=mem_page_faults_rate,json=memPageFaultsRate" json:"mem_page_faults_rate,omitempty"`
	MajorPageFaults     int64   `protobuf:"varint,15,opt,name=major_page_faults,json=majorPageFaults" json:"major_page_faults,omitempty"`
	MajorPageFaultsRate float64 `protobuf:"fixed64,16,opt,name=major_page_faults_rate,json=majorPageFaultsRate" json:"major_page_faults_rate,omitempty"`
	// Network stats in Kb.
	NetRx           int64   `protobuf:"varint,17,opt,name=net_rx,json=netRx" json:"net_rx,omitempty"`
	NetRxErrors     int64   `protobuf:"varint,18,opt,name=net_rx_errors,json=netRxErrors" json:"net_rx_errors,omitempty"`
	NetRxErrorsRate float64 `protobuf:"fixed64,19,opt,name=net_rx_errors_rate,json=netRxErrorsRate" json:"net_rx_errors_rate,omitempty"`
	NetRxRate       float64 `protobuf:"fixed64,20,opt,name=net_rx_rate,json=netRxRate" json:"net_rx_rate,omitempty"`
	NetTx           int64   `protobuf:"varint,21,opt,name=net_tx,json=netTx" json:"net_tx,omitempty"`
	NetTxErrors     int64   `protobuf:"varint,22,opt,name=net_tx_errors,json=netTxErrors" json:"net_tx_errors,omitempty"`
	NetTxErrorsRate float64 `protobuf:"fixed64,23,opt,name=net_tx_errors_rate,json=netTxErrorsRate" json:"net_tx_errors_rate,omitempty"`
	NetTxRate       float64 `protobuf:"fixed64,24,opt,name=net_tx_rate,json=netTxRate" json:"net_tx_rate,omitempty"`
}

func (m *TaskStats) Reset()                    { *m = TaskStats{} }
func (m *TaskStats) String() string            { return proto.CompactTextString(m) }
func (*TaskStats) ProtoMessage()               {}
func (*TaskStats) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *TaskStats) GetTaskId() uint64 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

func (m *TaskStats) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *TaskStats) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *TaskStats) GetCpuLimit() int64 {
	if m != nil {
		return m.CpuLimit
	}
	return 0
}

func (m *TaskStats) GetCpuRequest() int64 {
	if m != nil {
		return m.CpuRequest
	}
	return 0
}

func (m *TaskStats) GetCpuUsage() int64 {
	if m != nil {
		return m.CpuUsage
	}
	return 0
}

func (m *TaskStats) GetMemLimit() int64 {
	if m != nil {
		return m.MemLimit
	}
	return 0
}

func (m *TaskStats) GetMemRequest() int64 {
	if m != nil {
		return m.MemRequest
	}
	return 0
}

func (m *TaskStats) GetMemUsage() int64 {
	if m != nil {
		return m.MemUsage
	}
	return 0
}

func (m *TaskStats) GetMemRss() int64 {
	if m != nil {
		return m.MemRss
	}
	return 0
}

func (m *TaskStats) GetMemCache() int64 {
	if m != nil {
		return m.MemCache
	}
	return 0
}

func (m *TaskStats) GetMemWorkingSet() int64 {
	if m != nil {
		return m.MemWorkingSet
	}
	return 0
}

func (m *TaskStats) GetMemPageFaults() int64 {
	if m != nil {
		return m.MemPageFaults
	}
	return 0
}

func (m *TaskStats) GetMemPageFaultsRate() float64 {
	if m != nil {
		return m.MemPageFaultsRate
	}
	return 0
}

func (m *TaskStats) GetMajorPageFaults() int64 {
	if m != nil {
		return m.MajorPageFaults
	}
	return 0
}

func (m *TaskStats) GetMajorPageFaultsRate() float64 {
	if m != nil {
		return m.MajorPageFaultsRate
	}
	return 0
}

func (m *TaskStats) GetNetRx() int64 {
	if m != nil {
		return m.NetRx
	}
	return 0
}

func (m *TaskStats) GetNetRxErrors() int64 {
	if m != nil {
		return m.NetRxErrors
	}
	return 0
}

func (m *TaskStats) GetNetRxErrorsRate() float64 {
	if m != nil {
		return m.NetRxErrorsRate
	}
	return 0
}

func (m *TaskStats) GetNetRxRate() float64 {
	if m != nil {
		return m.NetRxRate
	}
	return 0
}

func (m *TaskStats) GetNetTx() int64 {
	if m != nil {
		return m.NetTx
	}
	return 0
}

func (m *TaskStats) GetNetTxErrors() int64 {
	if m != nil {
		return m.NetTxErrors
	}
	return 0
}

func (m *TaskStats) GetNetTxErrorsRate() float64 {
	if m != nil {
		return m.NetTxErrorsRate
	}
	return 0
}

func (m *TaskStats) GetNetTxRate() float64 {
	if m != nil {
		return m.NetTxRate
	}
	return 0
}

func init() {
	proto.RegisterType((*TaskStats)(nil), "firmament.TaskStats")
}

func init() { proto.RegisterFile("task_stats.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0x5f, 0x6f, 0xd3, 0x30,
	0x14, 0xc5, 0x15, 0xb6, 0xa5, 0xcd, 0x2d, 0xa5, 0xab, 0xf7, 0xa7, 0x16, 0x20, 0xa8, 0xf6, 0x80,
	0x2a, 0x90, 0xe0, 0x61, 0x1f, 0x01, 0x81, 0x84, 0xc4, 0x03, 0xca, 0x82, 0x78, 0x8c, 0x4c, 0x76,
	0x97, 0x85, 0xce, 0x49, 0xb0, 0x6f, 0x44, 0x3f, 0x1c, 0x1f, 0x0e, 0xf9, 0xba, 0x31, 0x49, 0xe1,
	0xf1, 0x9e, 0xf3, 0x3b, 0xc7, 0xe7, 0xc5, 0x70, 0x4a, 0xca, 0x6e, 0x73, 0x4b, 0x8a, 0xec, 0xdb,
	0xd6, 0x34, 0xd4, 0x88, 0xe4, 0xae, 0x32, 0x5a, 0x69, 0xac, 0xe9, 0xea, 0x77, 0x0c, 0x49, 0xa6,
	0xec, 0xf6, 0xc6, 0xd9, 0x62, 0x05, 0x13, 0x86, 0xab, 0x5b, 0x19, 0xad, 0xa3, 0xcd, 0x71, 0x1a,
	0xbb, 0xf3, 0xd3, 0xad, 0x78, 0x0a, 0xd3, 0xfb, 0xc6, 0x52, 0xad, 0x34, 0xca, 0x47, 0xeb, 0x68,
	0x93, 0xa4, 0xe1, 0x16, 0xcf, 0x21, 0xa1, 0x4a, 0xa3, 0x25, 0xa5, 0x5b, 0x79, 0xc4, 0xb1, 0xbf,
	0x82, 0x78, 0x06, 0x49, 0xd1, 0x76, 0xf9, 0x43, 0xa5, 0x2b, 0x92, 0xc7, 0xeb, 0x68, 0x73, 0x94,
	0x4e, 0x8b, 0xb6, 0xfb, 0xec, 0x6e, 0xf1, 0x12, 0x66, 0xce, 0x34, 0xf8, 0xb3, 0x43, 0x4b, 0xf2,
	0x84, 0x6d, 0x28, 0xda, 0x2e, 0xf5, 0x4a, 0x9f, 0xee, 0xac, 0x2a, 0x51, 0xc6, 0x21, 0xfd, 0xd5,
	0xdd, 0xce, 0xd4, 0xa8, 0xf7, 0xd5, 0x13, 0x6f, 0x6a, 0xd4, 0xa1, 0xda, 0x99, 0x7d, 0xf5, 0xd4,
	0x57, 0x6b, 0xd4, 0x83, 0x6a, 0x07, 0xf8, 0xea, 0x24, 0xa4, 0x7d, 0xf5, 0x0a, 0x26, 0x9c, 0xb6,
	0x56, 0x02, 0x5b, 0xb1, 0x4b, 0x5a, 0xdb, 0xa7, 0x0a, 0x55, 0xdc, 0xa3, 0x9c, 0x85, 0xd4, 0x7b,
	0x77, 0x8b, 0x57, 0xb0, 0x70, 0xe6, 0xaf, 0xc6, 0x6c, 0xab, 0xba, 0xcc, 0x2d, 0x92, 0x7c, 0xcc,
	0xc8, 0x5c, 0xa3, 0xfe, 0xe6, 0xd5, 0x1b, 0xa4, 0x9e, 0x6b, 0x55, 0x89, 0xf9, 0x9d, 0xea, 0x1e,
	0xc8, 0xca, 0x79, 0xe0, 0xbe, 0xa8, 0x12, 0x3f, 0xb2, 0x28, 0xde, 0xc1, 0xf9, 0x01, 0x97, 0x1b,
	0x45, 0x28, 0x9f, 0xac, 0xa3, 0x4d, 0x94, 0x2e, 0x47, 0x70, 0xaa, 0x08, 0xc5, 0x6b, 0x58, 0x6a,
	0xf5, 0xa3, 0x31, 0xa3, 0xea, 0x05, 0x57, 0x2f, 0xd8, 0x18, 0x94, 0x5f, 0xc3, 0xe5, 0x3f, 0xac,
	0xaf, 0x3f, 0xe5, 0xfa, 0xb3, 0x83, 0x00, 0x3f, 0x70, 0x01, 0x71, 0x8d, 0x94, 0x9b, 0x9d, 0x5c,
	0x72, 0xeb, 0x49, 0x8d, 0x94, 0xee, 0xc4, 0x15, 0xcc, 0xbd, 0x9c, 0xa3, 0x31, 0x8d, 0xb1, 0x52,
	0xb0, 0x3b, 0x63, 0xf7, 0x03, 0x4b, 0xe2, 0x0d, 0x88, 0x11, 0xe3, 0xdf, 0x3a, 0xe3, 0xb7, 0x16,
	0x03, 0x90, 0xdf, 0x79, 0x01, 0xb3, 0x3d, 0xcc, 0xd4, 0x39, 0x53, 0x09, 0x53, 0xc3, 0x1d, 0xb4,
	0x93, 0x17, 0x61, 0x47, 0x16, 0x76, 0x50, 0xd8, 0x71, 0x19, 0x76, 0x64, 0x07, 0x3b, 0x68, 0xbc,
	0x63, 0x15, 0x76, 0x64, 0xff, 0xd9, 0x41, 0xfb, 0x1d, 0x32, 0xec, 0xc8, 0x78, 0xc7, 0xf7, 0x98,
	0x3f, 0xd4, 0xf5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x76, 0x69, 0x75, 0x02, 0x64, 0x03, 0x00,
	0x00,
}
