// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/common"
	oci_monitoring "github.com/oracle/oci-go-sdk/monitoring"
)

func AlarmHistoryCollectionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularAlarmHistoryCollection,
		Schema: map[string]*schema.Schema{
			"alarm_historytype": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"alarm_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"timestamp_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"timestamp_less_than": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"entries": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"summary": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"timestamp": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"timestamp_triggered": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"is_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func readSingularAlarmHistoryCollection(d *schema.ResourceData, m interface{}) error {
	sync := &AlarmHistoryCollectionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).monitoringClient

	return ReadResource(sync)
}

type AlarmHistoryCollectionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_monitoring.MonitoringClient
	Res    *oci_monitoring.GetAlarmHistoryResponse
}

func (s *AlarmHistoryCollectionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *AlarmHistoryCollectionDataSourceCrud) Get() error {
	request := oci_monitoring.GetAlarmHistoryRequest{}

	if alarmHistorytype, ok := s.D.GetOkExists("alarm_historytype"); ok {
		request.AlarmHistorytype = oci_monitoring.GetAlarmHistoryAlarmHistorytypeEnum(alarmHistorytype.(string))
	}

	if alarmId, ok := s.D.GetOkExists("alarm_id"); ok {
		tmp := alarmId.(string)
		request.AlarmId = &tmp
	}

	if timestampGreaterThanOrEqualTo, ok := s.D.GetOkExists("timestamp_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timestampGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimestampGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timestampLessThan, ok := s.D.GetOkExists("timestamp_less_than"); ok {
		tmp, err := time.Parse(time.RFC3339, timestampLessThan.(string))
		if err != nil {
			return err
		}
		request.TimestampLessThan = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "monitoring")

	response, err := s.Client.GetAlarmHistory(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *AlarmHistoryCollectionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	entries := []interface{}{}
	for _, item := range s.Res.Entries {
		entries = append(entries, AlarmHistoryEntryToMap(item))
	}
	s.D.Set("entries", entries)

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	return nil
}

func AlarmHistoryEntryToMap(obj oci_monitoring.AlarmHistoryEntry) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Summary != nil {
		result["summary"] = string(*obj.Summary)
	}

	if obj.Timestamp != nil {
		result["timestamp"] = obj.Timestamp.String()
	}

	if obj.TimestampTriggered != nil {
		result["timestamp_triggered"] = obj.TimestampTriggered.String()
	}

	return result
}
