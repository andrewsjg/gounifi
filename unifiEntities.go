package gounifi

//AuthResponse - Response back from an auth request. The data field is always empty.
type AuthResponse struct {
	Meta struct {
		Rc string `json:"rc"`
	} `json:"meta"`
	Data []interface{}
}

//SiteHealth - Contains data from all configured subsystems. Requires parsing to determine if subsytem data is returned
//			   in the API data.
type SiteHealth struct {
	Meta struct {
		Rc string `json:"rc"`
	} `json:"meta"`
	Data []struct {
		Subsystem       string   `json:"subsystem"`
		NumUser         int      `json:"num_user,omitempty"`
		NumGuest        int      `json:"num_guest,omitempty"`
		NumIot          int      `json:"num_iot,omitempty"`
		TxBytesR        int      `json:"tx_bytes-r,omitempty"`
		RxBytesR        int      `json:"rx_bytes-r,omitempty"`
		Status          string   `json:"status"`
		NumAp           int      `json:"num_ap,omitempty"`
		NumAdopted      int      `json:"num_adopted,omitempty"`
		NumDisabled     int      `json:"num_disabled,omitempty"`
		NumDisconnected int      `json:"num_disconnected,omitempty"`
		NumPending      int      `json:"num_pending,omitempty"`
		NumGw           int      `json:"num_gw,omitempty"`
		WanIP           string   `json:"wan_ip,omitempty"`
		Gateways        []string `json:"gateways,omitempty"`
		Netmask         string   `json:"netmask,omitempty"`
		Nameservers     []string `json:"nameservers,omitempty"`
		NumSta          int      `json:"num_sta,omitempty"`
		GwMac           string   `json:"gw_mac,omitempty"`
		GwName          string   `json:"gw_name,omitempty"`
		GwSystemStats   struct {
			CPU    string `json:"cpu"`
			Mem    string `json:"mem"`
			Uptime string `json:"uptime"`
		} `json:"gw_system-stats,omitempty"`
		GwVersion             string  `json:"gw_version,omitempty"`
		Latency               int     `json:"latency,omitempty"`
		Uptime                int     `json:"uptime,omitempty"`
		Drops                 int     `json:"drops,omitempty"`
		XputUp                float64 `json:"xput_up,omitempty"`
		XputDown              float64 `json:"xput_down,omitempty"`
		SpeedtestStatus       string  `json:"speedtest_status,omitempty"`
		SpeedtestLastrun      int     `json:"speedtest_lastrun,omitempty"`
		SpeedtestPing         int     `json:"speedtest_ping,omitempty"`
		LanIP                 string  `json:"lan_ip,omitempty"`
		NumSw                 int     `json:"num_sw,omitempty"`
		RemoteUserEnabled     bool    `json:"remote_user_enabled,omitempty"`
		RemoteUserNumActive   int     `json:"remote_user_num_active,omitempty"`
		RemoteUserNumInactive int     `json:"remote_user_num_inactive,omitempty"`
		RemoteUserRxBytes     int     `json:"remote_user_rx_bytes,omitempty"`
		RemoteUserTxBytes     int     `json:"remote_user_tx_bytes,omitempty"`
		RemoteUserRxPackets   int     `json:"remote_user_rx_packets,omitempty"`
		RemoteUserTxPackets   int     `json:"remote_user_tx_packets,omitempty"`
		SiteToSiteEnabled     bool    `json:"site_to_site_enabled,omitempty"`
	} `json:"data"`
}

//SysInfo - Information about the controller
type SysInfo struct {
	Meta struct {
		Rc string `json:"rc"`
	} `json:"meta"`
	Data []struct {
		Timezone                                 string   `json:"timezone"`
		Autobackup                               bool     `json:"autobackup"`
		Build                                    string   `json:"build"`
		Version                                  string   `json:"version"`
		PreviousVersion                          string   `json:"previous_version"`
		DebugMgmt                                string   `json:"debug_mgmt"`
		DebugSystem                              string   `json:"debug_system"`
		DebugDevice                              string   `json:"debug_device"`
		DebugSdn                                 string   `json:"debug_sdn"`
		DataRetentionDays                        int      `json:"data_retention_days"`
		DataRetentionTimeInHoursFor5MinutesScale int      `json:"data_retention_time_in_hours_for_5minutes_scale"`
		DataRetentionTimeInHoursForHourlyScale   int      `json:"data_retention_time_in_hours_for_hourly_scale"`
		DataRetentionTimeInHoursForDailyScale    int      `json:"data_retention_time_in_hours_for_daily_scale"`
		DataRetentionTimeInHoursForMonthlyScale  int      `json:"data_retention_time_in_hours_for_monthly_scale"`
		DataRetentionTimeInHoursForOthers        int      `json:"data_retention_time_in_hours_for_others"`
		UpdateAvailable                          bool     `json:"update_available"`
		UpdateDownloaded                         bool     `json:"update_downloaded"`
		LiveChat                                 string   `json:"live_chat"`
		StoreEnabled                             string   `json:"store_enabled"`
		Hostname                                 string   `json:"hostname"`
		Name                                     string   `json:"name"`
		IPAddrs                                  []string `json:"ip_addrs"`
		InformPort                               int      `json:"inform_port"`
		HTTPSPort                                int      `json:"https_port"`
		OverrideInformHost                       bool     `json:"override_inform_host"`
		ImageMapsUseGoogleEngine                 bool     `json:"image_maps_use_google_engine"`
		RadiusDisconnectRunning                  bool     `json:"radius_disconnect_running"`
		FacebookWifiRegistered                   bool     `json:"facebook_wifi_registered"`
		SsoAppID                                 string   `json:"sso_app_id"`
		SsoAppSec                                string   `json:"sso_app_sec"`
		UnsupportedDeviceCount                   int      `json:"unsupported_device_count"`
		UnifiGoEnabled                           bool     `json:"unifi_go_enabled"`
		DefaultSiteDeviceAuthPasswordAlert       bool     `json:"default_site_device_auth_password_alert"`
	} `json:"data"`
}

