
# Modem Status

URL: `http://192.168.0.1/goform/goform_get_cmd_process?multi_data=1&isTest=false&sms_received_flag_flag=0&sts_received_flag_flag=0&cmd=modem_main_state%2Cpin_status%2Cloginfo%2Cnew_version_state%2Ccurrent_upgrade_state%2Cis_mandatory%2Csms_received_flag%2Csts_received_flag%2Csignalbar%2Cnetwork_type%2Cnetwork_provider%2Cppp_status%2CEX_SSID1%2Cex_wifi_status%2CEX_wifi_profile%2Cm_ssid_enable%2Csms_unread_num%2CRadioOff%2Csimcard_roam%2Clan_ipaddr%2Cstation_mac%2Cbattery_charging%2Cbattery_vol_percent%2Cbattery_pers%2Cspn_display_flag%2Cplmn_display_flag%2Cspn_name_data%2Cspn_b1_flag%2Cspn_b2_flag%2Crealtime_tx_bytes%2Crealtime_rx_bytes%2Crealtime_time%2Crealtime_tx_thrpt%2Crealtime_rx_thrpt%2Ctotal_tx_bytes%2Ctotal_rx_bytes%2Ctotal_time%2Cmonthly_rx_bytes%2Cmonthly_tx_bytes%2Cmonthly_time%2Cdate_month%2Cdata_volume_limit_switch%2Cdata_volume_limit_size%2Cdata_volume_alert_percent%2Cdata_volume_limit_unit%2Croam_setting_option%2Cupg_roam_switch&_=1405307976294'`


## Connected (4G)

```
{
  "modem_main_state": "modem_init_complete",
  "pin_status": "0",
  "loginfo": "ok",
  "new_version_state": "0",
  "current_upgrade_state": "",
  "is_mandatory": "",
  "sms_received_flag": "",
  "sts_received_flag": "",
  "signalbar": "5",
  "network_type": "LTE",
  "network_provider": "Telstra",
  "ppp_status": "ppp_connected",
  "EX_SSID1": "",
  "ex_wifi_status": "",
  "EX_wifi_profile": "",
  "m_ssid_enable": "",
  "sms_unread_num": "3",
  "sms_dev_unread_num": "3",
  "sms_sim_unread_num": "0",
  "RadioOff": "",
  "simcard_roam": "Home",
  "lan_ipaddr": "192.168.0.1",
  "station_mac": "",
  "battery_charging": "",
  "battery_vol_percent": "",
  "battery_pers": "",
  "spn_display_flag": "0",
  "plmn_display_flag": "1",
  "spn_name_data": "00540065006C0073007400720061",
  "spn_b1_flag": "0",
  "spn_b2_flag": "1",
  "realtime_tx_bytes": "123064",
  "realtime_rx_bytes": "229674",
  "realtime_time": "720",
  "realtime_tx_thrpt": "0",
  "realtime_rx_thrpt": "0",
  "total_tx_bytes": "67138992",
  "total_rx_bytes": "350813936",
  "total_time": "47118",
  "monthly_rx_bytes": "350813936",
  "monthly_tx_bytes": "67138992",
  "monthly_time": "47118",
  "date_month": "0",
  "data_volume_limit_switch": "0",
  "data_volume_limit_size": "",
  "data_volume_alert_percent": "",
  "data_volume_limit_unit": "",
  "roam_setting_option": "off",
  "upg_roam_switch": "0"
}
```



## Disconnected

```
{
  "modem_main_state": "modem_init_complete",
  "pin_status": "0",
  "loginfo": "ok",
  "new_version_state": "0",
  "current_upgrade_state": "",
  "is_mandatory": "",
  "sms_received_flag": "",
  "sts_received_flag": "",
  "signalbar": "5",
  "network_type": "LTE",
  "network_provider": "Telstra",
  "ppp_status": "ppp_disconnected",
  "EX_SSID1": "",
  "ex_wifi_status": "",
  "EX_wifi_profile": "",
  "m_ssid_enable": "",
  "sms_unread_num": "3",
  "sms_dev_unread_num": "3",
  "sms_sim_unread_num": "0",
  "RadioOff": "",
  "simcard_roam": "Home",
  "lan_ipaddr": "192.168.0.1",
  "station_mac": "",
  "battery_charging": "",
  "battery_vol_percent": "",
  "battery_pers": "",
  "spn_display_flag": "0",
  "plmn_display_flag": "1",
  "spn_name_data": "00540065006C0073007400720061",
  "spn_b1_flag": "0",
  "spn_b2_flag": "1",
  "realtime_tx_bytes": "0",
  "realtime_rx_bytes": "0",
  "realtime_time": "0",
  "realtime_tx_thrpt": "0",
  "realtime_rx_thrpt": "0",
  "total_tx_bytes": "67565344",
  "total_rx_bytes": "352340114",
  "total_time": "49044",
  "monthly_rx_bytes": "352340114",
  "monthly_tx_bytes": "67565344",
  "monthly_time": "49044",
  "date_month": "0",
  "data_volume_limit_switch": "0",
  "data_volume_limit_size": "",
  "data_volume_alert_percent": "",
  "data_volume_limit_unit": "",
  "roam_setting_option": "off",
  "upg_roam_switch": "0"
}
```

