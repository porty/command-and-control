
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

## Other JSON endpoints

```
http://192.168.0.1/goform/goform_get_cmd_process?isTest=false&cmd=wifi_coverage%2Cm_ssid_enable%2Cimei%2Cweb_version%2Cwa_inner_version%2Chardware_version%2CMAX_Access_num%2CSSID1%2Cm_SSID%2Cm_HideSSID%2Cm_MAX_Access_num%2Clan_ipaddr%2Cmac_address%2Cmsisdn%2CLocalDomain%2Cwan_ipaddr%2Cipv6_wan_ipaddr%2Cipv6_pdp_type%2Cpdp_type%2Cppp_status%2Csim_imsi%2Crssi%2Crscp%2Clte_rsrp%2Cnetwork_type&multi_data=1&_=1405508956431
```

```
{
  "wifi_coverage": "",
  "m_ssid_enable": "",
  "imei": "866948014610847",
  "web_version": "MF823_T03",
  "wa_inner_version": "MF823_T03",
  "hardware_version": "MF823-2.0.0",
  "MAX_Access_num": "",
  "SSID1": "",
  "m_SSID": "",
  "m_HideSSID": "",
  "m_MAX_Access_num": "",
  "lan_ipaddr": "192.168.0.1",
  "mac_address": "",
  "msisdn": "",
  "LocalDomain": "m.home",
  "wan_ipaddr": "10.96.78.97",
  "ipv6_wan_ipaddr": "3930:0000:0000:0000:948a:b820:4e9a:6a26",
  "ipv6_pdp_type": "IP",
  "pdp_type": "IP",
  "ppp_status": "ppp_connected",
  "sim_imsi": "505013457711705",
  "rssi": "",
  "rscp": "",
  "lte_rsrp": "-101",
  "network_type": "LTE"
}
```

## Custom JSON request

```
curl "http://192.168.0.1/goform/goform_get_cmd_process?isTest=false&multi_data=1&cmd=imei%2Csim_imsi%2Cwan_ipaddr%2Cipv6_wan_ipaddr%2Cppp_status%2Cnetwork_type%2Csignalbar%2Cnet_select"
```

```
{
  "imei": "866948014610847",
  "sim_imsi": "505013457711705",
  "wan_ipaddr": "10.96.78.97",
  "ipv6_wan_ipaddr": "3930:0000:0000:0000:948a:b820:4e9a:6a26",
  "ppp_status": "ppp_connected",
  "network_type": "LTE",
  "signalbar": "3"
  "net_select":"Only_LTE"
}
```

## Network Selection Info

```
http://192.168.0.1/goform/goform_get_cmd_process?isTest=false&cmd=current_network_mode%2Cm_netselect_save%2Cnet_select_mode%2Cm_netselect_contents%2Cnet_select%2Cppp_status%2Cmodem_main_state&multi_data=1&_=1405514324767
```

```
{
  "current_network_mode": "",
  "m_netselect_save": "",
  "net_select_mode": "auto_select",
  "m_netselect_contents": "",
  "net_select": "NETWORK_auto",
  "ppp_status": "ppp_connected",
  "modem_main_state": "modem_init_complete"
}
```

## Changing Network Preference

### To 4G Only

```
POST /goform/goform_set_cmd_process
Accept:application/json, text/javascript, */*; q=0.01
Content-Length:69
Content-Type:application/x-www-form-urlencoded; charset=UTF-8
Origin:http://192.168.0.1
Host:192.168.0.1

isTest=false&goformId=SET_BEARER_PREFERENCE&BearerPreference=Only_LTE


```

# Route dump example

```
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         192.168.0.1     0.0.0.0         UG    0      0        0 usb0
192.168.0.0     0.0.0.0         255.255.255.0   U     0      0        0 usb0
192.168.1.0     0.0.0.0         255.255.255.0   U     0      0        0 eth0
```