//ActiveClients - List of all active clients on the site
type ActiveClients struct {
	Meta struct {
		Rc string `json:"rc"`
	} `json:"meta"`
	Clients []struct {
		SiteID              string `json:"site_id"`
		AssocTime           int    `json:"assoc_time"`
		LatestAssocTime     int    `json:"latest_assoc_time"`
		Oui                 string `json:"oui"`
		UserID              string `json:"user_id"`
		ID                  string `json:"_id"`
		Mac                 string `json:"mac"`
		IsGuest             bool   `json:"is_guest"`
		FirstSeen           int    `json:"first_seen"`
		LastSeen            int    `json:"last_seen"`
		IsWired             bool   `json:"is_wired"`
		Hostname            string `json:"hostname,omitempty"`
		UsergroupID         string `json:"usergroup_id,omitempty"`
		Name                string `json:"name,omitempty"`
		Noted               bool   `json:"noted,omitempty"`
		FingerprintOverride bool   `json:"fingerprint_override,omitempty"`
		DevIDOverride       int    `json:"dev_id_override,omitempty"`
		Blocked             bool   `json:"blocked,omitempty"`
		UptimeByUap         int    `json:"_uptime_by_uap,omitempty"`
		LastSeenByUap       int    `json:"_last_seen_by_uap,omitempty"`
		IsGuestByUap        bool   `json:"_is_guest_by_uap,omitempty"`
		ApMac               string `json:"ap_mac,omitempty"`
		Channel             int    `json:"channel,omitempty"`
		Radio               string `json:"radio,omitempty"`
		RadioName           string `json:"radio_name,omitempty"`
		Essid               string `json:"essid,omitempty"`
		Bssid               string `json:"bssid,omitempty"`
		PowersaveEnabled    bool   `json:"powersave_enabled,omitempty"`
		Is11R               bool   `json:"is_11r,omitempty"`
		Ccq                 int    `json:"ccq,omitempty"`
		Rssi                int    `json:"rssi,omitempty"`
		Noise               int    `json:"noise,omitempty"`
		Signal              int    `json:"signal,omitempty"`
		TxRate              int    `json:"tx_rate,omitempty"`
		RxRate              int    `json:"rx_rate,omitempty"`
		TxPower             int    `json:"tx_power,omitempty"`
		Idletime            int    `json:"idletime,omitempty"`
		IP                  string `json:"ip"`
		DhcpendTime         int    `json:"dhcpend_time,omitempty"`
		Satisfaction        int    `json:"satisfaction"`
		Anomalies           int    `json:"anomalies,omitempty"`
		Vlan                int    `json:"vlan,omitempty"`
		RadioProto          string `json:"radio_proto,omitempty"`
		Uptime              int    `json:"uptime"`
		TxBytes             int    `json:"tx_bytes"`
		RxBytes             int    `json:"rx_bytes"`
		TxPackets           int    `json:"tx_packets"`
		TxRetries           int    `json:"tx_retries"`
		WifiTxAttempts      int    `json:"wifi_tx_attempts"`
		RxPackets           int    `json:"rx_packets"`
		BytesR              int    `json:"bytes-r"`
		TxBytesR            int    `json:"tx_bytes-r"`
		RxBytesR            int    `json:"rx_bytes-r"`
		Authorized          bool   `json:"authorized"`
		QosPolicyApplied    bool   `json:"qos_policy_applied"`
		UptimeByUsw         int    `json:"_uptime_by_usw"`
		LastSeenByUsw       int    `json:"_last_seen_by_usw"`
		IsGuestByUsw        bool   `json:"_is_guest_by_usw"`
		SwMac               string `json:"sw_mac"`
		SwDepth             int    `json:"sw_depth"`
		SwPort              int    `json:"sw_port"`
		Network             string `json:"network"`
		NetworkID           string `json:"network_id"`
		UptimeByUgw         int    `json:"_uptime_by_ugw"`
		LastSeenByUgw       int    `json:"_last_seen_by_ugw"`
		IsGuestByUgw        bool   `json:"_is_guest_by_ugw"`
		GwMac               string `json:"gw_mac"`
		WiredTxBytes        int    `json:"wired-tx_bytes,omitempty"`
		WiredRxBytes        int    `json:"wired-rx_bytes,omitempty"`
		WiredTxPackets      int    `json:"wired-tx_packets,omitempty"`
		WiredRxPackets      int    `json:"wired-rx_packets,omitempty"`
		WiredTxBytesR       int    `json:"wired-tx_bytes-r,omitempty"`
		WiredRxBytesR       int    `json:"wired-rx_bytes-r,omitempty"`
		UseFixedip          bool   `json:"use_fixedip,omitempty"`
		FixedIP             string `json:"fixed_ip,omitempty"`
		DevCat              int    `json:"dev_cat,omitempty"`
		DevFamily           int    `json:"dev_family,omitempty"`
		DevID               int    `json:"dev_id,omitempty"`
		OsClass             int    `json:"os_class,omitempty"`
		OsName              int    `json:"os_name,omitempty"`
		DevVendor           int    `json:"dev_vendor,omitempty"`
		Note                string `json:"note,omitempty"`
	} `json:"data"`
}

// KnownClients - List of all configured/known clients on the site
type KnownClients struct {
	Meta struct {
		Rc string `json:"rc"`
	} `json:"meta"`
	Clients []struct {
		ID                  string `json:"_id"`
		Mac                 string `json:"mac"`
		SiteID              string `json:"site_id"`
		Oui                 string `json:"oui,omitempty"`
		IsGuest             bool   `json:"is_guest,omitempty"`
		FirstSeen           int    `json:"first_seen,omitempty"`
		LastSeen            int    `json:"last_seen,omitempty"`
		IsWired             bool   `json:"is_wired,omitempty"`
		Hostname            string `json:"hostname,omitempty"`
		Blocked             bool   `json:"blocked,omitempty"`
		FingerprintOverride bool   `json:"fingerprint_override,omitempty"`
		DevIDOverride       int    `json:"dev_id_override,omitempty"`
		UsergroupID         string `json:"usergroup_id,omitempty"`
		Name                string `json:"name,omitempty"`
		Noted               bool   `json:"noted,omitempty"`
		UseFixedip          bool   `json:"use_fixedip,omitempty"`
		NetworkID           string `json:"network_id,omitempty"`
		FixedIP             string `json:"fixed_ip,omitempty"`
		Note                string `json:"note,omitempty"`
	} `json:"data"`
}