## Forced 3G

```
{
  "modem_main_state": "modem_init_complete",
  "pin_status": "0",
  "loginfo": "ok",
  "new_version_state": "0",
  "current_upgrade_state": "",
  "is_mandatory": "",
  "sms_received_flag": "",
  "sts_received_flag": "",
  "signalbar": "5",
  "network_type": "UMTS",
  "network_provider": "Telstra",
  "ppp_status": "ppp_connected",
  "EX_SSID1": "",
  "ex_wifi_status": "",
  "EX_wifi_profile": "",
  "m_ssid_enable": "",
  "sms_unread_num": "3",
  "sms_dev_unread_num": "3",
  "sms_sim_unread_num": "0",
  "RadioOff": "",
  "simcard_roam": "Home",
  "lan_ipaddr": "192.168.0.1",
  "station_mac": "",
  "battery_charging": "",
  "battery_vol_percent": "",
  "battery_pers": "",
  "spn_display_flag": "0",
  "plmn_display_flag": "1",
  "spn_name_data": "00540065006C0073007400720061",
  "spn_b1_flag": "0",
  "spn_b2_flag": "1",
  "realtime_tx_bytes": "0",
  "realtime_rx_bytes": "0",
  "realtime_time": "15",
  "realtime_tx_thrpt": "0",
  "realtime_rx_thrpt": "0",
  "total_tx_bytes": "67565344",
  "total_rx_bytes": "352340114",
  "total_time": "49071",
  "monthly_rx_bytes": "352340114",
  "monthly_tx_bytes": "67565344",
  "monthly_time": "49071",
  "date_month": "0",
  "data_volume_limit_switch": "0",
  "data_volume_limit_size": "",
  "data_volume_alert_percent": "",
  "data_volume_limit_unit": "",
  "roam_setting_option": "off",
  "upg_roam_switch": "0"
}
```

## Forced 2G

```
{
  "modem_main_state": "modem_init_complete",
  "pin_status": "0",
  "loginfo": "ok",
  "new_version_state": "0",
  "current_upgrade_state": "",
  "is_mandatory": "",
  "sms_received_flag": "",
  "sts_received_flag": "",
  "signalbar": "5",
  "network_type": "EDGE",
  "network_provider": "Telstra",
  "ppp_status": "ppp_connected",
  "EX_SSID1": "",
  "ex_wifi_status": "",
  "EX_wifi_profile": "",
  "m_ssid_enable": "",
  "sms_unread_num": "3",
  "sms_dev_unread_num": "3",
  "sms_sim_unread_num": "0",
  "RadioOff": "",
  "simcard_roam": "Home",
  "lan_ipaddr": "192.168.0.1",
  "station_mac": "",
  "battery_charging": "",
  "battery_vol_percent": "",
  "battery_pers": "",
  "spn_display_flag": "0",
  "plmn_display_flag": "1",
  "spn_name_data": "00540065006C0073007400720061",
  "spn_b1_flag": "0",
  "spn_b2_flag": "1",
  "realtime_tx_bytes": "0",
  "realtime_rx_bytes": "0",
  "realtime_time": "9",
  "realtime_tx_thrpt": "0",
  "realtime_rx_thrpt": "0",
  "total_tx_bytes": "67565344",
  "total_rx_bytes": "352340114",
  "total_time": "49128",
  "monthly_rx_bytes": "352340114",
  "monthly_tx_bytes": "67565344",
  "monthly_time": "49128",
  "date_month": "0",
  "data_volume_limit_switch": "0",
  "data_volume_limit_size": "",
  "data_volume_alert_percent": "",
  "data_volume_limit_unit": "",
  "roam_setting_option": "off",
  "upg_roam_switch": "0"
}
```


# Route dump example

```
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         192.168.0.1     0.0.0.0         UG    0      0        0 usb0
192.168.0.0     0.0.0.0         255.255.255.0   U     0      0        0 usb0
192.168.1.0     0.0.0.0         255.255.255.0   U     0      0        0 eth0
```
