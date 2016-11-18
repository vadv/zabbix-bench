#!/usr/bin/env bash

TEMPLATE_NAME="zabbix-bench"
GROUP_NAME="Linux servers"
CLIENT_PREFIX="client"
CLIENTS=40000

echo "<?xml version=\"1.0\" encoding=\"UTF-8\"?>
<zabbix_export>
    <version>3.0</version>
    <date>2016-11-14T15:56:40Z</date>
    <groups>
        <group>
            <name>${GROUP_NAME}</name>
        </group>
    </groups>
    <hosts>"

for I in $(seq 1 1 $CLIENTS)
do
    echo "
        <host>
            <host>$CLIENT_PREFIX-$I</host>
            <name>$CLIENT_PREFIX-$I</name>
            <description/>
            <proxy/>
            <status>0</status>
            <ipmi_authtype>-1</ipmi_authtype>
            <ipmi_privilege>2</ipmi_privilege>
            <ipmi_username/>
            <ipmi_password/>
            <tls_connect>1</tls_connect>
            <tls_accept>1</tls_accept>
            <tls_issuer/>
            <tls_subject/>
            <tls_psk_identity/>
            <tls_psk/>
            <templates>
                <template>
                    <name>$TEMPLATE_NAME</name>
                </template>
            </templates>
            <groups>
                <group>
                    <name>${GROUP_NAME}</name>
                </group>
            </groups>
            <interfaces>
                <interface>
                    <default>1</default>
                    <type>1</type>
                    <useip>1</useip>
                    <ip>127.0.0.1</ip>
                    <dns/>
                    <port>10050</port>
                    <bulk>1</bulk>
                    <interface_ref>if1</interface_ref>
                </interface>
            </interfaces>
            <applications/>
            <items/>
            <discovery_rules/>
            <macros/>
            <inventory/>
        </host>"
done

echo "
    </hosts>
</zabbix_export>
"