//SiteSettings - Detailed site settings
type SiteSettings struct {
	Meta struct {
		Rc string `json:"rc"`
	} `json:"meta"`
	Data []struct {
		ID                                       string        `json:"_id"`
		Key                                      string        `json:"key"`
		Name                                     string        `json:"name,omitempty"`
		Hostname                                 string        `json:"hostname,omitempty"`
		Port                                     interface{}   `json:"port,omitempty"`
		FingerbankKey                            string        `json:"fingerbank_key,omitempty"`
		Discoverable                             bool          `json:"discoverable,omitempty"`
		DataRetentionTimeEnabled                 bool          `json:"data_retention_time_enabled,omitempty"`
		DataRetentionTimeInHoursFor5MinutesScale int           `json:"data_retention_time_in_hours_for_5minutes_scale,omitempty"`
		DataRetentionTimeInHoursForHourlyScale   int           `json:"data_retention_time_in_hours_for_hourly_scale,omitempty"`
		DataRetentionTimeInHoursForDailyScale    int           `json:"data_retention_time_in_hours_for_daily_scale,omitempty"`
		DataRetentionTimeInHoursForMonthlyScale  int           `json:"data_retention_time_in_hours_for_monthly_scale,omitempty"`
		DataRetentionTimeInHoursForOthers        int           `json:"data_retention_time_in_hours_for_others,omitempty"`
		TimeSeriesPerClientStatsEnabled          bool          `json:"time_series_per_client_stats_enabled,omitempty"`
		AutobackupTimezone                       string        `json:"autobackup_timezone,omitempty"`
		AutobackupEnabled                        bool          `json:"autobackup_enabled,omitempty"`
		AutobackupDays                           int           `json:"autobackup_days,omitempty"`
		AutobackupCronExpr                       string        `json:"autobackup_cron_expr,omitempty"`
		StoreEnabled                             string        `json:"store_enabled,omitempty"`
		LiveChat                                 string        `json:"live_chat,omitempty"`
		AutobackupMaxFiles                       int           `json:"autobackup_max_files,omitempty"`
		MinimumUsableHdSpace                     int           `json:"minimum_usable_hd_space,omitempty"`
		LiveUpdates                              string        `json:"live_updates,omitempty"`
		EnableAnalytics                          bool          `json:"enable_analytics,omitempty"`
		AnalyticsDisapprovedFor                  string        `json:"analytics_disapproved_for,omitempty"`
		BackupToCloudEnabled                     bool          `json:"backup_to_cloud_enabled,omitempty"`
		OverrideInformHost                       bool          `json:"override_inform_host,omitempty"`
		SiteID                                   string        `json:"site_id,omitempty"`
		Ugw3Wan2Enabled                          bool          `json:"ugw3_wan2_enabled,omitempty"`
		Interval                                 int           `json:"interval,omitempty"`
		Enabled                                  bool          `json:"enabled,omitempty"`
		DhcpSnoop                                bool          `json:"dhcp_snoop,omitempty"`
		Community                                string        `json:"community,omitempty"`
		NtpServer1                               string        `json:"ntp_server_1,omitempty"`
		NtpServer2                               string        `json:"ntp_server_2,omitempty"`
		NtpServer3                               string        `json:"ntp_server_3,omitempty"`
		NtpServer4                               string        `json:"ntp_server_4,omitempty"`
		Code                                     string        `json:"code,omitempty"`
		FingerprintingEnabled                    bool          `json:"fingerprintingEnabled,omitempty"`
		Timezone                                 string        `json:"timezone,omitempty"`
		AdvancedFeatureEnabled                   bool          `json:"advanced_feature_enabled,omitempty"`
		UnifiIdpEnabled                          bool          `json:"unifi_idp_enabled,omitempty"`
		XMgmtKey                                 string        `json:"x_mgmt_key,omitempty"`
		XSSHAuthPasswordEnabled                  bool          `json:"x_ssh_auth_password_enabled,omitempty"`
		XSSHBindWildcard                         bool          `json:"x_ssh_bind_wildcard,omitempty"`
		XSSHEnabled                              bool          `json:"x_ssh_enabled,omitempty"`
		XSSHPassword                             string        `json:"x_ssh_password,omitempty"`
		XSSHUsername                             string        `json:"x_ssh_username,omitempty"`
		XSSHSha512Passwd                         string        `json:"x_ssh_sha512passwd,omitempty"`
		LedEnabled                               bool          `json:"led_enabled,omitempty"`
		AlertEnabled                             bool          `json:"alert_enabled,omitempty"`
		XSSHKeys                                 []interface{} `json:"x_ssh_keys,omitempty"`
		AutoUpgrade                              bool          `json:"auto_upgrade,omitempty"`
		UplinkType                               string        `json:"uplink_type,omitempty"`
		XMeshEssid                               string        `json:"x_mesh_essid,omitempty"`
		XMeshPsk                                 string        `json:"x_mesh_psk,omitempty"`
		IP                                       string        `json:"ip,omitempty"`
		Debug                                    bool          `json:"debug,omitempty"`
		ThisController                           bool          `json:"this_controller,omitempty"`
		ThisControllerEncryptedOnly              bool          `json:"this_controller_encrypted_only,omitempty"`
		NetconsoleEnabled                        bool          `json:"netconsole_enabled,omitempty"`
		NetconsoleHost                           string        `json:"netconsole_host,omitempty"`
		NetconsolePort                           string        `json:"netconsole_port,omitempty"`
		AuthPort                                 int           `json:"auth_port,omitempty"`
		AcctPort                                 int           `json:"acct_port,omitempty"`
		InterimUpdateInterval                    int           `json:"interim_update_interval,omitempty"`
		ConfigureWholeNetwork                    bool          `json:"configure_whole_network,omitempty"`
		TunneledReply                            bool          `json:"tunneled_reply,omitempty"`
		XSecret                                  string        `json:"x_secret,omitempty"`
		Download                                 int           `json:"download,omitempty"`
		Upload                                   int           `json:"upload,omitempty"`
		Brightness                               int           `json:"brightness,omitempty"`
		IdleTimeout                              int           `json:"idle_timeout,omitempty"`
		Sync                                     bool          `json:"sync,omitempty"`
		TouchEvent                               bool          `json:"touch_event,omitempty"`
		Auth                                     string        `json:"auth,omitempty"`
		RedirectHTTPS                            bool          `json:"redirect_https,omitempty"`
		RestrictedSubnet1                        string        `json:"restricted_subnet_1,omitempty"`
		RestrictedSubnet2                        string        `json:"restricted_subnet_2,omitempty"`
		RestrictedSubnet3                        string        `json:"restricted_subnet_3,omitempty"`
		PortalCustomizedLogoEnabled              bool          `json:"portal_customized_logo_enabled,omitempty"`
		PortalCustomizedTextColor                string        `json:"portal_customized_text_color,omitempty"`
		PortalCustomizedTitle                    string        `json:"portal_customized_title,omitempty"`
		PortalCustomizedButtonColor              string        `json:"portal_customized_button_color,omitempty"`
		PortalEnabled                            bool          `json:"portal_enabled,omitempty"`
		PortalCustomizedLanguages                []string      `json:"portal_customized_languages,omitempty"`
		PortalCustomizedBgImageTile              bool          `json:"portal_customized_bg_image_tile,omitempty"`
		PortalCustomizedButtonTextColor          string        `json:"portal_customized_button_text_color,omitempty"`
		PortalCustomizedBgImageEnabled           bool          `json:"portal_customized_bg_image_enabled,omitempty"`
		PortalCustomizedBoxLinkColor             string        `json:"portal_customized_box_link_color,omitempty"`
		PortalCustomizedLinkColor                string        `json:"portal_customized_link_color,omitempty"`
		XPassword                                string        `json:"x_password,omitempty"`
		RedirectEnabled                          bool          `json:"redirect_enabled,omitempty"`
		PortalCustomizedBgColor                  string        `json:"portal_customized_bg_color,omitempty"`
		VoucherEnabled                           bool          `json:"voucher_enabled,omitempty"`
		PortalCustomizedBoxOpacity               int           `json:"portal_customized_box_opacity,omitempty"`
		FacebookWifiGwName                       string        `json:"facebook_wifi_gw_name,omitempty"`
		Expire                                   int           `json:"expire,omitempty"`
		PortalCustomizedBoxTextColor             string        `json:"portal_customized_box_text_color,omitempty"`
		TemplateEngine                           string        `json:"template_engine,omitempty"`
		PortalCustomizedBoxColor                 string        `json:"portal_customized_box_color,omitempty"`
		RedirectURL                              string        `json:"redirect_url,omitempty"`
		RadiusAuthType                           string        `json:"radius_auth_type,omitempty"`
		RestrictedSubnet4                        string        `json:"restricted_subnet_4,omitempty"`
		PortalCustomizedTos                      string        `json:"portal_customized_tos,omitempty"`
		PortalCustomizedWelcomeText              string        `json:"portal_customized_welcome_text,omitempty"`
		Default                                  bool          `json:"default,omitempty"`
		ChannelsNg                               []string      `json:"channels_ng,omitempty"`
		CronExpr                                 string        `json:"cron_expr,omitempty"`
		ExcludeDevices                           []interface{} `json:"exclude_devices,omitempty"`
		HtModesNa                                []string      `json:"ht_modes_na,omitempty"`
		HtModesNg                                []string      `json:"ht_modes_ng,omitempty"`
		Radios                                   []string      `json:"radios,omitempty"`
		Optimize                                 []string      `json:"optimize,omitempty"`
		ChannelsNa                               []string      `json:"channels_na,omitempty"`
		FtpModule                                bool          `json:"ftp_module,omitempty"`
		GreModule                                bool          `json:"gre_module,omitempty"`
		H323Module                               bool          `json:"h323_module,omitempty"`
		PptpModule                               bool          `json:"pptp_module,omitempty"`
		SipModule                                bool          `json:"sip_module,omitempty"`
		TftpModule                               bool          `json:"tftp_module,omitempty"`
		BroadcastPing                            bool          `json:"broadcast_ping,omitempty"`
		ReceiveRedirects                         bool          `json:"receive_redirects,omitempty"`
		SendRedirects                            bool          `json:"send_redirects,omitempty"`
		SynCookies                               bool          `json:"syn_cookies,omitempty"`
		OffloadAccounting                        bool          `json:"offload_accounting,omitempty"`
		OffloadSch                               bool          `json:"offload_sch,omitempty"`
		OffloadL2Blocking                        bool          `json:"offload_l2_blocking,omitempty"`
		MdnsEnabled                              bool          `json:"mdns_enabled,omitempty"`
		UpnpEnabled                              bool          `json:"upnp_enabled,omitempty"`
		UpnpNatPmpEnabled                        bool          `json:"upnp_nat_pmp_enabled,omitempty"`
		UpnpSecureMode                           bool          `json:"upnp_secure_mode,omitempty"`
		MssClamp                                 string        `json:"mss_clamp,omitempty"`
		DhcpdHostfileUpdate                      bool          `json:"dhcpd_hostfile_update,omitempty"`
		DhcpdUseDnsmasq                          bool          `json:"dhcpd_use_dnsmasq,omitempty"`
		GeoIPFilteringEnabled                    bool          `json:"geo_ip_filtering_enabled,omitempty"`
		GeoIPFilteringBlock                      string        `json:"geo_ip_filtering_block,omitempty"`
		GeoIPFilteringCountries                  string        `json:"geo_ip_filtering_countries,omitempty"`
		GeoIPFilteringTrafficDirection           string        `json:"geo_ip_filtering_traffic_direction,omitempty"`
		IcmpTimeout                              int           `json:"icmp_timeout,omitempty"`
		OtherTimeout                             int           `json:"other_timeout,omitempty"`
		TCPCloseTimeout                          int           `json:"tcp_close_timeout,omitempty"`
		TCPCloseWaitTimeout                      int           `json:"tcp_close_wait_timeout,omitempty"`
		TCPEstablishedTimeout                    int           `json:"tcp_established_timeout,omitempty"`
		TCPFinWaitTimeout                        int           `json:"tcp_fin_wait_timeout,omitempty"`
		TCPLastAckTimeout                        int           `json:"tcp_last_ack_timeout,omitempty"`
		TCPSynRecvTimeout                        int           `json:"tcp_syn_recv_timeout,omitempty"`
		TCPSynSentTimeout                        int           `json:"tcp_syn_sent_timeout,omitempty"`
		TCPTimeWaitTimeout                       int           `json:"tcp_time_wait_timeout,omitempty"`
		UDPOtherTimeout                          int           `json:"udp_other_timeout,omitempty"`
		UDPStreamTimeout                         int           `json:"udp_stream_timeout,omitempty"`
		FirewallWanDefaultLog                    bool          `json:"firewall_wan_default_log,omitempty"`
		FirewallLanDefaultLog                    bool          `json:"firewall_lan_default_log,omitempty"`
		FirewallGuestDefaultLog                  bool          `json:"firewall_guest_default_log,omitempty"`
		UpnpWanInterface                         string        `json:"upnp_wan_interface,omitempty"`
		DhcpRelayHopCount                        int           `json:"dhcp_relay_hop_count,omitempty"`
		DhcpRelayMaxSize                         int           `json:"dhcp_relay_max_size,omitempty"`
		DhcpRelayPort                            int           `json:"dhcp_relay_port,omitempty"`
		DhcpRelayAgentsPackets                   string        `json:"dhcp_relay_agents_packets,omitempty"`
	} `json:"data"`
}

