#!/usr/bin/env bash

TEMPLATE_NAME="zabbix-bench"
APPLICATION_NAME="zabbix-bench"
METRIC_PREFIX="metric"
METRICS=200

echo "<?xml version=\"1.0\" encoding=\"UTF-8\"?>
<zabbix_export>
    <version>3.0</version>
    <date>2016-11-14T15:57:58Z</date>
    <groups>
        <group>
            <name>Linux servers</name>
        </group>
        <group>
            <name>Templates</name>
        </group>
    </groups>
    <templates>
        <template>
            <template>$TEMPLATE_NAME</template>
            <name>$TEMPLATE_NAME</name>
            <description/>
            <groups>
                <group>
                    <name>Linux servers</name>
                </group>
                <group>
                    <name>Templates</name>
                </group>
            </groups>
            <applications>
                <application>
                    <name>$APPLICATION_NAME</name>
                </application>
            </applications>
            <items>"
for I in $(seq 1 1 $METRICS)
do
    echo "
                <item>
                    <name>$METRIC_PREFIX-$I</name>
                    <type>2</type>
                    <snmp_community/>
                    <multiplier>0</multiplier>
                    <snmp_oid/>
                    <key>$METRIC_PREFIX-$I</key>
                    <delay>0</delay>
                    <history>90</history>
                    <trends>365</trends>
                    <status>0</status>
                    <value_type>3</value_type>
                    <allowed_hosts/>
                    <units/>
                    <delta>0</delta>
                    <snmpv3_contextname/>
                    <snmpv3_securityname/>
                    <snmpv3_securitylevel>0</snmpv3_securitylevel>
                    <snmpv3_authprotocol>0</snmpv3_authprotocol>
                    <snmpv3_authpassphrase/>
                    <snmpv3_privprotocol>0</snmpv3_privprotocol>
                    <snmpv3_privpassphrase/>
                    <formula>1</formula>
                    <delay_flex/>
                    <params/>
                    <ipmi_sensor/>
                    <data_type>0</data_type>
                    <authtype>0</authtype>
                    <username/>
                    <password/>
                    <publickey/>
                    <privatekey/>
                    <port/>
                    <description/>
                    <inventory_link>0</inventory_link>
                    <applications>
                        <application>
                            <name>$APPLICATION_NAME</name>
                        </application>
                    </applications>
                    <valuemap/>
                    <logtimefmt/>
                </item>"
done

echo "
            </items>
            <discovery_rules/>
            <macros/>
            <templates/>
            <screens/>
        </template>
    </templates>
</zabbix_export>
"