//Routing - Data about configured routes
type Routing struct {
	Meta struct {
		Rc string `json:"rc"`
	} `json:"meta"`
	Data []struct {
		Nh []struct {
			Intf   string `json:"intf"`
			Metric string `json:"metric"`
			T      string `json:"t"`
			Via    string `json:"via"`
		} `json:"nh"`
		Pfx string `json:"pfx"`
	} `json:"data"`
}

//FirewallRules - User defined firewall rules
type FirewallRules struct {
	Meta struct {
		Rc string `json:"rc"`
	} `json:"meta"`
	Data []struct {
		ID                    string        `json:"_id"`
		Ruleset               string        `json:"ruleset"`
		RuleIndex             int           `json:"rule_index"`
		Name                  string        `json:"name"`
		Enabled               bool          `json:"enabled"`
		Action                string        `json:"action"`
		ProtocolMatchExcepted bool          `json:"protocol_match_excepted"`
		Logging               bool          `json:"logging"`
		StateNew              bool          `json:"state_new"`
		StateEstablished      bool          `json:"state_established"`
		StateInvalid          bool          `json:"state_invalid"`
		StateRelated          bool          `json:"state_related"`
		Ipsec                 string        `json:"ipsec"`
		SrcFirewallgroupIds   []interface{} `json:"src_firewallgroup_ids"`
		SrcMacAddress         string        `json:"src_mac_address"`
		DstFirewallgroupIds   []interface{} `json:"dst_firewallgroup_ids"`
		DstAddress            string        `json:"dst_address"`
		SrcAddress            string        `json:"src_address"`
		Protocol              string        `json:"protocol"`
		IcmpTypename          string        `json:"icmp_typename"`
		SrcNetworkconfID      string        `json:"src_networkconf_id"`
		SrcNetworkconfType    string        `json:"src_networkconf_type"`
		DstNetworkconfID      string        `json:"dst_networkconf_id"`
		DstNetworkconfType    string        `json:"dst_networkconf_type"`
		SiteID                string        `json:"site_id"`
		DstPort               string        `json:"dst_port,omitempty"`
	} `json:"data"`
}

//FirewallGroups - Information about firewall groups
type FirewallGroups struct {
	Meta struct {
		Rc string `json:"rc"`
	} `json:"meta"`
	Data []struct {
		ID           string   `json:"_id"`
		Name         string   `json:"name"`
		GroupType    string   `json:"group_type"`
		GroupMembers []string `json:"group_members"`
		SiteID       string   `json:"site_id"`
	} `json:"data"`
}

//WLANConf - Wireless LAN Configurations
type WLANConf struct {
	Meta struct {
		Rc string `json:"rc"`
	} `json:"meta"`
	Data []struct {
		ID                        string        `json:"_id"`
		XIappKey                  string        `json:"x_iapp_key"`
		WpaMode                   string        `json:"wpa_mode"`
		MinrateNaAdvertisingRates bool          `json:"minrate_na_advertising_rates"`
		DtimNa                    int           `json:"dtim_na"`
		IsGuest                   bool          `json:"is_guest"`
		MinrateNaEnabled          bool          `json:"minrate_na_enabled"`
		MinrateNgAdvertisingRates bool          `json:"minrate_ng_advertising_rates"`
		Enabled                   bool          `json:"enabled"`
		MacFilterPolicy           string        `json:"mac_filter_policy"`
		Security                  string        `json:"security"`
		WepIdx                    int           `json:"wep_idx"`
		GroupRekey                int           `json:"group_rekey"`
		MinrateNgEnabled          bool          `json:"minrate_ng_enabled"`
		MinrateNgDataRateKbps     int           `json:"minrate_ng_data_rate_kbps"`
		WpaEnc                    string        `json:"wpa_enc"`
		WlangroupID               string        `json:"wlangroup_id"`
		BcFilterEnabled           bool          `json:"bc_filter_enabled"`
		XPassphrase               string        `json:"x_passphrase"`
		MinrateNaBeaconRateKbps   int           `json:"minrate_na_beacon_rate_kbps"`
		UsergroupID               string        `json:"usergroup_id"`
		MacFilterList             []interface{} `json:"mac_filter_list"`
		MinrateNaMgmtRateKbps     int           `json:"minrate_na_mgmt_rate_kbps"`
		DtimMode                  string        `json:"dtim_mode"`
		Schedule                  []interface{} `json:"schedule"`
		MinrateNgBeaconRateKbps   int           `json:"minrate_ng_beacon_rate_kbps"`
		MinrateNgMgmtRateKbps     int           `json:"minrate_ng_mgmt_rate_kbps"`
		BcFilterList              []interface{} `json:"bc_filter_list"`
		Name                      string        `json:"name"`
		SiteID                    string        `json:"site_id"`
		MinrateNaDataRateKbps     int           `json:"minrate_na_data_rate_kbps"`
		HideSsid                  bool          `json:"hide_ssid"`
		MacFilterEnabled          bool          `json:"mac_filter_enabled"`
		DtimNg                    int           `json:"dtim_ng"`
		Vlan                      string        `json:"vlan"`
		VlanEnabled               bool          `json:"vlan_enabled"`
		No2GhzOui                 bool          `json:"no2ghz_oui"`
		MinrateNgCckRatesEnabled  bool          `json:"minrate_ng_cck_rates_enabled"`
		NameCombineEnabled        bool          `json:"name_combine_enabled"`
		ScheduleEnabled           bool          `json:"schedule_enabled,omitempty"`
		ScheduleReversed          bool          `json:"schedule_reversed,omitempty"`
		McastenhanceEnabled       bool          `json:"mcastenhance_enabled,omitempty"`
		FastRoamingEnabled        bool          `json:"fast_roaming_enabled,omitempty"`
		UapsdEnabled              bool          `json:"uapsd_enabled,omitempty"`
		NameCombineSuffix         string        `json:"name_combine_suffix,omitempty"`
		RadiusMacAuthEnabled      bool          `json:"radius_mac_auth_enabled,omitempty"`
		RadiusMacaclFormat        string        `json:"radius_macacl_format,omitempty"`
		RadiusMacaclEmptyPassword bool          `json:"radius_macacl_empty_password,omitempty"`
		RadiusDasEnabled          bool          `json:"radius_das_enabled,omitempty"`
	} `json:"data"`
}

//RogueAPs - AP's not part of the network that have been seen
type RogueAPs struct {
	Meta struct {
		Rc string `json:"rc"`
	} `json:"meta"`
	Data []struct {
		ID         string `json:"_id"`
		ApMac      string `json:"ap_mac"`
		Bssid      string `json:"bssid"`
		SiteID     string `json:"site_id"`
		Age        int    `json:"age"`
		Band       string `json:"band"`
		Bw         int    `json:"bw"`
		CenterFreq int    `json:"center_freq"`
		Channel    int    `json:"channel"`
		Essid      string `json:"essid"`
		Freq       int    `json:"freq"`
		IsAdhoc    bool   `json:"is_adhoc"`
		IsUbnt     bool   `json:"is_ubnt"`
		Noise      int    `json:"noise"`
		Rssi       int    `json:"rssi"`
		RssiAge    int    `json:"rssi_age"`
		Security   string `json:"security"`
		Signal     int    `json:"signal"`
		Radio      string `json:"radio"`
		RadioName  string `json:"radio_name"`
		LastSeen   int    `json:"last_seen"`
		ReportTime int    `json:"report_time"`
		IsRogue    bool   `json:"is_rogue"`
		Oui        string `json:"oui"`
	} `json:"data"`
}

//PortProfiles - Configured port profiles
type PortProfiles struct {
	Meta struct {
		Rc string `json:"rc"`
	} `json:"meta"`
	Data []struct {
		ID                  string `json:"_id"`
		SiteID              string `json:"site_id"`
		Name                string `json:"name"`
		Forward             string `json:"forward"`
		AttrHiddenID        string `json:"attr_hidden_id,omitempty"`
		AttrHidden          bool   `json:"attr_hidden,omitempty"`
		AttrNoDelete        bool   `json:"attr_no_delete,omitempty"`
		AttrNoEdit          bool   `json:"attr_no_edit"`
		NativeNetworkconfID string `json:"native_networkconf_id,omitempty"`
		Isolation           bool   `json:"isolation,omitempty"`
	} `json:"data"`
}

//RadiusProfiles - Configured Radius Profiles
type RadiusProfiles struct {
	Meta struct {
		Rc string `json:"rc"`
	} `json:"meta"`
	Data []struct {
		ID          string `json:"_id"`
		AuthServers []struct {
			IP      string `json:"ip"`
			Port    int    `json:"port"`
			XSecret string `json:"x_secret"`
		} `json:"auth_servers"`
		Name             string        `json:"name"`
		SiteID           string        `json:"site_id"`
		AcctServers      []interface{} `json:"acct_servers"`
		AttrNoDelete     bool          `json:"attr_no_delete"`
		AttrHiddenID     string        `json:"attr_hidden_id"`
		UseUsgAuthServer bool          `json:"use_usg_auth_server"`
	} `json:"data"`
}

//RadiusAccounts -  The details of Radius accounts
type RadiusAccounts struct {
	Meta struct {
		Rc string `json:"rc"`
	} `json:"meta"`
	Data []struct {
		ID               string `json:"_id"`
		Name             string `json:"name"`
		XPassword        string `json:"x_password"`
		TunnelType       int    `json:"tunnel_type"`
		TunnelMediumType int    `json:"tunnel_medium_type"`
		Vlan             string `json:"vlan"`
		SiteID           string `json:"site_id"`
	} `json:"data"`
}

//PortForwardRules - Details of configured port forwarding rules
type PortForwardRules struct {
	Meta struct {
		Rc string `json:"rc"`
	} `json:"meta"`
	Data []struct {
		ID      string `json:"_id"`
		Enabled bool   `json:"enabled"`
		Name    string `json:"name"`
		DstPort string `json:"dst_port"`
		Fwd     string `json:"fwd"`
		FwdPort string `json:"fwd_port"`
		Log     bool   `json:"log,omitempty"`
		Src     string `json:"src"`
		Proto   string `json:"proto"`
		SiteID  string `json:"site_id"`
	} `json:"data"`
}

//RFChannels - Details of the wireless channels available and in use
type RFChannels struct {
	Meta struct {
		Rc string `json:"rc"`
	} `json:"meta"`
	Data []struct {
		ChannelsNa       []int  `json:"channels_na"`
		ChannelsNa160    []int  `json:"channels_na_160"`
		ChannelsNa40     []int  `json:"channels_na_40"`
		ChannelsNa40Bcm  []int  `json:"channels_na_40_bcm"`
		ChannelsNa80     []int  `json:"channels_na_80"`
		ChannelsNa80Bcm  []int  `json:"channels_na_80_bcm"`
		ChannelsNaBcm    []int  `json:"channels_na_bcm"`
		ChannelsNaDfs    []int  `json:"channels_na_dfs"`
		ChannelsNaIndoor []int  `json:"channels_na_indoor"`
		ChannelsNg       []int  `json:"channels_ng"`
		ChannelsNg40     []int  `json:"channels_ng_40"`
		ChannelsNg40Bcm  []int  `json:"channels_ng_40_bcm"`
		ChannelsNgBcm    []int  `json:"channels_ng_bcm"`
		Code             string `json:"code"`
		Key              string `json:"key"`
		Name             string `json:"name"`
	} `json:"data"`
}

//CountryCodes - Dictionary of country codes used
type CountryCodes struct {
	Meta struct {
		Rc string `json:"rc"`
	} `json:"meta"`
	Data []struct {
		Code string `json:"code"`
		Name string `json:"name"`
		Key  string `json:"key"`
	} `json:"data"`
}

//LoggedInUser - User currently loggedin
type LoggedInUser struct {
	Meta struct {
		Rc string `json:"rc"`
	} `json:"meta"`
	Data []struct {
		Name                      string        `json:"name"`
		Email                     string        `json:"email"`
		SiteID                    string        `json:"site_id"`
		SiteName                  string        `json:"site_name"`
		SiteRole                  string        `json:"site_role"`
		SitePermissions           []interface{} `json:"site_permissions"`
		SuperSitePermissions      []interface{} `json:"super_site_permissions"`
		LastSiteID                string        `json:"last_site_id"`
		RequiresNewPassword       bool          `json:"requires_new_password"`
		IsSuper                   bool          `json:"is_super"`
		DeviceID                  string        `json:"device_id"`
		AdminID                   string        `json:"admin_id"`
		EmailAlertEnabled         bool          `json:"email_alert_enabled"`
		EmailAlertGroupingEnabled bool          `json:"email_alert_grouping_enabled"`
		EmailAlertGroupingDelay   int           `json:"email_alert_grouping_delay"`
		PushAlertEnabled          bool          `json:"push_alert_enabled"`
		IsProfessionalInstaller   bool          `json:"is_professional_installer"`
		HTMLEmailEnabled          bool          `json:"html_email_enabled"`
		UISettings                struct {
			NeverCheckForUpdate   bool   `json:"neverCheckForUpdate"`
			StatisticsPrefferedTZ string `json:"statisticsPrefferedTZ"`
			StatisticsPreferBps   bool   `json:"statisticsPreferBps"`
			DeviceStatusList      []struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"deviceStatusList"`
			DashboardConfig struct {
				LastActiveDashboardID string `json:"lastActiveDashboardId"`
				Dashboards            struct {
					FiveD1694851A9Cf2195F548752 struct {
						Order int `json:"order"`
					} `json:"5d1694851a9cf2195f548752"`
					FiveF12D85Ae358B452Cd0Eb762 struct {
					} `json:"5f12d85ae358b452cd0eb762"`
				} `json:"dashboards"`
			} `json:"dashboardConfig"`
			Preferences struct {
				AlertsPosition                  string `json:"alertsPosition"`
				AllowHiddenDashboardModules     bool   `json:"allowHiddenDashboardModules"`
				BrowserLogLevel                 string `json:"browserLogLevel"`
				BypassAutoFindDevices           bool   `json:"bypassAutoFindDevices"`
				BypassConfirmAdoptAndUpgrade    bool   `json:"bypassConfirmAdoptAndUpgrade"`
				BypassConfirmBlock              bool   `json:"bypassConfirmBlock"`
				BypassConfirmRestart            bool   `json:"bypassConfirmRestart"`
				BypassConfirmUpgrade            bool   `json:"bypassConfirmUpgrade"`
				DateFormat                      string `json:"dateFormat"`
				EnableNewSettings               bool   `json:"enableNewSettings"`
				IsAppDark                       bool   `json:"isAppDark"`
				IsPropertyPanelFixed            bool   `json:"isPropertyPanelFixed"`
				IsRegularGraphForAirViewEnabled bool   `json:"isRegularGraphForAirViewEnabled"`
				IsResponsive                    bool   `json:"isResponsive"`
				IsSettingsDark                  bool   `json:"isSettingsDark"`
				IsUndockedByDefault             bool   `json:"isUndockedByDefault"`
				NoWhatsNew                      bool   `json:"noWhatsNew"`
				PropertyPanelCollapse           bool   `json:"propertyPanelCollapse"`
				PropertyPanelMultiMode          bool   `json:"propertyPanelMultiMode"`
				RefreshButtonEnabled            bool   `json:"refreshButtonEnabled"`
				RefreshRate                     string `json:"refreshRate"`
				RefreshRateRememberAll          bool   `json:"refreshRateRememberAll"`
				RowsPerPage                     int    `json:"rowsPerPage"`
				ShowAllPanelActions             bool   `json:"showAllPanelActions"`
				TimeFormat                      string `json:"timeFormat"`
				Use24HourTime                   bool   `json:"use24HourTime"`
				UseBrowserTheme                 bool   `json:"useBrowserTheme"`
				UseSettingsPanelView            bool   `json:"useSettingsPanelView"`
				WebsocketEnabled                bool   `json:"websocketEnabled"`
				WithStickyTableActions          bool   `json:"withStickyTableActions"`
			} `json:"preferences"`
		} `json:"ui_settings"`
	} `json:"data"`
}

//NetworkConfig - Network configurations
type NetworkConfig struct {
	Meta struct {
		Rc string `json:"rc"`
	} `json:"meta"`
	ConfigItems []NetworkConfigItem `json:"data"`
}

//NetworkConfigItem - A network configuration item
type NetworkConfigItem struct {
	ID                     string `json:"_id"`
	IsNat                  bool   `json:"is_nat,omitempty"`
	DhcpdDNSEnabled        bool   `json:"dhcpd_dns_enabled,omitempty"`
	Purpose                string `json:"purpose"`
	DhcpdLeasetime         int    `json:"dhcpd_leasetime,omitempty"`
	DhcpdGatewayEnabled    bool   `json:"dhcpd_gateway_enabled,omitempty"`
	DhcpdTimeOffsetEnabled bool   `json:"dhcpd_time_offset_enabled,omitempty"`
	DhcpdStart             string `json:"dhcpd_start,omitempty"`
	DhcpRelayEnabled       bool   `json:"dhcp_relay_enabled,omitempty"`
	DhcpdStop              string `json:"dhcpd_stop,omitempty"`
	Enabled                bool   `json:"enabled,omitempty"`
	DomainName             string `json:"domain_name,omitempty"`
	DhcpdEnabled           bool   `json:"dhcpd_enabled,omitempty"`
	IPSubnet               string `json:"ip_subnet,omitempty"`
	Networkgroup           string `json:"networkgroup,omitempty"`
	Name                   string `json:"name"`
	SiteID                 string `json:"site_id"`
	AttrNoDelete           bool   `json:"attr_no_delete,omitempty"`
	AttrHiddenID           string `json:"attr_hidden_id,omitempty"`
	VlanEnabled            bool   `json:"vlan_enabled,omitempty"`
	Ipv6InterfaceType      string `json:"ipv6_interface_type,omitempty"`
	DhcpdDNS1              string `json:"dhcpd_dns_1,omitempty"`
	DhcpdDNS2              string `json:"dhcpd_dns_2,omitempty"`
	DhcpdDNS3              string `json:"dhcpd_dns_3,omitempty"`
	LteLanEnabled          bool   `json:"lte_lan_enabled,omitempty"`
	UpnpLanEnabled         bool   `json:"upnp_lan_enabled,omitempty"`
	Ipv6PdStart            string `json:"ipv6_pd_start,omitempty"`
	Ipv6PdStop             string `json:"ipv6_pd_stop,omitempty"`
	IgmpSnooping           bool   `json:"igmp_snooping,omitempty"`
	Vlan                   string `json:"vlan,omitempty"`
	DhcpdWpadURL           string `json:"dhcpd_wpad_url,omitempty"`
	DhcpdBootEnabled       bool   `json:"dhcpd_boot_enabled,omitempty"`
	DhcpdNtpEnabled        bool   `json:"dhcpd_ntp_enabled,omitempty"`
	DhcpdTftpServer        string `json:"dhcpd_tftp_server,omitempty"`
	DhcpdUnifiController   string `json:"dhcpd_unifi_controller,omitempty"`
	DhcpguardEnabled       bool   `json:"dhcpguard_enabled,omitempty"`
	DhcpdWinsEnabled       bool   `json:"dhcpd_wins_enabled,omitempty"`
	Ipv6RaEnabled          bool   `json:"ipv6_ra_enabled,omitempty"`
	VpnType                string `json:"vpn_type,omitempty"`
	XIpsecPreSharedKey     string `json:"x_ipsec_pre_shared_key,omitempty"`
	RadiusprofileID        string `json:"radiusprofile_id,omitempty"`
	RequireMschapv2        bool   `json:"require_mschapv2,omitempty"`
	ExposedToSiteVpn       bool   `json:"exposed_to_site_vpn,omitempty"`
	WanNetworkgroup        string `json:"wan_networkgroup,omitempty"`
	WanType                string `json:"wan_type,omitempty"`
	WanIP                  string `json:"wan_ip,omitempty"`
	WanNetmask             string `json:"wan_netmask,omitempty"`
	WanGateway             string `json:"wan_gateway,omitempty"`
	WanDNS1                string `json:"wan_dns1,omitempty"`
	WanDNS2                string `json:"wan_dns2,omitempty"`
	WanTypeV6              string `json:"wan_type_v6,omitempty"`
	ReportWanEvent         bool   `json:"report_wan_event,omitempty"`
	WanLoadBalanceType     string `json:"wan_load_balance_type,omitempty"`
	WanLoadBalanceWeight   int    `json:"wan_load_balance_weight,omitempty"`
	WanVlanEnabled         bool   `json:"wan_vlan_enabled,omitempty"`
	WanVlan                int    `json:"wan_vlan,omitempty"`
	WanEgressQos           string `json:"wan_egress_qos,omitempty"`
	WanSmartqEnabled       bool   `json:"wan_smartq_enabled,omitempty"`
}
